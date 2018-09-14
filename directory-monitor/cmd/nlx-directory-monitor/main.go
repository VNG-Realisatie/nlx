package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-db/dbversion"
	"go.nlx.io/nlx/directory-monitor"
)

var options struct {
	NLXRootCert     string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem"`
	MonitorCertFile string `long:"tls-monitor-cert" env:"TLS_MONITOR_CERT" description:"Absolute or relative path to the Monitor cert .pem"`
	MonitorKeyFile  string `long:"tls-monitor-key" env:"TLS_MONITOR_KEY" description:"Absolute or relative path to the Monitor key .pem"`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	logoptions.LogOptions
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

	// Setup new zap loggerv
	config := options.LogOptions.ZapConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	defer func() { // TODO(GeertJohan): #205 make this a common/process exitFunc?
		syncErr := logger.Sync()
		if syncErr != nil {
			// notify the user that proper logging has failed
			fmt.Fprintf(os.Stderr, "failed to sync zap logger: %v\n", syncErr)
			// don't exit when we're in a panic
			if p := recover(); p != nil {
				panic(p)
			}
			os.Exit(1)
		}
	}()

	process.Setup(logger)

	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}
	db.MapperFunc(xstrings.ToSnakeCase)

	dbversion.WaitUntilLatestDirectoryDBVersion(logger, db.DB)

	caCertPool, err := orgtls.LoadRootCert(options.NLXRootCert)
	if err != nil {
		logger.Fatal("failed to load root cert", zap.Error(err))
	}
	certKeyPair, err := tls.LoadX509KeyPair(options.MonitorCertFile, options.MonitorKeyFile)
	if err != nil {
		logger.Fatal("failed to load x509 keypair for monitor", zap.Error(err))
	}

	err = monitor.RunHealthChecker(logger, db, caCertPool, certKeyPair)
	if err != nil {
		logger.Fatal("failed to run monitor healthchecker", zap.Error(err))
	}
}
