// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"fmt"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"xojoc.pw/useragent"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
)

func (s *ManagementService) ListAuditLogs(ctx context.Context, _ *emptypb.Empty) (*api.ListAuditLogsResponse, error) {
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

	for i, model := range models {
		actionType, err := convertAuditLogActionTypeFromDatabaseToModel(model.ActionType)
		if err != nil {
			return nil, err
		}

		createdAt, err := types.TimestampProto(model.CreatedAt)
		if err != nil {
			return nil, err
		}

		parsedUserAgent := useragent.Parse(model.UserAgent)

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
			match := re.FindStringSubmatch(model.UserAgent)

			if match != nil {
				operatingSystem = match[1]
			}
		}

		convertedRecords[i] = &api.AuditLogRecord{
			Id:              model.ID,
			Action:          actionType,
			User:            model.Username,
			OperatingSystem: operatingSystem,
			Browser:         browser,
			Client:          client,
			Organization:    model.Organization,
			Service:         model.Service,
			CreatedAt:       createdAt,
		}
	}

	return convertedRecords, nil
}

//nolint:gocyclo // we need this many action types for the auditlog
func convertAuditLogActionTypeFromDatabaseToModel(actionType auditlog.ActionType) (api.AuditLogRecord_ActionType, error) {
	switch actionType {
	case auditlog.LoginSuccess:
		return api.AuditLogRecord_loginSuccess, nil
	case auditlog.LoginFail:
		return api.AuditLogRecord_loginFail, nil
	case auditlog.LogoutSuccess:
		return api.AuditLogRecord_logout, nil
	case auditlog.IncomingAccesRequestAccept:
		return api.AuditLogRecord_incomingAccessRequestAccept, nil
	case auditlog.IncomingAccesRequestReject:
		return api.AuditLogRecord_incomingAccessRequestReject, nil
	case auditlog.AccessGrantRevoke:
		return api.AuditLogRecord_accessGrantRevoke, nil
	case auditlog.OutgoingAccessRequestCreate:
		return api.AuditLogRecord_outgoingAccessRequestCreate, nil
	case auditlog.ServiceCreate:
		return api.AuditLogRecord_serviceCreate, nil
	case auditlog.ServiceUpdate:
		return api.AuditLogRecord_serviceUpdate, nil
	case auditlog.ServiceDelete:
		return api.AuditLogRecord_serviceDelete, nil
	case auditlog.OrganizationSettingsUpdate:
		return api.AuditLogRecord_organizationSettingsUpdate, nil
	case auditlog.OrganizationInsightConfigurationUpdate:
		return api.AuditLogRecord_organizationInsightConfigurationUpdate, nil

	default:
		return 0, fmt.Errorf("unable to convert audit log action type '%s'", actionType)
	}
}
