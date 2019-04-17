// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/tlsconfig"

	"github.com/pkg/errors"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (o *Outway) ListenAndServe(process *process.Process, address string) error {
	server := &http.Server{
		Addr:    address,
		Handler: o,
	}

	process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		return server.Shutdown(localCtx)
	})

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return errors.Wrap(err, "failed to run http server")
	}

	o.wg.Wait() // Wait until all async jobs will finish
	return nil
}

// ListenAndServeTLS is a blocking function that listens on provided tcp address to handle requests.
func (o *Outway) ListenAndServeTLS(process *process.Process, address string, certFile, keyFile string) error {
	server := &http.Server{
		Addr:      address,
		Handler:   o,
		TLSConfig: tlsconfig.Defaults(),
	}

	process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		return server.Shutdown(localCtx)
	})

	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil && err != http.ErrServerClosed {
		return errors.Wrap(err, "failed to run http server")
	}

	o.wg.Wait() // Wait until all async jobs will finish
	return nil
}

func createHTTPTransport(tlsConfig *tls.Config) *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}
}

// ServeHTTP handles requests from the organization to the outway, it selects the correct service backend and lets it handle the request further.
func (o *Outway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		o.handleTunneling(w, r)
	} else {
		o.handleDirect(w, r)
	}
}
