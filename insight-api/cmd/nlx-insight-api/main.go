// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.nlx.io/nlx/insight-api/irma"

	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/version"
	insightapi "go.nlx.io/nlx/insight-api"
	"go.nlx.io/nlx/insight-api/config"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	logoptions.LogOptions

	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8080" description:"Adress for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	IRMAJWTRSASignPrivateKeyDER  string `long:"irma-jwt-rsa-sign-private-key-der" env:"IRMA_JWT_RSA_SIGN_PRIVATE_KEY_DER" required:"true" description:"PEM RSA private key to sign requests for irma api server"`
	IRMAJWTRSAVerifyPublicKeyDER string `long:"irma-jwt-rsa-verify-public-key-der" env:"IRMA_JWT_RSA_VERIFY_PUBLIC_KEY_DER" required:"true" description:"PEM RSA public key to verify results from irma api server"`

	InsightConfig string `long:"insight-config" env:"INSIGHT_CONFIG" default:"insight-config.toml" description:"Location of the insight config toml file"`

	CertFile string `long:"tls-cert" env:"TLS_CERT" description:"Absolute or relative path to the cert .pem"`
	KeyFile  string `long:"tls-key" env:"TLS_KEY" description:"Absolute or relative path to the key .pem"`
}

func main() {
	if args := parseArgs(); args == nil {
		return
	}

	// Setup new zap logger
	zapConfig := options.LogOptions.ZapConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	process := process.NewProcess(logger)

	insightConfig, err := config.LoadInsightConfig(logger, options.InsightConfig)
	if err != nil {
		logger.Fatal("error loading insight config", zap.Error(err))

	}

	db, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(2)
	db.MapperFunc(xstrings.ToSnakeCase)

	process.CloseGracefully(db.Close)

	dbversion.WaitUntilLatestTxlogDBVersion(logger, db.DB)

	insightLogFetcher, err := insightapi.NewInsightDatabase(db)
	if err != nil {
		logger.Fatal("error creating log fetcher", zap.Error(err))
	}

	irmaHandler := irma.NewJWTGenerator()

	insightAPI, err := insightapi.NewInsightAPI(logger, insightConfig, irmaHandler, insightLogFetcher, options.IRMAJWTRSASignPrivateKeyDER, options.IRMAJWTRSAVerifyPublicKeyDER)
	if err != nil {
		logger.Fatal("error creating insightAPI", zap.Error(err))
	}

	server := &http.Server{
		Addr:    options.ListenAddress,
		Handler: insightAPI,
	}

	process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		return server.Shutdown(localCtx)
	})

	if len(options.KeyFile) > 0 {
		err = server.ListenAndServeTLS(options.CertFile, options.KeyFile)
	} else {
		err = server.ListenAndServe()
	}

	if err != http.ErrServerClosed {
		logger.Fatal("error listen and serverinsightAPI", zap.Error(err))
	}

}

func parseArgs() []string {
	// Parse options
	args, err := flags.Parse(&options)
	if err != nil {
		if et, ok := err.(*flags.Error); ok {
			if et.Type == flags.ErrHelp {
				return nil
			}
		}
		log.Fatalf("error parsing flags: %v", err)
	}
	if len(args) > 0 {
		log.Fatalf("unexpected arguments: %v", args)
	}

	return args
}
