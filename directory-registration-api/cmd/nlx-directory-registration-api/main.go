// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"

	"go.nlx.io/nlx/common/tlsconfig"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/directory-db/dbversion"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/directory-registration-api/registrationservice"
)

var options struct {
	ListenAddress     string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:443" description:"Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
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
	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	process := process.NewProcess(logger)

	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}
	db.MapperFunc(xstrings.ToSnakeCase)
	process.CloseGracefully(db.Close)

	dbversion.WaitUntilLatestDirectoryDBVersion(logger, db.DB)

	caCertPool, err := orgtls.LoadRootCert(options.NLXRootCert)
	if err != nil {
		logger.Fatal("failed to load root cert", zap.Error(err))
	}
	certKeyPair, err := tls.LoadX509KeyPair(options.DirectoryCertFile, options.DirectoryKeyFile)
	if err != nil {
		logger.Fatal("failed to load x509 keypair for directory registration api", zap.Error(err))
	}

	registrationService, err := registrationservice.NewRegisterInwayHandler(db, logger, caCertPool, certKeyPair)
	if err != nil {
		logger.Fatal("failed to create new directory registration service", zap.Error(err))
	}

	// setup zap connection for global grpc logging
	grpc_zap.ReplaceGrpcLogger(logger)
	// prepare grpc server options
	serverTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{certKeyPair}, // using the grpc server's own cert to connect to it, perhaps find a way for the http/json gateway to bypass TLS locally
		ClientCAs:    caCertPool,
		NextProtos:   []string{"h2"},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	tlsconfig.ApplyDefaults(serverTLSConfig)
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
	process.CloseGracefully(func() error {
		grpcServer.GracefulStop()
		return nil
	})
	if err := grpcServer.Serve(listen); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal("error serving", zap.Error(err))
		}
	}
}
