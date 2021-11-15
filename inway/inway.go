// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/monitoring"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/inway/grpcproxy"
	"go.nlx.io/nlx/inway/plugins"
	"go.nlx.io/nlx/management-api/api"
)

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-]{1,100}$`)

type Organization struct {
	SerialNumber string
	Name         string
}

type Inway struct {
	name                        string
	organization                Organization
	address                     string
	listenManagementAddress     string
	isOrganizationInway         bool
	orgCertBundle               *common_tls.CertificateBundle
	logger                      *zap.Logger
	serverTLS                   *http.Server
	monitoringService           *monitoring.Service
	managementClient            api.ManagementClient
	managementProxy             *grpcproxy.Proxy
	directoryRegistrationClient registrationapi.DirectoryRegistrationClient
	plugins                     []plugins.Plugin
	services                    map[string]*plugins.Service
	servicesLock                sync.RWMutex
}

type Params struct {
	Context                     context.Context
	Logger                      *zap.Logger
	Txlogger                    transactionlog.TransactionLogger
	ManagementClient            api.ManagementClient
	ManagementProxy             *grpcproxy.Proxy
	Name                        string
	Address                     string
	MonitoringAddress           string
	ListenManagementAddress     string
	OrgCertBundle               *common_tls.CertificateBundle
	DirectoryRegistrationClient registrationapi.DirectoryRegistrationClient
}

func NewInway(params *Params) (*Inway, error) {
	logger := params.Logger

	if logger == nil {
		logger = zap.NewNop()
	}

	if !nameRegex.MatchString(params.Name) {
		return nil, errors.New("a valid name is required (alphanumeric & dashes, max. 100 characters)")
	}

	orgCert := params.OrgCertBundle.Certificate()

	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	err := addressIsInOrgCert(params.Address, orgCert)
	if err != nil {
		return nil, err
	}

	if params.Context == nil {
		return nil, errors.New("context is nil. needed to close gracefully")
	}

	organizationName := orgCert.Subject.Organization[0]
	organizationSerialNumber := orgCert.Subject.SerialNumber

	logger.Info("loaded certificates for inway", zap.String("inway-organization-serial-number", organizationSerialNumber), zap.String("inway-organization-name", organizationName))

	i := &Inway{
		logger: logger.With(zap.String("inway-organization-serial-number", organizationName)),
		organization: Organization{
			SerialNumber: organizationSerialNumber,
			Name:         organizationName,
		},
		listenManagementAddress:     params.ListenManagementAddress,
		address:                     params.Address,
		orgCertBundle:               params.OrgCertBundle,
		managementClient:            params.ManagementClient,
		managementProxy:             params.ManagementProxy,
		directoryRegistrationClient: params.DirectoryRegistrationClient,
		services:                    map[string]*plugins.Service{},
		servicesLock:                sync.RWMutex{},
		plugins: []plugins.Plugin{
			plugins.NewAuthenticationPlugin(),
			plugins.NewDelegationPlugin(),
			plugins.NewAuthorizationPlugin(),
			plugins.NewLogRecordPlugin(organizationSerialNumber, params.Txlogger),
		},
	}

	// setup monitoring service
	i.monitoringService, err = monitoring.NewMonitoringService(params.MonitoringAddress, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create monitoring service")
	}

	if params.Name != "" {
		i.name = params.Name
	} else {
		i.name = getFingerPrint(orgCert.Raw)
	}

	return i, nil
}

func addressIsInOrgCert(address string, orgCert *x509.Certificate) error {
	hostname := address

	if strings.Contains(hostname, ":") {
		host, _, err := net.SplitHostPort(address)
		if err != nil {
			return errors.Wrapf(err, "failed to parse address hostname from '%s'", address)
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

	return errors.Errorf("'%s' is not in the list of DNS names of the certificate, %v", address, orgCert.DNSNames)
}

func getFingerPrint(rawCert []byte) string {
	rawSum := sha256.Sum256(rawCert)
	bytes := make([]byte, sha256.Size)

	for i, b := range rawSum {
		bytes[i] = b
	}

	return base64.URLEncoding.EncodeToString(bytes)
}
