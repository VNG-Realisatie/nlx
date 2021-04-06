// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/monitoring"
	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/inway/grpcproxy"
	"go.nlx.io/nlx/management-api/api"
	external_api "go.nlx.io/nlx/management-api/api/external"
)

// Inway handles incoming requests and holds a list of registered ServiceEndpoints.
// The Inway is responsible for selecting the correct ServiceEndpoint for an incoming request.
type Inway struct {
	name                        string
	organizationName            string
	selfAddress                 string
	managementAPIAddress        string
	listenManagementAddress     string
	orgCertBundle               *common_tls.CertificateBundle
	certBundle                  *common_tls.CertificateBundle
	logger                      *zap.Logger
	process                     *process.Process
	serverTLS                   *http.Server
	monitoringService           *monitoring.Service
	txlogger                    transactionlog.TransactionLogger
	managementClient            api.ManagementClient
	managementProxy             *grpcproxy.Proxy
	directoryRegistrationClient registrationapi.DirectoryRegistrationClient
	plugins                     []plugins.Plugin
	services                    map[string]*plugins.Service
	servicesLock                sync.RWMutex
}

// NewInway creates and prepares a new Inway.
func NewInway(
	ctx context.Context,
	logger *zap.Logger,
	txlogger transactionlog.TransactionLogger,
	name,
	selfAddress string,
	monitoringAddress string,
	managementAPIAddress string,
	listenManagementAddress string,
	orgCertBundle *common_tls.CertificateBundle,
	certBundle *common_tls.CertificateBundle,
	directoryRegistrationAddress string,
) (*Inway, error) {
	orgCert := orgCertBundle.Certificate()

	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	err := selfAddressIsInOrgCert(selfAddress, orgCert)
	if err != nil {
		return nil, err
	}

	if ctx == nil {
		return nil, errors.New("context is nil. needed to close gracefully")
	}

	organizationName := orgCert.Subject.Organization[0]
	logger.Info("loaded certificates for inway", zap.String("inway-organization-name", organizationName))

	i := &Inway{
		logger:           logger.With(zap.String("inway-organization-name", organizationName)),
		organizationName: organizationName,
		txlogger:         txlogger,
		selfAddress:      selfAddress,
		orgCertBundle:    orgCertBundle,
		services:         map[string]*plugins.Service{},
		servicesLock:     sync.RWMutex{},
		plugins: []plugins.Plugin{
			plugins.NewAuthenticationPlugion(),
			plugins.NewDelegationPlugin(),
			plugins.NewAuthorizationPlugin(),
			plugins.NewLogRecordPlugin(),
		},
	}

	// setup monitoring service
	i.monitoringService, err = monitoring.NewMonitoringService(monitoringAddress, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create monitoring service")
	}

	if name != "" {
		i.name = name
	} else {
		i.name = getFingerPrint(orgCert.Raw)
	}

	directoryDialCredentials := credentials.NewTLS(orgCertBundle.TLSConfig())
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}

	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(nlxversion.NewGRPCContext(ctx, "inway"), 1*time.Minute)
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryRegistrationAddress, directoryDialOptions...)

	defer directoryConnCtxCancel()

	if err != nil {
		logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}

	i.directoryRegistrationClient = registrationapi.NewDirectoryRegistrationClient(directoryConn)

	logger.Info("directory registration client setup complete", zap.String("directory-address", directoryRegistrationAddress))

	creds := credentials.NewTLS(certBundle.TLSConfig())

	connCtx, connCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer connCtxCancel()

	i.logger.Info("creating management api connection", zap.String("management api address", managementAPIAddress))
	conn, err := grpc.DialContext(connCtx, managementAPIAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	i.managementClient = api.NewManagementClient(conn)

	p, err := grpcproxy.New(context.TODO(), i.logger, managementAPIAddress, i.orgCertBundle, certBundle)
	if err != nil {
		return nil, err
	}

	p.RegisterService(external_api.GetAccessRequestServiceDesc())
	p.RegisterService(external_api.GetDelegationServiceDesc())

	i.managementProxy = p

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

func (i *Inway) announceToDirectory(ctx context.Context) {
	expBackOff := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Factor: 2,
		Max:    20 * time.Second,
	}

	sleepDuration := 10 * time.Second

	for {
		select {
		case <-ctx.Done():
			i.logger.Info("stopping directory announcement")
			return
		case <-time.After(sleepDuration):
			ctx := context.Background()
			protoServiceDetails := []*registrationapi.RegisterInwayRequest_RegisterService{}

			for _, service := range i.ServicesMap.GetServices() {
				protoServiceDetails = append(protoServiceDetails, &registrationapi.RegisterInwayRequest_RegisterService{
					Name:                        service.Name,
					Internal:                    service.Internal,
					DocumentationUrl:            service.DocumentationURL,
					ApiSpecificationDocumentUrl: service.APISpecificationDocumentURL,
					InsightApiUrl:               service.InsightAPIURL,
					IrmaApiUrl:                  service.IrmaAPIURL,
					PublicSupportContact:        service.PublicSupportContact,
					TechSupportContact:          service.TechSupportContact,
					OneTimeCosts:                service.OneTimeCosts,
					MonthlyCosts:                service.MonthlyCosts,
					RequestCosts:                service.RequestCosts,
				})
			}

			resp, err := i.directoryRegistrationClient.RegisterInway(
				nlxversion.NewGRPCContext(ctx, "inway"),
				&registrationapi.RegisterInwayRequest{
					InwayAddress: i.selfAddress,
					Services:     protoServiceDetails,
				},
			)
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
}
