// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
)

type ManagementService struct {
	api.UnimplementedDirectoryServer
	api.UnimplementedManagementServer
	external.UnimplementedAccessRequestServiceServer
	external.UnimplementedDelegationServiceServer

	logger                     *zap.Logger
	mainProcess                *process.Process
	directoryClient            directory.Client
	orgCert                    *tls.CertificateBundle
	configDatabase             database.ConfigDatabase
	txlogDatabase              txlogdb.TxlogDatabase
	auditLogger                auditlog.Logger
	createManagementClientFunc func(context.Context, string, *tls.CertificateBundle) (management.Client, error)
}

func NewManagementService(logger *zap.Logger, mainProcess *process.Process, directoryClient directory.Client, orgCert *tls.CertificateBundle, configDatabase database.ConfigDatabase, txlogDatabase txlogdb.TxlogDatabase, auditLogger auditlog.Logger, createManagementClientFunc func(context.Context, string, *tls.CertificateBundle) (management.Client, error)) *ManagementService {
	return &ManagementService{
		logger:                     logger,
		orgCert:                    orgCert,
		mainProcess:                mainProcess,
		directoryClient:            directoryClient,
		configDatabase:             configDatabase,
		txlogDatabase:              txlogDatabase,
		auditLogger:                auditLogger,
		createManagementClientFunc: createManagementClientFunc,
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
