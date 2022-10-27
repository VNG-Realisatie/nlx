// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package auditlog

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/pkg/database"
)

type PostgresLogger struct {
	database database.ConfigDatabase
	logger   *zap.Logger
}

type recordMetadata struct {
	Delegatee            *string `json:"delegatee,omitempty"`
	Reference            *string `json:"reference,omitempty"`
	InwayName            *string `json:"inwayName,omitempty"`
	OutwayName           *string `json:"outwayName,omitempty"`
	PublicKeyFingerprint *string `json:"publicKeyFingerprint,omitempty"`
}

func NewPostgresLogger(configDatabase database.ConfigDatabase, logger *zap.Logger) Logger {
	return &PostgresLogger{
		database: configDatabase,
		logger:   logger,
	}
}

func (a *PostgresLogger) List(ctx context.Context, limit int) ([]*Record, error) {
	auditLogRecords, err := a.database.ListAuditLogRecords(ctx, limit)
	if err != nil {
		a.logger.Error("error retrieving audit log records from database", zap.Error(err))
		return nil, errors.New("database error")
	}

	return convertAuditLogRecordsFromDatabase(auditLogRecords)
}

func (a *PostgresLogger) SetAsSucceeded(ctx context.Context, id int64) error {
	return a.database.SetAuditLogAsSucceeded(ctx, id)
}

func convertAuditLogRecordsFromDatabase(records []*database.AuditLog) ([]*Record, error) {
	convertedRecords := make([]*Record, len(records))

	for i, record := range records {
		convertedRecord := &Record{
			ID:           record.ID,
			ActionType:   ActionType(record.ActionType),
			Username:     record.UserName,
			UserAgent:    record.UserAgent,
			Services:     make([]RecordService, len(record.Services)),
			CreatedAt:    record.CreatedAt,
			HasSucceeded: record.HasSucceeded,
		}

		if record.Data.Valid {
			data := &recordMetadata{}

			if err := json.Unmarshal([]byte(record.Data.String), data); err != nil {
				return nil, err
			}

			convertedRecord.Data = &RecordData{
				Delegatee:            data.Delegatee,
				Reference:            data.Reference,
				InwayName:            data.InwayName,
				OutwayName:           data.OutwayName,
				PublicKeyFingerprint: data.PublicKeyFingerprint,
			}
		}

		convertedRecords[i] = convertedRecord

		for j, service := range record.Services {
			convertedRecords[i].Services[j] = RecordService{
				Organization: RecordServiceOrganization{
					SerialNumber: service.Organization.SerialNumber,
					Name:         service.Organization.Name,
				},
				Service: service.Service,
			}
		}
	}

	return convertedRecords, nil
}

func (a *PostgresLogger) LoginSuccess(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLog{
		UserAgent:    userAgent,
		UserName:     userName,
		ActionType:   database.LoginSuccess,
		HasSucceeded: true,
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
		UserAgent:    userAgent,
		UserName:     userName,
		ActionType:   database.LogoutSuccess,
		HasSucceeded: true,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) IncomingAccessRequestAccept(ctx context.Context, userName, userAgent, organizationSerialNumber, organizationName, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: database.AuditLogServiceOrganization{
					SerialNumber: organizationSerialNumber,
					Name:         organizationName,
				},
				Service: service,
			},
		},
		ActionType:   database.IncomingAccessRequestAccept,
		HasSucceeded: true,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) IncomingAccessRequestReject(ctx context.Context, userName, userAgent, organizationSerialNumber, organizationName, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: database.AuditLogServiceOrganization{
					SerialNumber: organizationSerialNumber,
					Name:         organizationName,
				},
				Service: service,
			},
		},
		ActionType:   database.IncomingAccessRequestReject,
		HasSucceeded: true,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) AccessGrantRevoke(ctx context.Context, userName, userAgent, organizationSerialNumber, organizationName, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: database.AuditLogServiceOrganization{
					SerialNumber: organizationSerialNumber,
					Name:         organizationName,
				},
				Service: service,
			},
		},
		ActionType:   database.AccessGrantRevoke,
		HasSucceeded: true,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OutgoingAccessRequestCreate(ctx context.Context, userName, userAgent, organizationSerialNumber, service string) error {
	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: database.AuditLogServiceOrganization{
					SerialNumber: organizationSerialNumber,
				},
				Service: service,
			},
		},
		ActionType:   database.OutgoingAccessRequestCreate,
		HasSucceeded: true,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

//nolint:dupl // looks the same as OutgoingAccessRequestTerminate but has a different actionType
func (a *PostgresLogger) OutgoingAccessRequestWithdraw(ctx context.Context, userName, userAgent, organizationSerialNumber, service, publicKeyFingerprint string) (int64, error) {
	auditLogRecordID, err := a.outgoingAccessRequestAuditLog(ctx, database.OutgoingAccessRequestWithdraw, userName, userAgent, organizationSerialNumber, service, publicKeyFingerprint)
	if err != nil {
		return 0, err
	}

	return int64(auditLogRecordID), err
}

//nolint:dupl // looks the same as OutgoingAccessRequestWithdraw but has a different actionType
func (a *PostgresLogger) AccessTerminate(ctx context.Context, userName, userAgent, organizationSerialNumber, service, publicKeyFingerprint string) (int64, error) {
	auditLogRecordID, err := a.outgoingAccessRequestAuditLog(ctx, database.AccessGrantTerminate, userName, userAgent, organizationSerialNumber, service, publicKeyFingerprint)
	if err != nil {
		return 0, err
	}

	return int64(auditLogRecordID), err
}

func (a *PostgresLogger) outgoingAccessRequestAuditLog(ctx context.Context, actionType database.AuditLogActionType, userName, userAgent, organizationSerialNumber, service, publicKeyFingerprint string) (uint64, error) {
	metaData := &recordMetadata{
		PublicKeyFingerprint: &publicKeyFingerprint,
	}

	data, err := json.Marshal(metaData)
	if err != nil {
		return 0, err
	}

	record := &database.AuditLog{
		UserAgent: userAgent,
		UserName:  userName,
		Services: []database.AuditLogService{
			{
				Organization: database.AuditLogServiceOrganization{
					SerialNumber: organizationSerialNumber,
				},
				Service: service,
			},
		},
		Data:       sql.NullString{String: string(data), Valid: true},
		ActionType: actionType,
	}

	record, err = a.database.CreateAuditLogRecord(ctx, record)

	return record.ID, err
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
		ActionType:   database.ServiceCreate,
		HasSucceeded: true,
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
		ActionType:   database.ServiceUpdate,
		HasSucceeded: true,
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
		ActionType:   database.ServiceDelete,
		HasSucceeded: true,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OrderCreate(ctx context.Context, userName, userAgent, delegatee string, services []RecordService) error {
	createLog := &recordMetadata{
		Delegatee: &delegatee,
	}

	data, err := json.Marshal(createLog)
	if err != nil {
		return err
	}

	record := &database.AuditLog{
		UserAgent:    userAgent,
		UserName:     userName,
		Services:     make([]database.AuditLogService, len(services)),
		Data:         sql.NullString{String: string(data), Valid: true},
		ActionType:   database.OrderCreate,
		HasSucceeded: true,
	}

	for i, service := range services {
		record.Services[i] = database.AuditLogService{
			Organization: database.AuditLogServiceOrganization{
				SerialNumber: service.Organization.SerialNumber,
				Name:         service.Organization.Name,
			},
			Service: service.Service,
		}
	}

	_, err = a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OrderOutgoingUpdate(ctx context.Context, userName, userAgent, delegatee, orderReference string, services []RecordService) error {
	updateLog := &recordMetadata{
		Delegatee: &delegatee,
		Reference: &orderReference,
	}

	data, err := json.Marshal(updateLog)
	if err != nil {
		return err
	}

	record := &database.AuditLog{
		UserAgent:    userAgent,
		UserName:     userName,
		Services:     make([]database.AuditLogService, len(services)),
		Data:         sql.NullString{String: string(data), Valid: true},
		ActionType:   database.OrderOutgoingUpdate,
		HasSucceeded: true,
	}

	for i, service := range services {
		record.Services[i] = database.AuditLogService{
			Organization: database.AuditLogServiceOrganization{
				SerialNumber: service.Organization.SerialNumber,
				Name:         service.Organization.Name,
			},
			Service: service.Service,
		}
	}

	_, err = a.database.CreateAuditLogRecord(ctx, record)

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
		UserAgent:  userAgent,
		UserName:   userName,
		ActionType: database.OrderOutgoingRevoke,
		Data: sql.NullString{
			Valid:  true,
			String: string(data),
		},
		HasSucceeded: true,
	}

	_, err = a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OrganizationSettingsUpdate(ctx context.Context, userName, userAgent string) error {
	record := &database.AuditLog{
		UserAgent:    userAgent,
		UserName:     userName,
		ActionType:   database.OrganizationSettingsUpdate,
		HasSucceeded: true,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) InwayDelete(ctx context.Context, userName, userAgent, inwayName string) error {
	metaData := &recordMetadata{
		InwayName: &inwayName,
	}

	data, err := json.Marshal(metaData)
	if err != nil {
		return err
	}

	record := &database.AuditLog{
		UserAgent:  userAgent,
		UserName:   userName,
		ActionType: database.InwayDelete,
		Data: sql.NullString{
			Valid:  true,
			String: string(data),
		},
		HasSucceeded: true,
	}

	_, err = a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) OutwayDelete(ctx context.Context, userName, userAgent, outwayName string) error {
	metaData := &recordMetadata{
		OutwayName: &outwayName,
	}

	data, err := json.Marshal(metaData)
	if err != nil {
		return err
	}

	record := &database.AuditLog{
		UserAgent:  userAgent,
		UserName:   userName,
		ActionType: database.OutwayDelete,
		Data: sql.NullString{
			Valid:  true,
			String: string(data),
		},
		HasSucceeded: true,
	}

	_, err = a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresLogger) AcceptTermsOfService(ctx context.Context, userName, userAgent string) error {
	_, err := a.database.CreateAuditLogRecord(ctx, &database.AuditLog{
		UserAgent:    userAgent,
		UserName:     userName,
		ActionType:   database.AcceptTermsOfService,
		HasSucceeded: true,
	})

	return err
}
