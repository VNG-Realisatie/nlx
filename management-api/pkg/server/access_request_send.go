// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

func (s *ManagementService) SendAccessRequest(ctx context.Context, req *api.SendAccessRequestRequest) (*api.SendAccessRequestResponse, error) {
	err := s.authorize(ctx, permissions.SendOutgoingAccessRequest)
	if err != nil {
		return nil, err
	}

	userInfo, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))

		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	fingerprint, err := tls.PemPublicKeyFingerprint([]byte(req.PublicKeyPEM))
	if err != nil {
		s.logger.Error("invalid public key format", zap.Error(err))

		return nil, status.Error(codes.Internal, "invalid public key format")
	}

	address, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, req.OrganizationSerialNumber)
	if err != nil {
		s.logger.Error("unable to retrieve organization inway proxy address", zap.String("organization-serial-number", req.OrganizationSerialNumber), zap.Error(err))

		return nil, status.Error(codes.Internal, "internal")
	}

	s.logger.Info("got organization inway address", zap.String("inway proxy address", address))

	client, err := s.createManagementClientFunc(ctx, address, s.orgCert)
	if err != nil {
		s.logger.Error("unable to setup management client", zap.String("address", address), zap.Error(err))

		return nil, status.Error(codes.Internal, "internal")
	}

	requestAccessResponse, err := client.RequestAccess(ctx, &external.RequestAccessRequest{
		ServiceName:  req.ServiceName,
		PublicKeyPem: req.PublicKeyPEM,
	})
	if err != nil {
		return nil, status.Error(codes.Aborted, "failed to request access, please retry")
	}

	ar := &database.OutgoingAccessRequest{
		Organization: database.Organization{
			SerialNumber: req.OrganizationSerialNumber,
		},
		ReferenceID:          uint(requestAccessResponse.ReferenceId),
		ServiceName:          req.ServiceName,
		PublicKeyPEM:         req.PublicKeyPEM,
		PublicKeyFingerprint: fingerprint,
		State:                database.OutgoingAccessRequestReceived,
	}

	request, err := s.configDatabase.CreateOutgoingAccessRequest(ctx, ar)
	if err != nil {
		if errors.Is(err, database.ErrActiveAccessRequest) {
			return nil, status.Errorf(codes.AlreadyExists, "there is already an active access request")
		}

		s.logger.Error("unable to create outgoing access request", zap.Error(err))

		return nil, status.Error(codes.Internal, "internal")
	}

	err = s.auditLogger.OutgoingAccessRequestCreate(ctx, userInfo.Email, userInfo.UserAgent, req.OrganizationSerialNumber, req.ServiceName)
	if err != nil {
		s.logger.Error("unable to create audit log", zap.Error(err))

		return nil, status.Error(codes.Internal, "internal")
	}

	return &api.SendAccessRequestResponse{
		OutgoingAccessRequest: convertOutgoingAccessRequest(request),
	}, nil
}
