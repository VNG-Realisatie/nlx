// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"hash/crc64"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory/directoryapi"
)

// Outway handles requests from inside the organization
type Outway struct {
	organizationName string // the organization running this outway

	tlsOptions orgtls.TLSOptions
	tlsRoots   *x509.CertPool

	logger *zap.Logger

	txlogger transactionlog.TransactionLogger

	directoryClient directoryapi.DirectoryClient

	requestFlake *sonyflake.Sonyflake
	ecmaTable    *crc64.Table

	servicesLock sync.RWMutex
	services     map[string]HTTPService // services mapped by <organizationName>.<serviceName>, PoC shortcut in the absence of directory
}

// NewOutway creates a new Outway and sets it up to handle requests.
func NewOutway(logger *zap.Logger, logdb *sqlx.DB, tlsOptions orgtls.TLSOptions, directoryAddress string) (*Outway, error) {
	// load certs and get organization name from cert
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		logger.Fatal("failed to load tls certs", zap.Error(err))
	}
	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}
	organizationName := orgCert.Subject.Organization[0]

	o := &Outway{
		logger:           logger.With(zap.String("outway-organization-name", organizationName)),
		organizationName: organizationName,

		tlsOptions: tlsOptions,
		tlsRoots:   roots,

		requestFlake: sonyflake.NewSonyflake(sonyflake.Settings{}),
		ecmaTable:    crc64.MakeTable(crc64.ECMA),
	}

	// setup transactionlog
	if logdb == nil {
		o.txlogger = transactionlog.NewDiscardTransactionLogger()
	} else {
		o.txlogger, err = transactionlog.NewPostgresTransactionLogger(logdb, transactionlog.DirectionOut)
		if err != nil {
			return nil, errors.Wrap(err, "failed to setup transactionlog")
		}
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
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryAddress, directoryDialOptions...)
	if err != nil {
		logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}
	o.directoryClient = directoryapi.NewDirectoryClient(directoryConn)
	err = o.updateServiceList()
	if err != nil {
		return nil, errors.Wrap(err, "failed to update internal service directory")
	}
	go func() {
		for {
			time.Sleep(5 * time.Second)
			err := o.updateServiceList()
			if err != nil {
				o.logger.Warn("failed to update the service list", zap.Error(err))
			}
		}
	}()

	return o, nil
}

func (o *Outway) updateServiceList() error {
	services := make(map[string]HTTPService)
	resp, err := o.directoryClient.ListServices(context.Background(), &directoryapi.ListServicesRequest{})
	if err != nil {
		return errors.Wrap(err, "failed to fetch services from directory")
	}
	for _, service := range resp.Services {
		// create the service
		s, err := NewSimpleHTTPService(o.logger, o.tlsRoots, o.tlsOptions.OrgCertFile, o.tlsOptions.OrgKeyFile, service.OrganizationName, service.ServiceName, service.InwayAddresses)
		if err != nil {
			if err == errNoInwaysAvailable {
				// skip, we just pretend this service is not existing now
				// TODO(GeertJohan): #208 post-poc we should have the service available but with no inways attached. This can best be done once we have inway loadbalancing.
				continue
			}
			o.logger.Fatal("failed to create new service", zap.String("service-organization-name", service.OrganizationName), zap.String("service-name", service.ServiceName), zap.Error(err))
		}
		services[s.FullName()] = s
	}
	o.servicesLock.Lock()
	defer o.servicesLock.Unlock()
	o.services = services
	return nil
}
