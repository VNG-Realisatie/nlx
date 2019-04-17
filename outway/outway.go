// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"hash/crc64"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
)

// Outway handles requests from inside the organization
type Outway struct {
	wg               *sync.WaitGroup
	organizationName string // the organization running this outway

	tlsOptions orgtls.TLSOptions
	tlsRoots   *x509.CertPool

	logger *zap.Logger

	txlogger transactionlog.TransactionLogger

	directoryInspectionClient inspectionapi.DirectoryInspectionClient

	requestFlake *sonyflake.Sonyflake
	ecmaTable    *crc64.Table

	headersStripList *http.Header

	authorizationSettings *authSettings
	authorizationClient   http.Client

	servicesLock sync.RWMutex
	services     map[string]HTTPService // services mapped by <organizationName>.<serviceName>
}

// NewOutway creates a new Outway and sets it up to handle requests.
func NewOutway(process *process.Process, logger *zap.Logger, logdb *sqlx.DB, tlsOptions orgtls.TLSOptions, directoryInspectionAddress, authServiceURL, authCAPath string) (*Outway, error) {
	// load certs and get organization name from cert
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		logger.Fatal("failed to load tls certs", zap.Error(err))
	}
	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	organizationName := orgCert.Subject.Organization[0]
	logger.Info("loaded certificates for outway", zap.String("outway-organization-name", organizationName))

	o := &Outway{
		wg:               &sync.WaitGroup{},
		logger:           logger.With(zap.String("outway-organization-name", organizationName)),
		organizationName: organizationName,

		tlsOptions: tlsOptions,
		tlsRoots:   roots,

		requestFlake: sonyflake.NewSonyflake(sonyflake.Settings{}),
		ecmaTable:    crc64.MakeTable(crc64.ECMA),
	}

	if len(authServiceURL) > 0 {
		if len(authCAPath) == 0 {
			return nil, fmt.Errorf("authorization service URL set but no CA for authorization provided")
		}
		url, err := url.Parse(authServiceURL)
		if err != nil {
			return nil, err
		}

		if url.Scheme != "https" {
			return nil, fmt.Errorf("scheme of authorization service URL is not 'https'")
		}
		o.authorizationSettings = &authSettings{
			serviceURL: fmt.Sprintf("%s/auth", authServiceURL),
		}

		o.authorizationSettings.ca, err = orgtls.LoadRootCert(authCAPath)
		if err != nil {
			return nil, err
		}

		o.authorizationClient = http.Client{
			Transport: createHTTPTransport(&tls.Config{
				RootCAs: o.authorizationSettings.ca}),
		}
	}

	// setup transactionlog
	if logdb == nil {
		logger.Info("logging to transaction log disabled")
		o.txlogger = transactionlog.NewDiscardTransactionLogger()
	} else {
		o.txlogger, err = transactionlog.NewPostgresTransactionLogger(logger, logdb, transactionlog.DirectionOut)
		if err != nil {
			return nil, errors.Wrap(err, "failed to setup transactionlog")
		}
		logger.Info("transaction logger created")
	}

	orgKeypair, err := tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read tls keypair")
	}
	directoryDialCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{orgKeypair},
		RootCAs:      roots,
	})
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}
	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer directoryConnCtxCancel()
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryInspectionAddress, directoryDialOptions...)
	if err != nil {
		logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}
	o.directoryInspectionClient = inspectionapi.NewDirectoryInspectionClient(directoryConn)
	logger.Info("directory inspection client setup complete", zap.String("directory-inspection-address", directoryInspectionAddress))
	err = o.updateServiceList(process)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update internal service directory")
	}

	go o.keepServiceListUpToDate(process)
	return o, nil
}
