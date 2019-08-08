// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/tlsconfig"
	directory_http "go.nlx.io/nlx/directory-inspection-api/http"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
)

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

// runServer is a blocking function which sets up the grpc and http/json server and runs them on a single address/port.
func runServer(
	p *process.Process,
	log *zap.Logger,
	address,
	addressPlain string,
	caCertPool *x509.CertPool,
	certKeyPair *tls.Certificate,
	inspectionService inspectionapi.DirectoryInspectionServer,
	httpServer *directory_http.Server,
) {

	// setup zap connection for global grpc logging
	grpc_zap.ReplaceGrpcLogger(log)

	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			log.Warn("recovered from a panic in a grpc request handler", zap.ByteString("stack", debug.Stack()))
			return status.Errorf(codes.Internal, "%s", p)
		}),
	}

	// prepare grpc server options
	opts := []grpc.ServerOption{
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(log),
			grpc_recovery.StreamServerInterceptor(recoveryOptions...),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(recoveryOptions...),
		),
	}

	// start grpc server and attach directory service
	grpcServer := grpc.NewServer(opts...)
	inspectionapi.RegisterDirectoryInspectionServer(grpcServer, inspectionService)

	// setup client credentials for grpc gateway
	gatewayDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				Certificates: []tls.Certificate{*certKeyPair}, // using the grpc server's own cert to connect to it, perhaps find a way for the http/json gateway to bypass TLS locally
				RootCAs:      caCertPool,
				// This is a local connection; hostname won't match
				InsecureSkipVerify: true, //nolint:gosec
			}),
		),
	}

	// setup grpc gateway and attach to main mux
	gatewayMux := runtime.NewServeMux()
	err := inspectionapi.RegisterDirectoryInspectionHandlerFromEndpoint(
		context.Background(),
		gatewayMux,
		address,
		gatewayDialOptions,
	)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}

	httpServer.Mount("/", gatewayMux)

	// Start HTTPS server
	// let server handle connections on the TLS Listener
	HTTPSHandler := &http.Server{
		Addr: address,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*certKeyPair},
			NextProtos:   []string{"h2"},
			ClientCAs:    caCertPool,
			ClientAuth:   tls.RequireAndVerifyClientCert,
		},
		Handler: newGRPCSplitterHandlerFunc(grpcServer, httpServer),
	}
	tlsconfig.ApplyDefaults(HTTPSHandler.TLSConfig)

	go func() {
		err = HTTPSHandler.ListenAndServeTLS("", "") // Key and Cert is empty because we provided them in TLSConfig
		if err != nil {
			if err != http.ErrServerClosed {
				log.Error("ListenAndServe for HTTPS failed", zap.Error(err))
			}
		}
	}()

	// Start plain HTTP server, during the PoC this is proxied by k8s ingress which adds TLS using letsencrypt.
	HTTPHandler := &http.Server{
		Addr: addressPlain,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("proto:%d path:%s\n", r.ProtoMajor, r.URL.Path)
			httpServer.ServeHTTP(w, r)
		}),
	}
	// TODO: #206 When directory has a separate storage/backing, the inspection API should actually become a separate process.
	go func() {
		if err := HTTPHandler.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Error("ListenAndServe plain HTTP failed", zap.Error(err))
			}
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(2)
	p.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak
		err := HTTPSHandler.Shutdown(localCtx)
		wg.Done()
		return err
	})
	p.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak
		err := HTTPHandler.Shutdown(localCtx)
		wg.Done()
		return err
	})

	wg.Wait()
}
