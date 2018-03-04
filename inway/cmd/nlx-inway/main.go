// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	flags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/VNG-Realisatie/nlx/common/orgtls"
	"github.com/VNG-Realisatie/nlx/common/process"
	"github.com/VNG-Realisatie/nlx/inway"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:2018" description:"Adress for the inway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryAddress string `long:"directory-address" env:"DIRECTORY_ADDRESS" description:"Address for the directory where this inway can register it's services" required:"true"`

	SelfAddress string `long:"self-address" env:"SELF_ADDRESS" description:"The address that outways can use to reach me" required:"true"`

	ServiceConfig string `long:"service-config" env:"SERVICE_CONFIG" default:"service-config.toml" description:"Location of the service config toml file"`

	orgtls.TLSOptions
}

// ServiceConfig is the top-level for the service configuration file.
type ServiceConfig struct {
	Services map[string]ServiceDetails
}

// ServiceDetails holds the details for a single service definition.
type ServiceDetails struct {
	Address string
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
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	defer func() { // TODO(GeertJohan): make this a common/process exitFunc?
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

	serviceConfig := &ServiceConfig{}
	tomlMetaData, err := toml.DecodeFile(options.ServiceConfig, serviceConfig)
	if err != nil {
		logger.Fatal("failed to load service config", zap.Error(err))
	}
	if len(tomlMetaData.Undecoded()) > 0 {
		logger.Fatal("unsupported values in toml", zap.String("key", strings.Join(tomlMetaData.Undecoded()[0], ">")))
	}

	iw, err := inway.NewInway(logger, options.SelfAddress, options.TLSOptions, options.DirectoryAddress)
	if err != nil {
		logger.Fatal("cannot setup inway", zap.Error(err))
	}

	for serviceName, serviceDetails := range serviceConfig.Services {
		endpoint, err := inway.NewHTTPServiceEndpoint(logger, serviceName, serviceDetails.Address)
		if err != nil {
			logger.Fatal("failed to create service", zap.Error(err))
		}
		iw.AddServiceEndpoint(endpoint)
	}
	// Listen on the address provided in the options
	err = iw.ListenAndServeTLS(options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
