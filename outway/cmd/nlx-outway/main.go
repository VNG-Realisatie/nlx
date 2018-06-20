// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"log"

	"github.com/huandu/xstrings"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/logdb/logdbversion"
	"go.nlx.io/nlx/outway"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:80" description:"Adress for the outway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryAddress string `long:"directory-address" env:"DIRECTORY_ADDRESS" description:"Address for the directory where this outway can fetch the service list" required:"true"`

	PostgresDSN string `long:"postgres-dsn" env:"POSTGRES_DSN" default:"postgres://postgres:postgres@postgres/nlx_logdb?sslmode=disable" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`

	logoptions.LogOptions
	orgtls.TLSOptions
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
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	defer func() {
		syncErr := logger.Sync()
		if syncErr != nil {
			log.Fatalf("failed to sync zap logger: %v", syncErr)
		}
	}()

	process.Setup(logger)

	logDB, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		logger.Fatal("could not open connection to postgres", zap.Error(err))
	}
	logDB.MapperFunc(xstrings.ToSnakeCase)

	logdbversion.WaitUntilLatestVersion(logger, logDB.DB)

	// Create new outway and provide it with a hardcoded service.
	ow, err := outway.NewOutway(logger, logDB, options.TLSOptions, options.DirectoryAddress)
	if err != nil {
		logger.Fatal("failed to setup outway", zap.Error(err))
	}
	// Listen on the address provided in the options
	err = ow.ListenAndServe(options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
