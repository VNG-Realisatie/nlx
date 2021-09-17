// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/huandu/xstrings"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/cmd"
	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/logoptions"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/outway"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8080" description:"Address for the outway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenHTTPS   bool   `long:"listen-https" env:"LISTEN_HTTPS" description:"Enable HTTPS on the ListenAddress" required:"false"`

	MonitoringAddress string `long:"monitoring-address" env:"MONITORING_ADDRESS" default:"127.0.0.1:8081" description:"Address for the outway monitoring endpoints to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryInspectionAddress string `long:"directory-inspection-address" env:"DIRECTORY_INSPECTION_ADDRESS" description:"Address for the directory where this outway can fetch the service list" required:"true"`

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-termChan
		cancel()
	}()

	// Setup new zap logger
	config := options.LogOptions.ZapConfig()
	logger, err := config.Build()

	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.OrgKeyFile); errValidate != nil {
		logger.Warn("invalid organization key permissions", zap.Error(errValidate), zap.String("file-path", options.OrgCertFile))
	}

	orgCert, err := common_tls.NewBundleFromFiles(options.OrgCertFile, options.OrgKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
	}

	var logDB *sqlx.DB
	if !options.DisableLogdb {
		logDB, err = sqlx.Open("postgres", options.PostgresDSN)
		if err != nil {
			logger.Fatal("could not open connection to postgres", zap.Error(err))
		}

		logDB.SetConnMaxLifetime(5 * time.Minute)
		logDB.SetMaxOpenConns(100)
		logDB.SetMaxIdleConns(100)
		logDB.MapperFunc(xstrings.ToSnakeCase)

		common_db.WaitForLatestDBVersion(logger, logDB.DB, dbversion.LatestTxlogDBVersion)
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

	creds := credentials.NewTLS(cert.TLSConfig())

	conn, err := grpc.DialContext(context.TODO(), options.ManagementAPIAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Fatal("failed to connect to Management API", zap.Error(err))
	}

	client := api.NewManagementClient(conn)

	ow, err := outway.NewOutway(
		ctx,
		logger,
		logDB,
		client,
		options.MonitoringAddress,
		orgCert,
		options.DirectoryInspectionAddress,
		options.AuthorizationServiceAddress,
		options.AuthorizationCA,
		options.UseAsHTTPProxy)

	if err != nil {
		logger.Fatal("failed to start outway", zap.Error(err))
	}

	go func() {
		err = ow.RunServer(options.ListenAddress, serverCertificate)
		if err != nil {
			logger.Fatal("error running outway", zap.Error(err))
		}
	}()

	<-ctx.Done()

	gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ow.Shutdown(gracefulCtx)
	conn.Close()
	logDB.Close()
}
