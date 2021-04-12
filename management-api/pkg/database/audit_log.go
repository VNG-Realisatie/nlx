// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"time"

	"gorm.io/gorm/clause"
)

type AuditLogActionType string

const (
	LoginSuccess                           AuditLogActionType = "login_success"
	LoginFail                              AuditLogActionType = "login_fail"
	LogoutSuccess                          AuditLogActionType = "logout_success"
	IncomingAccesRequestAccept             AuditLogActionType = "incoming_access_request_accept"
	IncomingAccesRequestReject             AuditLogActionType = "incoming_access_request_reject"
	AccessGrantRevoke                      AuditLogActionType = "access_grant_revoke"
	OutgoingAccessRequestCreate            AuditLogActionType = "outgoing_access_request_create"
	ServiceCreate                          AuditLogActionType = "service_create"
	ServiceUpdate                          AuditLogActionType = "service_update"
	ServiceDelete                          AuditLogActionType = "service_delete"
	OrderCreate                            AuditLogActionType = "order_create"
	OrganizationSettingsUpdate             AuditLogActionType = "organization_settings_update"
	OrganizationInsightConfigurationUpdate AuditLogActionType = "organization_insight_configuration_update"
)

type AuditLog struct {
	ID         uint64 `gorm:"primarykey;column:audit_log_id;"`
	UserName   string
	ActionType AuditLogActionType
	UserAgent  string
	Delegatee  string
	Services   []AuditLogService
	CreatedAt  time.Time
}

func (a *AuditLog) TableName() string {
	return "nlx_management.audit_logs"
}

type AuditLogService struct {
	AuditLogID   uint64
	Service      string
	Organization string
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

	result := tx.Commit()
	if result.Error != nil {
		return nil, result.Error
	}

	return auditLog, nil
}

func (db *PostgresConfigDatabase) ListAuditLogRecords(ctx context.Context) ([]*AuditLog, error) {
	auditLogs := []*AuditLog{}

	if err := db.DB.
		WithContext(ctx).
		Preload("Services").
		Order("created_at desc").
		Find(&auditLogs).Error; err != nil {
		return nil, err
	}

	return auditLogs, nil
}
