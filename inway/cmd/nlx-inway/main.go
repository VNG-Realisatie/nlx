// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"github.com/huandu/xstrings"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/inway"
	"go.nlx.io/nlx/inway/config"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8443" description:"Address for the inway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryRegistrationAddress string `long:"directory-registration-address" env:"DIRECTORY_REGISTRATION_ADDRESS" description:"Address for the directory where this inway can register it's services" required:"true"`

	DisableLogdb bool `long:"disable-logdb" env:"DISABLE_LOGDB" description:"Disable logdb connections"`

	SelfAddress string `long:"self-address" env:"SELF_ADDRESS" description:"The address that outways can use to reach me" required:"true"`

	ServiceConfig string `long:"service-config" env:"SERVICE_CONFIG" default:"service-config.toml" description:"Location of the service config toml file"`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	logoptions.LogOptions
	orgtls.TLSOptions
}

func main() { //nolint
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
	zapConfig := options.LogOptions.ZapConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))
	xprocess := process.NewProcess(logger)

	logger.Info("starting inway", zap.String("directory-registration-address", options.DirectoryRegistrationAddress))
	serviceConfig := config.LoadServiceConfig(logger, options.ServiceConfig)
	var logDB *sqlx.DB
	if !options.DisableLogdb {
		logDB, err = sqlx.Open("postgres", options.PostgresDSN)
		if err != nil {
			logger.Fatal("could not open connection to postgres", zap.Error(err))
		}
		logDB.SetConnMaxLifetime(5 * time.Minute)
		logDB.SetMaxIdleConns(2)
		logDB.MapperFunc(xstrings.ToSnakeCase)

		dbversion.WaitUntilLatestTxlogDBVersion(logger, logDB.DB)
		xprocess.CloseGracefully(logDB.Close)
	}

	iw, err := inway.NewInway(logger, logDB, options.SelfAddress, options.TLSOptions, options.DirectoryRegistrationAddress, serviceConfig)
	if err != nil {
		logger.Fatal("cannot setup inway", zap.Error(err))
	}

	// TODO: Issue #403
	for serviceName, serviceDetails := range serviceConfig.Services { //nolint
		logger.Info("loaded service from service-config.toml", zap.String("service-name", serviceName))
		logger.Debug("service configuration details", zap.String("service-name", serviceName), zap.String("endpoint-url", serviceDetails.EndpointURL),
			zap.String("root-ca-path", serviceDetails.CACertPath), zap.String("authorization-model", serviceDetails.AuthorizationModel),
			zap.String("irma-api-url", serviceDetails.IrmaAPIURL), zap.String("insight-api-url", serviceDetails.InsightAPIURL),
			zap.String("api-spec-url", serviceDetails.APISpecificationDocumentURL), zap.Bool("internal", serviceDetails.Internal),
			zap.String("public-support-contact", serviceDetails.PublicSupportContact), zap.String("tech-support-contact", serviceDetails.TechSupportContact))
		var rootCrt *x509.CertPool
		if len(serviceDetails.CACertPath) > 0 {
			rootCrt, err = orgtls.LoadRootCert(serviceDetails.CACertPath)
			if err != nil {
				logger.Fatal("Unable to load ca certificate for inway", zap.Error(err))
			}
		}
		endpoint, errr := iw.NewHTTPServiceEndpoint(logger, serviceName, serviceDetails.EndpointURL, &tls.Config{RootCAs: rootCrt})
		if errr != nil {
			logger.Fatal("failed to create service", zap.Error(err))
		}
		switch serviceDetails.AuthorizationModel {
		case "none", "":
			endpoint.SetAuthorizationPublic()
		case "whitelist":
			endpoint.SetAuthorizationWhitelist(serviceDetails.AuthorizationWhitelist)
		default:
			logger.Fatal(fmt.Sprintf(`invalid authorization model "%s" for service "%s"`, serviceDetails.AuthorizationModel, serviceName))
		}
		err = iw.AddServiceEndpoint(xprocess, endpoint, serviceDetails)
		if err != nil {
			logger.Fatal(fmt.Sprintf(`adding endpoint "%s"`, err))
		}
	}

	// Listen on the address provided in the options
	err = iw.ListenAndServeTLS(xprocess, options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
