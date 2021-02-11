// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"time"

	"go.nlx.io/nlx/common/diagnostics"
)

// ConfigDatabase is the interface for a configuration database
type ConfigDatabase interface {
	GetUser(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, email string, roleNames []string) (*User, error)

	ListServices(ctx context.Context) ([]*Service, error)
	GetService(ctx context.Context, name string) (*Service, error)
	CreateService(ctx context.Context, service *Service) error
	CreateServiceWithInways(ctx context.Context, service *Service, inwayNames []string) error
	UpdateServiceWithInways(ctx context.Context, service *Service, inwayNames []string) error
	UpdateService(ctx context.Context, service *Service) error
	DeleteService(ctx context.Context, name string) error

	ListInways(ctx context.Context) ([]*Inway, error)
	GetInway(ctx context.Context, name string) (*Inway, error)
	CreateInway(ctx context.Context, inway *Inway) error
	UpdateInway(ctx context.Context, inway *Inway) error
	DeleteInway(ctx context.Context, name string) error

	ListAllOutgoingAccessRequests(ctx context.Context) ([]*OutgoingAccessRequest, error)
	ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*OutgoingAccessRequest, error)
	GetOutgoingAccessRequest(ctx context.Context, id uint) (*OutgoingAccessRequest, error)
	GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*OutgoingAccessRequest, error)
	CreateOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) (*OutgoingAccessRequest, error)
	UpdateOutgoingAccessRequestState(ctx context.Context, id uint, state OutgoingAccessRequestState, referenceID uint, err *diagnostics.ErrorDetails) error
	TakePendingOutgoingAccessRequest(ctx context.Context) (*OutgoingAccessRequest, error)
	UnlockOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) error

	ListAllIncomingAccessRequests(ctx context.Context) ([]*IncomingAccessRequest, error)
	ListIncomingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*IncomingAccessRequest, error)
	GetLatestIncomingAccessRequest(ctx context.Context, organizationName, serviceName string) (*IncomingAccessRequest, error)
	GetIncomingAccessRequestCountByService(ctx context.Context) (map[string]int, error)
	GetIncomingAccessRequest(ctx context.Context, id uint) (*IncomingAccessRequest, error)
	CreateIncomingAccessRequest(ctx context.Context, accessRequest *IncomingAccessRequest) (*IncomingAccessRequest, error)
	UpdateIncomingAccessRequestState(ctx context.Context, id uint, state IncomingAccessRequestState) error

	CreateAccessGrant(ctx context.Context, accessRequest *IncomingAccessRequest) (*AccessGrant, error)
	RevokeAccessGrant(ctx context.Context, id uint, revokedAt time.Time) (*AccessGrant, error)
	ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*AccessGrant, error)
	GetLatestAccessGrantForService(ctx context.Context, organizationName, serviceName string) (*AccessGrant, error)

	CreateAccessProof(ctx context.Context, accessRequest *OutgoingAccessRequest) (*AccessProof, error)
	RevokeAccessProof(ctx context.Context, id uint, revokedAt time.Time) (*AccessProof, error)
	GetLatestAccessProofForService(ctx context.Context, organizationName, serviceName string) (*AccessProof, error)
	GetAccessProofForOutgoingAccessRequest(ctx context.Context, accessRequestID uint) (*AccessProof, error)

	GetSettings(ctx context.Context) (*Settings, error)
	PutOrganizationInway(ctx context.Context, inwayID *uint) (*Settings, error)
	PutInsightConfiguration(ctx context.Context, irmaServerURL, insightAPIURL string) (*Settings, error)

	CreateAuditLogRecord(ctx context.Context, auditLogRecord *AuditLogRecord) (*AuditLogRecord, error)
	ListAuditLogRecords(ctx context.Context) ([]*AuditLogRecord, error)
}
