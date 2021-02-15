// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"fmt"
	"regexp"

	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"xojoc.pw/useragent"

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

		parsedUserAgent := useragent.Parse(databaseModel.UserAgent)

		operatingSystem := ""
		browser := ""
		client := "nlxctl"

		if parsedUserAgent != nil {
			operatingSystem = parsedUserAgent.OS

			if parsedUserAgent.Type == useragent.Browser {
				browser = parsedUserAgent.Name
				client = "NLX Management"
			}
		} else {
			re := regexp.MustCompile(`.*\(([a-zA-Z ]*)\)$`)
			match := re.FindStringSubmatch(databaseModel.UserAgent)

			if match != nil {
				operatingSystem = match[1]
			}
		}

		convertedRecords[i] = &api.AuditLogRecord{
			Id:              databaseModel.ID,
			Action:          actionType,
			User:            fmt.Sprintf("%d", databaseModel.UserID),
			OperatingSystem: operatingSystem,
			Browser:         browser,
			Client:          client,
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
