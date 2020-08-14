// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/directory-db/dbversion"
	"go.nlx.io/nlx/directory-inspection-api/http"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
	"go.nlx.io/nlx/directory-inspection-api/statsservice"
)

var options struct {
	ListenAddress      string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8443" description:"Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenAddressPlain string `long:"listen-address-plain" env:"LISTEN_ADDRESS_PLAIN" default:"0.0.0.0:8080" description:"Address for the directory to listen on using plain HTTP. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

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

	// Setup new zap logger
	config := options.LogOptions.ZapConfig()
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	mainProcess := process.NewProcess(logger)

	directoryDatabase, err := database.NewPostgreSQLDirectoryDatabase(options.PostgresDSN, mainProcess, logger)
	if err != nil {
		logger.Fatal("failed to setup postgresql directory database:", zap.Error(err))
	}
	log.Printf("created the directory database: %v", directoryDatabase)

	caCertPool, err := orgtls.LoadRootCert(options.NLXRootCert)
	if err != nil {
		logger.Fatal("failed to load root cert", zap.Error(err))
	}
	certKeyPair, err := tls.LoadX509KeyPair(options.DirectoryCertFile, options.DirectoryKeyFile)
	if err != nil {
		logger.Fatal("failed to load x509 keypair for directory inspection api", zap.Error(err))
	}

	directoryService := inspectionservice.New(logger, directoryDatabase)
	if err != nil {
		logger.Fatal("failed to create new directory inspection service", zap.Error(err))
	}

	// NOTE: remove creation of DB later on, once the statsservice also uses the directoryDatabase
	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(2)
	db.MapperFunc(xstrings.ToSnakeCase)

	mainProcess.CloseGracefully(db.Close)

	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

	statsService, err := statsservice.New(logger, db)
	if err != nil {
		logger.Fatal("failed to create new stats service", zap.Error(err))
	}

	httpServer := http.NewServer(db, caCertPool, &certKeyPair, logger)

	runServer(mainProcess, logger, options.ListenAddress, options.ListenAddressPlain, caCertPool, &certKeyPair, directoryService, statsService, httpServer)
}
