// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/outway"
	"go.nlx.io/nlx/management-api/pkg/txlog"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
)

type ManagementService struct {
	api.UnimplementedDirectoryServer
	api.UnimplementedManagementServer
	api.UnimplementedTXLogServer
	external.UnimplementedAccessRequestServiceServer
	external.UnimplementedDelegationServiceServer

	logger                     *zap.Logger
	directoryClient            directory.Client
	txlogClient                txlog.Client
	orgCert                    *tls.CertificateBundle
	internalCert               *tls.CertificateBundle
	configDatabase             database.ConfigDatabase
	txlogDatabase              txlogdb.TxlogDatabase
	auditLogger                auditlog.Logger
	createManagementClientFunc func(context.Context, string, *tls.CertificateBundle) (management.Client, error)
	createOutwayClientFunc     func(context.Context, string, *tls.CertificateBundle) (outway.Client, error)
}

func NewManagementService(logger *zap.Logger, directoryClient directory.Client, txlogClient txlog.Client, orgCert, internalCert *tls.CertificateBundle, configDatabase database.ConfigDatabase, txlogDatabase txlogdb.TxlogDatabase, auditLogger auditlog.Logger, createManagementClientFunc func(context.Context, string, *tls.CertificateBundle) (management.Client, error), createOutwayClientFunc func(context.Context, string, *tls.CertificateBundle) (outway.Client, error)) *ManagementService {
	return &ManagementService{
		logger:                     logger,
		orgCert:                    orgCert,
		internalCert:               internalCert,
		directoryClient:            directoryClient,
		txlogClient:                txlogClient,
		configDatabase:             configDatabase,
		txlogDatabase:              txlogDatabase,
		auditLogger:                auditLogger,
		createManagementClientFunc: createManagementClientFunc,
		createOutwayClientFunc:     createOutwayClientFunc,
	}
}

func errIsNotFound(err error) bool {
	return errors.Is(err, database.ErrNotFound)
}

type auditLogInfoFromGRPC struct {
	username  string
	userEmail string
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

	userEmail := ""
	if len(md.Get("user-email")) > 0 {
		userEmail = md.Get("user-email")[0]
	}

	userAgent := ""
	if len(md.Get("grpcgateway-user-agent")) > 0 {
		userAgent = md.Get("grpcgateway-user-agent")[0]
	}

	return &auditLogInfoFromGRPC{
		username:  username,
		userEmail: userEmail,
		userAgent: userAgent,
	}, nil
}

func retrieveUserFromContext(ctx context.Context) (*domain.User, error) {
	user := ctx.Value(domain.UserKey)
	if user == nil {
		return nil, fmt.Errorf("no user in context")
	}

	convertedUser, ok := user.(*domain.User)
	if !ok {
		return nil, fmt.Errorf("user value in context is not a valid api user")
	}

	return convertedUser, nil
}
