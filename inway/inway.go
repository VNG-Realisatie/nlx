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
	"net"
	"net/http"
	"os"
	"strings"
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

	"go.nlx.io/nlx/common/monitoring"
	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/pkg/configapi"
)

// Inway handles incoming requests and holds a list of registered ServiceEndpoints.
// The Inway is responsible for selecting the correct ServiceEndpoint for an incoming request.
type Inway struct {
	logger           *zap.Logger
	organizationName string

	selfAddress string
	roots       *x509.CertPool
	orgKeyPair  *tls.Certificate

	name string

	process *process.Process

	serverTLS *http.Server

	serviceEndpointsLock sync.RWMutex
	serviceEndpoints     map[string]ServiceEndpoint
	stopInwayChannel     chan struct{}

	monitoringService *monitoring.Service

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
	monitoringAddress string,
	tlsOptions orgtls.TLSOptions,
	directoryRegistrationAddress string) (*Inway, error) {
	// parse tls certificate
	roots, orgKeyPair, err := orgtls.Load(tlsOptions)
	if err != nil {
		return nil, err
	}

	orgCert := orgKeyPair.Leaf

	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	err = selfAddressIsInOrgCert(selfAddress, orgCert)
	if err != nil {
		return nil, err
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
		orgKeyPair:  orgKeyPair,

		process: mainProcess,

		serviceEndpoints: make(map[string]ServiceEndpoint),
		stopInwayChannel: make(chan struct{}),
	}

	// setup monitoring service
	i.monitoringService, err = monitoring.NewMonitoringService(monitoringAddress, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create monitoring service")
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

func selfAddressIsInOrgCert(selfAddress string, orgCert *x509.Certificate) error {
	hostname := selfAddress

	if strings.Contains(hostname, ":") {
		host, _, err := net.SplitHostPort(selfAddress)
		if err != nil {
			return errors.Wrapf(err, "Failed to parse selfAddress hostname from '%s'", selfAddress)
		}

		hostname = host
	}

	if hostname == orgCert.Subject.CommonName {
		return nil
	}

	for _, dnsName := range orgCert.DNSNames {
		if hostname == dnsName {
			return nil
		}
	}

	return errors.Errorf("'%s' is not in the list of DNS names of the certificate, %v", selfAddress, orgCert.DNSNames)
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
	i.monitoringService.SetNotReady()
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

// hostname returns the hostname of this inway
func (i *Inway) hostname() string {
	h, err := os.Hostname()

	if err != nil {
		i.logger.Warn("failed to get inway hostname", zap.Error(err))
		return ""
	}

	return h
}
