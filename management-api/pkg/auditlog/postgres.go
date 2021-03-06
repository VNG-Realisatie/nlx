// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package auditlog

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/pkg/database"
)

type PostgresLogger struct {
	database database.ConfigDatabase
	logger   *zap.Logger
}

func NewPostgresLogger(configDatabase database.ConfigDatabase, logger *zap.Logger) Logger {
	return &PostgresLogger{
		database: configDatabase,
		logger:   logger,
	}
}

func (a *PostgresLogger) ListAll(ctx context.Context) ([]*Record, error) {
	auditLogRecords, err := a.database.ListAuditLogRecords(ctx)
	if err != nil {
		a.logger.Error("error retrieving audit log records from database", zap.Error(err))
		return nil, errors.New("database error")
	}

	convertedAuditLogRecords, err := convertAuditLogRecordsFromDatabase(auditLogRecords)
	if err != nil {
		a.logger.Error("error converting audit log records from database", zap.Error(err))
		return nil, errors.New("error converting audit log records")
	}

	return convertedAuditLogRecords, nil
}

func convertAuditLogRecordsFromDatabase(models []*database.AuditLogRecord) ([]*Record, error) {
	convertedRecords := make([]*Record, len(models))

	for i, databaseModel := range models {
		actionType, err := convertDatabaseRecordActionTypeToModel(databaseModel.ActionType)
		if err != nil {
			return nil, err
		}

		convertedRecords[i] = &Record{
			ID:           databaseModel.ID,
			ActionType:   actionType,
			Username:     databaseModel.UserName,
			UserAgent:    databaseModel.UserAgent,
			Organization: databaseModel.Organization,
			Service:      databaseModel.Service,
			CreatedAt:    databaseModel.CreatedAt,
		}
	}

	return convertedRecords, nil
}

//nolint:gocyclo // we need this many auditlog types
func convertDatabaseRecordActionTypeToModel(action database.AuditLogActionType) (ActionType, error) {
	switch action {
	case database.LoginSuccess:
		return LoginSuccess, nil
	case database.LogoutSuccess:
		return LogoutSuccess, nil
	case database.LoginFail:
		return LoginFail, nil
	case database.IncomingAccesRequestAccept:
		return IncomingAccesRequestAccept, nil
	case database.IncomingAccesRequestReject:
		return IncomingAccesRequestReject, nil
	case database.AccessGrantRevoke:
		return AccessGrantRevoke, nil
	case database.OutgoingAccessRequestCreate:
		return OutgoingAccessRequestCreate, nil
	case database.ServiceCreate:
		return ServiceCreate, nil
	case database.ServiceUpdate:
		return ServiceUpdate, nil
	case database.ServiceDelete:
		return ServiceDelete, nil
	case database.OrganizationSettingsUpdate:
		return OrganizationSettingsUpdate, nil
	case database.OrganizationInsightConfigurationUpdate:
		return OrganizationInsightConfigurationUpdate, nil
	default:
		return "", fmt.Errorf("failed to convert action type, unknown action '%s'", action)
	}
}

func (a *PostgresLogger) LoginSuccess(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserName:   userName,
		ActionType: database.LoginSuccess,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) LoginFail(ctx context.Context, userAgent string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		ActionType: database.LoginFail,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) LogoutSuccess(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserName:   userName,
		ActionType: database.LogoutSuccess,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) IncomingAccessRequestAccept(ctx context.Context, userName, userAgent, organization, service string) error {
	record := &database.AuditLogRecord{
		UserAgent:    userAgent,
		UserName:     userName,
		Organization: organization,
		Service:      service,
		ActionType:   database.IncomingAccesRequestAccept,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) IncomingAccessRequestReject(ctx context.Context, userName, userAgent, organization, service string) error {
	record := &database.AuditLogRecord{
		UserAgent:    userAgent,
		UserName:     userName,
		Organization: organization,
		Service:      service,
		ActionType:   database.IncomingAccesRequestReject,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) AccessGrantRevoke(ctx context.Context, userName, userAgent, organization, service string) error {
	record := &database.AuditLogRecord{
		UserAgent:    userAgent,
		UserName:     userName,
		Organization: organization,
		Service:      service,
		ActionType:   database.AccessGrantRevoke,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OutgoingAccessRequestCreate(ctx context.Context, userName, userAgent, organization, service string) error {
	record := &database.AuditLogRecord{
		UserAgent:    userAgent,
		UserName:     userName,
		Organization: organization,
		Service:      service,
		ActionType:   database.OutgoingAccessRequestCreate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) ServiceCreate(ctx context.Context, userName, userAgent, serviceName string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserName:   userName,
		Service:    serviceName,
		ActionType: database.ServiceCreate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) ServiceUpdate(ctx context.Context, userName, userAgent, serviceName string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserName:   userName,
		Service:    serviceName,
		ActionType: database.ServiceUpdate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) ServiceDelete(ctx context.Context, userName, userAgent, serviceName string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserName:   userName,
		Service:    serviceName,
		ActionType: database.ServiceDelete,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OrganizationSettingsUpdate(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLogRecord{
		UserAgent: userAgent,

		UserName:   userName,
		ActionType: database.OrganizationSettingsUpdate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OrganizationInsightConfigurationUpdate(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserName:   userName,
		ActionType: database.OrganizationInsightConfigurationUpdate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}
