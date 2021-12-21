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

	"go.nlx.io/nlx/common/cmd"
	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	pgadapter "go.nlx.io/nlx/txlog-api/adapters/storage/postgres"
	"go.nlx.io/nlx/txlog-api/pkg/api"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	ListenAddress      string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8443" description:"Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenAddressPlain string `long:"listen-address-plain" env:"LISTEN_ADDRESS_PLAIN" default:"127.0.0.1:8080" description:"Address for the directory to listen on using plain HTTP. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	logoptions.LogOptions
	cmd.TLSOptions
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

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.KeyFile); errValidate != nil {
		logger.Warn("invalid internal PKI key permissions", zap.Error(errValidate), zap.String("file-path", options.KeyFile))
	}

	certificate, err := common_tls.NewBundleFromFiles(options.CertFile, options.KeyFile, options.RootCertFile)
	if err != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
	}

	db, err := pgadapter.NewPostgreSQLConnection(options.PostgresDSN)
	if err != nil {
		logger.Fatal("can not create db connection:", zap.Error(err))
	}

	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestTxlogDBVersion)

	storage, err := pgadapter.New(logger, db)
	if err != nil {
		logger.Fatal("failed to setup postgresql txlog database:", zap.Error(err))
	}

	server, err := api.NewAPI(logger, certificate, storage)
	if err != nil {
		logger.Fatal("could not start server", zap.Error(err))
	}

	go func() {
		err = server.ListenAndServe(options.ListenAddress, options.ListenAddressPlain)
		if err != nil {
			logger.Fatal("could not listen and serve", zap.Error(err))
		}
	}()

	p.Wait()

	logger.Info("starting graceful shutdown")

	gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err = server.Shutdown(gracefulCtx)
	if err != nil {
		logger.Error("could not shutdown server", zap.Error(err))
	}

	err = storage.Shutdown()
	if err != nil {
		logger.Error("could not shutdown storage", zap.Error(err))
	}
}
