// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package storage

import (
	"context"
	"errors"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-api/domain"
)

var (
	ErrDuplicateAddress     = errors.New("another inway is already registered with this address")
	ErrNoInwayWithAddress   = errors.New("no inway found for address")
	ErrOrganizationNotFound = errors.New("no organization found")
)

type Repository interface {
	RegisterInway(*domain.Inway) error
	GetInway(name, organizationSerialNumber string) (*domain.Inway, error)

	RegisterService(*domain.Service) error
	GetService(id uint) (*domain.Service, error)

	SetOrganizationInway(ctx context.Context, organizationSerialNumber, inwayAddress string) error
	ClearOrganizationInway(ctx context.Context, organizationSerialNumber string) error
	ClearIfSetAsOrganizationInway(ctx context.Context, organizationSerialNumber, inwayAddress string) error
	GetOrganizationInwayAddress(ctx context.Context, organizationSerialNumber string) (string, error)

	ListServices(ctx context.Context, organizationSerialNumber string) ([]*domain.Service, error)
	RegisterOutwayVersion(ctx context.Context, version nlxversion.Version) error
	ListOrganizations(ctx context.Context) ([]*domain.Organization, error)
	ListVersionStatistics(ctx context.Context) ([]*domain.VersionStatistics, error)

	Shutdown() error
}
