// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/environment"
	"go.nlx.io/nlx/management-api/pkg/oidc"
	"go.nlx.io/nlx/management-api/pkg/server"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

// API handles incoming requests and authenticates them
type API struct {
	logger          *zap.Logger
	environment     *environment.Environment
	cert            *common_tls.CertificateBundle
	orgCert         *common_tls.CertificateBundle
	process         *process.Process
	mux             *runtime.ServeMux
	grpcServer      *grpc.Server
	authenticator   *oidc.Authenticator
	directoryClient directory.Client
}

const (
	directoryClientDialTimeout = 1 * time.Minute
	singleElementArrayLength   = 1
)

// NewAPI creates and prepares a new API
//nolint:gocyclo // parameter validation
func NewAPI(logger *zap.Logger, mainProcess *process.Process, cert, orgCert *common_tls.CertificateBundle, etcdConnectionString, directoryRegistrationAddress, directoryEndpointURL string, authenticator *oidc.Authenticator) (*API, error) {
	if mainProcess == nil {
		return nil, errors.New("process argument is nil. needed to close gracefully")
	}

	if len(orgCert.Certificate().Subject.Organization) != singleElementArrayLength {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	if etcdConnectionString == "" {
		return nil, errors.New("etcd connection string is not configured")
	}

	if directoryRegistrationAddress == "" {
		return nil, errors.New("directory registration address is not configured")
	}

	if directoryEndpointURL == "" {
		return nil, errors.New("directory endpoint URL is not configured")
	}

	if authenticator == nil {
		return nil, errors.New("authenticator is not configured")
	}

	directoryRegistrationClient, err := newDirectoryRegistrationClient(orgCert, directoryRegistrationAddress)
	if err != nil {
		logger.Fatal("failed to setup directory client", zap.Error(err))
	}

	db := newDatabase(logger, mainProcess, etcdConnectionString)
	managementService := server.NewManagementService(logger, mainProcess, directoryRegistrationClient, db)

	grpcServer := newGRPCServer(logger, cert)

	api.RegisterManagementServer(grpcServer, managementService)
	e := &environment.Environment{
		OrganizationName: orgCert.Certificate().Subject.Organization[0],
	}

	directoryClient, err := directory.NewClient(directoryEndpointURL)
	if err != nil {
		return nil, err
	}

	directoryService := server.NewDirectoryService(logger, e, directoryClient, db)

	api.RegisterDirectoryServer(grpcServer, directoryService)

	api := &API{
		logger:          logger.With(zap.String("api-organization-name", e.OrganizationName)),
		environment:     e,
		cert:            cert,
		orgCert:         orgCert,
		grpcServer:      grpcServer,
		process:         mainProcess,
		mux:             runtime.NewServeMux(),
		authenticator:   authenticator,
		directoryClient: directoryClient,
	}

	return api, nil
}

func newGRPCServer(logger *zap.Logger, cert *common_tls.CertificateBundle) *grpc.Server {
	// setup zap connection for global grpc logging
	grpc_zap.ReplaceGrpcLogger(logger)

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

func newDatabase(logger *zap.Logger, mainProcess *process.Process, etcdConnectionString string) database.ConfigDatabase {
	c := clock.RealClock{}

	db, err := database.NewEtcdConfigDatabase(logger, mainProcess, strings.Split(etcdConnectionString, ","), c)
	if err != nil {
		logger.Fatal("failed to setup database", zap.Error(err))
	}

	return db
}

func newDirectoryRegistrationClient(cert *common_tls.CertificateBundle, directoryRegistrationAddress string) (registrationapi.DirectoryRegistrationClient, error) {
	directoryDialCredentials := credentials.NewTLS(cert.TLSConfig())
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}
	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(nlxversion.NewContext("management-api"), directoryClientDialTimeout)
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryRegistrationAddress, directoryDialOptions...)

	defer directoryConnCtxCancel()

	if err != nil {
		return nil, err
	}

	directoryRegistrationClient := registrationapi.NewDirectoryRegistrationClient(directoryConn)

	return directoryRegistrationClient, nil
}
