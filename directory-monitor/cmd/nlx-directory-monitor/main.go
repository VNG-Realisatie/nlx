// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"log"
	"time"

	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	monitor "go.nlx.io/nlx/directory-monitor"
)

var options struct {
	NLXRootCert                string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem"`
	MonitorCertFile            string `long:"tls-monitor-cert" env:"TLS_MONITOR_CERT" description:"Absolute or relative path to the Monitor cert .pem"`
	MonitorKeyFile             string `long:"tls-monitor-key" env:"TLS_MONITOR_KEY" description:"Absolute or relative path to the Monitor key .pem"`
	TTLOfflineService          int    `long:"ttl-offline-service" env:"TTL_OFFLINE_SERVICE" description:"Time, in seconds, a service can be offline before being removed from the directory" required:"true"`
	DatabaseMaxOpenConnections int    `long:"db-max-open-connections" env:"DB_MAX_OPEN_CONNECTIONS" description:"Maximum number of open connections allowed to the database" default:"25" required:"true"`
	DatabaseMaxIdleConnections int    `long:"db-max-idle-connections" env:"DB_MAX_IDLE_CONNECTIONS" description:"Maximum number of idle connections allowed to the database" default:"5" required:"true"`
	DatabaseConnectionTimeout  int    `long:"db-connection-timeout" env:"DB_CONNECTION_TIMEOUT" description:"Database connection timeout in seconds" default:"300" required:"true"`
	PostgresDSN                string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	logoptions.LogOptions
}

func main() {
	parseOptions()

	p := process.NewProcess()

	logger := initLogger()

	logger.Debug("opening database connection", zap.Int("maxIdleConnections", options.DatabaseMaxIdleConnections), zap.Int("maxOpenConnections", options.DatabaseMaxOpenConnections), zap.Int("connectionTimeout in seconds", options.DatabaseConnectionTimeout))

	db, err := monitor.InitDatabase(
		&monitor.DBConnectionArgs{
			DSN:                options.PostgresDSN,
			MaxIdleConnections: options.DatabaseMaxIdleConnections,
			MaxOpenConnections: options.DatabaseMaxOpenConnections,
			ConnectionTimeout:  time.Duration(options.DatabaseConnectionTimeout) * time.Second},
	)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.MonitorKeyFile); errValidate != nil {
		logger.Warn("invalid private key permissions", zap.Error(errValidate), zap.String("file-path", options.MonitorKeyFile))
	}

	certificate, err := common_tls.NewBundleFromFiles(options.MonitorCertFile, options.MonitorKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading certificate", zap.Error(err))
	}

	logger.Debug("starting health checker", zap.Int("ttlOfflineService", options.TTLOfflineService))

	healthChecker := monitor.New(logger, certificate)

	go func() {
		err = healthChecker.Run(db, options.PostgresDSN, options.TTLOfflineService)
		if err != nil && err != context.DeadlineExceeded {
			logger.Fatal("failed to run monitor healthchecker", zap.Error(err))
		}
	}()

	p.Wait()

	logger.Info("starting graceful shutdown")

	err = healthChecker.Shutdown()
	if err != nil {
		logger.Error("could not shutdown health checker", zap.Error(err))
	}

	err = db.Close()
	if err != nil {
		logger.Error("could not shutdown db", zap.Error(err))
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
}

func initLogger() *zap.Logger {
	config := options.LogOptions.ZapConfig()

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	return logger
}
