// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) ListAuditLogs(ctx context.Context, _ *types.Empty) (*api.ListAuditLogsResponse, error) {
	auditLogRecords, err := s.configDatabase.ListAuditLogRecords(ctx)
	if err != nil {
		s.logger.Error("error retrieving audit log records from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	convertedAuditLogRecords, err := convertFromDatabaseAuditLogRecords(auditLogRecords)
	if err != nil {
		s.logger.Error("error converting audit log records from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "error converting audit log records")
	}

	return &api.ListAuditLogsResponse{
		AuditLogs: convertedAuditLogRecords,
	}, nil
}

func convertFromDatabaseAuditLogRecords(models []*database.AuditLogRecord) ([]*api.AuditLogRecord, error) {
	convertedRecords := make([]*api.AuditLogRecord, len(models))

	for i, databaseModel := range models {
		createdAt, err := types.TimestampProto(databaseModel.CreatedAt)
		if err != nil {
			return nil, err
		}

		actionType, err := auditLogActionTypeToProto(databaseModel.ActionType)
		if err != nil {
			return nil, err
		}

		convertedRecords[i] = &api.AuditLogRecord{
			Id:           databaseModel.ID,
			Action:       actionType,
			User:         fmt.Sprintf("%d", databaseModel.UserID),
			UserAgent:    databaseModel.UserAgent,
			Organization: databaseModel.Organization,
			Service:      databaseModel.Service,
			CreatedAt:    createdAt,
		}
	}
	return convertedRecords, nil
}

func auditLogActionTypeToProto(actionType database.AuditLogActionType) (api.AuditLogRecord_ActionType, error) {
    switch(actionType) {
	case database.LoginSuccess:
		return api.AuditLogRecord_loginSuccess, nil 
    default:		
		return 0, fmt.Errorf("unable to convert database audit log action type '%s' to api audit log action type", actionType)
	}

}
