// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm/clause"
)

type AuditLogActionType string

const (
	LoginSuccess                AuditLogActionType = "login_success"
	LoginFail                   AuditLogActionType = "login_fail"
	LogoutSuccess               AuditLogActionType = "logout_success"
	IncomingAccessRequestAccept AuditLogActionType = "incoming_access_request_accept"
	IncomingAccessRequestReject AuditLogActionType = "incoming_access_request_reject"
	AccessGrantRevoke           AuditLogActionType = "access_grant_revoke"
	OutgoingAccessRequestCreate AuditLogActionType = "outgoing_access_request_create"
	ServiceCreate               AuditLogActionType = "service_create"
	ServiceUpdate               AuditLogActionType = "service_update"
	ServiceDelete               AuditLogActionType = "service_delete"
	OrderCreate                 AuditLogActionType = "order_create"
	OrderOutgoingUpdate         AuditLogActionType = "order_outgoing_update"
	OrderOutgoingRevoke         AuditLogActionType = "order_outgoing_revoke"
	OrganizationSettingsUpdate  AuditLogActionType = "organization_settings_update"
	InwayDelete                 AuditLogActionType = "inway_delete"
	OutwayDelete                AuditLogActionType = "outway_delete"
	AcceptTermsOfService        AuditLogActionType = "accept_terms_of_service"
)

type AuditLog struct {
	ID         uint64
	UserName   string
	ActionType AuditLogActionType
	UserAgent  string

	Data      sql.NullString
	Services  []AuditLogService
	CreatedAt time.Time
}

func (a *AuditLog) TableName() string {
	return "nlx_management.audit_logs"
}

type AuditLogServiceOrganization struct {
	SerialNumber string
	Name         string
}

type AuditLogService struct {
	AuditLogID   uint64
	Service      string
	Organization AuditLogServiceOrganization `gorm:"embedded;embeddedPrefix:organization_"`
	CreatedAt    time.Time
}

func (a *AuditLogService) TableName() string {
	return "nlx_management.audit_logs_services"
}

func (db *PostgresConfigDatabase) CreateAuditLogRecord(ctx context.Context, auditLog *AuditLog) (*AuditLog, error) {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{DB: tx}

	if err := dbWithTx.DB.
		WithContext(ctx).
		Omit(clause.Associations).
		Create(auditLog).Error; err != nil {
		return nil, err
	}

	if len(auditLog.Services) > 0 {
		auditLogServices := []AuditLogService{}

		for _, service := range auditLog.Services {
			auditLogServices = append(auditLogServices, AuditLogService{
				AuditLogID:   auditLog.ID,
				Organization: service.Organization,
				Service:      service.Service,
			})
		}

		if err := dbWithTx.DB.
			WithContext(ctx).
			Model(AuditLogService{}).
			Create(auditLogServices).Error; err != nil {
			return nil, err
		}
	}

	result := tx.Commit()
	if result.Error != nil {
		return nil, result.Error
	}

	return auditLog, nil
}

func (db *PostgresConfigDatabase) ListAuditLogRecords(ctx context.Context, limit int) ([]*AuditLog, error) {
	auditLogs := []*AuditLog{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Services").
		Order("created_at desc").
		Limit(limit).
		Find(&auditLogs).Error; err != nil {
		return nil, err
	}

	return auditLogs, nil
}
