// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package auditlog

import (
	"context"

	"go.nlx.io/nlx/management-api/api"
)

type Logger interface {
	ListAll(ctx context.Context) ([]*api.AuditLogRecord, error)
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
