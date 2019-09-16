// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"

	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/directory-db/dbversion"
	monitor "go.nlx.io/nlx/directory-monitor"
	"go.uber.org/zap"
)

var options struct {
	NLXRootCert       string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem"`
	MonitorCertFile   string `long:"tls-monitor-cert" env:"TLS_MONITOR_CERT" description:"Absolute or relative path to the Monitor cert .pem"`
	MonitorKeyFile    string `long:"tls-monitor-key" env:"TLS_MONITOR_KEY" description:"Absolute or relative path to the Monitor key .pem"`
	TTLOfflineService int    `long:"ttl-offline-service" env:"TTL_OFFLINE_SERVICE" description:"Days a service can be offline before being removed from the directory" required:"true"`
	PostgresDSN       string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	logoptions.LogOptions
}

func main() {
	parseOptions()

	logger := initLogger()
	proc := process.NewProcess(logger)

	db, err := monitor.InitDatabase(options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}

	proc.CloseGracefully(db.Close)
	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

	caCertPool, err := orgtls.LoadRootCert(options.NLXRootCert)
	if err != nil {
		logger.Fatal("failed to load root cert", zap.Error(err))
	}
	certKeyPair, err := tls.LoadX509KeyPair(options.MonitorCertFile, options.MonitorKeyFile)
	if err != nil {
		logger.Fatal("failed to load x509 keypair for monitor", zap.Error(err))
	}

	logger.Debug("starting health checker", zap.Int("ttlOfflineService", options.TTLOfflineService))
	err = monitor.RunHealthChecker(proc, logger, db, options.PostgresDSN, caCertPool, &certKeyPair, options.TTLOfflineService)
	if err != nil && err != context.DeadlineExceeded {
		logger.Fatal("failed to run monitor healthchecker", zap.Error(err))
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
