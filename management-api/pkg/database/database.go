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
}
