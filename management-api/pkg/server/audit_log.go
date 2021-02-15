// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
)

func (s *ManagementService) ListAuditLogs(ctx context.Context, _ *types.Empty) (*api.ListAuditLogsResponse, error) {
	auditLogs, err := s.auditLogger.ListAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve audit logs")
	}

	responseModels, err := convertAuditLogModelToResponseAuditLog(auditLogs)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to convert audit log records to response models")
	}

	return &api.ListAuditLogsResponse{
		AuditLogs: responseModels,
	}, nil
}

func convertAuditLogModelToResponseAuditLog(models []*auditlog.Record) ([]*api.AuditLogRecord, error) {
	convertedRecords := make([]*api.AuditLogRecord, len(models))

	for i, databaseModel := range models {
		actionType, err := convertAuditLogActionTypeFromDatabaseToModel(databaseModel.ActionType)
		if err != nil {
			return nil, err
		}

		createdAt, err := types.TimestampProto(databaseModel.CreatedAt)
		if err != nil {
			return nil, err
		}

		convertedRecords[i] = &api.AuditLogRecord{
			Id:              databaseModel.ID,
			Action:          actionType,
			User:            fmt.Sprintf("%d", databaseModel.UserID),
			OperatingSystem: "",
			Browser:         "",
			Client:          "",
			Organization:    databaseModel.Organization,
			Service:         databaseModel.Service,
			CreatedAt:       createdAt,
		}
	}

	return convertedRecords, nil
}

func convertAuditLogActionTypeFromDatabaseToModel(actionType auditlog.ActionType) (api.AuditLogRecord_ActionType, error) {
	switch actionType {
	case auditlog.LoginSuccess:
		return api.AuditLogRecord_loginSuccess, nil
	case auditlog.LogoutSuccess:
		return api.AuditLogRecord_logout, nil
	default:
		return 0, fmt.Errorf("unable to convert audit log action type '%s'", actionType)
	}
}
