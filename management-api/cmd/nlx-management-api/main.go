// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"log"

	flags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/daos"
	"go.nlx.io/nlx/management-api/repositories"
	"go.nlx.io/nlx/management-api/session"
)

var options struct {
	ListenAddress    string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8080" description:"Address for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ConfigAPIAddress string `long:"config-api-address" env:"CONFIG_API_ADDRESS" description:"Address of the config API. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	logoptions.LogOptions
	orgtls.TLSOptions
	daos.AccountCSVOptions
	session.SessionstoreOptions
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
	zapConfig := options.LogOptions.ZapConfig()
	logger, err := zapConfig.Build()

	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	logger.Info("starting management api", zap.String("listen-address", options.ListenAddress))

	mainProcess := process.NewProcess(logger)

	accountCSV, err := daos.NewAccountCSV(options.AccountCSVOptions.CsvFileName)
	if err != nil {
		logger.Fatal("cannot load accounts csv file", zap.String("csv file", options.AccountCSVOptions.CsvFileName), zap.Error(err))
	}

	accountRepository, err := repositories.NewAccount(accountCSV)
	if err != nil {
		logger.Fatal("cannot load accounts repository", zap.Error(err))
	}

	sessionstore := session.NewSessionstoreImpl(logger, options.SessionstoreOptions, accountRepository)

	authorizer := session.NewAuthorizer()

	a, err := api.NewAPI(logger, mainProcess, options.TLSOptions, options.ConfigAPIAddress, sessionstore, authorizer)
	if err != nil {
		logger.Fatal("cannot setup management api", zap.Error(err))
	}

	// Listen on the address provided in the options
	err = a.ListenAndServe(options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
