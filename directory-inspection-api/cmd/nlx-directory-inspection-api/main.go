// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/directory-db/dbversion"
	"go.nlx.io/nlx/directory-inspection-api/http"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/directory-inspection-api/pkg/inspectionservice"
)

var options struct {
	ListenAddress      string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8443" description:"Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenAddressPlain string `long:"listen-address-plain" env:"LISTEN_ADDRESS_PLAIN" default:"127.0.0.1:8080" description:"Address for the directory to listen on using plain HTTP. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	NLXRootCert       string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem"`
	DirectoryCertFile string `long:"tls-directory-cert" env:"TLS_DIRECTORY_CERT" description:"Absolute or relative path to the Directory cert .pem"`
	DirectoryKeyFile  string `long:"tls-directory-key" env:"TLS_DIRECTORY_KEY" description:"Absolute or relative path to the Directory key .pem"`

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

func main() {
	err := parseArgs()
	if err != nil {
		return
	}

	p := process.NewProcess()

	config := options.LogOptions.ZapConfig()

	logger, err := config.Build()
	if err != nil {
		// nolint:gocritic // exitAfterDefer: we know that defers aren't called after Fatalf
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	directoryDatabase, err := database.NewPostgreSQLDirectoryDatabase(options.PostgresDSN, logger)
	if err != nil {
		logger.Fatal("failed to setup postgresql directory database:", zap.Error(err))
	}

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.DirectoryKeyFile); errValidate != nil {
		logger.Warn("invalid directory key permissions", zap.Error(errValidate), zap.String("file-path", options.DirectoryKeyFile))
	}

	certificate, err := common_tls.NewBundleFromFiles(options.DirectoryCertFile, options.DirectoryKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading certificate", zap.Error(err))
	}

	directoryService := inspectionservice.New(logger, directoryDatabase, getOrganisationNameFromRequest)

	// NOTE: remove creation of DB later on, once the statsservice also uses the directoryDatabase
	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}

	setDBOptions(db)

	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

	httpServer := http.NewServer(db, certificate, logger)

	server, err := NewServer(logger, options.ListenAddress, options.ListenAddressPlain, certificate, directoryService, httpServer)
	if err != nil {
		logger.Fatal("could not start server", zap.Error(err))
	}

	p.Wait()

	logger.Info("starting graceful shutdown")

	gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	shutdown(gracefulCtx, logger, server, directoryDatabase, db)
}

func shutdown(ctx context.Context, logger *zap.Logger, server *Server, directoryDB database.DirectoryDatabase, db *sqlx.DB) {
	err := server.Shutdown(ctx)
	if err != nil {
		logger.Error("could not shutdown server", zap.Error(err))
	}

	err = directoryDB.Shutdown()
	if err != nil {
		logger.Error("could not shutdown directory database", zap.Error(err))
	}

	err = db.Close()
	if err != nil {
		logger.Error("could not shutdown db", zap.Error(err))
	}
}

func getOrganisationNameFromRequest(ctx context.Context) (string, error) {
	peerContext, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("failed to obtain peer from context")
	}

	tlsInfo := peerContext.AuthInfo.(credentials.TLSInfo)

	return tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0], nil
}

func setDBOptions(db *sqlx.DB) {
	const (
		MaxIdleConnections = 2
		FiveMinutes        = 5 * time.Minute
	)

	db.SetConnMaxLifetime(FiveMinutes)
	db.SetMaxIdleConns(MaxIdleConnections)
	db.MapperFunc(xstrings.ToSnakeCase)
}
