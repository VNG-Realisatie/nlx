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

	"github.com/VNG-Realisatie/nlx/common/orgtls"
	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"github.com/pkg/errors"
	"go.uber.org/zap"
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

	serviceEndpointsLock sync.RWMutex
	serviceEndpoints     map[string]ServiceEndpoint

	directoryClient directoryapi.DirectoryClient
}

// NewInway creates and prepares a new Inway.
func NewInway(logger *zap.Logger, selfAddress string, tlsOptions orgtls.TLSOptions, directoryAddress string) (*Inway, error) {
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		logger.Fatal("failed to load tls certs", zap.Error(err))
	}
	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
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

		serviceEndpoints: make(map[string]ServiceEndpoint),
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
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryAddress, directoryDialOptions...)
	defer directoryConnCtxCancel()
	if err != nil {
		logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}
	i.directoryClient = directoryapi.NewDirectoryClient(directoryConn)
	return i, nil
}

// AddServiceEndpoint adds an ServiceEndpoint to the inway's internal registry.
func (i *Inway) AddServiceEndpoint(s ServiceEndpoint, documentationURL string) error {
	i.serviceEndpointsLock.Lock()
	defer i.serviceEndpointsLock.Unlock()
	if _, exists := i.serviceEndpoints[s.ServiceName()]; exists {
		return errors.New("service endpoint for a service with the same name has already been registered")
	}
	i.serviceEndpoints[s.ServiceName()] = s
	i.announceToDirectory(s, documentationURL)
	return nil
}

func (i *Inway) announceToDirectory(s ServiceEndpoint, documentationURL string) {
	for {
		resp, err := i.directoryClient.RegisterInway(context.Background(), &directoryapi.RegisterInwayRequest{
			InwayAddress:     i.selfAddress,
			ServiceNames:     []string{s.ServiceName()},
			DocumentationUrl: documentationURL,
		})
		if err != nil {

			if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
				i.logger.Info("waiting for director...")
				time.Sleep(100 * time.Millisecond)
				continue
			}
			i.logger.Error("failed to register to directory", zap.Error(err))
		}
		if resp != nil && resp.Error != "" {
			i.logger.Error(fmt.Sprintf("failed to register to directory: %s", resp.Error))
		}
		break
	}
}
