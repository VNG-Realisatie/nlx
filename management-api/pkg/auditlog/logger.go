// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package auditlog

import (
	"context"
	"time"
)

type Record struct {
	ID           uint64
	Username     string
	ActionType   ActionType
	UserAgent    string
	Organization string
	Delegatee    string
	Services     []RecordService
	CreatedAt    time.Time
}

type RecordService struct {
	Organization string
	Service      string
}

type ActionType string

const (
	LoginSuccess                           ActionType = "login_success"
	LoginFail                              ActionType = "login_fail"
	LogoutSuccess                          ActionType = "logout_success"
	IncomingAccesRequestAccept             ActionType = "incoming_access_request_accept"
	IncomingAccesRequestReject             ActionType = "incoming_access_request_reject"
	AccessGrantRevoke                      ActionType = "access_grant_revoke"
	OutgoingAccessRequestCreate            ActionType = "outgoing_access_request_create"
	ServiceCreate                          ActionType = "service_create"
	ServiceUpdate                          ActionType = "service_update"
	ServiceDelete                          ActionType = "service_delete"
	OrderCreate                            ActionType = "order_create"
	OrganizationSettingsUpdate             ActionType = "organization_settings_update"
	OrganizationInsightConfigurationUpdate ActionType = "organization_insight_configuration_update"
)

type Logger interface {
	ListAll(ctx context.Context) ([]*Record, error)

	LoginSuccess(ctx context.Context, userName, userAgent string) error
	LoginFail(ctx context.Context, userAgent string) error
	LogoutSuccess(ctx context.Context, userName, userAgent string) error
	IncomingAccessRequestAccept(ctx context.Context, userName, userAgent, organization, service string) error
	IncomingAccessRequestReject(ctx context.Context, userName, userAgent, organization, service string) error
	AccessGrantRevoke(ctx context.Context, userName, userAgent, organization, serviceName string) error
	OutgoingAccessRequestCreate(ctx context.Context, userName, userAgent, organization, service string) error
	OrderCreate(ctx context.Context, userName, userAgent, delegatee string, services []RecordService) error
	ServiceCreate(ctx context.Context, userName, userAgent, serviceName string) error
	ServiceUpdate(ctx context.Context, userName, userAgent, serviceName string) error
	ServiceDelete(ctx context.Context, userName, userAgent, serviceName string) error
	OrganizationSettingsUpdate(ctx context.Context, userName, userAgent string) error
	OrganizationInsightConfigurationUpdate(ctx context.Context, userName, userAgent string) error
}
