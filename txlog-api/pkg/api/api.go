// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain/txlog/storage"
	"go.nlx.io/nlx/txlog-api/pkg/server"
)

// API handles incoming requests and authenticates them
type API struct {
	logger     *zap.Logger
	cert       *common_tls.CertificateBundle
	mux        *runtime.ServeMux
	grpcServer *grpc.Server
	httpServer *http.Server
	storage    storage.Repository
}

// NewAPI creates and prepares a new API
//nolint:gocyclo // parameter validation
func NewAPI(logger *zap.Logger, cert *common_tls.CertificateBundle, s storage.Repository) (*API, error) {
	if logger == nil {
		return nil, errors.New("logger is required")
	}

	if cert == nil {
		return nil, errors.New("cert is required")
	}

	if s == nil {
		return nil, errors.New("storage is required")
	}

	txlogService := server.NewTXLogService(
		logger,
		s,
	)

	grpcServer := newGRPCServer(logger, cert)

	api.RegisterTXLogServer(grpcServer, txlogService)

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

	a := &API{
		logger:     logger,
		cert:       cert,
		grpcServer: grpcServer,
		mux:        mux,
		storage:    s,
	}

	return a, nil
}

func newGRPCServer(logger *zap.Logger, cert *common_tls.CertificateBundle) *grpc.Server {
	// setup zap connection for global grpc logging
	// grpc_zap.ReplaceGrpcLogger(logger)

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
