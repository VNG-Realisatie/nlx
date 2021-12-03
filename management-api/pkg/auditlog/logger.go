// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package auditlog

import (
	"context"
	"time"
)

type Record struct {
	ID         uint64
	Username   string
	ActionType ActionType
	UserAgent  string
	Delegatee  string
	Services   []RecordService
	CreatedAt  time.Time
	Data       *RecordData
}

type RecordData struct {
	Delegator *string
	Delegatee *string
	Reference *string

	InwayName *string
}

type RecordServiceOrganization struct {
	SerialNumber string
	Name         string
}

type RecordService struct {
	Organization RecordServiceOrganization
	Service      string
}

type ActionType string

const (
	LoginSuccess                ActionType = "login_success"
	LoginFail                   ActionType = "login_fail"
	LogoutSuccess               ActionType = "logout_success"
	IncomingAccesRequestAccept  ActionType = "incoming_access_request_accept"
	IncomingAccesRequestReject  ActionType = "incoming_access_request_reject"
	AccessGrantRevoke           ActionType = "access_grant_revoke"
	OutgoingAccessRequestCreate ActionType = "outgoing_access_request_create"
	ServiceCreate               ActionType = "service_create"
	ServiceUpdate               ActionType = "service_update"
	ServiceDelete               ActionType = "service_delete"
	OrderCreate                 ActionType = "order_create"
	OrderOutgoingUpdate         ActionType = "order_outgoing_update"
	OrderOutgoingRevoke         ActionType = "order_outgoing_revoke"
	OrganizationSettingsUpdate  ActionType = "organization_settings_update"
	InwayDelete                 ActionType = "inway_delete"
)

type Logger interface {
	ListAll(ctx context.Context) ([]*Record, error)

	LoginSuccess(ctx context.Context, userName, userAgent string) error
	LoginFail(ctx context.Context, userAgent string) error
	LogoutSuccess(ctx context.Context, userName, userAgent string) error
	IncomingAccessRequestAccept(ctx context.Context, userName, userAgent, organizationSerialNumber, organizationName, service string) error
	IncomingAccessRequestReject(ctx context.Context, userName, userAgent, organizationSerialNumber, organizationName, service string) error
	AccessGrantRevoke(ctx context.Context, userName, userAgent, organizationSerialNumber, organizationName, serviceName string) error
	OutgoingAccessRequestCreate(ctx context.Context, userName, userAgent, organizationSerialNumber, service string) error
	OrderCreate(ctx context.Context, userName, userAgent, delegatee string, services []RecordService) error
	OrderOutgoingUpdate(ctx context.Context, userName, userAgent, delegatee string, orderReference string, services []RecordService) error
	OrderOutgoingRevoke(ctx context.Context, userName, userAgent, delegatee, reference string) error
	ServiceCreate(ctx context.Context, userName, userAgent, serviceName string) error
	ServiceUpdate(ctx context.Context, userName, userAgent, serviceName string) error
	ServiceDelete(ctx context.Context, userName, userAgent, serviceName string) error
	OrganizationSettingsUpdate(ctx context.Context, userName, userAgent string) error
	InwayDelete(ctx context.Context, userName, userAgent, inwayName string) error
}
