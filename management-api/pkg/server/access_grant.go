// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server

import (
	context "context"
	"errors"
	"time"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) ListAccessGrantsForService(ctx context.Context, req *api.ListAccessGrantsForServiceRequest) (*api.ListAccessGrantsForServiceResponse, error) {
	_, err := s.configDatabase.GetService(ctx, req.ServiceName)
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "service not found")
		}

		s.logger.Error("error fetching service", zap.String("name", req.ServiceName), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	accessRequests, err := s.configDatabase.ListAccessGrantsForService(ctx, req.ServiceName)
	if err != nil {
		s.logger.Error("fetching access grants", zap.String("service name", req.ServiceName), zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response := &api.ListAccessGrantsForServiceResponse{
		AccessGrants: make([]*api.AccessGrant, len(accessRequests)),
	}

	for i, accessGrant := range accessRequests {
		responseAccessGrant, err := convertAccessGrant(accessGrant)
		if err != nil {
			s.logger.Error(
				"converting access grant",
				zap.Uint("id", accessGrant.ID),
				zap.String("service", accessGrant.IncomingAccessRequest.Service.Name),
				zap.Error(err),
			)

			return nil, status.Error(codes.Internal, "error converting access grant")
		}

		response.AccessGrants[i] = responseAccessGrant
	}

	return response, nil
}

func (s *ManagementService) RevokeAccessGrant(ctx context.Context, req *api.RevokeAccessGrantRequest) (*api.AccessGrant, error) {
	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.AccessGrantRevoke(ctx, userInfo.username, userInfo.userAgent, req.OrganizationName, req.ServiceName)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	accessGrant, err := s.configDatabase.RevokeAccessGrant(ctx, uint(req.AccessGrantID), time.Now())
	if err != nil {
		if errors.Is(err, database.ErrAccessGrantAlreadyRevoked) {
			s.logger.Warn("access grant is already revoked")
			return nil, status.Error(codes.AlreadyExists, "access grant is already revoked")
		}

		if errIsNotFound(err) {
			s.logger.Warn("access grant not found")
			return nil, status.Error(codes.NotFound, "access grant not found")
		}

		s.logger.Error("revoking access grant", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	apiAccessGrant, err := convertAccessGrant(accessGrant)
	if err != nil {
		s.logger.Error(
			"converting access grant",
			zap.Uint("id", accessGrant.ID),
			zap.String("service", accessGrant.IncomingAccessRequest.Service.Name),
			zap.Error(err),
		)

		return nil, status.Error(codes.Internal, "error converting access grant")
	}

	return apiAccessGrant, nil
}

func convertAccessGrant(accessGrant *database.AccessGrant) (*api.AccessGrant, error) {
	createdAt, err := types.TimestampProto(accessGrant.CreatedAt)
	if err != nil {
		return nil, err
	}

	var revokedAt *types.Timestamp

	if accessGrant.RevokedAt.Valid {
		revokedAt, err = types.TimestampProto(accessGrant.RevokedAt.Time)
		if err != nil {
			return nil, err
		}
	}

	return &api.AccessGrant{
		Id:                   uint64(accessGrant.ID),
		OrganizationName:     accessGrant.IncomingAccessRequest.OrganizationName,
		ServiceName:          accessGrant.IncomingAccessRequest.Service.Name,
		PublicKeyFingerprint: accessGrant.IncomingAccessRequest.PublicKeyFingerprint,
		CreatedAt:            createdAt,
		RevokedAt:            revokedAt,
	}, nil
}
