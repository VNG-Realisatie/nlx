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

	ListAllOutgoingAccessRequests(ctx context.Context) ([]*OutgoingAccessRequest, error)
	ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*OutgoingAccessRequest, error)
	GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*OutgoingAccessRequest, error)
	ListAllLatestOutgoingAccessRequests(ctx context.Context) (map[string]*OutgoingAccessRequest, error)
	LockOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) error
	UnlockOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) error
	CreateOutgoingAccessRequest(ctx context.Context, accessRequest *OutgoingAccessRequest) (*OutgoingAccessRequest, error)
	UpdateOutgoingAccessRequestState(ctx context.Context, accessRequest *OutgoingAccessRequest, state AccessRequestState) error
	WatchOutgoingAccessRequests(ctx context.Context, output chan *OutgoingAccessRequest)

	ListAllIncomingAccessRequests(ctx context.Context) ([]*IncomingAccessRequest, error)
	ListIncomingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*IncomingAccessRequest, error)
	GetLatestIncomingAccessRequest(ctx context.Context, organizationName, serviceName string) (*IncomingAccessRequest, error)
	ListAllLatestIncomingAccessRequests(ctx context.Context) (map[string]*IncomingAccessRequest, error)
	GetIncomingAccessRequest(ctx context.Context, id string) (*IncomingAccessRequest, error)
	CreateIncomingAccessRequest(ctx context.Context, accessRequest *IncomingAccessRequest) (*IncomingAccessRequest, error)
	UpdateIncomingAccessRequestState(ctx context.Context, accessRequest *IncomingAccessRequest, state AccessRequestState) error

	CreateAccessGrant(ctx context.Context, accessRequest *IncomingAccessRequest) (*AccessGrant, error)
	ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*AccessGrant, error)

	GetSettings(ctx context.Context) (*Settings, error)
	UpdateSettings(ctx context.Context, settings *Settings) error
}
