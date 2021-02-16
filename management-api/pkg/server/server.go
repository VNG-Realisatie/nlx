package server

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
)

// ManagementService handles all requests for the config api
type ManagementService struct {
	logger          *zap.Logger
	configDatabase  database.ConfigDatabase
	mainProcess     *process.Process
	directoryClient directory.Client
	orgCert         *tls.CertificateBundle
	auditLogger     auditlog.Logger
}

// New creates new ManagementService
func NewManagementService(logger *zap.Logger, mainProcess *process.Process, directoryClient directory.Client, orgCert *tls.CertificateBundle, configDatabase database.ConfigDatabase, auditLogger auditlog.Logger) *ManagementService {
	return &ManagementService{
		configDatabase:  configDatabase,
		logger:          logger,
		orgCert:         orgCert,
		mainProcess:     mainProcess,
		directoryClient: directoryClient,
		auditLogger:     auditLogger,
	}
}

func errIsNotFound(err error) bool {
	return errors.Is(err, database.ErrNotFound)
}

type auditLogInfoFromGRPC struct {
	username  string
	userAgent string
}

func retrieveUserInfoFromGRPCContext(ctx context.Context) (*auditLogInfoFromGRPC, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not extract metadata from context")
	}

	username := ""
	if len(md.Get("username")) > 0 {
		username = md.Get("username")[0]
	}

	userAgent := ""
	if len(md.Get("grpcgateway-user-agent")) > 0 {
		userAgent = md.Get("grpcgateway-user-agent")[0]
	}

	return &auditLogInfoFromGRPC{
		username:  username,
		userAgent: userAgent,
	}, nil
}
