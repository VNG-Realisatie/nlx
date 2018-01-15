// Copyright 2018 VNG Realisatie. All rights reserved.
// Use of this source code is governed by the EUPL
// license that can be found in the LICENSE.md file.

package main

import (
	"log"

	flags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/VNG-Realisatie/nlx/inway"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"localhost:2018" description:"Adress for the inway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
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

	// Create new inway and provide it with a hardcoded service.
	iw := inway.NewInway(logger, "DemoProviderOrganization")
	// NOTE: hardcoded service endpoint because we don't have any other means to register endpoints yet
	echoServiceEndpoint, err := inway.NewHTTPServiceEndpoint(logger, "PostmanEcho", "https://postman-echo.com/")
	if err != nil {
		logger.Fatal("failed to create PostmanEcho service", zap.Error(err))
	}
	iw.AddServiceEndpoint(echoServiceEndpoint)
	// Listen on the address provided in the options
	err = iw.ListenAndServe(options.ListenAddress)
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
