// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

func (s *ManagementService) SendAccessRequest(ctx context.Context, req *api.SendAccessRequestRequest) (*api.SendAccessRequestResponse, error) {
	err := s.authorize(ctx, permissions.SendOutgoingAccessRequest)
	if err != nil {
		return nil, err
	}

	userInfo, userAgent, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))

		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	fingerprint, err := tls.PemPublicKeyFingerprint([]byte(req.PublicKeyPem))
	if err != nil {
		s.logger.Error("invalid public key format", zap.Error(err))

		return nil, status.Error(codes.Internal, "invalid public key format")
	}

	var (
		state            = database.OutgoingAccessRequestReceived
		errorCause       = ""
		referenceID uint = 0
	)

	requestAccessResponse, err := requestAccess(&requestAccessArgs{
		ctx:                        ctx,
		dc:                         s.directoryClient,
		orgCert:                    s.orgCert,
		createManagementClientFunc: s.createManagementClientFunc,
		organizationSerialNumber:   req.OrganizationSerialNumber,
		serviceName:                req.ServiceName,
		publicKeyPEM:               req.PublicKeyPem,
	})
	if err != nil {
		state = database.OutgoingAccessRequestFailed
		errorCause = "The organization is not available."
	} else {
		referenceID = uint(requestAccessResponse.ReferenceId)
	}

	request, err := createOutgoingAccessRequest(&createOutgoingAccessRequestArgs{
		ctx:                      ctx,
		cd:                       s.configDatabase,
		l:                        s.logger,
		al:                       s.auditLogger,
		userInfo:                 userInfo,
		userAgent:                userAgent,
		organizationSerialNumber: req.OrganizationSerialNumber,
		serviceName:              req.ServiceName,
		publicKeyPEM:             req.PublicKeyPem,
		publicKeyFingerprint:     fingerprint,
		referenceID:              referenceID,
		state:                    state,
		errorCause:               errorCause,
		createdAt:                s.clock.Now(),
		updatedAt:                s.clock.Now(),
	})
	if err != nil {
		s.logger.Error("failed to create outgoing access request", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal")
	}

	return &api.SendAccessRequestResponse{
		OutgoingAccessRequest: convertOutgoingAccessRequest(request),
	}, nil
}

type requestAccessArgs struct {
	ctx                        context.Context
	dc                         directory.Client
	orgCert                    *tls.CertificateBundle
	createManagementClientFunc func(context.Context, string, *tls.CertificateBundle) (management.Client, error)
	organizationSerialNumber   string
	serviceName                string
	publicKeyPEM               string
}

func requestAccess(args *requestAccessArgs) (*external.RequestAccessResponse, error) {
	address, err := args.dc.GetOrganizationInwayProxyAddress(args.ctx, args.organizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve organization inway proxy address: %e", err)
	}

	client, err := args.createManagementClientFunc(args.ctx, address, args.orgCert)
	if err != nil {
		return nil, fmt.Errorf("unable to setup management client: %e", err)
	}

	requestAccessResponse, err := client.RequestAccess(args.ctx, &external.RequestAccessRequest{
		ServiceName:  args.serviceName,
		PublicKeyPem: args.publicKeyPEM,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to request access: %e", err)
	}

	return requestAccessResponse, nil
}

type createOutgoingAccessRequestArgs struct {
	ctx                      context.Context
	cd                       database.ConfigDatabase
	l                        *zap.Logger
	al                       auditlog.Logger
	userInfo                 *domain.User
	userAgent                string
	organizationSerialNumber string
	serviceName              string
	publicKeyPEM             string
	publicKeyFingerprint     string
	referenceID              uint
	errorCause               string
	state                    database.OutgoingAccessRequestState
	createdAt                time.Time
	updatedAt                time.Time
}

func createOutgoingAccessRequest(args *createOutgoingAccessRequestArgs) (*database.OutgoingAccessRequest, error) {
	ar := &database.OutgoingAccessRequest{
		Organization: database.Organization{
			SerialNumber: args.organizationSerialNumber,
		},
		ReferenceID:          args.referenceID,
		ServiceName:          args.serviceName,
		PublicKeyPEM:         args.publicKeyPEM,
		PublicKeyFingerprint: args.publicKeyFingerprint,
		State:                args.state,
		ErrorCause:           args.errorCause,
		CreatedAt:            args.createdAt,
		UpdatedAt:            args.updatedAt,
	}

	request, err := args.cd.CreateOutgoingAccessRequest(args.ctx, ar)
	if err != nil {
		return nil, fmt.Errorf("database error: %e", err)
	}

	err = args.al.OutgoingAccessRequestCreate(args.ctx, args.userInfo.Email, args.userAgent, args.organizationSerialNumber, args.serviceName)
	if err != nil {
		return nil, fmt.Errorf("unable to create audit log: %e", err)
	}

	return request, nil
}
