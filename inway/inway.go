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
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/monitoring"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/inway/grpcproxy"
	"go.nlx.io/nlx/inway/plugins"
	"go.nlx.io/nlx/management-api/api"
)

var nameRegex = regexp.MustCompile(`^[a-zA-Z0-9-]{1,100}$`)

const (
	timeOut               = 30 * time.Second
	keepAlive             = 30 * time.Second
	maxIdleCons           = 100
	IdleConnTimeout       = 20 * time.Second
	TLSHandshakeTimeout   = 10 * time.Second
	ExpectContinueTimeout = 1 * time.Second
)

type Organization struct {
	SerialNumber string
	Name         string
}

type Inway struct {
	name                            string
	organization                    Organization
	address                         string
	managementAPIProxyAddress       string
	listenAddressManagementAPIProxy string
	isOrganizationInway             bool
	orgCertBundle                   *common_tls.CertificateBundle
	logger                          *zap.Logger
	serverTLS                       *http.Server
	monitoringService               *monitoring.Service
	managementClient                api.ManagementServiceClient
	managementProxy                 *grpcproxy.Proxy
	directoryClient                 directoryapi.DirectoryClient
	plugins                         []plugins.Plugin
	services                        map[string]*plugins.Service
	servicesLock                    sync.RWMutex
}

type Params struct {
	Context                         context.Context
	Logger                          *zap.Logger
	Txlogger                        transactionlog.TransactionLogger
	ManagementClient                api.ManagementServiceClient
	ManagementProxy                 *grpcproxy.Proxy
	Name                            string
	Address                         string
	ManagementAPIProxyAddress       string
	MonitoringAddress               string
	ListenAddressManagementAPIProxy string
	OrgCertBundle                   *common_tls.CertificateBundle
	DirectoryClient                 directoryapi.DirectoryClient
	AuthServiceURL                  string
	AuthCAPath                      string
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
		logger: logger.With(zap.String("inway-organization-serial-number", organizationSerialNumber)),
		organization: Organization{
			SerialNumber: organizationSerialNumber,
			Name:         organizationName,
		},
		listenAddressManagementAPIProxy: params.ListenAddressManagementAPIProxy,
		address:                         params.Address,
		managementAPIProxyAddress:       params.ManagementAPIProxyAddress,
		orgCertBundle:                   params.OrgCertBundle,
		managementClient:                params.ManagementClient,
		managementProxy:                 params.ManagementProxy,
		directoryClient:                 params.DirectoryClient,
		services:                        map[string]*plugins.Service{},
		servicesLock:                    sync.RWMutex{},
		plugins: []plugins.Plugin{
			plugins.NewAuthenticationPlugin(),
			plugins.NewDelegationPlugin(),
		},
	}

	authorizationPlugin, err := configureAuthorizationPlugin(params.AuthCAPath, params.AuthServiceURL)
	if err != nil {
		return nil, err
	}

	i.plugins = append(i.plugins, []plugins.Plugin{
		authorizationPlugin,
		plugins.NewLogRecordPlugin(organizationSerialNumber, params.Txlogger),
	}...,
	)

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

func configureAuthorizationPlugin(authCAPath, authServiceURL string) (*plugins.AuthorizationPlugin, error) {
	if authServiceURL == "" {
		return plugins.NewAuthorizationPlugin(&plugins.NewAuthorizationPluginArgs{}), nil
	}

	if authCAPath == "" {
		return nil, fmt.Errorf("authorization service URL set but no CA for authorization provided")
	}

	authURL, err := url.Parse(authServiceURL)
	if err != nil {
		return nil, err
	}

	if authURL.Scheme != "https" {
		return nil, errors.New("scheme of authorization service URL is not 'https'")
	}

	ca, _, err := common_tls.NewCertPoolFromFile(authCAPath)
	if err != nil {
		return nil, err
	}

	tlsConfig := common_tls.NewConfig(common_tls.WithTLS12())
	tlsConfig.RootCAs = ca

	return plugins.NewAuthorizationPlugin(&plugins.NewAuthorizationPluginArgs{
		CA:         ca,
		ServiceURL: authURL.String(),
		AuthorizationClient: &http.Client{
			Transport: createHTTPTransport(tlsConfig),
		},
		AuthServerEnabled: true,
	}), nil
}

func createHTTPTransport(tlsConfig *tls.Config) *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeOut,
			KeepAlive: keepAlive,
		}).DialContext,
		MaxIdleConns:          maxIdleCons,
		IdleConnTimeout:       IdleConnTimeout,
		TLSHandshakeTimeout:   TLSHandshakeTimeout,
		ExpectContinueTimeout: ExpectContinueTimeout,
		TLSClientConfig:       tlsConfig,
	}
}
