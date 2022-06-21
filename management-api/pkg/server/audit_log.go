// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"
	"fmt"
	"regexp"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"xojoc.pw/useragent"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
)

var actionTypes = map[auditlog.ActionType]api.AuditLogRecord_ActionType{
	auditlog.LoginSuccess:                api.AuditLogRecord_loginSuccess,
	auditlog.LoginFail:                   api.AuditLogRecord_loginFail,
	auditlog.LogoutSuccess:               api.AuditLogRecord_logout,
	auditlog.IncomingAccesRequestAccept:  api.AuditLogRecord_incomingAccessRequestAccept,
	auditlog.IncomingAccesRequestReject:  api.AuditLogRecord_incomingAccessRequestReject,
	auditlog.AccessGrantRevoke:           api.AuditLogRecord_accessGrantRevoke,
	auditlog.OutgoingAccessRequestCreate: api.AuditLogRecord_outgoingAccessRequestCreate,
	auditlog.ServiceCreate:               api.AuditLogRecord_serviceCreate,
	auditlog.ServiceUpdate:               api.AuditLogRecord_serviceUpdate,
	auditlog.ServiceDelete:               api.AuditLogRecord_serviceDelete,
	auditlog.OrganizationSettingsUpdate:  api.AuditLogRecord_organizationSettingsUpdate,
	auditlog.OrderCreate:                 api.AuditLogRecord_orderCreate,
	auditlog.OrderOutgoingRevoke:         api.AuditLogRecord_orderOutgoingRevoke,
	auditlog.InwayDelete:                 api.AuditLogRecord_inwayDelete,
	auditlog.OrderOutgoingUpdate:         api.AuditLogRecord_orderOutgoingUpdate,
	auditlog.AcceptTermsOfService:        api.AuditLogRecord_acceptTermsOfService,
}

func (s *ManagementService) ListAuditLogs(ctx context.Context, _ *emptypb.Empty) (*api.ListAuditLogsResponse, error) {
	organizations, err := s.directoryClient.ListOrganizations(ctx, &emptypb.Empty{})
	if err != nil {
		s.logger.Error("failed to retrieve organizations from directory", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to retrieve audit logs")
	}

	oinToOrgNameHash := convertOrganizationsToHash(organizations)

	auditLogs, err := s.auditLogger.ListAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve audit logs")
	}

	responseModels, err := convertAuditLogModelToResponseAuditLog(auditLogs, oinToOrgNameHash)
	if err != nil {
		s.logger.Error("failed to convert audit log records to response models", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to convert audit log records to response models")
	}

	return &api.ListAuditLogsResponse{
		AuditLogs: responseModels,
	}, nil
}

func convertAuditLogModelToResponseAuditLog(records []*auditlog.Record, oinToOrgNameHash map[string]string) ([]*api.AuditLogRecord, error) {
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

		metadata := convertAuditLogMetadataFromDatabaseToModel(record.Data, oinToOrgNameHash)

		convertedRecords[i] = &api.AuditLogRecord{
			Id:              record.ID,
			Action:          actionType,
			User:            record.Username,
			OperatingSystem: operatingSystem,
			Browser:         browser,
			Client:          client,
			CreatedAt:       createdAt,
			Services:        make([]*api.AuditLogRecord_Service, len(record.Services)),
			Data:            metadata,
		}

		for j, service := range record.Services {
			convertedRecords[i].Services[j] = &api.AuditLogRecord_Service{
				Organization: &api.Organization{
					SerialNumber: service.Organization.SerialNumber,
					Name:         oinToOrgNameHash[service.Organization.SerialNumber],
				},
				Service: service.Service,
			}
		}
	}

	return convertedRecords, nil
}

func convertOrganizationsToHash(organizations *directoryapi.ListOrganizationsResponse) map[string]string {
	result := map[string]string{
		"": "",
	}

	for _, organization := range organizations.Organizations {
		result[organization.SerialNumber] = organization.Name
	}

	return result
}

func convertAuditLogActionTypeFromDatabaseToModel(actionType auditlog.ActionType) (api.AuditLogRecord_ActionType, error) {
	value, ok := actionTypes[actionType]
	if !ok {
		return 0, fmt.Errorf("unable to convert audit log action type '%s'", actionType)
	}

	return value, nil
}

func convertAuditLogMetadataFromDatabaseToModel(data *auditlog.RecordData, oinToOrgNameHash map[string]string) *api.AuditLogRecordMetadata {
	var metadata *api.AuditLogRecordMetadata
	if data != nil {
		metadata = &api.AuditLogRecordMetadata{}

		if data.Delegatee != nil {
			metadata.Delegatee = &api.Organization{
				SerialNumber: *data.Delegatee,
				Name:         oinToOrgNameHash[*data.Delegatee],
			}
		}

		if data.Delegator != nil {
			metadata.Delegator = &api.Organization{
				SerialNumber: *data.Delegator,
				Name:         oinToOrgNameHash[*data.Delegator],
			}
		}

		if data.Reference != nil {
			metadata.Reference = *data.Reference
		}

		if data.InwayName != nil {
			metadata.InwayName = *data.InwayName
		}
	}

	return metadata
}
