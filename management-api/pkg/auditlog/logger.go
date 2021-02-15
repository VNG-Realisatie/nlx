// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package auditlog

import (
	"context"
	"time"
)

type Record struct {
	ID           uint64
	UserID       uint
	ActionType   ActionType
	UserAgent    string
	Organization string
	Service      string
	CreatedAt    time.Time
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
	OrganizationSettingsUpdate             ActionType = "organization_settings_update"
	OrganizationInsightConfigurationUpdate ActionType = "organization_insight_configuration_update"
)

type Logger interface {
	ListAll(ctx context.Context) ([]*Record, error)

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
