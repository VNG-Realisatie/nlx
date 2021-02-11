// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
)

func (s *ManagementService) ListAuditLogs(ctx context.Context, _ *types.Empty) (*api.ListAuditLogsResponse, error) {
	auditLogs, err := s.auditLogger.ListAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve audit logs error")
	}

	return &api.ListAuditLogsResponse{
		AuditLogs: auditLogs,
	}, nil
}
