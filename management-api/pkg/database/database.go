// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"time"

	"go.nlx.io/nlx/management-api/domain"
)

type ConfigDatabase interface {
	GetUser(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, email, password string, roleNames []string) (id uint, error error)
	VerifyUserCredentials(ctx context.Context, email, password string) (bool, error)

	ListServices(ctx context.Context) ([]*Service, error)
	GetService(ctx context.Context, name string) (*Service, error)
	CreateService(ctx context.Context, service *Service) error
	CreateServiceWithInways(ctx context.Context, service *Service, inwayNames []string) error
	UpdateServiceWithInways(ctx context.Context, service *Service, inwayNames []string) error
	UpdateService(ctx context.Context, service *Service) error
	DeleteService(ctx context.Context, serviceName, organizationSerialNumber string) error

	ListInways(ctx context.Context) ([]*Inway, error)
	GetInway(ctx context.Context, name string) (*Inway, error)
	RegisterInway(ctx context.Context, inway *Inway) error
	UpdateInway(ctx context.Context, inway *Inway) error
	DeleteInway(ctx context.Context, name string) error

	ListOutways(ctx context.Context) ([]*Outway, error)
	GetOutway(ctx context.Context, name string) (*Outway, error)
	GetOutwaysByPublicKeyFingerprint(ctx context.Context, publicKeyFingerprint string) ([]*Outway, error)
	RegisterOutway(ctx context.Context, outway *Outway) error
	DeleteOutway(ctx context.Context, name string) error

	GetFingerprintOfPublicKeys(ctx context.Context) ([]string, error)

	GetOutgoingAccessRequest(ctx context.Context, id uint) (*OutgoingAccessRequest, error)
	GetLatestOutgoingAccessRequest(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) (*OutgoingAccessRequest, error)
	ListLatestOutgoingAccessRequests(ctx context.Context, organizationSerialNumber, serviceName string) ([]*OutgoingAccessRequest, error)
	ListAllLatestOutgoingAccessRequests(ctx context.Context) ([]*OutgoingAccessRequest, error)
	CreateOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) (*OutgoingAccessRequest, error)
	UpdateOutgoingAccessRequestState(ctx context.Context, id uint, state OutgoingAccessRequestState) error
	DeleteOutgoingAccessRequests(ctx context.Context, organizationSerialNumber, serviceName string) error
	DeleteOutgoingAccessRequest(ctx context.Context, id uint) error

	ListIncomingAccessRequests(ctx context.Context, serviceName string) ([]*IncomingAccessRequest, error)
	GetLatestIncomingAccessRequest(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) (*IncomingAccessRequest, error)
	GetIncomingAccessRequestCountByService(ctx context.Context) (map[string]int, error)
	GetIncomingAccessRequest(ctx context.Context, id uint) (*IncomingAccessRequest, error)
	CreateIncomingAccessRequest(ctx context.Context, accessRequest *IncomingAccessRequest) (*IncomingAccessRequest, error)
	UpdateIncomingAccessRequestState(ctx context.Context, id uint, state IncomingAccessRequestState) error
	DeleteIncomingAccessRequest(ctx context.Context, id uint) error

	CreateAccessGrant(ctx context.Context, accessRequest *IncomingAccessRequest) (*AccessGrant, error)
	RevokeAccessGrant(ctx context.Context, id uint, revokedAt time.Time) (*AccessGrant, error)
	GetAccessGrantIDForIncomingAccessRequest(ctx context.Context, accessRequestID uint) (uint, error)
	GetAccessGrant(ctx context.Context, id uint) (*AccessGrant, error)
	ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*AccessGrant, error)
	GetLatestAccessGrantForService(ctx context.Context, organizationSerialNumber, serviceName, publicKeyFingerprint string) (*AccessGrant, error)
	TerminateAccessGrant(ctx context.Context, id uint, terminatedAt time.Time) error

	CreateAccessProof(ctx context.Context, accessRequestOutgoingID uint) (*AccessProof, error)
	RevokeAccessProof(ctx context.Context, id uint, revokedAt time.Time) (*AccessProof, error)
	GetAccessProofForOutgoingAccessRequest(ctx context.Context, accessRequestID uint) (*AccessProof, error)
	GetAccessProofs(ctx context.Context, accessProofIDs []uint64) ([]*AccessProof, error)
	TerminateAccessProof(ctx context.Context, id uint, terminatedAt time.Time) error

	GetSettings(ctx context.Context) (*domain.Settings, error)
	UpdateSettings(ctx context.Context, settings *domain.Settings) error

	CreateAuditLogRecord(ctx context.Context, auditLogRecord *AuditLog) (uint64, error)
	ListAuditLogRecords(ctx context.Context, limit int) ([]*AuditLog, error)
	SetAuditLogAsSucceeded(ctx context.Context, id int64) error

	CreateOutgoingOrder(ctx context.Context, order *CreateOutgoingOrder) error
	UpdateOutgoingOrder(ctx context.Context, order *UpdateOutgoingOrder) error
	GetOutgoingOrderByReference(ctx context.Context, reference string) (*OutgoingOrder, error)
	ListOutgoingOrders(ctx context.Context) ([]*OutgoingOrder, error)
	ListOutgoingOrdersByOrganization(ctx context.Context, organizationSerialNumber string) ([]*OutgoingOrder, error)
	RevokeOutgoingOrderByReference(ctx context.Context, delegatee, reference string, revokedAt time.Time) error

	ListIncomingOrders(ctx context.Context) ([]*domain.IncomingOrder, error)
	SynchronizeOrders(ctx context.Context, orders []*IncomingOrder) error

	ListPermissions(ctx context.Context) ([]*domain.Permission, error)

	GetTermsOfServiceStatus(ctx context.Context) (*domain.TermsOfServiceStatus, error)
	AcceptTermsOfService(ctx context.Context, username string, createdAt time.Time) (alreadyAccepted bool, error error)
}
