// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"crypto/tls"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/cloudflare/cfssl/log"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	common_tls "go.nlx.io/nlx/common/tls"
	directory_http "go.nlx.io/nlx/directory-inspection-api/http"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
)

type Server struct {
	httpServer  *http.Server
	httpsServer *http.Server
}

// newGRPCSplitterHandlerFunc returns an http.Handler that delegates gRPC connections to grpcServer
// and all other connections to otherHandler.
func newGRPCSplitterHandlerFunc(
	grpcServer,
	otherHandler http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// gRPC connection detected when HTTP protocol is version 2 and content-type is application/grpc
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func NewServer(
	logger *zap.Logger,
	address,
	addressPlain string,
	certificate *common_tls.CertificateBundle,
	inspectionService inspectionapi.DirectoryInspectionServer,
	httpServer *directory_http.Server) (*Server, error) {
	server := &Server{}

	// setup zap connection for global grpc logging
	grpc_zap.ReplaceGrpcLogger(logger)

	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Warn("recovered from a panic in a grpc request handler", zap.ByteString("stack", debug.Stack()))
			return status.Errorf(codes.Internal, "%s", p)
		}),
	}

	// prepare grpc server options
	opts := []grpc.ServerOption{
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

	// start grpc server and attach directory service
	grpcServer := grpc.NewServer(opts...)
	inspectionapi.RegisterDirectoryInspectionServer(grpcServer, inspectionService)

	tlsConfig := certificate.TLSConfig()
	tlsConfig.InsecureSkipVerify = true //nolint:gosec // local connection; hostname won't match

	gatewayDialOptions, gatewayMux := setupGateway(tlsConfig)

	err := inspectionapi.RegisterDirectoryInspectionHandlerFromEndpoint(
		context.Background(),
		gatewayMux,
		address,
		gatewayDialOptions,
	)
	if err != nil {
		return nil, err
	}

	httpServer.Mount("/", gatewayMux)

	// Start HTTPS server
	// let server handle connections on the TLS Listener
	tlsConfig = certificate.TLSConfig(certificate.WithTLSClientAuth())
	tlsConfig.NextProtos = []string{"h2"}

	server.httpsServer = &http.Server{
		Addr:      address,
		TLSConfig: tlsConfig,
		Handler:   newGRPCSplitterHandlerFunc(grpcServer, httpServer),
	}

	// Start plain HTTP server, during the PoC this is proxied by k8s ingress which adds TLS using letsencrypt.
	server.httpServer = &http.Server{
		Addr: addressPlain,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Incoming request", zap.Int("proto", r.ProtoMajor), zap.String("path", r.URL.Path))
			httpServer.ServeHTTP(w, r)
		}),
	}

	startHTTPServers(server.httpsServer, server.httpServer)

	return server, err
}

func setupGateway(tlsConfig *tls.Config) ([]grpc.DialOption, *runtime.ServeMux) {
	// setup client credentials for grpc gateway
	gatewayDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(tlsConfig),
		),
	}

	// setup grpc gateway and attach to main mux
	gatewayMetadata := func(context.Context, *http.Request) metadata.MD {
		return metadata.New(map[string]string{
			"grpcgateway-internal": "true",
		})
	}

	gatewayMux := runtime.NewServeMux(
		runtime.WithMetadata(gatewayMetadata),
		// Change the default behavior of marshaling to JSON
		// Emit empty fields by default
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{},
			},
		}),
	)

	return gatewayDialOptions, gatewayMux
}

func startHTTPServers(httpsServer, httpServer *http.Server) {
	go func() {
		err := httpsServer.ListenAndServeTLS("", "") // Key and Cert is empty because we provided them in TLSConfig
		if err != nil {
			if err != http.ErrServerClosed {
				log.Error("ListenAndServe for HTTPS failed", zap.Error(err))
			}
		}
	}()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Error("ListenAndServe plain HTTP failed", zap.Error(err))
			}
		}
	}()
}

func closeHTTPServer(s *http.Server) error {
	localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := s.Shutdown(localCtx)
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

const amountOfHTTPServers = 2

func (s *Server) Shutdown() error {
	wg := sync.WaitGroup{}
	wg.Add(amountOfHTTPServers)

	errChan := make(chan error)

	go func() {
		defer wg.Done()

		err := closeHTTPServer(s.httpServer)
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()

		err := closeHTTPServer(s.httpsServer)
		if err != nil {
			errChan <- err
		}
	}()

	wg.Wait()

	if len(errChan) > 0 {
		return <-errChan
	}

	return nil
}
