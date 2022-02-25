// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"runtime/debug"
	"sync"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/common/monitoring"
	"go.nlx.io/nlx/common/nlxversion"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	managementapi "go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/outway/api"
	"go.nlx.io/nlx/outway/pkg/server"
	"go.nlx.io/nlx/outway/plugins"
)

type loggerHTTPHandler func(logger *zap.Logger, w http.ResponseWriter, r *http.Request)

type Organization struct {
	serialNumber string
	name         string
}

type Outway struct {
	name                string
	ctx                 context.Context
	wg                  *sync.WaitGroup
	organization        *Organization // the organization running this outway
	orgCert             *common_tls.CertificateBundle
	logger              *zap.Logger
	txlogger            transactionlog.TransactionLogger
	directoryClient     directoryapi.DirectoryClient
	httpServer          *http.Server
	grpcServer          *grpc.Server
	monitorService      *monitoring.Service
	requestHTTPHandler  loggerHTTPHandler
	forwardingHTTPProxy *httputil.ReverseProxy
	servicesLock        sync.RWMutex
	servicesHTTP        map[string]HTTPService
	servicesDirectory   map[string]*directoryapi.ListServicesResponse_Service
	plugins             []plugins.Plugin
}

type NewOutwayArgs struct {
	Name              string
	Ctx               context.Context
	Logger            *zap.Logger
	Txlogger          transactionlog.TransactionLogger
	ManagementClient  managementapi.ManagementClient
	MonitoringAddress string
	OrgCert           *common_tls.CertificateBundle
	DirectoryClient   directoryapi.DirectoryClient
	AuthServiceURL    string
	AuthCAPath        string
	UseAsHTTPProxy    bool
}

func (o *Outway) configureAuthorizationPlugin(authCAPath, authServiceURL string) (*plugins.AuthorizationPlugin, error) {
	if authServiceURL == "" {
		return nil, nil
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

	return plugins.NewAuthorizationPlugin(
		ca,
		fmt.Sprintf("%s/auth", authServiceURL),
		http.Client{
			Transport: createHTTPTransport(tlsConfig),
		},
	), nil
}

func (o *Outway) startDirectoryInspector() error {
	// update service for the first time
	err := o.updateServiceList()
	if err != nil {
		o.logger.Error("failed to update the service list from directory on startup.", zap.Error(err))
		return err
	}

	go o.keepServiceListUpToDate()

	return nil
}

func NewOutway(args *NewOutwayArgs) (*Outway, error) {
	cert := args.OrgCert.Certificate()

	if len(cert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	err := common_tls.ValidateSerialNumber(cert.Subject.SerialNumber)
	if err != nil {
		return nil, fmt.Errorf("validation error for subject serial number from cert: %s", err)
	}

	organizationName := cert.Subject.Organization[0]
	organizationSerialNumber := cert.Subject.SerialNumber

	o := &Outway{
		ctx: args.Ctx,
		wg:  &sync.WaitGroup{},
		logger: args.Logger.With(
			zap.String("outway-organization-name", organizationName),
			zap.String("outway-organization-serialnumber", organizationSerialNumber)),
		txlogger: args.Txlogger,
		organization: &Organization{
			serialNumber: cert.Subject.SerialNumber,
			name:         organizationName,
		},
		orgCert: args.OrgCert,

		servicesHTTP:      make(map[string]HTTPService),
		servicesDirectory: make(map[string]*directoryapi.ListServicesResponse_Service),
	}

	if args.UseAsHTTPProxy {
		o.requestHTTPHandler = o.handleHTTPRequestAsProxy
		o.forwardingHTTPProxy = newForwardingProxy()
	} else {
		o.requestHTTPHandler = o.handleHTTPRequest
	}

	authorizationPlugin, err := o.configureAuthorizationPlugin(args.AuthCAPath, args.AuthServiceURL)
	if err != nil {
		return nil, err
	}

	o.monitorService, err = monitoring.NewMonitoringService(args.MonitoringAddress, args.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create monitoring service")
	}

	o.plugins = []plugins.Plugin{
		plugins.NewDelegationPlugin(args.ManagementClient),
		plugins.NewLogRecordPlugin(o.organization.serialNumber, o.txlogger),
		plugins.NewStripHeadersPlugin(o.organization.serialNumber),
	}

	if authorizationPlugin != nil {
		o.plugins = append(o.plugins, authorizationPlugin)
	}

	if args.Name != "" {
		o.name = args.Name
	} else {
		o.name = getFingerPrint(args.OrgCert.Certificate().Raw)
	}

	if args.DirectoryClient == nil {
		return nil, errors.New("directory client must be not nil")
	}

	o.directoryClient = args.DirectoryClient

	outwayService := server.NewOutwayService(
		o.logger,
		args.OrgCert,
	)

	grpcServer := newGRPCServer(o.logger, args.OrgCert)

	api.RegisterOutwayServer(grpcServer, outwayService)

	o.grpcServer = grpcServer

	return o, nil
}

func (o *Outway) Run() error {
	go o.announceToDirectory(context.Background())

	err := o.startDirectoryInspector()
	if err != nil {
		return err
	}

	return nil
}

func newForwardingProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	return &httputil.ReverseProxy{
		Director: director,
	}
}

func newGRPCServer(logger *zap.Logger, cert *common_tls.CertificateBundle) *grpc.Server {
	// setup zap connection for global grpc logging
	// grpc_zap.ReplaceGrpcLogger(logger)

	tlsConfig := cert.TLSConfig(cert.WithTLSClientAuth())
	transportCredentials := credentials.NewTLS(tlsConfig)

	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Warn("recovered from a panic in a grpc request handler", zap.ByteString("stack", debug.Stack()))
			return status.Error(codes.Internal, fmt.Sprintf("%s", p))
		}),
	}

	opts := []grpc.ServerOption{
		grpc.Creds(transportCredentials),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(recoveryOptions...),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(recoveryOptions...),
		),
	}

	return grpc.NewServer(opts...)
}

func (o *Outway) keepServiceListUpToDate() {
	o.wg.Add(1)
	defer o.wg.Done()

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
		case <-o.ctx.Done():
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

func serviceKey(s *directoryapi.ListServicesResponse_Service) string {
	return s.Organization.SerialNumber + "." + s.Name
}

func (o *Outway) createService(
	serviceToImplement *directoryapi.ListServicesResponse_Service,
) {
	// Look for healthy inwayaddresses unless it there is only one
	// known address.
	moreEndpoints := len(serviceToImplement.Inways) > 1
	inways := []directoryapi.Inway{}

	for _, inway := range serviceToImplement.Inways {
		inwayAddress := inway.Address

		healthy := inway.State == directoryapi.Inway_UP

		if healthy {
			// we want to use only healthy endpoints.
			// if there is only one unhealthy endpoint then
			// we use that one endpoint anyway which is useful
			// for testing / setup purposes.

			// Cloned this way to avoid the following error:
			// govet: copylocks: call of append copies lock value: go.nlx.io/nlx/directory-api/directoryapi.Inway contains google.golang.org/protobuf/internal/impl.MessageState contains sync.Mutex
			inways = append(inways, directoryapi.Inway{
				Address: inway.Address,
				State:   inway.State,
			})

			continue
		}

		if !healthy && moreEndpoints {
			o.logger.Info("ignoring unhealthy inway endpoint. we have healthy ones.",
				zap.String("unhealthy endpoint", inwayAddress),
				zap.String("service-organization-serial-number", serviceToImplement.Organization.SerialNumber),
				zap.String("service-organization-name", serviceToImplement.Organization.Name))
			continue
		}

		if !healthy && !moreEndpoints {
			o.logger.Info(
				"inway might not be healthy / reachable by directory / behind firewall",
				zap.String("service-organization-serial-number", serviceToImplement.Organization.SerialNumber),
				zap.String("service-organization-name", serviceToImplement.Organization.Name),
				zap.String("service-name", serviceToImplement.Name),
				zap.String("inway address", inwayAddress),
			)

			// Cloned this way to avoid the following error:
			// govet: copylocks: call of append copies lock value: go.nlx.io/nlx/directory-api/directoryapi.Inway contains google.golang.org/protobuf/internal/impl.MessageState contains sync.Mutex
			inways = append(inways, directoryapi.Inway{
				Address: inway.Address,
				State:   inway.State,
			})

			continue
		}
	}

	rrlbService, err := NewRoundRobinLoadBalancedHTTPService(
		o.logger,
		o.orgCert,
		serviceToImplement.Organization.SerialNumber,
		serviceToImplement.Name,
		inways,
	)
	if err != nil {
		if err == errNoInwaysAvailable {
			o.logger.Debug(
				"service exists but there are no inwayaddresses available",
				zap.String("service-organization-serial-number", serviceToImplement.Organization.SerialNumber),
				zap.String("service-organization-name", serviceToImplement.Organization.Name),
				zap.String("service-name", serviceToImplement.Name))

			return
		}

		o.logger.Error(
			"failed to create new service",
			zap.String("service-organization-serial-number", serviceToImplement.Organization.SerialNumber),
			zap.String("service-organization-name", serviceToImplement.Organization.Name),
			zap.String("service-name", serviceToImplement.Name),
			zap.Error(err))

		return
	}

	service := rrlbService

	o.logger.Debug(
		"implemented service",
		zap.String("service-name", serviceToImplement.Name),
		zap.String("service-organization-serial-number", serviceToImplement.Organization.SerialNumber),
		zap.String("service-organization-name", serviceToImplement.Organization.Name),
	)

	o.servicesLock.Lock()
	o.servicesHTTP[service.FullName()] = service
	o.servicesLock.Unlock()
}

func (o *Outway) updateServiceList() error {
	if o.servicesDirectory == nil {
		o.servicesDirectory = make(map[string]*directoryapi.ListServicesResponse_Service)
	}

	if o.servicesHTTP == nil {
		o.servicesHTTP = make(map[string]HTTPService)
	}

	ctx := context.TODO()

	resp, err := o.directoryClient.ListServices(nlxversion.NewGRPCContext(ctx, "outway"), &emptypb.Empty{})
	if err != nil {
		return errors.Wrap(err, "failed to fetch services from directory")
	}

	// keep track of currently known directory services.
	servicesToKeep := make(map[string]bool)

	for _, serviceToImplement := range resp.Services {
		o.logger.Debug(
			"directory listed service",
			zap.String("service-name", serviceToImplement.Name),
			zap.String("service-organization-serial-number", serviceToImplement.Organization.SerialNumber),
			zap.String("service-organization-name", serviceToImplement.Organization.Name))

		if len(serviceToImplement.Inways) == 0 {
			o.logger.Debug(
				"directory listed service missing inways for:",
				zap.String("service-name", serviceToImplement.Name),
				zap.String("service-organization-serial-number", serviceToImplement.Organization.SerialNumber),
				zap.String("service-organization-name", serviceToImplement.Organization.Name),
			)

			continue
		}

		serviceKey := serviceKey(serviceToImplement)
		_, exists := o.servicesHTTP[serviceKey]

		// if HttpService is used/created before update
		// httpService on changes.
		if exists {
			inwaysFromDirectory := o.servicesDirectory[serviceKey].Inways

			changed := !reflect.DeepEqual(inwaysFromDirectory, serviceToImplement.Inways)

			if changed {
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

	o.monitorService.SetReady()

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
				delete(o.servicesHTTP, serviceKey)
			}
			o.servicesLock.Unlock()
		}
	}
}

func (o *Outway) getService(organizationSerialNumber, service string) HTTPService {
	serviceKey := organizationSerialNumber + "." + service

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

func getFingerPrint(rawCert []byte) string {
	rawSum := sha256.Sum256(rawCert)
	bytes := make([]byte, sha256.Size)

	for i, b := range rawSum {
		bytes[i] = b
	}

	return base64.URLEncoding.EncodeToString(bytes)
}
