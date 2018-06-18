// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"fmt"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:2018" description:"Adress for the inway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryAddress string `long:"directory-address" env:"DIRECTORY_ADDRESS" description:"Address for the directory where this inway can register it's services" required:"true"`

	SelfAddress string `long:"self-address" env:"SELF_ADDRESS" description:"The address that outways can use to reach me" required:"true"`

	ServiceConfig string `long:"service-config" env:"SERVICE_CONFIG" default:"service-config.toml" description:"Location of the service config toml file"`

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
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	defer func() { // TODO(GeertJohan): #205 make this a common/process exitFunc?
		syncErr := logger.Sync()
		if syncErr != nil {
			// notify the user that proper logging has failed
			fmt.Fprintf(os.Stderr, "failed to sync zap logger: %v\n", syncErr)
			// don't exit when we're in a panic
			if p := recover(); p != nil {
				panic(p)
			}
			os.Exit(1)
		}
	}()

	process.Setup(logger)

	serviceConfig := loadServiceConfig(logger, options.ServiceConfig)

	iw, err := inway.NewInway(logger, options.SelfAddress, options.TLSOptions, options.DirectoryAddress)
	if err != nil {
		logger.Fatal("cannot setup inway", zap.Error(err))
	}

	for serviceName, serviceDetails := range serviceConfig.Services {
		endpoint, err := inway.NewHTTPServiceEndpoint(logger, serviceName, serviceDetails.Address)
		if err != nil {
			logger.Fatal("failed to create service", zap.Error(err))
		}
		switch serviceDetails.AuthorizationModel {
		case "none", "":
			endpoint.SetAuthorizationPublic()
		case "whitelist":
			endpoint.SetAuthorizationWhitelist(serviceDetails.AuthorizationWhitelist)
		default:
			logger.Fatal(fmt.Sprintf(`invalid authorization model "%s" for service "%s"`, serviceDetails.AuthorizationModel, serviceName))
		}
		iw.AddServiceEndpoint(endpoint, serviceDetails.DocumentationURL)
	}
	// Listen on the address provided in the options
	err = iw.ListenAndServeTLS(options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
