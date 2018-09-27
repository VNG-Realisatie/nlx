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
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory/directoryapi"
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

	serviceEndpointsLock sync.RWMutex
	serviceEndpoints     map[string]ServiceEndpoint

	txlogger transactionlog.TransactionLogger

	directoryClient directoryapi.DirectoryClient
}

// NewInway creates and prepares a new Inway.
func NewInway(logger *zap.Logger, logdb *sqlx.DB, selfAddress string, tlsOptions orgtls.TLSOptions, directoryAddress string, serviceConfig *config.ServiceConfig) (*Inway, error) {
	// parse tls certificate
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		logger.Fatal("failed to load tls certs", zap.Error(err))
	}
	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}
	organizationName := orgCert.Subject.Organization[0]

	i := &Inway{
		logger:           logger.With(zap.String("inway-organization-name", organizationName)),
		organizationName: organizationName,

		selfAddress: selfAddress,
		roots:       roots,
		orgCertFile: tlsOptions.OrgCertFile,
		orgKeyFile:  tlsOptions.OrgKeyFile,

		serviceConfig: serviceConfig,

		serviceEndpoints: make(map[string]ServiceEndpoint),
	}

	// setup transactionlog
	if logdb == nil {
		i.txlogger = transactionlog.NewDiscardTransactionLogger()
	} else {
		i.txlogger, err = transactionlog.NewPostgresTransactionLogger(logger, logdb, transactionlog.DirectionIn)
		if err != nil {
			return nil, errors.Wrap(err, "failed to setup transactionlog")
		}
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
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryAddress, directoryDialOptions...)
	defer directoryConnCtxCancel()
	if err != nil {
		logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}
	i.directoryClient = directoryapi.NewDirectoryClient(directoryConn)
	return i, nil
}

// AddServiceEndpoint adds an ServiceEndpoint to the inway's internal registry.
func (i *Inway) AddServiceEndpoint(s ServiceEndpoint, documentationURL string, apiSpecificationType string) error {
	i.serviceEndpointsLock.Lock()
	defer i.serviceEndpointsLock.Unlock()
	if _, exists := i.serviceEndpoints[s.ServiceName()]; exists {
		return errors.New("service endpoint for a service with the same name has already been registered")
	}
	i.serviceEndpoints[s.ServiceName()] = s
	i.announceToDirectory(s, documentationURL, apiSpecificationType)
	return nil
}

func (i *Inway) announceToDirectory(s ServiceEndpoint, documentationURL string, apiSpecificationType string) {
	go func() {
		for {
			resp, err := i.directoryClient.RegisterInway(context.Background(), &directoryapi.RegisterInwayRequest{
				InwayAddress: i.selfAddress,
				Services: []*directoryapi.RegisterInwayRequest_RegisterService{
					{
						Name:                 s.ServiceName(),
						DocumentationUrl:     documentationURL,
						ApiSpecificationType: apiSpecificationType,
					},
				},
			})
			if err != nil {
				if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
					i.logger.Info("waiting for directory...")
					time.Sleep(1 * time.Second)
					continue
				}
				i.logger.Error("failed to register to directory", zap.Error(err))
			}
			if resp != nil && resp.Error != "" {
				i.logger.Error(fmt.Sprintf("failed to register to directory: %s", resp.Error))
			}

			// sleep 10 seconds before re-registering
			time.Sleep(10 * time.Second)
		}
	}()
}
