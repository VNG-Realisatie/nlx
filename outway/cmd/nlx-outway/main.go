// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/huandu/xstrings"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/cmd"
	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/common/version"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/outway"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	Name string `long:"name" env:"OUTWAY_NAME" description:"ServiceName of the outway. Every outway should have a unique name within the organization." required:"true"`

	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8080" description:"Address for the outway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenHTTPS   bool   `long:"listen-https" env:"LISTEN_HTTPS" description:"Enable HTTPS on the ListenAddress" required:"false"`

	MonitoringAddress string `long:"monitoring-address" env:"MONITORING_ADDRESS" default:"127.0.0.1:8081" description:"Address for the outway monitoring endpoints to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryInspectionAddress string `long:"directory-inspection-address" env:"DIRECTORY_INSPECTION_ADDRESS" description:"Address for the directory where this outway can fetch the service list"`
	DirectoryAddress           string `long:"directory-address" env:"DIRECTORY_ADDRESS" description:"Address for the directory where this outway can fetch the service list"`

	UseAsHTTPProxy       bool   `long:"use-as-http-proxy" env:"USE_AS_HTTP_PROXY" description:"An experimental flag which when true makes the outway function as an HTTP proxy"`
	ManagementAPIAddress string `long:"management-api-address" env:"MANAGEMENT_API_ADDRESS" description:"The address of the NLX Management API" required:"true"`

	DisableLogdb bool   `long:"disable-logdb" env:"DISABLE_LOGDB" description:"Disable logdb connections"`
	PostgresDSN  string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	AuthorizationServiceAddress string `long:"authorization-service-address" env:"AUTHORIZATION_SERVICE_ADDRESS" description:"Address of the authorization service. If set calls will go through the authorization service before being send to the inway"`
	AuthorizationCA             string `long:"authorization-root-ca" env:"AUTHORIZATION_ROOT_CA" description:"absolute path to root CA used to verify auth service certifcate"`

	ServerCertFile string `long:"tls-server-cert" env:"TLS_SERVER_CERT" description:"Path to a cert .pem, used for the HTTPS server" required:"false"`
	ServerKeyFile  string `long:"tls-server-key" env:"TLS_SERVER_KEY" description:"Path the a key .pem, used for the HTTPS server" required:"false"`

	logoptions.LogOptions
	cmd.TLSOrgOptions
	cmd.TLSOptions
}

// nolint:funlen,gocyclo // this is the main function
func main() {
	parseOptions()

	p := process.NewProcess()

	// Setup new zap logger
	config := options.LogOptions.ZapConfig()
	logger, err := config.Build()

	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))
	txlogger, logDB := setupTransactionLogger(logger, options.DisableLogdb)

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.OrgKeyFile); errValidate != nil {
		logger.Warn("invalid organization key permissions", zap.Error(errValidate), zap.String("file-path", options.OrgCertFile))
	}

	orgCert, err := common_tls.NewBundleFromFiles(options.OrgCertFile, options.OrgKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
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

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.KeyFile); errValidate != nil {
		logger.Warn("invalid internal PKI key permissions", zap.Error(errValidate), zap.String("file-path", options.KeyFile))
	}

	cert, err := common_tls.NewBundleFromFiles(options.CertFile, options.KeyFile, options.RootCertFile)
	if err != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
	}

	publicKeyPEM, err := orgCert.PublicKeyPEM()
	if err != nil {
		logger.Fatal("unable to get public key pem from certificate TLS files", zap.Error(err))
	}

	creds := credentials.NewTLS(cert.TLSConfig())

	conn, err := grpc.DialContext(context.TODO(), options.ManagementAPIAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Fatal("failed to connect to Management API", zap.Error(err))
	}

	client := api.NewManagementClient(conn)

	_, err = client.RegisterOutway(context.TODO(), &api.RegisterOutwayRequest{
		Name:         options.Name,
		PublicKeyPEM: publicKeyPEM,
		Version:      version.BuildVersion,
	})
	if err != nil {
		logger.Fatal("failed to register outway in Management API", zap.Error(err))
	}

	directoryDialCredentials := credentials.NewTLS(orgCert.TLSConfig())
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}

	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(nlxversion.NewGRPCContext(context.Background(), "outway"), 1*time.Minute)
	directoryConn, err := grpc.DialContext(directoryConnCtx, options.DirectoryAddress, directoryDialOptions...)

	defer directoryConnCtxCancel()

	if err != nil {
		logger.Fatal("failed to setup connection to directory registration api", zap.Error(err))
	}

	directoryClient := directoryapi.NewDirectoryClient(directoryConn)

	ow, err := outway.NewOutway(&outway.NewOutwayArgs{
		Name:              options.Name,
		Ctx:               context.Background(),
		Logger:            logger,
		Txlogger:          txlogger,
		ManagementClient:  client,
		MonitoringAddress: options.MonitoringAddress,
		OrgCert:           orgCert,
		DirectoryClient:   directoryClient,
		AuthServiceURL:    options.AuthorizationServiceAddress,
		AuthCAPath:        options.AuthorizationCA,
		UseAsHTTPProxy:    options.UseAsHTTPProxy,
	})

	if err != nil {
		logger.Fatal("failed to start outway", zap.Error(err))
	}

	go func() {
		err = ow.Run()
		if err != nil {
			logger.Fatal("error running outway", zap.Error(err))
		}

		err = ow.RunServer(options.ListenAddress, serverCertificate)
		if err != nil {
			logger.Fatal("error running outway", zap.Error(err))
		}
	}()

	p.Wait()

	logger.Info("starting graceful shutdown")

	gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ow.Shutdown(gracefulCtx)
	conn.Close()

	if logDB != nil {
		logDB.Close()
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

func setupTransactionLogger(logger *zap.Logger, disabled bool) (transactionlog.TransactionLogger, *sqlx.DB) {
	if disabled {
		logger.Info("logging to transaction-log disabled")
		return transactionlog.NewDiscardTransactionLogger(), nil
	}

	logDB, err := setupDatabase(logger)
	if err != nil {
		logger.Fatal("failed to setup database", zap.Error(err))
	}

	postgresTxLogger, err := transactionlog.NewPostgresTransactionLogger(logger, logDB, transactionlog.DirectionOut)
	if err != nil {
		logger.Fatal("failed to setup transactionlog", zap.Error(err))
	}

	logger.Info("transaction logger created")

	return postgresTxLogger, logDB
}

func setupDatabase(logger *zap.Logger) (*sqlx.DB, error) {
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
	logDB.MapperFunc(xstrings.ToSnakeCase)

	common_db.WaitForLatestDBVersion(logger, logDB.DB, dbversion.LatestTxlogDBVersion)

	return logDB, nil
}
