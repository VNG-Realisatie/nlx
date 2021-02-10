// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package auditlog

import (
	"context"

	"go.nlx.io/nlx/management-api/pkg/database"
)

type AuditLog interface {
	LoginSuccess(ctx context.Context, userID uint, userAgent string) error
	LoginFail(ctx context.Context, userID uint, userAgent string) error
	LogoutSuccess(ctx context.Context, userID uint, userAgent string) error
	IncomingAccessRequestAccept(ctx context.Context, userID uint, userAgent, organization, service string) error
	IncomingAccessRequestReject(ctx context.Context, userID uint, userAgent, organization, service string) error
	AccessGrantRevoke(ctx context.Context, userID uint, userAgent, organization, service string) error
	OutgoingAccessRequestCreate(ctx context.Context, userID uint, userAgent, organization, service string) error
	ServiceCreate(ctx context.Context, userID uint, userAgent string) error
	ServiceUpdate(ctx context.Context, userID uint, userAgent string) error
	ServiceDelete(ctx context.Context, userID uint, userAgent string) error
	OrganizationSettingsUpdate(ctx context.Context, userID uint, userAgent string) error
	OrganizationInsightConfigurationUpdate(ctx context.Context, userID uint, userAgent string) error
}

type PostgresAuditLog struct {
	database database.ConfigDatabase
}

func NewPostgressAuditLog(configDatabase database.ConfigDatabase) AuditLog {
	return &PostgresAuditLog{
		database: configDatabase,
	}
}

func (a *PostgresAuditLog) LoginSuccess(ctx context.Context, userID uint, userAgent string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		ActionType: database.LoginSuccess,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)

	return err
}

func (a *PostgresAuditLog) LoginFail(ctx context.Context, userID uint, userAgent string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		ActionType: database.LoginFail,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}

func (a *PostgresAuditLog) LogoutSuccess(ctx context.Context, userID uint, userAgent string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		ActionType: database.LogoutSuccess,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}

func (a *PostgresAuditLog) IncomingAccessRequestAccept(ctx context.Context, userID uint, userAgent, organization, service string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		Organization: organization,
		Service: service,
		ActionType: database.IncomingAccesRequestAccept,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}

func (a *PostgresAuditLog) IncomingAccessRequestReject(ctx context.Context, userID uint, userAgent, organization, service string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		Organization: organization,
		Service: service,
		ActionType: database.IncomingAccesRequestReject,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}

func (a *PostgresAuditLog) AccessGrantRevoke(ctx context.Context, userID uint, userAgent, organization, service string) error {
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		Organization: organization,
		Service: service,
		ActionType: database.AccessGrantRevoke,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}


func (a *PostgresAuditLog) OutgoingAccessRequestCreate(ctx context.Context, userID uint, userAgent, organization, service string) error{
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		Organization: organization,
		Service: service,
		ActionType: database.OutgoingAccessRequestCreate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}


func (a *PostgresAuditLog) ServiceCreate(ctx context.Context, userID uint, userAgent string) error{
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		ActionType: database.ServiceCreate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}


func (a *PostgresAuditLog) ServiceUpdate(ctx context.Context, userID uint, userAgent string) error{
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		ActionType: database.ServiceUpdate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}


func (a *PostgresAuditLog) ServiceDelete(ctx context.Context, userID uint, userAgent string) error{
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		ActionType: database.ServiceDelete,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}

func (a *PostgresAuditLog) OrganizationSettingsUpdate(ctx context.Context, userID uint, userAgent string) error{
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		ActionType: database.OrganizationSettingsUpdate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}

func (a *PostgresAuditLog) OrganizationInsightConfigurationUpdate(ctx context.Context, userID uint, userAgent string) error{
	record := &database.AuditLogRecord{
		UserAgent:  userAgent,
		UserID:     userID,
		ActionType: database.OrganizationInsightConfigurationUpdate,
	}

	_, err := a.database.CreateAuditLogRecord(ctx, record)
	return err
}

