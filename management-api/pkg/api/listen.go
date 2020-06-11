// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/management-api/pkg/configapi"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (a *API) ListenAndServe(address, configAddress string) error {
	g, ctx := errgroup.WithContext(context.Background())

	listen, err := net.Listen("tcp", configAddress)
	if err != nil {
		log.Fatal("failed to create listener", zap.Error(err))
	}

	g.Go(func() error {
		return a.grpcServer.Serve(listen)
	})

	// setup client credentials for grpc gateway
	gatewayDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				Certificates:       []tls.Certificate{*a.orgCertKeyPair},
				RootCAs:            a.roots,
				InsecureSkipVerify: true, //nolint:gosec // this is an internal connection to itself
			}),
		),
	}

	err = configapi.RegisterConfigApiHandlerFromEndpoint(ctx, a.mux, configAddress, gatewayDialOptions)
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", health)
	r.Mount("/oidc", a.authenticator.Routes())
	r.Mount("/api", a.authenticator.OnlyAuthenticated(a.mux))
	r.Mount("/api/v1/directory", directoryRoutes(a))
	r.Mount("/api/v1/environment", environmentRoutes(a))

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	// ErrServerClosed is more info message than error
	g.Go(server.ListenAndServe)

	shutDownComplete := make(chan struct{})

	a.process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak
		sherr := server.Shutdown(localCtx)
		close(shutDownComplete)
		return sherr
	})

	// Listener will return immediately on Shutdown call.
	// So we need to wait until all open connections will be closed gracefully
	<-shutDownComplete

	err = g.Wait()

	if err != http.ErrServerClosed {
		return errors.Wrap(err, "failed to run http server")
	}

	return nil
}

type healthResponse struct {
	Status string `json:"status"`
}

func health(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, healthResponse{Status: "ok"})
}

// ServeHTTP handles a specific HTTP request
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
