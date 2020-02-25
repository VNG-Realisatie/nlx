// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"go.nlx.io/nlx/common/nlxversion"

	"github.com/jmoiron/sqlx"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

// Inway handles incoming requests and holds a list of registered ServiceEndpoints.
// The Inway is responsible for selecting the correct ServiceEndpoint for an incoming request.
type Inway struct {
	logger           *zap.Logger
	organizationName string

	selfAddress string
	roots       *x509.CertPool
	orgCertFile string
	orgKeyFile  string

	name string

	process *process.Process

	serviceEndpointsLock sync.RWMutex
	serviceEndpoints     map[string]ServiceEndpoint
	stopInwayChannel     chan struct{}

	txlogger transactionlog.TransactionLogger

	configAPIClient configapi.ConfigApiClient

	directoryRegistrationClient registrationapi.DirectoryRegistrationClient
}

// NewInway creates and prepares a new Inway.
func NewInway(
	logger *zap.Logger,
	logDB *sqlx.DB,
	mainProcess *process.Process,
	name,
	selfAddress string,
	tlsOptions orgtls.TLSOptions,
	directoryRegistrationAddress string) (*Inway, error) {
	// parse tls certificate
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		return nil, err
	}
	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	if selfAddress == "" {
		logger.Info("\n\n inway selfaddress is empty \n\n")
	}

	if mainProcess == nil {
		return nil, errors.New("process argument is nil. needed to close gracefully")
	}

	organizationName := orgCert.Subject.Organization[0]
	logger.Info("loaded certificates for inway", zap.String("inway-organization-name", organizationName))
	i := &Inway{
		logger:           logger.With(zap.String("inway-organization-name", organizationName)),
		organizationName: organizationName,

		selfAddress: selfAddress,
		roots:       roots,
		orgCertFile: tlsOptions.OrgCertFile,
		orgKeyFile:  tlsOptions.OrgKeyFile,

		process: mainProcess,

		serviceEndpoints: make(map[string]ServiceEndpoint),
		stopInwayChannel: make(chan struct{}),
	}

	// setup transactionlog
	if logDB == nil {
		logger.Info("logging to transaction-log disabled")
		i.txlogger = transactionlog.NewDiscardTransactionLogger()
	} else {
		i.txlogger, err = transactionlog.NewPostgresTransactionLogger(logger, logDB, transactionlog.DirectionIn)
		if err != nil {
			return nil, errors.Wrap(err, "failed to setup transactionlog")
		}
		logger.Info("transaction logger created")
	}

	// setup directory client
	orgKeypair, err := tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read tls keypair")
	}

	if name != "" {
		i.name = name
	} else {
		i.name = getFingerPrint(orgKeypair.Certificate[0])
	}

	mainProcess.CloseGracefully(func() error {
		i.stop()
		return nil
	})

	directoryDialCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{orgKeypair},
		RootCAs:      roots,
	})
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}
	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(nlxversion.NewContext("inway"), 1*time.Minute)
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryRegistrationAddress, directoryDialOptions...)
	defer directoryConnCtxCancel()
	if err != nil {
		logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}
	i.directoryRegistrationClient = registrationapi.NewDirectoryRegistrationClient(directoryConn)
	logger.Info("directory registration client setup complete", zap.String("directory-address", directoryRegistrationAddress))
	return i, nil
}

func getFingerPrint(rawCert []byte) string {
	rawSum := sha256.Sum256(rawCert)
	bytes := make([]byte, sha256.Size)
	for i, b := range rawSum {
		bytes[i] = b
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

// stop will stop the announcement of services and the config retrieval process (if a configAPI is configured)
func (i *Inway) stop() {
	close(i.stopInwayChannel)
}

func (i *Inway) announceToDirectory(s ServiceEndpoint) {
	go func() {
		expBackOff := &backoff.Backoff{
			Min:    100 * time.Millisecond,
			Factor: 2,
			Max:    20 * time.Second,
		}

		sleepDuration := 10 * time.Second
		for {
			select {
			case <-i.stopInwayChannel:
				i.logger.Info("stopping directory announcement", zap.String("service-name", s.ServiceName()))
				return
			case <-time.After(sleepDuration):
				serviceDetails := s.ServiceDetails()
				resp, err := i.directoryRegistrationClient.RegisterInway(nlxversion.NewContext("inway"), &registrationapi.RegisterInwayRequest{
					InwayAddress: i.selfAddress,
					Services: []*registrationapi.RegisterInwayRequest_RegisterService{
						{
							Name:                        s.ServiceName(),
							Internal:                    serviceDetails.Internal,
							DocumentationUrl:            serviceDetails.DocumentationURL,
							ApiSpecificationDocumentUrl: serviceDetails.APISpecificationDocumentURL,
							InsightApiUrl:               serviceDetails.InsightAPIURL,
							IrmaApiUrl:                  serviceDetails.IrmaAPIURL,
							PublicSupportContact:        serviceDetails.PublicSupportContact,
							TechSupportContact:          serviceDetails.TechSupportContact,
						},
					},
				})
				if err != nil {
					if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
						i.logger.Info("waiting for directory...", zap.Error(err))
						sleepDuration = expBackOff.Duration()
						continue
					}
					i.logger.Error("failed to register to directory", zap.Error(err))
				}
				if resp != nil && resp.Error != "" {
					i.logger.Error(fmt.Sprintf("failed to register to directory: %s", resp.Error))
				}
				i.logger.Info("directory registration successful")
				sleepDuration = 10 * time.Second
				expBackOff.Reset()
			}

		}
	}()
}
