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

	common_db "go.nlx.io/nlx/common/db"
	nlxhttp "go.nlx.io/nlx/common/http"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	pgadapter "go.nlx.io/nlx/directory-api/adapters/storage/postgres"
	"go.nlx.io/nlx/directory-api/http"
	"go.nlx.io/nlx/directory-api/pkg/directory"
	"go.nlx.io/nlx/directory-db/dbversion"
)

var options struct {
	ListenAddress      string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8443" description:"Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenAddressPlain string `long:"listen-address-plain" env:"LISTEN_ADDRESS_PLAIN" default:"127.0.0.1:8080" description:"Address for the directory to listen on using plain HTTP. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	NLXRootCert        string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem"`
	DirectoryCertFile  string `long:"tls-directory-cert" env:"TLS_DIRECTORY_CERT" description:"Absolute or relative path to the Directory cert .pem"`
	DirectoryKeyFile   string `long:"tls-directory-key" env:"TLS_DIRECTORY_KEY" description:"Absolute or relative path to the Directory key .pem"`

	TermsOfServiceURL string `long:"terms-of-service-url" env:"TERMS_OF_SERVICE_URL" description:"optional link to the terms of service"`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	logoptions.LogOptions
}

func parseArgs() error {
	args, err := flags.Parse(&options)
	if err != nil {
		if et, ok := err.(*flags.Error); ok {
			if et.Type == flags.ErrHelp {
				return err
			}
		}

		log.Fatalf("error parsing flags: %v", err)
	}

	if len(args) > 0 {
		log.Fatalf("unexpected arguments: %v", args)
	}

	return nil
}

func newZapLogger() *zap.Logger {
	config := options.LogOptions.ZapConfig()

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	return logger
}

func main() {
	err := parseArgs()
	if err != nil {
		return
	}

	p := process.NewProcess()

	logger := newZapLogger()

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.DirectoryKeyFile); errValidate != nil {
		logger.Warn("invalid directory key permissions", zap.Error(errValidate), zap.String("file-path", options.DirectoryKeyFile))
	}

	certificate, err := common_tls.NewBundleFromFiles(options.DirectoryCertFile, options.DirectoryKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading certificate", zap.Error(err))
	}

	db, err := pgadapter.NewPostgreSQLConnection(options.PostgresDSN)
	if err != nil {
		logger.Fatal("can not create db connection:", zap.Error(err))
	}

	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

	storage, err := pgadapter.New(logger, db)
	if err != nil {
		logger.Fatal("failed to setup postgresql directory database:", zap.Error(err))
	}

	httpClient := nlxhttp.NewHTTPClient(certificate)
	directoryService := directory.New(
		logger,
		options.TermsOfServiceURL,
		storage,
		httpClient,
		common_tls.GetOrganizationInfoFromRequest,
	)

	httpServer := http.NewServer(db, certificate, logger)

	server, err := NewServer(&NewServerArgs{
		Logger:                       logger,
		Address:                      options.ListenAddress,
		AddressPlain:                 options.ListenAddressPlain,
		Certificate:                  certificate,
		DirectoryService:             directoryService,
		DirectoryRegistrationService: directoryService,
		DirectoryInspectionService:   directoryService,
		HTTPServer:                   httpServer,
	})
	if err != nil {
		logger.Fatal("could not start server", zap.Error(err))
	}

	p.Wait()

	logger.Info("starting graceful shutdown")

	gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err = server.Shutdown(gracefulCtx)
	if err != nil {
		logger.Error("could not shutdown server", zap.Error(err))
	}

	err = db.Close()
	if err != nil {
		logger.Error("could not shutdown db", zap.Error(err))
	}
}
