// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"errors"
	"fmt"
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

	logger.Info(fmt.Sprintf("created the directory database: %v", directoryDatabase))

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

	const (
		MaxIdleConnections = 2
		FiveMinutes        = 5 * time.Minute
	)

	db.SetConnMaxLifetime(FiveMinutes)
	db.SetMaxIdleConns(MaxIdleConnections)
	db.MapperFunc(xstrings.ToSnakeCase)

	mainProcess.CloseGracefully(db.Close)

	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

	httpServer := http.NewServer(db, certificate, logger)

	runServer(mainProcess, logger, options.ListenAddress, options.ListenAddressPlain, certificate, directoryService, httpServer)
}

func getOrganisationNameFromRequest(ctx context.Context) (string, error) {
	peerContext, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("failed to obtain peer from context")
	}

	tlsInfo := peerContext.AuthInfo.(credentials.TLSInfo)

	return tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0], nil
}
