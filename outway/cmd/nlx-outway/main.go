// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"log"

	flags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/VNG-Realisatie/nlx/outway"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"localhost:12018" description:"Adress for the outway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
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

	// Create new outway and provide it with a hardcoded service.
	ow := outway.NewOutway(logger, "DemoRequesterOrganization")
	// hardcoded service+inway because we don't have directory up yet
	echoService, err := outway.NewService(logger, "DemoProviderOrganization", "PostmanEcho", "http://localhost:2018")
	if err != nil {
		logger.Fatal("failed to create new PostmanEcho service", zap.Error(err))
	}
	ow.AddService(echoService)
	// Listen on the address provided in the options
	err = ow.ListenAndServe(options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
