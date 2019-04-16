// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"log"

	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/inway"
)

var options struct {
	ConfigAPIAAddress string `long:"config-api-address" env:"CONFIG_API_ADDRESS" description:"Address of the config-api" required:"true"`
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
	zapConfig := options.LogOptions.ZapConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}
	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))
	process := process.NewProcess(logger)

	iw, err := inway.NewInway(logger, options.TLSOptions, options.ConfigAPIAAddress)
	if err != nil {
		logger.Fatal("cannot setup inway", zap.Error(err))
	}

	if err := iw.Start(process); err != nil {
		log.Fatal(err)
	}

}
