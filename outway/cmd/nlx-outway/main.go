// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/cmd"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/strings"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/common/version"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/outway"
)

var options struct {
	Name                        string `long:"name" env:"OUTWAY_NAME" description:"Name of the outway. Every outway should have a unique name within the organization." required:"true"`
	ListenAddress               string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8080" description:"Address for the outway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenHTTPS                 bool   `long:"listen-https" env:"LISTEN_HTTPS" description:"Enable HTTPS on the ListenAddress" required:"false"`
	ListenAddressAPI            string `long:"listen-address-api" env:"LISTEN_ADDRESS_API" default:"127.0.0.1:8082" description:"Address for the outway api server to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	AddressAPI                  string `long:"self-address-api" env:"SELF_ADDRESS_API" description:"The address that the management API can use to reach the api of the outway" required:"true"`
	MonitoringAddress           string `long:"monitoring-address" env:"MONITORING_ADDRESS" default:"127.0.0.1:8081" description:"Address for the outway monitoring endpoints to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	DirectoryInspectionAddress  string `long:"directory-inspection-address" env:"DIRECTORY_INSPECTION_ADDRESS" description:"Address for the directory where this outway can fetch the service list"`
	DirectoryAddress            string `long:"directory-address" env:"DIRECTORY_ADDRESS" description:"Address for the directory where this outway can fetch the service list"`
	UseAsHTTPProxy              bool   `long:"use-as-http-proxy" env:"USE_AS_HTTP_PROXY" description:"An experimental flag which when true makes the outway function as an HTTP proxy"`
	ManagementAPIAddress        string `long:"management-api-address" env:"MANAGEMENT_API_ADDRESS" description:"The address of the NLX Management API" required:"true"`
	DisableLogdb                bool   `long:"disable-logdb" env:"DISABLE_LOGDB" description:"Disable logdb connections"`
	PostgresDSN                 string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`
	TxLogAPIAddress             string `long:"tx-log-api-address" env:"TX_LOG_API_ADDRESS" description:"The address of the transaction log API" required:"false"`
	AuthorizationServiceAddress string `long:"authorization-service-address" env:"AUTHORIZATION_SERVICE_ADDRESS" description:"Address of the authorization service. If set calls will go through the authorization service before being send to the inway"`
	AuthorizationCA             string `long:"authorization-root-ca" env:"AUTHORIZATION_ROOT_CA" description:"absolute path to root CA used to verify auth service certificate"`
	ServerCertFile              string `long:"tls-server-cert" env:"TLS_SERVER_CERT" description:"Path to a cert .pem, used for the HTTPS server" required:"false"`
	ServerKeyFile               string `long:"tls-server-key" env:"TLS_SERVER_KEY" description:"Path the a key .pem, used for the HTTPS server" required:"false"`

	logoptions.LogOptions
	cmd.TLSOrgOptions
	cmd.TLSOptions
}

func main() {
	parseOptions()

	p := process.NewProcess()

	// Setup new zap logger
	logger, err := options.LogOptions.ZapConfig().Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.OrgKeyFile); errValidate != nil {
		logger.Warn("invalid organization key permissions", zap.Error(errValidate), zap.String("file-path", options.OrgCertFile))
	}

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.KeyFile); errValidate != nil {
		logger.Warn("invalid internal PKI key permissions", zap.Error(errValidate), zap.String("file-path", options.KeyFile))
	}

	orgCert, err := common_tls.NewBundleFromFiles(options.OrgCertFile, options.OrgKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("unable to load organization certificate and key", zap.Error(err))
	}

	cert, err := common_tls.NewBundleFromFiles(options.CertFile, options.KeyFile, options.RootCertFile)
	if err != nil {
		logger.Fatal("unable to load internal PKI certificate and key", zap.Error(err))
	}

	var txLogger transactionlog.TransactionLogger
	if options.DisableLogdb {
		txLogger = transactionlog.NewDiscardTransactionLogger()
	} else {
		txLogger, err = setupTransactionLogger(logger, options.PostgresDSN, options.TxLogAPIAddress, cert)
		if err != nil {
			logger.Fatal("unable to setup the transaction logger", zap.Error(err))
		}
	}

	var serverCertificate *tls.Certificate

	if options.ListenHTTPS {
		if options.ServerCertFile == "" || options.ServerKeyFile == "" {
			logger.Fatal("server certificate and key are required")
		}

		cert, certErr := tls.LoadX509KeyPair(options.ServerCertFile, options.ServerKeyFile)
		if certErr != nil {
			logger.Fatal("failed to load server certificate", zap.Error(err))
		}

		serverCertificate = &cert
	}

	managementAPIConn, err := grpc.DialContext(context.TODO(), options.ManagementAPIAddress, grpc.WithTransportCredentials(credentials.NewTLS(cert.TLSConfig())))
	if err != nil {
		logger.Fatal("failed to connection to the management api", zap.Error(err))
	}

	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(nlxversion.NewGRPCContext(context.Background(), "outway"), 1*time.Minute)
	defer directoryConnCtxCancel()

	directoryConn, err := grpc.DialContext(directoryConnCtx, options.DirectoryAddress,
		grpc.WithTransportCredentials(credentials.NewTLS(orgCert.TLSConfig())),
	)
	if err != nil {
		logger.Fatal("failed to setup connection to the directory api", zap.Error(err))
	}

	ow, err := outway.New(&outway.NewOutwayArgs{
		Name:                options.Name,
		AddressAPI:          options.AddressAPI,
		Ctx:                 context.Background(),
		Logger:              logger,
		Txlogger:            txLogger,
		ManagementAPIClient: api.NewManagementClient(managementAPIConn),
		MonitoringAddress:   options.MonitoringAddress,
		OrgCert:             orgCert,
		InternalCert:        cert,
		DirectoryClient:     directoryapi.NewDirectoryClient(directoryConn),
		AuthServiceURL:      options.AuthorizationServiceAddress,
		AuthCAPath:          options.AuthorizationCA,
		UseAsHTTPProxy:      options.UseAsHTTPProxy,
	})
	if err != nil {
		logger.Fatal("failed to initialize the outway", zap.Error(err))
	}

	ctxAnnouncementsCancel, cancelAnnouncements := context.WithCancel(context.Background())
	go func() {
		err = ow.Run(ctxAnnouncementsCancel)
		if err != nil {
			logger.Fatal("error running outway", zap.Error(err))
		}

		err = ow.RunServer(options.ListenAddress, options.ListenAddressAPI, serverCertificate)
		if err != nil {
			logger.Fatal("error running outway server", zap.Error(err))
		}
	}()

	p.Wait()

	logger.Info("starting graceful shutdown")
	cancelAnnouncements()

	gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ow.Shutdown(gracefulCtx)

	err = managementAPIConn.Close()
	if err != nil {
		logger.Error("could not close management API grpc connection", zap.Error(err))
	}

	err = directoryConn.Close()
	if err != nil {
		logger.Error("could not close directory API grpc connection", zap.Error(err))
	}

	err = txLogger.Close()
	if err != nil {
		logger.Error("could not close log db", zap.Error(err))
	}
}

func parseOptions() {
	args, err := flags.Parse(&options)
	if err != nil {
		if et, ok := err.(*flags.Error); ok {
			if et.Type == flags.ErrHelp {
				return
			}
		}

		log.Fatalf("error parsing flags: %v", err)
	}

	if len(args) > 0 {
		log.Fatalf("unexpected arguments: %v", args)
	}

	// Deprecated flags
	if options.DirectoryInspectionAddress != "" {
		log.Println("Flag --directory-inspection-address has been deprecated, please use '--directory-address' instead")
	}

	if options.DirectoryAddress == "" {
		if options.DirectoryInspectionAddress == "" {
			log.Fatal(errors.New("directory-address is required"))
		}

		options.DirectoryAddress = options.DirectoryInspectionAddress
	}
}

func setupTransactionLogger(logger *zap.Logger, postgresDSN, txLogAPIAddress string, certificateBundle *common_tls.CertificateBundle) (transactionlog.TransactionLogger, error) {
	if postgresDSN != "" && txLogAPIAddress != "" {
		return nil, fmt.Errorf("cannot configure both postgresDSN and txlogAPIAddress")
	}

	if postgresDSN != "" {
		logDB, err := setupDatabase()
		if err != nil {
			return nil, err
		}

		return transactionlog.NewPostgresTransactionLogger(logger, logDB, transactionlog.DirectionOut)
	}

	if txLogAPIAddress != "" {
		return transactionlog.NewAPITransactionLogger(&transactionlog.NewAPITransactionLoggerArgs{
			Logger:       logger,
			Direction:    transactionlog.DirectionOut,
			APIAddress:   txLogAPIAddress,
			InternalCert: certificateBundle,
		})
	}

	return transactionlog.NewDiscardTransactionLogger(), nil
}

func setupDatabase() (*sqlx.DB, error) {
	var logDB *sqlx.DB

	logDB, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		return nil, errors.Wrap(err, "could not open connection to postgres")
	}

	var (
		connMaxLifetime = 5 * time.Minute
		maxOpenConns    = 100
		maxIdleConns    = 100
	)

	logDB.SetConnMaxLifetime(connMaxLifetime)
	logDB.SetMaxOpenConns(maxOpenConns)
	logDB.SetMaxIdleConns(maxIdleConns)
	logDB.MapperFunc(strings.ToSnakeCase)

	return logDB, nil
}
