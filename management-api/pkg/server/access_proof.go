// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"

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

	createdAt, err := ptypes.TimestampProto(grant.CreatedAt)
	if err != nil {
		s.logger.Error("failed to parse created at time", zap.Error(err))

		return nil, status.Error(codes.Internal, "data error")
	}

	revokedAt, err := ptypes.TimestampProto(grant.RevokedAt.Time)
	if err != nil {
		s.logger.Error("failed to parse revoked at time", zap.Error(err))

		return nil, status.Error(codes.Internal, "data error")
	}

	return &api.AccessProof{
		Id:               uint64(grant.ID),
		AccessRequestId:  uint64(grant.IncomingAccessRequest.ID),
		OrganizationName: grant.IncomingAccessRequest.OrganizationName,
		ServiceName:      grant.IncomingAccessRequest.Service.Name,
		CreatedAt:        createdAt,
		RevokedAt:        revokedAt,
	}, nil
}
