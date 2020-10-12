// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/huandu/xstrings"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/cmd"
	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/outway"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8080" description:"Address for the outway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenHTTPS   bool   `long:"listen-https" env:"LISTEN_HTTPS" description:"Enable HTTPS on the ListenAddress" required:"false"`

	MonitoringAddress string `long:"monitoring-address" env:"MONITORING_ADDRESS" default:"0.0.0.0:8081" description:"Address for the outway monitoring endpoints to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryInspectionAddress string `long:"directory-inspection-address" env:"DIRECTORY_INSPECTION_ADDRESS" description:"Address for the directory where this outway can fetch the service list" required:"true"`

	UseAsHTTPProxy bool `long:"use-as-http-proxy" env:"USE_AS_HTTP_PROXY" description:"An experimental flag which when true makes the outway function as an HTTP proxy"`

	DisableLogdb bool   `long:"disable-logdb" env:"DISABLE_LOGDB" description:"Disable logdb connections"`
	PostgresDSN  string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	AuthorizationServiceAddress string `long:"authorization-service-address" env:"AUTHORIZATION_SERVICE_ADDRESS" description:"Address of the authorization service. If set calls will go through the authorization service before being send to the inway"`
	AuthorizationCA             string `long:"authorization-root-ca" env:"AUTHORIZATION_ROOT_CA" description:"absolute path to root CA used to verify auth service certifcate"`

	ServerCertFile string `long:"tls-server-cert" env:"TLS_SERVER_CERT" description:"Path to a cert .pem, used for the HTTPS server" required:"false"`
	ServerKeyFile  string `long:"tls-server-key" env:"TLS_SERVER_KEY" description:"Path the a key .pem, used for the HTTPS server" required:"false"`

	logoptions.LogOptions
	cmd.TLSOrgOptions
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
}

func main() {
	parseOptions()

	// Setup new zap logger
	config := options.LogOptions.ZapConfig()
	logger, err := config.Build()

	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	orgCert, err := common_tls.NewBundleFromFiles(options.OrgCertFile, options.OrgKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
	}

	mainProcess := process.NewProcess(logger)

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
		mainProcess.CloseGracefully(logDB.Close)
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

	// Create new outway and provide it with a hardcoded service.
	ow, err := outway.NewOutway(
		logger,
		logDB,
		mainProcess,
		options.MonitoringAddress,
		orgCert,
		options.DirectoryInspectionAddress,
		options.AuthorizationServiceAddress,
		options.AuthorizationCA,
		options.UseAsHTTPProxy)

	if err != nil {
		logger.Fatal("failed to setup outway", zap.Error(err))
	}

	err = ow.RunServer(options.ListenAddress, serverCertificate)
	if err != nil {
		logger.Fatal("error running outway", zap.Error(err))
	}
}
