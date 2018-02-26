// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"log"

	flags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/VNG-Realisatie/nlx/common/orgtls"
	"github.com/VNG-Realisatie/nlx/common/process"
	"github.com/VNG-Realisatie/nlx/outway"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:12018" description:"Adress for the outway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`

	DirectoryAddress string `long:"directory-address" env:"DIRECTORY_ADDRESS" description:"Address for the directory where this outway can fetch the service list" required:"true"`

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
	config := zap.NewDevelopmentConfig()
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

	// Create new outway and provide it with a hardcoded service.
	ow, err := outway.NewOutway(logger, options.TLSOptions, options.DirectoryAddress)
	if err != nil {
		logger.Fatal("failed to setup outway", zap.Error(err))
	}
	// Listen on the address provided in the options
	err = ow.ListenAndServe(options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
