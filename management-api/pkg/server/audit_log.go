// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	context "context"
	"github.com/gogo/protobuf/types"
	"go.nlx.io/nlx/management-api/api"
)

func (s *ManagementService) ListAuditLogs(ctx context.Context, _ *types.Empty) (*api.ListAuditLogsResponse, error) {
	response := &api.ListAuditLogsResponse{
		AuditLogs: []*api.AuditLogRecord{
			{
				Action:    api.AuditLogRecord_login,
				CreatedAt: types.TimestampNow(),
				User:      "Dummy User",
			},
			{
				Action:    api.AuditLogRecord_logout,
				CreatedAt: types.TimestampNow(),
				User:      "Dummy User",
			},
			{
				Action:       api.AuditLogRecord_rejectIncomingAccessRequest,
				CreatedAt:    types.TimestampNow(),
				User:         "Dummy User",
				Organization: "Dummy Organization",
			},
		},
	}

	return response, nil
}
