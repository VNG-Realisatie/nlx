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
	"google.golang.org/protobuf/types/known/timestamppb"
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

func convertAuditLogModelToResponseAuditLog(records []*auditlog.Record) ([]*api.AuditLogRecord, error) {
	convertedRecords := make([]*api.AuditLogRecord, len(records))

	for i, record := range records {
		actionType, err := convertAuditLogActionTypeFromDatabaseToModel(record.ActionType)
		if err != nil {
			return nil, err
		}

		createdAt := timestamppb.New(record.CreatedAt)
		parsedUserAgent := useragent.Parse(record.UserAgent)

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
			match := re.FindStringSubmatch(record.UserAgent)

			if match != nil {
				operatingSystem = match[1]
			}
		}

		convertedRecords[i] = &api.AuditLogRecord{
			Id:              record.ID,
			Action:          actionType,
			User:            record.Username,
			OperatingSystem: operatingSystem,
			Browser:         browser,
			Client:          client,
			CreatedAt:       createdAt,
			Services:        make([]*api.AuditLogRecord_Service, len(record.Services)),
		}

		for j, service := range record.Services {
			convertedRecords[i].Services[j] = &api.AuditLogRecord_Service{
				Organization: service.Organization,
				Service:      service.Service,
			}
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
	case auditlog.OrderCreate:
		return api.AuditLogRecord_orderCreate, nil
	default:
		return 0, fmt.Errorf("unable to convert audit log action type '%s'", actionType)
	}
}
