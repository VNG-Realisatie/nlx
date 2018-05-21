package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	flags "github.com/svent/go-flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/VNG-Realisatie/nlx/common/logoptions"
	"github.com/VNG-Realisatie/nlx/common/orgtls"
	"github.com/VNG-Realisatie/nlx/common/process"
	"github.com/VNG-Realisatie/nlx/db/dbversion"
	"github.com/VNG-Realisatie/nlx/directory/directoryservice"
)

var options struct {
	ListenAddress      string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:1984" description:"Adress for the inway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenAddressPlain string `long:"listen-address-plain" env:"LISTEN_ADDRESS_PLAIN" default:"0.0.0.0:1985" description:"Adress for the inway to listen on using plain HTTP. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	NLXRootCert       string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem"`
	DirectoryCertFile string `long:"tls-directory-cert" env:"TLS_DIRECTORY_CERT" description:"Absolute or relative path to the Directory cert .pem"`
	DirectoryKeyFile  string `long:"tls-directory-key" env:"TLS_DIRECTORY_KEY" description:"Absolute or relative path to the Directory key .pem"`

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

	dbversion.WaitUntilLatestVersion(logger, db.DB)

	caCertPool, err := orgtls.LoadRootCert(options.NLXRootCert)
	if err != nil {
		logger.Fatal("failed to load root cert", zap.Error(err))
	}
	certKeyPair, err := tls.LoadX509KeyPair(options.DirectoryCertFile, options.DirectoryKeyFile)
	if err != nil {
		logger.Fatal("failed to load x509 keypair for directory", zap.Error(err))
	}

	directoryService, err := directoryservice.New(logger, db)
	if err != nil {
		logger.Fatal("failed to create new directory service", zap.Error(err))
	}

	runServer(logger, options.ListenAddress, options.ListenAddressPlain, caCertPool, certKeyPair, directoryService)
}
