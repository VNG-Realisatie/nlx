// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// RunServer is a blocking function that listens on provided tcp address to handle requests.
func (i *Inway) RunServer(address, managementAddress string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/.nlx/api-spec-doc/", i.handleAPISpecDocRequest)
	serveMux.HandleFunc("/.nlx/health/", i.handleHealthRequest)
	serveMux.Handle("/.nlx/", http.NotFoundHandler())
	serveMux.HandleFunc("/", i.handleProxyRequest)

	config := i.orgCertBundle.TLSConfig(i.orgCertBundle.WithTLSClientAuth())
	i.serverTLS = &http.Server{
		Addr:      address,
		Handler:   serveMux,
		TLSConfig: config,
	}

	errorChannel := make(chan error)

	go func() {
		if err := i.serverTLS.ListenAndServeTLS("", ""); err != nil {
			if err != http.ErrServerClosed {
				errorChannel <- errors.Wrap(err, "error listening on TLS server")
			}
		}
	}()

	go func() {
		if err := i.monitoringService.Start(); err != nil {
			errorChannel <- errors.Wrap(err, "error listening on monitoring service")
		}
	}()

	if i.managementProxy != nil {
		i.logger.Info("management proxy: starting")

		go func() {
			i.logger.Info("management proxy: listening on %v", zap.String("management-address", managementAddress))

			l, err := net.Listen("tcp", managementAddress)
			if err != nil {
				errorChannel <- errors.Wrap(err, "listen on management-address")
			}

			if err := i.managementProxy.Serve(l); err != nil {
				errorChannel <- errors.Wrap(err, "management proxy")
			}
		}()
	}

	i.process.CloseGracefully(func() error {
		i.shutDown()
		return nil
	})

	err := <-errorChannel

	i.shutDown()

	return err
}

func (i *Inway) shutDown() {
	i.monitoringService.SetNotReady()

	localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak

	err := i.serverTLS.Shutdown(localCtx)
	if err != nil {
		i.logger.Error("error shutting down server tls", zap.Error(err))
	}

	if i.managementProxy != nil {
		i.managementProxy.Stop()
	}

	err = i.monitoringService.Stop()
	if err != nil {
		i.logger.Error("error shutting down monitoring service", zap.Error(err))
	}
}
