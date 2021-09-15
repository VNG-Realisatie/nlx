// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/sign"
	"github.com/cloudflare/cfssl/signer"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	"go.nlx.io/nlx/ca-certportal/server"
	"go.nlx.io/nlx/common/logoptions"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-termChan
		cancel()
	}()

	// Setup new zap logger
	config := options.LogOptions.ZapConfig()

	logger, err := config.Build()
	if err != nil {
		//nolint:gocritic // we know that defer will not run after Fatalf
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	certSigner := func() (signer.Signer, error) {
		signer, signErr := sign.SignerFromConfig(cli.Config{
			Remote: options.CAHost,
		})
		if signErr != nil {
			logger.Error("failed to sign from config", zap.Error(signErr))
			return nil, signErr
		}

		return signer, nil
	}

	cp := server.NewCertPortal(logger, certSigner, options.ListenAddress)

	go func() {
		err = cp.Run()
		if err != nil {
			logger.Fatal("error running cert portal", zap.Error(err))
		}
	}()

	<-ctx.Done()

	err = cp.Shutdown()
	if err != nil {
		logger.Error("shutdown cert portal", zap.Error(err))
	}
}
