// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/environment"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/outway"
	"go.nlx.io/nlx/management-api/pkg/server"
	"go.nlx.io/nlx/management-api/pkg/txlog"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
)

// API handles incoming requests and authenticates them
type API struct {
	logger          *zap.Logger
	environment     *environment.Environment
	cert            *common_tls.CertificateBundle
	orgCert         *common_tls.CertificateBundle
	mux             *runtime.ServeMux
	grpcServer      *grpc.Server
	httpServer      *http.Server
	authenticator   Authenticator
	directoryClient directory.Client
	txlogClient     txlog.Client
	configDatabase  database.ConfigDatabase
}

type NewAPIArgs struct {
	DB               database.ConfigDatabase
	TXlogDB          txlogdb.TxlogDatabase
	Logger           *zap.Logger
	InternalCert     *common_tls.CertificateBundle
	OrgCert          *common_tls.CertificateBundle
	DirectoryAddress string
	TXLogAddress     string
	Authenticator    Authenticator
	AuditLogger      auditlog.Logger
}

type clock struct{}

func (c *clock) Now() time.Time {
	return time.Now()
}

type Authenticator interface {
	MountRoutes(router chi.Router)
	OnlyAuthenticated(h http.Handler) http.Handler
	UnaryServerInterceptor(configDatabase database.ConfigDatabase, getUserFromDatabase func(ctx context.Context, configDatabase database.ConfigDatabase, email string) (*domain.User, error)) grpc.UnaryServerInterceptor
	StreamServerInterceptor(configDatabase database.ConfigDatabase, getUserFromDatabase func(ctx context.Context, configDatabase database.ConfigDatabase, email string) (*domain.User, error)) grpc.StreamServerInterceptor
}

//nolint:gocyclo // parameter validation
func NewAPI(args *NewAPIArgs) (*API, error) {
	if args.DB == nil {
		return nil, errors.New("database is not configured")
	}

	if len(args.OrgCert.Certificate().Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	if args.DirectoryAddress == "" {
		return nil, errors.New("directory address is not configured")
	}

	if args.Authenticator == nil {
		return nil, errors.New("authenticator is not configured")
	}

	directoryClient, err := directory.NewClient(context.TODO(), args.DirectoryAddress, args.OrgCert)
	if err != nil {
		args.Logger.Fatal("failed to setup directory client", zap.Error(err))
	}

	var txlogClient txlog.Client
	if args.TXLogAddress != "" {
		txlogClient, err = txlog.NewClient(context.TODO(), args.TXLogAddress, args.InternalCert)
		if err != nil {
			args.Logger.Fatal("failed to setup txlog client", zap.Error(err))
		}
	}

	managementService := server.NewManagementService(
		&server.NewManagementServiceArgs{
			Logger:                     args.Logger,
			DirectoryClient:            directoryClient,
			TxlogClient:                txlogClient,
			OrgCert:                    args.OrgCert,
			InternalCert:               args.InternalCert,
			ConfigDatabase:             args.DB,
			TxlogDatabase:              args.TXlogDB,
			AuditLogger:                args.AuditLogger,
			CreateManagementClientFunc: management.NewClient,
			CreateOutwayClientFunc:     outway.NewClient,
			Clock:                      &clock{},
		})

	grpcServer := newGRPCServer(args.Logger, args.InternalCert, args.DB, args.Authenticator)

	api.RegisterManagementServer(grpcServer, managementService)
	external.RegisterAccessRequestServiceServer(grpcServer, managementService)
	external.RegisterDelegationServiceServer(grpcServer, managementService)

	e := &environment.Environment{
		OrganizationSerialNumber: args.OrgCert.Certificate().Subject.SerialNumber,
		OrganizationName:         args.OrgCert.Certificate().Subject.Organization[0],
	}

	directoryService := server.NewDirectoryService(args.Logger, e, directoryClient, args.DB)

	api.RegisterDirectoryServer(grpcServer, directoryService)

	if args.TXLogAddress != "" {
		txlogService := server.NewTXLogService(args.Logger, txlogClient, directoryClient)
		api.RegisterTXLogServer(grpcServer, txlogService)
	}

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
		}),
		// Detect HTTP headers with user information and include the data in the gRPC calls.
		// This data is needed for auditlogging
		runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher),
	)

	a := &API{
		logger:          args.Logger,
		environment:     e,
		cert:            args.InternalCert,
		orgCert:         args.OrgCert,
		grpcServer:      grpcServer,
		mux:             mux,
		authenticator:   args.Authenticator,
		directoryClient: directoryClient,
		txlogClient:     txlogClient,
		configDatabase:  args.DB,
	}

	return a, nil
}

func newGRPCServer(logger *zap.Logger, cert *common_tls.CertificateBundle, db database.ConfigDatabase, authenticator Authenticator) *grpc.Server {
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
			authenticator.StreamServerInterceptor(db, getUserFromDatabase),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(recoveryOptions...),
			authenticator.UnaryServerInterceptor(db, getUserFromDatabase),
		),
	}

	return grpc.NewServer(opts...)
}

func getUserFromDatabase(ctx context.Context, configDatabase database.ConfigDatabase, email string) (*domain.User, error) {
	user, err := configDatabase.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Errorf(codes.Internal, "user %q not found in database while authorizing user", email)
		}

		return nil, status.Error(codes.Internal, "could not retrieve user from database while authorizing user")
	}

	return user, nil
}
