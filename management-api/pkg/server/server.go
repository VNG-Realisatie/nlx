// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"

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

type Clock interface {
	Now() time.Time
}

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
	clock                      Clock
	createManagementClientFunc func(context.Context, string, *tls.CertificateBundle) (management.Client, error)
	createOutwayClientFunc     func(context.Context, string, *tls.CertificateBundle) (outway.Client, error)
}

type NewManagementServiceArgs struct {
	Logger                     *zap.Logger
	DirectoryClient            directory.Client
	TxlogClient                txlog.Client
	OrgCert                    *tls.CertificateBundle
	InternalCert               *tls.CertificateBundle
	ConfigDatabase             database.ConfigDatabase
	TxlogDatabase              txlogdb.TxlogDatabase
	AuditLogger                auditlog.Logger
	CreateManagementClientFunc func(context.Context, string, *tls.CertificateBundle) (management.Client, error)
	CreateOutwayClientFunc     func(context.Context, string, *tls.CertificateBundle) (outway.Client, error)
	Clock                      Clock
}

func NewManagementService(args *NewManagementServiceArgs) *ManagementService {
	return &ManagementService{
		logger:                     args.Logger,
		orgCert:                    args.OrgCert,
		internalCert:               args.InternalCert,
		directoryClient:            args.DirectoryClient,
		txlogClient:                args.TxlogClient,
		configDatabase:             args.ConfigDatabase,
		txlogDatabase:              args.TxlogDatabase,
		auditLogger:                args.AuditLogger,
		createManagementClientFunc: args.CreateManagementClientFunc,
		createOutwayClientFunc:     args.CreateOutwayClientFunc,
		clock:                      args.Clock,
	}
}

func errIsNotFound(err error) bool {
	return errors.Is(err, database.ErrNotFound)
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
