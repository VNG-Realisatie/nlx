// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

	common_db "go.nlx.io/nlx/common/db"
	nlxhttp "go.nlx.io/nlx/common/http"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/directory-db/dbversion"
	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/pkg/registrationservice"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

var options struct {
	ListenAddress     string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8443" description:"Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	NLXRootCert       string `long:"tls-nlx-root-cert" env:"TLS_NLX_ROOT_CERT" description:"Absolute or relative path to the NLX CA root cert .pem"`
	DirectoryCertFile string `long:"tls-directory-cert" env:"TLS_DIRECTORY_CERT" description:"Absolute or relative path to the Directory cert .pem"`
	DirectoryKeyFile  string `long:"tls-directory-key" env:"TLS_DIRECTORY_KEY" description:"Absolute or relative path to the Directory key .pem"`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	logoptions.LogOptions
}

func parseArgs() error {
	// Parse options
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

const (
	dbConnectionMaxLifetime = 5 * time.Minute
	dbMaxIdleConnections    = 2
)

func setupDB(logger *zap.Logger) *sqlx.DB {
	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}

	db.SetConnMaxLifetime(dbConnectionMaxLifetime)
	db.SetMaxIdleConns(dbMaxIdleConnections)
	db.MapperFunc(xstrings.ToSnakeCase)

	return db
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

	logger := newZapLogger()

	mainProcess := process.NewProcess(logger)
	db := setupDB(logger)
	mainProcess.CloseGracefully(db.Close)

	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

	certificate, err := common_tls.NewBundleFromFiles(options.DirectoryCertFile, options.DirectoryKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading certificate", zap.Error(err))
	}

	directoryDatabase, err := database.NewPostgreSQLDirectoryDatabase(options.PostgresDSN, mainProcess, logger)
	if err != nil {
		logger.Fatal("failed to setup postgresql directory database:", zap.Error(err))
	}

	httpClient := nlxhttp.NewHTTPClient(certificate)
	registrationService := registrationservice.New(logger, directoryDatabase, httpClient, getOrganisationNameFromRequest)

	// setup zap connection for global grpc logging
	grpc_zap.ReplaceGrpcLogger(logger)
	// prepare grpc server options
	serverTLSConfig := certificate.TLSConfig(certificate.WithTLSClientAuth())
	serverTLSConfig.NextProtos = []string{"h2"}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewTLS(serverTLSConfig)),
	}

	// start grpc server and attach directory service
	grpcServer := grpc.NewServer(opts...)
	registrationapi.RegisterDirectoryRegistrationServer(grpcServer, registrationService)

	listen, err := net.Listen("tcp", options.ListenAddress)
	if err != nil {
		log.Fatal("failed to create listener", zap.Error(err))
	}

	mainProcess.CloseGracefully(func() error {
		grpcServer.GracefulStop()
		return nil
	})

	if err := grpcServer.Serve(listen); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal("error serving", zap.Error(err))
		}
	}
}

func getOrganisationNameFromRequest(ctx context.Context) (string, error) {
	orgPeer, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("failed to obtain peer from context")
	}

	tlsInfo := orgPeer.AuthInfo.(credentials.TLSInfo)
	if len(tlsInfo.State.VerifiedChains) == 0 {
		return "", errors.New("no valid TLS certificate chain found")
	}

	return tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0], nil
}
