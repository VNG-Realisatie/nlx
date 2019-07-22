// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"sync"
	"time"

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
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/inway/config"
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

	serviceConfig *config.ServiceConfig // This should be removed once we have centralized service management
	process       *process.Process

	serviceEndpointsLock sync.RWMutex
	serviceEndpoints     map[string]ServiceEndpoint

	txlogger transactionlog.TransactionLogger

	directoryRegistrationClient registrationapi.DirectoryRegistrationClient
}

// NewInway creates and prepares a new Inway.
func NewInway(
	logger *zap.Logger,
	logDB *sqlx.DB,
	mainProcess *process.Process,
	selfAddress string,
	tlsOptions orgtls.TLSOptions,
	directoryRegistrationAddress string,
	serviceConfig *config.ServiceConfig) (*Inway, error) {
	// parse tls certificate
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		logger.Fatal("failed to load tls certs", zap.Error(err))
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

		serviceConfig: serviceConfig,
		process:       mainProcess,

		serviceEndpoints: make(map[string]ServiceEndpoint),
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
	directoryDialCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{orgKeypair},
		RootCAs:      roots,
	})
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}
	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryRegistrationAddress, directoryDialOptions...)
	defer directoryConnCtxCancel()
	if err != nil {
		logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}
	i.directoryRegistrationClient = registrationapi.NewDirectoryRegistrationClient(directoryConn)
	logger.Info("directory registration client setup complete", zap.String("directory-address", directoryRegistrationAddress))
	return i, nil
}

// AddServiceEndpoint adds an ServiceEndpoint to the inway's internal registry.
func (i *Inway) AddServiceEndpoint(s ServiceEndpoint,
	serviceDetails config.ServiceDetails) error { //nolint
	if err := i.addServiceEndpointToMap(s); err != nil {
		return err
	}
	i.announceToDirectory(s, &serviceDetails)
	return nil
}

func (i *Inway) addServiceEndpointToMap(s ServiceEndpoint) error {
	i.serviceEndpointsLock.Lock()
	defer i.serviceEndpointsLock.Unlock()

	if _, exists := i.serviceEndpoints[s.ServiceName()]; exists {
		return errors.New("service endpoint for a service with the same name has already been registered")
	}
	i.serviceEndpoints[s.ServiceName()] = s
	return nil
}

func (i *Inway) announceToDirectory(s ServiceEndpoint, serviceDetails *config.ServiceDetails) {
	go func() {
		expBackOff := &backoff.Backoff{
			Min:    100 * time.Millisecond,
			Factor: 2,
			Max:    20 * time.Second,
		}
		shutDownComplete := make(chan struct{})
		if i.process != nil {
			i.process.CloseGracefully(func() error {
				close(shutDownComplete)
				return nil
			})
		}
		sleepDuration := 10 * time.Second
		for {
			select {
			case <-shutDownComplete:
				return
			case <-time.After(sleepDuration):
				resp, err := i.directoryRegistrationClient.RegisterInway(context.Background(), &registrationapi.RegisterInwayRequest{
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
