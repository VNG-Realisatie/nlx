// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/management-api/authorization"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (a *API) ListenAndServe(address string) error {
	ctx, cancel := context.WithCancel(context.Background())
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
	r.Use(a.authenticationManager.Middleware)

	r.Get("/health", heatlh)

	apiRouter := chi.NewRouter()
	apiRouter.Use(authorization.NewAuthorization(a.authorizer).Middleware)
	apiRouter.Mount("/auth", a.authenticationManager.Routes())
	apiRouter.Mount("/", a.mux)
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

type healthResponse struct {
	Status string `json:"status"`
}

func heatlh(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, healthResponse{Status: "ok"})
}

// ServeHTTP handles a specific HTTP request
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
