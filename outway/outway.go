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
	"reflect"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jpillora/backoff"
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
	process                   *process.Process

	requestFlake *sonyflake.Sonyflake
	ecmaTable    *crc64.Table

	// headersStripList *http.Header

	authorizationSettings *authSettings
	authorizationClient   http.Client

	servicesLock sync.RWMutex
	// services mapped by <organizationName>.<serviceName>
	// httpService
	servicesHTTP map[string]HTTPService
	// directory listing copy
	servicesDirectory map[string]*inspectionapi.ListServicesResponse_Service
}

func loadCertificates(logger *zap.Logger, tlsOptions orgtls.TLSOptions) (*x509.CertPool, string, error) {

	// load certs and get organization name from cert
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		msg := "cannot obtain organization name from self cert"
		logger.Fatal("failed to load tls certs "+msg, zap.Error(err))
		return nil, "", errors.New(msg)

	}
	if len(orgCert.Subject.Organization) != 1 {
		return nil, "", errors.New("cannot obtain organization name from self cert")
	}

	organizationName := orgCert.Subject.Organization[0]
	logger.Info("loaded certificates for outway", zap.String("outway-organization-name", organizationName))

	return roots, organizationName, nil
}

func (o *Outway) validateAuthURL(authCAPath, authServiceURL string) error {
	if authServiceURL == "" {
		return nil
	}

	if authCAPath == "" {
		return fmt.Errorf("authorization service URL set but no CA for authorization provided")
	}

	authURL, err := url.Parse(authServiceURL)
	if err != nil {
		return err
	}

	if authURL.Scheme != "https" {
		return errors.New("scheme of authorization service URL is not 'https'")
	}
	o.authorizationSettings = &authSettings{
		serviceURL: fmt.Sprintf("%s/auth", authServiceURL),
	}

	o.authorizationSettings.ca, err = orgtls.LoadRootCert(authCAPath)
	if err != nil {
		return err
	}

	o.authorizationClient = http.Client{
		Transport: createHTTPTransport(&tls.Config{
			RootCAs: o.authorizationSettings.ca}),
	}
	return nil
}

func (o *Outway) startDirectoryInspector(directoryInspectionAddress string) error {

	orgKeypair, err := tls.LoadX509KeyPair(o.tlsOptions.OrgCertFile, o.tlsOptions.OrgKeyFile)
	if err != nil {
		return errors.Wrap(err, "failed to read tls keypair")
	}
	directoryDialCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{orgKeypair},
		RootCAs:      o.tlsRoots,
	})
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}
	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer directoryConnCtxCancel()
	directoryConn, err := grpc.DialContext(
		directoryConnCtx, directoryInspectionAddress, directoryDialOptions...)
	if err != nil {
		o.logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}
	o.directoryInspectionClient = inspectionapi.NewDirectoryInspectionClient(directoryConn)
	o.logger.Info(
		"directory inspection client setup complete",
		zap.String("directory-inspection-address", directoryInspectionAddress))

	go o.keepServiceListUpToDate()
	return nil
}

// NewOutway creates a new Outway and sets it up to handle requests.
func NewOutway(
	logger *zap.Logger,
	logdb *sqlx.DB,
	mainProcess *process.Process,
	tlsOptions orgtls.TLSOptions,
	directoryInspectionAddress,
	authServiceURL,
	authCAPath string) (*Outway, error) {

	roots, organizationName, err := loadCertificates(logger, tlsOptions)
	if err != nil {
		return nil, err
	}

	if mainProcess == nil {
		return nil, errors.New("process argument is nil. needed enable to close gracefully")
	}

	o := &Outway{
		wg:               &sync.WaitGroup{},
		logger:           logger.With(zap.String("outway-organization-name", organizationName)),
		organizationName: organizationName,

		tlsOptions: tlsOptions,
		tlsRoots:   roots,
		process:    mainProcess,

		requestFlake:      sonyflake.NewSonyflake(sonyflake.Settings{}),
		ecmaTable:         crc64.MakeTable(crc64.ECMA),
		servicesHTTP:      make(map[string]HTTPService),
		servicesDirectory: make(map[string]*inspectionapi.ListServicesResponse_Service),
	}

	err = o.validateAuthURL(authCAPath, authServiceURL)
	if err != nil {
		return nil, err
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

	err = o.startDirectoryInspector(directoryInspectionAddress)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *Outway) keepServiceListUpToDate() {
	o.wg.Add(1)
	defer o.wg.Done()

	// update service for the first time
	err := o.updateServiceList()
	if err != nil {
		o.logger.Error("failed to update the service list from directory on startup.", zap.Error(err))
		o.process.ExitGracefully()
		return
	}

	// update service list every x seconds
	expBackOff := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Factor: 2,
		Max:    20 * time.Second,
	}

	const baseInterval = 30 * time.Second
	interval := baseInterval
	for {
		select {
		case <-o.process.ShutdownRequested:
			return
		case <-time.After(interval):
			err := o.updateServiceList()
			if err != nil {
				o.logger.Warn("failed to update the service list", zap.Error(err))
				// Change interval on retry
				interval = expBackOff.Duration()
			} else {
				interval = baseInterval // Resetting interval to base on success update
				expBackOff.Reset()
			}
		}
	}
}

func serviceKey(s *inspectionapi.ListServicesResponse_Service) string {
	return s.OrganizationName + "." + s.ServiceName
}

func (o *Outway) createService(
	serviceToImplement *inspectionapi.ListServicesResponse_Service,
) {

	// Look for healthy inwayaddresses unless it there is only one
	// known address.
	moreEndpoints := len(serviceToImplement.HealthyStates) > 1
	InwayAddresses := []string{}
	HealthyStates := []bool{}

	for i, healthy := range serviceToImplement.HealthyStates {
		inwayAddress := serviceToImplement.InwayAddresses[i]
		if healthy {
			// we want to use only healthy endpoints.
			// if there is only one unhealthy endpoint then
			// we use that one endpoint anyway which is useful
			// for testing / setup purposes.
			InwayAddresses = append(InwayAddresses, inwayAddress)
			HealthyStates = append(HealthyStates, healthy)
			continue
		}

		if !healthy && moreEndpoints {
			o.logger.Info("ignoring unhealthy inway endpoint. we have healthy ones.", zap.String("unhealthy endpoint", inwayAddress))
			continue
		}

		if !healthy && !moreEndpoints {
			o.logger.Info(
				"inway might not be healthy / reachable by directory / behind firewall",
				zap.String("service-organization-name", serviceToImplement.OrganizationName),
				zap.String("service-name", serviceToImplement.ServiceName),
				zap.String("inway address", inwayAddress),
			)
			InwayAddresses = append(InwayAddresses, inwayAddress)
			HealthyStates = append(HealthyStates, healthy)
			continue
		}
	}

	rrlbService, err := NewRoundRobinLoadBalancedHTTPService(
		o.logger,
		o.tlsRoots,
		o.tlsOptions.OrgCertFile,
		o.tlsOptions.OrgKeyFile,
		serviceToImplement.OrganizationName,
		serviceToImplement.ServiceName,
		InwayAddresses,
		HealthyStates,
	)
	if err != nil {
		if err == errNoInwaysAvailable {
			o.logger.Debug(
				"service exists but there are no inwayaddresses available",
				zap.String("service-organization-name", serviceToImplement.OrganizationName),
				zap.String("service-name", serviceToImplement.ServiceName))
			return
		}
		o.logger.Error(
			"failed to create new service",
			zap.String("service-organization-name", serviceToImplement.OrganizationName),
			zap.String("service-name", serviceToImplement.ServiceName),
			zap.Error(err))
		return
	}

	service := rrlbService
	o.logger.Debug(
		"implemented service",
		zap.String("service-name", serviceToImplement.ServiceName),
		zap.String("service-organization-name", serviceToImplement.OrganizationName),
	)

	o.servicesLock.Lock()
	o.servicesHTTP[service.FullName()] = service
	o.servicesLock.Unlock()

}

func (o *Outway) updateServiceList() error {
	if o.servicesDirectory == nil {
		o.servicesDirectory = make(map[string]*inspectionapi.ListServicesResponse_Service)
	}

	if o.servicesHTTP == nil {
		o.servicesHTTP = make(map[string]HTTPService)
	}

	resp, err := o.directoryInspectionClient.ListServices(context.Background(), &inspectionapi.ListServicesRequest{})
	if err != nil {
		return errors.Wrap(err, "failed to fetch services from directory")
	}

	// keep track of currently known directory services.
	servicesToKeep := make(map[string]bool)

	for _, serviceToImplement := range resp.Services {

		o.logger.Debug(
			"directory listed service",
			zap.String("service-name", serviceToImplement.ServiceName),
			zap.String("service-organization-name", serviceToImplement.OrganizationName))

		if len(serviceToImplement.InwayAddresses) == 0 {
			o.logger.Debug(
				"directory listed service missing inway addresses for:",
				zap.String("service-name", serviceToImplement.ServiceName),
				zap.String("service-organization-name", serviceToImplement.OrganizationName),
			)
			continue
		}

		serviceKey := serviceKey(serviceToImplement)
		_, exists := o.servicesHTTP[serviceKey]

		// if HttpService is used/created before update
		// httpService on changes.
		if exists {
			addressesChange := !reflect.DeepEqual(
				o.servicesDirectory[serviceKey].InwayAddresses,
				serviceToImplement.InwayAddresses)

			healthyChange := !reflect.DeepEqual(
				o.servicesDirectory[serviceKey].HealthyStates,
				serviceToImplement.HealthyStates)

			if addressesChange || healthyChange {
				o.createService(serviceToImplement)
			}
		}

		// update local cache directory list
		o.servicesLock.Lock()
		o.servicesDirectory[serviceKey] = serviceToImplement
		o.servicesLock.Unlock()

		servicesToKeep[serviceKey] = true

	}

	o.cleanUpservices(servicesToKeep)

	return nil
}

// cleanUpservices removes no longer advertised services
func (o *Outway) cleanUpservices(servicesToKeep map[string]bool) {

	for serviceKey := range o.servicesDirectory {
		_, exists := servicesToKeep[serviceKey]
		if !exists {
			o.servicesLock.Lock()
			// service is no longer active in directory.
			delete(o.servicesDirectory, serviceKey)
			_, exists = o.servicesHTTP[serviceKey]
			// remove HttpService if present.
			if exists {
				// remove http service.
				// TODO maybe leave it as cache?
				delete(o.servicesHTTP, serviceKey)
			}
			o.servicesLock.Unlock()
		}
	}
}

func (o *Outway) getService(organization, service string) HTTPService {
	serviceKey := organization + "." + service
	o.servicesLock.RLock()
	httpService := o.servicesHTTP[serviceKey]
	o.servicesLock.RUnlock()

	if httpService == nil {
		// create the HttpService if possible.
		directoryService := o.servicesDirectory[serviceKey]
		if directoryService != nil {
			o.createService(directoryService)
			o.servicesLock.RLock()
			httpService = o.servicesHTTP[serviceKey]
			o.servicesLock.RUnlock()
		}
	}
	return httpService
}
