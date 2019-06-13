// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"log"

	"github.com/huandu/xstrings"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/outway"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	ListenAddress    string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8080" description:"Address for the outway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenAddressTLS string `long:"listen-address-tls" env:"LISTEN_ADDRESS_TLS" default:"0.0.0.0:8443" description:"Address for the outway to listen on for TLS connections. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryInspectionAddress string `long:"directory-inspection-address" env:"DIRECTORY_INSPECTION_ADDRESS" description:"Address for the directory where this outway can fetch the service list" required:"true"`

	DisableLogdb bool   `long:"disable-logdb" env:"DISABLE_LOGDB" description:"Disable logdb connections"`
	PostgresDSN  string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	AuthorizationServiceAddress string `long:"authorization-service-address" env:"AUTHORIZATION_SERVICE_ADDRESS" description:"Address of the authorization service. If set calls will go through the authorization service before being send to the inway"`
	AuthorizationCA             string `long:"authorization-root-ca" env:"AUTHORIZATION_ROOT_CA" description:"absolute path to root CA used to verify auth service certifcate"`

	logoptions.LogOptions
	orgtls.TLSOptions
}

func main() {
	// Parse options
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

	// Setup new zap logger
	config := options.LogOptions.ZapConfig()
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	defer func() {
		syncErr := logger.Sync()
		if syncErr != nil {
			log.Fatalf("failed to sync zap logger: %v", syncErr)
		}
	}()

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	process := process.NewProcess(logger)

	var logDB *sqlx.DB
	if !options.DisableLogdb {
		logDB, err = sqlx.Open("postgres", options.PostgresDSN)
		if err != nil {
			logger.Fatal("could not open connection to postgres", zap.Error(err))
		}
		logDB.MapperFunc(xstrings.ToSnakeCase)
		process.CloseGracefully(logDB.Close)
		dbversion.WaitUntilLatestTxlogDBVersion(logger, logDB.DB)
	}

	// Create new outway and provide it with a hardcoded service.
	ow, err := outway.NewOutway(process, logger, logDB, options.TLSOptions, options.DirectoryInspectionAddress, options.AuthorizationServiceAddress, options.AuthorizationCA)
	if err != nil {
		logger.Fatal("failed to setup outway", zap.Error(err))
	}

	go func() {
		// Listen on the address provided in the options
		err = ow.ListenAndServeTLS(process, options.ListenAddressTLS, options.TLSOptions.OrgCertFile, options.TLSOptions.OrgKeyFile)
		if err != nil {
			logger.Fatal("failed to listen and serve", zap.Error(err))
		}
	}()

	// Listen on the address provided in the options
	err = ow.ListenAndServe(process, options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
