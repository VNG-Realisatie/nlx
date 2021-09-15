// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/scheduler"
)

func (a *API) ListenAndServe(address, configAddress string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := scheduler.NewOutgoingAccessRequestScheduler(
		a.logger,
		a.directoryClient,
		a.configDatabase,
		a.orgCert,
	)

	go s.Run(ctx)

	g, ctx := errgroup.WithContext(context.Background())

	listen, err := net.Listen("tcp", configAddress)
	if err != nil {
		return err
	}

	g.Go(func() error {
		return a.grpcServer.Serve(listen)
	})

	tlsConfig := a.cert.TLSConfig()
	tlsConfig.InsecureSkipVerify = true //nolint:gosec // this is an internal connection to itself

	// setup client credentials for grpc gateway
	gatewayDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(tlsConfig),
		),
	}

	err = api.RegisterManagementHandlerFromEndpoint(ctx, a.mux, configAddress, gatewayDialOptions)
	if err != nil {
		return err
	}

	err = api.RegisterDirectoryHandlerFromEndpoint(ctx, a.mux, configAddress, gatewayDialOptions)
	if err != nil {
		return err
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", health)
	r.Mount("/api", a.authenticator.OnlyAuthenticated(a.mux))
	r.Mount("/api/v1/environment", environmentRoutes(a))
	a.authenticator.MountRoutes(r)

	a.httpServer = &http.Server{
		Addr:    address,
		Handler: r,
	}

	g.Go(func() error {
		err = a.httpServer.ListenAndServe()
		if err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	return g.Wait()
}

func (a *API) Shutdown(ctx context.Context) error {
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		shutdownGrpcServer(ctx, a.grpcServer)
		return nil
	})

	g.Go(func() error {
		err := a.httpServer.Shutdown(ctx)
		if err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	return g.Wait()
}

func shutdownGrpcServer(ctx context.Context, s *grpc.Server) {
	stopped := make(chan struct{})

	go func() {
		s.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		s.Stop()
	case <-stopped:
		return
	}
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
