// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"crypto/tls"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"

	"go.nlx.io/nlx/common/tlsconfig"
)

// RunServer is a blocking function that listens on provided tcp address to handle requests.
func (i *Inway) RunServer(address string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/.nlx/api-spec-doc/", i.handleAPISpecDocRequest)
	serveMux.HandleFunc("/.nlx/health/", i.handleHealthRequest)
	serveMux.Handle("/.nlx/", http.NotFoundHandler())
	serveMux.HandleFunc("/", i.handleProxyRequest)

	config := &tls.Config{
		// only allow clients that present a cert signed by our root CA
		ClientCAs:    i.roots,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{*i.orgKeyPair},
	}

	tlsconfig.ApplyDefaults(config)

	server := &http.Server{
		Addr:      address,
		Handler:   serveMux,
		TLSConfig: config,
	}

	errorChannel := make(chan error)

	go func() {
		if err := server.ListenAndServeTLS("", ""); err != nil {
			if err != http.ErrServerClosed {
				errorChannel <- errors.Wrap(err, "error listening on TLS server")
			}
		}
	}()

	go func() {
		if err := i.monitoringService.Start(); err != nil {
			errorChannel <- errors.Wrap(err, "error listening on TLS server")
		}
	}()

	wg := sync.WaitGroup{}

	numberOfServers := 2
	wg.Add(numberOfServers)

	i.process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak
		err := server.Shutdown(localCtx)
		wg.Done()
		return err
	})

	i.process.CloseGracefully(func() error {
		err := i.monitoringService.Stop()
		wg.Done()
		return err
	})

	// Listener will return immediately on Shutdown call.
	// So we need to wait until all open connections will be closed gracefully

	return <-errorChannel
}
