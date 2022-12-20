// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/tabbed/pqtype"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
)

type AuditLogActionType string

const (
	LoginSuccess                  AuditLogActionType = "login_success"
	LoginFail                     AuditLogActionType = "login_fail"
	LogoutSuccess                 AuditLogActionType = "logout_success"
	IncomingAccessRequestAccept   AuditLogActionType = "incoming_access_request_accept"
	IncomingAccessRequestReject   AuditLogActionType = "incoming_access_request_reject"
	AccessGrantRevoke             AuditLogActionType = "access_grant_revoke"
	OutgoingAccessRequestCreate   AuditLogActionType = "outgoing_access_request_create"
	OutgoingAccessRequestWithdraw AuditLogActionType = "outgoing_access_request_withdraw"
	AccessGrantTerminate          AuditLogActionType = "access_terminate"
	ServiceCreate                 AuditLogActionType = "service_create"
	ServiceUpdate                 AuditLogActionType = "service_update"
	ServiceDelete                 AuditLogActionType = "service_delete"
	OrderCreate                   AuditLogActionType = "order_create"
	OrderOutgoingUpdate           AuditLogActionType = "order_outgoing_update"
	OrderOutgoingRevoke           AuditLogActionType = "order_outgoing_revoke"
	OrganizationSettingsUpdate    AuditLogActionType = "organization_settings_update"
	InwayDelete                   AuditLogActionType = "inway_delete"
	OutwayDelete                  AuditLogActionType = "outway_delete"
	AcceptTermsOfService          AuditLogActionType = "accept_terms_of_service"
)

type AuditLog struct {
	ID         uint64
	UserName   string
	ActionType AuditLogActionType
	UserAgent  string

	Data         sql.NullString
	Services     []AuditLogService
	HasSucceeded bool
	CreatedAt    time.Time
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

func (db *PostgresConfigDatabase) CreateAuditLogRecord(ctx context.Context, auditLog *AuditLog) (uint64, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return
			}

			log.Println(fmt.Printf("cannot rollback database transaction for creating an audit log: %e", err))
		}
	}()

	qtx := db.queries.WithTx(tx)

	userName := sql.NullString{}

	if auditLog.UserName != "" {
		userName.String = auditLog.UserName
		userName.Valid = true
	}

	data := pqtype.NullRawMessage{}

	if auditLog.Data.Valid {
		data.RawMessage = json.RawMessage(auditLog.Data.String)
		data.Valid = true
	}

	id, err := qtx.CreateAuditLog(ctx, &queries.CreateAuditLogParams{
		UserName:     userName,
		ActionType:   string(auditLog.ActionType),
		UserAgent:    auditLog.UserAgent,
		Data:         data,
		CreatedAt:    auditLog.CreatedAt,
		HasSucceeded: auditLog.HasSucceeded,
	})
	if err != nil {
		return 0, err
	}

	for _, service := range auditLog.Services {
		errService := qtx.CreateAuditLogService(ctx, &queries.CreateAuditLogServiceParams{
			AuditLogID: sql.NullInt64{
				Valid: true,
				Int64: id,
			},
			OrganizationName: sql.NullString{
				Valid:  true,
				String: service.Organization.Name,
			},
			OrganizationSerialNumber: service.Organization.SerialNumber,
			CreatedAt:                service.CreatedAt,
			Service: sql.NullString{
				Valid:  true,
				String: service.Service,
			},
		})
		if errService != nil {
			return 0, errService
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (db *PostgresConfigDatabase) ListAuditLogRecords(ctx context.Context, limit int) ([]*AuditLog, error) {
	rows, err := db.queries.ListAuditLogs(ctx, int32(limit))
	if err != nil {
		return nil, err
	}

	auditLogs := make([]*AuditLog, len(rows))

	for i, row := range rows {
		serviceRows, servicesErr := db.queries.ListAuditLogServices(ctx, sql.NullInt64{
			Valid: true,
			Int64: row.ID,
		})
		if servicesErr != nil {
			return nil, servicesErr
		}

		services := make([]AuditLogService, len(serviceRows))

		for j, service := range serviceRows {
			var serviceName string

			if service.Service.Valid {
				serviceName = service.Service.String
			}

			var organizationName string

			if service.OrganizationName.Valid {
				organizationName = service.OrganizationName.String
			}

			services[j] = AuditLogService{
				AuditLogID: uint64(row.ID),
				Service:    serviceName,
				Organization: AuditLogServiceOrganization{
					Name:         organizationName,
					SerialNumber: service.OrganizationSerialNumber,
				},
				CreatedAt: service.CreatedAt,
			}
		}

		data := sql.NullString{}

		if row.Data.Valid {
			marshaled, marshallErr := row.Data.RawMessage.MarshalJSON()
			if marshallErr != nil {
				return nil, marshallErr
			}

			data = sql.NullString{
				String: string(marshaled),
				Valid:  true,
			}
		}

		var userName string

		if row.UserName.Valid {
			userName = row.UserName.String
		}

		auditLogs[i] = &AuditLog{
			ID:           uint64(row.ID),
			UserName:     userName,
			ActionType:   AuditLogActionType(row.ActionType),
			UserAgent:    row.UserAgent,
			Data:         data,
			Services:     services,
			HasSucceeded: row.HasSucceeded,
			CreatedAt:    row.CreatedAt,
		}
	}

	return auditLogs, nil
}

func (db *PostgresConfigDatabase) SetAuditLogAsSucceeded(ctx context.Context, id int64) error {
	return db.queries.SetAuditLogAsSucceeded(ctx, id)
}
