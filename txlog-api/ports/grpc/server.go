// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package grpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/app"
)

type Server struct {
	api.UnimplementedTXLogServer
	app        *app.Application
	logger     *zap.Logger
	clock      Clock
	mux        *runtime.ServeMux
	service    *grpc.Server
	cert       *common_tls.CertificateBundle
	httpServer *http.Server
}

var readHeaderTimeout = 5 * time.Second

type Clock interface {
	Now() time.Time
}

func New(logger *zap.Logger, clock Clock, a *app.Application, cert *common_tls.CertificateBundle) *Server {
	grpcServer := newGRPCServer(logger, cert)

	mux := runtime.NewServeMux(
		// Change the default behavior of marshaling to JSON
		// Emit empty fields by default
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}))

	server := &Server{
		logger:  logger,
		clock:   clock,
		app:     a,
		mux:     mux,
		service: grpcServer,
		cert:    cert,
	}

	api.RegisterTXLogServer(grpcServer, server)

	return server
}

func newGRPCServer(logger *zap.Logger, cert *common_tls.CertificateBundle) *grpc.Server {
	tlsConfig := cert.TLSConfig(cert.WithTLSClientAuth())
	transportCredentials := credentials.NewTLS(tlsConfig)

	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Warn("recovered from a panic in a grpc request handler", zap.ByteString("stack", debug.Stack()))
			return status.Error(codes.Internal, fmt.Sprintf("%s", p))
		}),
	}

	opts := []grpc.ServerOption{
		grpc.Creds(transportCredentials),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(recoveryOptions...),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(recoveryOptions...),
		),
	}

	return grpc.NewServer(opts...)
}

func (s *Server) ListenAndServe(address, addressPlain string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	g.Go(func() error {
		return s.service.Serve(listen)
	})

	tlsConfig := s.cert.TLSConfig()
	tlsConfig.InsecureSkipVerify = true //nolint:gosec // this is an internal connection to itself

	// setup client credentials for grpc gateway
	gatewayDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(tlsConfig),
		),
	}

	err = api.RegisterTXLogHandlerFromEndpoint(ctx, s.mux, address, gatewayDialOptions)
	if err != nil {
		return err
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", health)
	r.Mount("/api", s.mux)

	s.httpServer = &http.Server{
		Addr:              addressPlain,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	g.Go(func() error {
		err = s.httpServer.ListenAndServe()
		if err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	return g.Wait()
}

func (s *Server) Shutdown(ctx context.Context) error {
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		shutdownGrpcServer(ctx, s.service)
		return nil
	})

	g.Go(func() error {
		err := s.httpServer.Shutdown(ctx)
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
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
