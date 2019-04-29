// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"

	"google.golang.org/grpc/peer"

	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc"

	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/txlog-db/dbversion"

	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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
	orgKeyPair  tls.Certificate
	orgCertFile string
	orgKeyFile  string

	serviceEndpointsLock sync.RWMutex
	serviceEndpoints     map[string]ServiceEndpoint

	txlogger transactionlog.TransactionLogger

	directoryRegistrationClient registrationapi.DirectoryRegistrationClient

	configAPIClient configapi.ConfigApiClient

	logDB *sqlx.DB

	server http.Server

	certFingerPrint string

	configChan chan config.InwayConfig
	stopChan   chan struct{}
}

// NewInway creates and prepares a new Inway.
func NewInway(logger *zap.Logger, tlsOptions orgtls.TLSOptions, configAPIAddress string) (*Inway, error) {
	// parse tls certificate
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		logger.Fatal("failed to load tls certs", zap.Error(err))
	}
	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	organizationName := orgCert.Subject.Organization[0]
	logger.Info("loaded certificates for inway", zap.String("inway-organization-name", organizationName))
	i := &Inway{
		logger:           logger.With(zap.String("inway-organization-name", organizationName)),
		organizationName: organizationName,
		roots:            roots,
		orgCertFile:      tlsOptions.OrgCertFile,
		orgKeyFile:       tlsOptions.OrgKeyFile,
		serviceEndpoints: make(map[string]ServiceEndpoint),
	}

	i.orgKeyPair, err = tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read tls keypair")
	}

	rawSum := sha1.Sum(i.orgKeyPair.Certificate[0])
	bytes := make([]byte, 20)
	for i, b := range rawSum {
		bytes[i] = b
	}

	i.certFingerPrint = base64.URLEncoding.EncodeToString(bytes)
	i.logger.Info("inway fingerprint", zap.String("fingerprint", i.certFingerPrint))

	dialCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{i.orgKeyPair},
		RootCAs:      roots,
	})
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time: 10 * time.Second,
		}),
	}

	i.logger.Info("dailing config API", zap.String("config-api address", configAPIAddress))
	con, err := grpc.Dial(configAPIAddress, dialOptions...)
	if err != nil {
		return nil, err
	}

	i.configAPIClient = configapi.NewConfigApiClient(con)

	return i, nil
}

// Start will start the inway
func (i *Inway) Start(p *process.Process) error {
	i.logger.Info("starting inway")
	stream := i.setupConfigStream()
	for {
		err := i.listenToStream(stream)
		if err != nil {
			stream = i.setupConfigStream()
		}
	}
	return nil
}

func (i *Inway) listenToStream(s configapi.ConfigApi_CreateConfigStreamClient) error {
	i.logger.Info("waiting for new config")
	msg, err := s.Recv()
	if err != nil {
		i.logger.Error("error receiving from stream. reconnecting", zap.Error(err))
		return err
	}

	i.logger.Info("received new config", zap.String("config", msg.Config))
	conf, err := i.parseConfig(msg.Config)
	if err != nil {
		i.logger.Error("error parsing config", zap.Error(err))
		return err
	}
	err = i.close()
	if err != nil {
		i.logger.Error("error closing inway for config change", zap.Error(err))
		return err
	}

	err = i.applyConfig(conf.Config)
	if err != nil {
		i.logger.Error("error applying config", zap.Error(err))
		return err
	}

	return nil
}

func (i *Inway) setupConfigStream() configapi.ConfigApi_CreateConfigStreamClient {
	expBackOff := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Factor: 2,
		Max:    20 * time.Second,
	}

	for {
		stream, err := i.configAPIClient.CreateConfigStream(context.Background(), &configapi.CreateConfigStreamRequest{
			ComponentName: i.certFingerPrint,
			ComponentKind: "inway",
		})

		if err != nil {
			if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
				i.logger.Info("waiting for config-api...", zap.Error(err))
				time.Sleep(expBackOff.Duration())
				continue
			}

			i.logger.Error("error setting up config API stream", zap.Error(err))
			time.Sleep(expBackOff.Duration())
			continue
		}

		organization, err := getOrganizationFromStream(stream)
		if err != nil {
			i.logger.Error("cannot parse organization from stream", zap.Error(err))
			stream.CloseSend()
			continue
		}

		if organization != i.organizationName {
			i.logger.Error("organization names do not match")
			stream.CloseSend()
			continue
		}

		return stream
	}

}

func getOrganizationFromStream(s configapi.ConfigApi_CreateConfigStreamClient) (string, error) {
	peer, ok := peer.FromContext(s.Context())
	if !ok {
		return "", fmt.Errorf("cannot parse peer from context")
	}

	tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
	if !ok {
		return "", fmt.Errorf("cannot parse tls info")
	}

	if len(tlsInfo.State.PeerCertificates) == 0 {
		return "", fmt.Errorf("no certificate")
	}

	if len(tlsInfo.State.PeerCertificates[0].Subject.Organization) == 0 {
		return "", fmt.Errorf("no organization set in certificate")
	}
	return tlsInfo.State.PeerCertificates[0].Subject.Organization[0], nil
}

func (i *Inway) parseConfig(c string) (*config.Config, error) {
	config, err := config.ParseConfig(i.logger, []byte(c))
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (i *Inway) close() error {
	if i.stopChan == nil {
		return nil
	}
	if i.logDB != nil {
		i.logDB.Close()
	}
	close(i.stopChan)
	return nil
}

func (i *Inway) applyConfig(config *config.InwayConfig) error {
	err := i.setupTransactionLog(config.TransactionLogDSN, config.DisableLogging)
	if err != nil {
		return err
	}

	err = i.setupDirectory(config.DirectoryAddress)
	if err != nil {
		return err
	}

	for _, api := range config.APIS {
		i.logger.Info("loaded API from config", zap.String("service-name", api.Name))
		i.logger.Debug("api configuration details", zap.String("service-name", api.Name), zap.String("endpoint-url", api.EndpointURL),
			zap.String("root-ca-path", api.CACertPath), zap.String("authorization-model", api.Authorization.Mode),
			zap.String("api-spec-url", api.APISpecificationDocumentURL), zap.Bool("internal", api.Internal),
			zap.String("public-support-contact", api.PublicSupportContact), zap.String("tech-support-contact", api.TechSupportContact))
		var rootCrt *x509.CertPool
		if len(api.CACertPath) > 0 {
			rootCrt, err = orgtls.LoadRootCert(api.CACertPath)
			if err != nil {
				return err
			}
		}
		endpoint, err := i.NewHTTPServiceEndpoint(i.logger, api.Name, api.EndpointURL, &tls.Config{RootCAs: rootCrt})
		if err != nil {
			return err
		}
		switch api.Authorization.Mode {
		case "none", "":
			endpoint.SetAuthorizationPublic()
		case "whitelist":
			endpoint.SetAuthorizationWhitelist(api.Authorization.Organizations)
		default:
			return fmt.Errorf(`invalid authorization model "%s" for service "%s"`, api.Authorization.Mode, api.Name)
		}

		i.AddServiceEndpoint(nil, endpoint, api)
	}

	go func() {
		i.logger.Info("Config applied. Starting to listen for new requests", zap.String("listen-address", config.ListenAddress))
		if err := i.ListenAndServeTLS(nil, config.ListenAddress); err != nil {
			if err != http.ErrServerClosed {
				i.logger.Fatal("listen & serve error", zap.Error(err))
			}
		}
	}()

	i.stopChan = make(chan struct{})
	return nil
}

func (i *Inway) setupTransactionLog(postgresDSN string, disableTxLog bool) error {
	if disableTxLog {
		i.logger.Info("logging to transaction-log disabled")
		i.txlogger = transactionlog.NewDiscardTransactionLogger()
		return nil
	}

	if !disableTxLog {
		var err error
		i.logDB, err = sqlx.Open("postgres", postgresDSN)
		if err != nil {
			return err
		}
		i.logDB.MapperFunc(xstrings.ToSnakeCase)
		dbversion.WaitUntilLatestTxlogDBVersion(i.logger, i.logDB.DB)
		i.txlogger, err = transactionlog.NewPostgresTransactionLogger(i.logger, i.logDB, transactionlog.DirectionIn)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Inway) setupDirectory(directoryAddress string) error {
	directoryDialCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{i.orgKeyPair},
		RootCAs:      i.roots,
	})
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}
	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryAddress, directoryDialOptions...)
	defer directoryConnCtxCancel()
	if err != nil {
		return err
	}
	i.directoryRegistrationClient = registrationapi.NewDirectoryRegistrationClient(directoryConn)
	i.logger.Info("directory registration client setup complete", zap.String("directory-address", directoryAddress))
	return err
}

// AddServiceEndpoint adds an ServiceEndpoint to the inway's internal registry.
func (i *Inway) AddServiceEndpoint(p *process.Process, s ServiceEndpoint, serviceDetails config.APIDetails) error {
	if err := i.addServiceEndpointToMap(s); err != nil {
		return err
	}
	i.announceToDirectory(p, s, serviceDetails)
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

func (i *Inway) announceToDirectory(p *process.Process, s ServiceEndpoint, serviceDetails config.APIDetails) {
	go func() {
		expBackOff := &backoff.Backoff{
			Min:    100 * time.Millisecond,
			Factor: 2,
			Max:    20 * time.Second,
		}
		sleepDuration := 10 * time.Second
		for {
			select {
			case <-i.stopChan:
				i.logger.Debug("received stop signal stopping directory announcement")
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
