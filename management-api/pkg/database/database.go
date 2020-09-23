// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
)

// ConfigDatabase is the interface for a configuration database
type ConfigDatabase interface {
	ListServices(ctx context.Context) ([]*Service, error)
	GetService(ctx context.Context, name string) (*Service, error)
	CreateService(ctx context.Context, service *Service) error
	UpdateService(ctx context.Context, name string, service *Service) error
	DeleteService(ctx context.Context, name string) error
	ListInways(ctx context.Context) ([]*Inway, error)
	GetInway(ctx context.Context, name string) (*Inway, error)
	CreateInway(ctx context.Context, inway *Inway) error
	UpdateInway(ctx context.Context, name string, inway *Inway) error
	DeleteInway(ctx context.Context, name string) error
	PutInsightConfiguration(ctx context.Context, configuration *InsightConfiguration) error
	GetInsightConfiguration(ctx context.Context) (*InsightConfiguration, error)

	ListAllOutgoingAccessRequests(ctx context.Context) ([]*AccessRequest, error)
	ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*AccessRequest, error)
	GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*AccessRequest, error)
	ListAllLatestOutgoingAccessRequests(ctx context.Context) (map[string]*AccessRequest, error)
	LockOutgoingAccessRequest(ctx context.Context, accessRequest *AccessRequest) error
	UnlockOutgoingAccessRequest(ctx context.Context, accessRequest *AccessRequest) error

	CreateAccessRequest(ctx context.Context, accessRequest *AccessRequest) (*AccessRequest, error)
	UpdateAccessRequestState(ctx context.Context, accessRequest *AccessRequest, state AccessRequestState) error
	WatchOutgoingAccessRequests(ctx context.Context, output chan *AccessRequest)

	CreateAccessGrant(ctx context.Context, accessGrant *AccessGrant) (*AccessGrant, error)
	ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*AccessGrant, error)

	GetSettings(ctx context.Context) (*Settings, error)
	UpdateSettings(ctx context.Context, settings *Settings) error
}
