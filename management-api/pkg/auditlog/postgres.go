// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package auditlog

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"
	"gopkg.in/square/go-jose.v2/json"

	"go.nlx.io/nlx/management-api/pkg/database"
)

type PostgresLogger struct {
	database database.ConfigDatabase
	logger   *zap.Logger
}

type recordMetadata struct {
	Delegatee *string `json:"delegatee,omitempty"`
	Reference *string `json:"reference,omitempty"`
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

	return convertAuditLogRecordsFromDatabase(auditLogRecords)
}

func convertAuditLogRecordsFromDatabase(records []*database.AuditLog) ([]*Record, error) {
	convertedRecords := make([]*Record, len(records))

	for i, record := range records {
		convertedRecord := &Record{
			ID:         record.ID,
			ActionType: ActionType(record.ActionType),
			Username:   record.UserName,
			UserAgent:  record.UserAgent,
			Delegatee:  record.Delegatee,
			Services:   make([]RecordService, len(record.Services)),
			CreatedAt:  record.CreatedAt,
		}

		if record.Data.Valid {
			data := &recordMetadata{}

			if err := json.Unmarshal([]byte(record.Data.String), data); err != nil {
				return nil, err
			}

			convertedRecord.Data = &RecordData{
				Delegatee: data.Delegatee,
				Reference: data.Reference,
			}
		}

		convertedRecords[i] = convertedRecord

		for j, service := range record.Services {
			convertedRecords[i].Services[j] = RecordService{
				Organization: service.Organization,
				Service:      service.Service,
			}
		}
	}

	return convertedRecords, nil
}

func (a *PostgresLogger) LoginSuccess(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLog{
		UserAgent:  userAgent,
		UserName:   userName,
		ActionType: database.LoginSuccess,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) LoginFail(ctx context.Context, userAgent string) error {
	record := &database.AuditLog{
		UserAgent:  userAgent,
		ActionType: database.LoginFail,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) LogoutSuccess(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLog{
		UserAgent:  userAgent,
		UserName:   userName,
		ActionType: database.LogoutSuccess,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) IncomingAccessRequestAccept(ctx context.Context, userName, userAgent, organization, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: organization,
				Service:      service,
			},
		},
		ActionType: database.IncomingAccesRequestAccept,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) IncomingAccessRequestReject(ctx context.Context, userName, userAgent, organization, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: organization,
				Service:      service,
			},
		},
		ActionType: database.IncomingAccesRequestReject,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) AccessGrantRevoke(ctx context.Context, userName, userAgent, organization, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: organization,
				Service:      service,
			},
		},
		ActionType: database.AccessGrantRevoke,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OutgoingAccessRequestCreate(ctx context.Context, userName, userAgent, organization, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: organization,
				Service:      service,
			},
		},
		ActionType: database.OutgoingAccessRequestCreate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) ServiceCreate(ctx context.Context, userName, userAgent, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Service: service,
			},
		},
		ActionType: database.ServiceCreate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) ServiceUpdate(ctx context.Context, userName, userAgent, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Service: service,
			},
		},
		ActionType: database.ServiceUpdate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) ServiceDelete(ctx context.Context, userName, userAgent, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Service: service,
			},
		},
		ActionType: database.ServiceDelete,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OrderCreate(ctx context.Context, userName, userAgent, delegatee string, services []RecordService) error {
	record := &database.AuditLog{
		UserAgent:  userAgent,
		UserName:   userName,
		Services:   make([]database.AuditLogService, len(services)),
		Delegatee:  delegatee,
		ActionType: database.OrderCreate,
	}

	for i, service := range services {
		record.Services[i] = database.AuditLogService{
			Organization: service.Organization,
			Service:      service.Service,
		}
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OrderOutgoingRevoke(ctx context.Context, userName, userAgent, delegatee, reference string) error {
	revokeLog := &recordMetadata{
		Delegatee: &delegatee,
		Reference: &reference,
	}

	data, err := json.Marshal(revokeLog)
	if err != nil {
		return err
	}

	record := &database.AuditLog{
		UserAgent: userAgent,

		UserName:   userName,
		ActionType: database.OrderOutgoingRevoke,

		Data: sql.NullString{
			Valid:  true,
			String: string(data),
		},
	}

	_, err = a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OrganizationSettingsUpdate(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,

		UserName:   userName,
		ActionType: database.OrganizationSettingsUpdate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}
