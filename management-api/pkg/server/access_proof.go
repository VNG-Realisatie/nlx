package server

import (
	context "context"
	"errors"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) GetAccessProof(ctx context.Context, req *external.GetAccessProofRequest) (*api.AccessProof, error) {
	md, err := s.parseProxyMetadata(ctx)
	if err != nil {
		s.logger.Error("failed to parse proxy metadata", zap.Error(err))

		return nil, err
	}

	grant, err := s.configDatabase.GetLatestAccessGrantForService(ctx, md.OrganizationName, req.ServiceName)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "access proof not found")
		}

		s.logger.Error("failed to get latest proof latest for service", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	createdAt, err := types.TimestampProto(grant.CreatedAt)
	if err != nil {
		s.logger.Error("failed to parse created at time", zap.Error(err))

		return nil, status.Error(codes.Internal, "data error")
	}

	revokedAt, err := types.TimestampProto(grant.RevokedAt)
	if err != nil {
		s.logger.Error("failed to parse revoked at time", zap.Error(err))

		return nil, status.Error(codes.Internal, "data error")
	}

	return &api.AccessProof{
		Id:               grant.ID,
		OrganizationName: grant.OrganizationName,
		ServiceName:      grant.ServiceName,
		CreatedAt:        createdAt,
		RevokedAt:        revokedAt,
	}, nil
}
