// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//nolint:dupl // looks the same as WithDrawOutgoingAccessRequest but writes a different audit-log
package server

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

//nolint:funlen,gocyclo // unable to shorten method
func (s *ManagementService) TerminateAccessProof(ctx context.Context, req *api.TerminateAccessProofRequest) (*api.TerminateAccessProofResponse, error) {
	err := s.authorize(ctx, permissions.TerminateAccess)
	if err != nil {
		return nil, err
	}

	userInfo, userAgent, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	outgoingAccessRequest, err := s.configDatabase.GetLatestOutgoingAccessRequest(ctx, req.OrganizationSerialNumber, req.ServiceName, req.PublicKeyFingerprint)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			auditLogRecordID, errAuditLogTerminate := s.auditLogger.AccessTerminate(ctx, userInfo.Email, userAgent, req.OrganizationSerialNumber, req.ServiceName, req.PublicKeyFingerprint)
			if errAuditLogTerminate != nil {
				return nil, status.Error(codes.Internal, "could not create audit log")
			}

			errAuditLogSucceeded := s.auditLogger.SetAsSucceeded(ctx, auditLogRecordID)
			if errAuditLogSucceeded != nil {
				return nil, status.Error(codes.Internal, "could not update audit log to succeeded")
			}

			return &api.TerminateAccessProofResponse{}, nil
		}

		s.logger.Error("could not retrieve outgoing access request from database", zap.Error(err))

		return nil, status.Errorf(codes.Internal, "internal")
	}

	auditLogRecordID, err := s.auditLogger.AccessTerminate(ctx, userInfo.Email, userAgent, req.OrganizationSerialNumber, req.ServiceName, req.PublicKeyFingerprint)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	address, err := s.directoryClient.GetOrganizationInwayProxyAddress(ctx, req.OrganizationSerialNumber)
	if err != nil {
		s.logger.Error("unable to retrieve organization inway proxy address", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal")
	}

	client, err := s.createManagementClientFunc(ctx, address, s.orgCert)
	if err != nil {
		return nil, fmt.Errorf("unable to setup management client: %e", err)
	}

	terminateResponse, err := client.TerminateAccess(ctx, &external.TerminateAccessRequest{
		PublicKeyFingerprint: req.PublicKeyFingerprint,
		ServiceName:          req.ServiceName,
	})
	if err != nil {
		s.logger.Error("could not terminate external access request", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal")
	}

	accessProof, err := s.configDatabase.GetAccessProofForOutgoingAccessRequest(ctx, outgoingAccessRequest.ID)
	if err != nil {
		s.logger.Error("unable to get access proof from database", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal")
	}

	err = s.configDatabase.TerminateAccessProof(ctx, accessProof.ID, terminateResponse.TerminatedAt.AsTime())
	if err != nil {
		s.logger.Error("could terminate access proof", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal")
	}

	err = s.auditLogger.SetAsSucceeded(ctx, auditLogRecordID)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not update audit log to succeeded")
	}

	return &api.TerminateAccessProofResponse{}, nil
}
