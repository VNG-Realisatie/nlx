// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"go.nlx.io/nlx/ca-certportal/server"
	"log"
	"net/http"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/sign"
	"github.com/cloudflare/cfssl/signer"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
)

var options struct {
	ListenAddress string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8080" description:"Address for the certportal to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	CAHost        string `long:"ca-host" env:"CA_HOST" default:"localhost" description:"The host of the certificate authority (CA)."`
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

	if options.CAHost == "" {
		log.Fatal("CA host option is empty")
	}

	// Setup new zap logger
	config := options.LogOptions.ZapConfig()

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	process.NewProcess(logger)

	// Create new certportal and provide it with a hardcoded service.
	cp := server.NewCertPortal(logger, func() (signer.Signer, error) {
		signer, signErr := sign.SignerFromConfig(cli.Config{
			Remote: options.CAHost,
		})
		if signErr != nil {
			return nil, signErr
		}

		return signer, nil
	})
	// Listen on the address provided in the options
	err = http.ListenAndServe(options.ListenAddress, cp.GetRouter())
	if err != nil {
		logger.Fatal("failed to listen and serve", zap.Error(err))
	}
}
