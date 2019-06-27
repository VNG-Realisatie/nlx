// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/tlsconfig"
)

// ListenAndServeTLS is a blocking function that listens on provided tcp address to handle requests.
func (i *Inway) ListenAndServeTLS(xprocess *process.Process, address string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/.nlx/api-spec-doc/", i.handleAPISpecDocRequest)
	serveMux.HandleFunc("/.nlx/health/", i.handleHealthRequest)
	serveMux.Handle("/.nlx/", http.NotFoundHandler())
	serveMux.HandleFunc("/", i.handleProxyRequest)
	server := &http.Server{
		Addr: address,
		TLSConfig: &tls.Config{
			// only allow clients that present a cert signed by our root CA
			ClientCAs:  i.roots,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
		Handler: serveMux,
	}
	tlsconfig.ApplyDefaults(server.TLSConfig)

	shutDownComplete := make(chan struct{})
	err := xprocess.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak
		err := server.Shutdown(localCtx)
		close(shutDownComplete)
		return err
	})

	if err != nil {
		return errors.Wrap(err, "failed to close gracefully")
	}

	// ErrServerClosed is more info message than error
	if err := server.ListenAndServeTLS(i.orgCertFile, i.orgKeyFile); err != nil {
		if err != http.ErrServerClosed {
			return errors.Wrap(err, "failed to run http server")
		}
	}

	// Listener will return immediately on Shutdown call.
	// So we need to wait until all open connections will be closed gracefully
	<-shutDownComplete
	return nil
}
