// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server

import (
	context "context"
	"errors"

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
				zap.String("id", accessGrant.ID),
				zap.String("service", accessGrant.ServiceName),
				zap.Error(err),
			)

			return nil, status.Error(codes.Internal, "error converting access grant")
		}

		response.AccessGrants[i] = responseAccessGrant
	}

	return response, nil
}

func (s *ManagementService) RevokeAccessGrant(ctx context.Context, req *api.RevokeAccessGrantRequest) (*api.AccessGrant, error) {
	accessGrant, err := s.configDatabase.RevokeAccessGrant(ctx, req.ServiceName, req.OrganizationName, req.AccessGrantID)
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
			zap.String("id", accessGrant.ID),
			zap.String("service", accessGrant.ServiceName),
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

	if accessGrant.Revoked() {
		revokedAt, err = types.TimestampProto(accessGrant.RevokedAt)
		if err != nil {
			return nil, err
		}
	}

	return &api.AccessGrant{
		Id:                   accessGrant.ID,
		OrganizationName:     accessGrant.OrganizationName,
		ServiceName:          accessGrant.ServiceName,
		PublicKeyFingerprint: accessGrant.PublicKeyFingerprint,
		CreatedAt:            createdAt,
		RevokedAt:            revokedAt,
	}, nil
}
