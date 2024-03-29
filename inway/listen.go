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

const readHeaderTimeout = time.Second * 60

func (i *Inway) Run(ctx context.Context, address string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/.nlx/api-spec-doc/", i.handleAPISpecDocRequest)
	serveMux.HandleFunc("/.nlx/health/", i.handleHealthRequest)
	serveMux.Handle("/.nlx/", http.NotFoundHandler())
	serveMux.HandleFunc("/", i.handleProxyRequest)

	config := i.orgCertBundle.TLSConfig(i.orgCertBundle.WithTLSClientAuth())

	i.serverTLS = &http.Server{
		Addr:              address,
		Handler:           serveMux,
		TLSConfig:         config,
		ReadHeaderTimeout: readHeaderTimeout,
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
			if err != http.ErrServerClosed {
				errorChannel <- errors.Wrap(err, "error listening on monitoring service")
			}
		}
	}()

	go i.startConfigurationPolling(ctx)

	go i.announceToDirectory(ctx)

	i.logger.Info("management proxy: starting")

	go func() {
		i.logger.Info("management proxy: listening", zap.String("management-address", i.listenAddressManagementAPIProxy))

		l, err := net.Listen("tcp", i.listenAddressManagementAPIProxy)
		if err != nil {
			errorChannel <- errors.Wrap(err, "listen on management-address")
		}

		if err := i.managementProxy.Serve(l); err != nil {
			if err != http.ErrServerClosed {
				errorChannel <- errors.Wrap(err, "management proxy")
			}
		}
	}()

	return <-errorChannel
}

func (i *Inway) Shutdown(ctx context.Context) error {
	i.monitoringService.SetNotReady()

	err := i.serverTLS.Shutdown(ctx)
	if err != http.ErrServerClosed {
		return err
	}

	i.managementProxy.Stop()

	if err := i.monitoringService.Stop(); err != nil {
		return err
	}

	return nil
}
