// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"strconv"
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
	"go.nlx.io/nlx/inway"
	"go.nlx.io/nlx/inway/config"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	ListenAddress           string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8443" description:"Address for the inway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenManagementAddress string `long:"listen-management-address" env:"LISTEN_MANAGEMENT_ADDRESS" description:"Address for the inway to listen on for management requests. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	MonitoringAddress string `long:"monitoring-address" env:"MONITORING_ADDRESS" default:"0.0.0.0:8081" description:"Address for the inway monitoring endpoints to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryRegistrationAddress string `long:"directory-registration-address" env:"DIRECTORY_REGISTRATION_ADDRESS" description:"Address for the directory where this inway can register it's services" required:"true"`

	DisableLogdb bool `long:"disable-logdb" env:"DISABLE_LOGDB" description:"Disable logdb connections"`

	ManagementAPIAddress string `long:"management-api-address" env:"MANAGEMENT_API_ADDRESS" description:"The address of the NLX Management API"`

	SelfAddress string `long:"self-address" env:"SELF_ADDRESS" description:"The address that outways can use to reach me" required:"true"`

	ServiceConfig string `long:"service-config" env:"SERVICE_CONFIG" default:"service-config.toml" description:"Location of the service config toml file"`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	InwayName string `long:"name" env:"INWAY_NAME" description:"Name of the inway"`

	logoptions.LogOptions
	cmd.TLSOrgOptions
	cmd.TLSOptions
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
	logger := setupLogger()
	mainProcess := process.NewProcess(logger)

	logDB := setupDatabase(logger, mainProcess)

	orgCert, err := common_tls.NewBundleFromFiles(options.OrgCertFile, options.OrgKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
	}

	iw, err := inway.NewInway(logger, logDB, mainProcess, options.InwayName, options.SelfAddress, options.MonitoringAddress, orgCert, options.DirectoryRegistrationAddress)
	if err != nil {
		logger.Fatal("cannot setup inway", zap.Error(err))
	}

	if len(options.ManagementAPIAddress) > 0 {
		logger.Info("management-api set")

		cert, mErr := common_tls.NewBundleFromFiles(options.CertFile, options.KeyFile, options.RootCertFile)
		if mErr != nil {
			logger.Fatal("loading TLS files", zap.Error(err))
		}

		err = iw.SetupManagementAPI(options.ManagementAPIAddress, cert)
		if err != nil {
			logger.Fatal("cannot configure management-api", zap.Error(err))
		}
		err = iw.StartConfigurationPolling()
		if err != nil {
			logger.Fatal("cannot retrieving inway configuration from the management-api", zap.Error(err))
		}
	} else {
		serviceConfig, err2 := config.LoadServiceConfig(options.ServiceConfig)
		if err2 != nil {
			if serviceConfig == nil {
				logger.Fatal("failed to load service config", zap.Error(err2))
			} else {
				logger.Warn("warning while loading service config", zap.Error(err))
			}
		}
		loadServices(logger, serviceConfig, iw)
	}

	managementAddress := options.ListenManagementAddress
	if managementAddress == "" {
		managementAddress, err = defaultManagementAddress(options.ListenAddress)
		if err != nil {
			logger.Fatal("unable to crete default management address", zap.Error(err))
		}
	}

	// Listen on the address provided in the options
	err = iw.RunServer(options.ListenAddress, managementAddress)
	if err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}
}

func defaultManagementAddress(address string) (string, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", err
	}

	intPort, err := strconv.Atoi(port)
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%v:%v", host, (intPort + 1)), nil
}

func setupLogger() *zap.Logger {
	zapConfig := options.LogOptions.ZapConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	logger.Info("starting inway", zap.String("directory-registration-address", options.DirectoryRegistrationAddress))
	return logger
}

func setupDatabase(logger *zap.Logger, mainProcess *process.Process) *sqlx.DB {
	var logDB *sqlx.DB
	if !options.DisableLogdb {
		var err error
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

	return logDB
}

func loadServices(logger *zap.Logger, serviceConfig *config.ServiceConfig, iw *inway.Inway) {
	if len(serviceConfig.Services) == 0 {
		logger.Warn("inway has 0 configured services")
	}

	// TODO: Issue #403
	serviceEndpoints := make([]inway.ServiceEndpoint, len(serviceConfig.Services))
	i := 0
	for serviceName := range serviceConfig.Services {
		serviceDetails := serviceConfig.Services[serviceName]
		logger.Info("loaded service from service-config.toml", zap.String("service-name", serviceName))
		logger.Debug("service configuration details", zap.String("service-name", serviceName), zap.String("endpoint-url", serviceDetails.EndpointURL),
			zap.String("root-ca-path", serviceDetails.CACertPath), zap.String("authorization-model", string(serviceDetails.AuthorizationModel)),
			zap.String("irma-api-url", serviceDetails.IrmaAPIURL), zap.String("insight-api-url", serviceDetails.InsightAPIURL),
			zap.String("api-spec-url", serviceDetails.APISpecificationDocumentURL), zap.Bool("internal", serviceDetails.Internal),
			zap.String("public-support-contact", serviceDetails.PublicSupportContact), zap.String("tech-support-contact", serviceDetails.TechSupportContact))

		var rootCAs *x509.CertPool
		var err error
		if len(serviceDetails.CACertPath) > 0 {
			rootCAs, _, err = common_tls.NewCertPoolFromFile(serviceDetails.CACertPath)
			if err != nil {
				logger.Fatal("Unable to load ca certificate for inway", zap.Error(err))
			}
		}

		endpoint, errr := iw.NewHTTPServiceEndpoint(serviceName, &serviceDetails, &tls.Config{
			RootCAs:    rootCAs,
			MinVersion: tls.VersionTLS12,
		})
		if errr != nil {
			logger.Fatal("failed to create service", zap.Error(err))
		}

		serviceEndpoints[i] = endpoint
		i++
	}

	err := iw.SetServiceEndpoints(serviceEndpoints)
	if err != nil {
		logger.Fatal(fmt.Sprintf(`error setting service endpoints "%s"`, err))
	}
}
