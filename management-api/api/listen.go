// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"go.nlx.io/nlx/management-api/authorization"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/config-api/configapi"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (a *API) ListenAndServe(address string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// setup client credentials for grpc gateway
	gatewayDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				Certificates: []tls.Certificate{*a.orgCertKeyPair},
				RootCAs:      a.roots,
			}),
		),
	}

	err := configapi.RegisterConfigApiHandlerFromEndpoint(ctx, a.mux, a.configAPIAddress, gatewayDialOptions)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(a.sessionstore.Middleware)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			panic(err)
		}
	})
	r.Mount("/session", a.sessionstore.Routes())

	apiRouter := chi.NewRouter()
	apiRouter.Use(authorization.NewAuthorization(a.authorizer).Middleware)
	r.Mount("/api", apiRouter)
	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	// ErrServerClosed is more info message than error
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			return errors.Wrap(err, "failed to run http server")
		}
	}

	shutDownComplete := make(chan struct{})
	a.process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak
		err := server.Shutdown(localCtx)
		close(shutDownComplete)
		return err
	})

	// Listener will return immediately on Shutdown call.
	// So we need to wait until all open connections will be closed gracefully
	<-shutDownComplete
	return nil
}

// ServeHTTP handles a specific HTTP request
func (a API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
