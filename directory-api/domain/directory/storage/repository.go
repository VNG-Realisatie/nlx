// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package storage

import (
	"context"
	"errors"

	"go.nlx.io/nlx/directory-api/domain"
)

var (
	ErrDuplicateAddress   = errors.New("another inway is already registered with this address")
	ErrNoInwayWithAddress = errors.New("no inway found for address")
	ErrNotFound           = errors.New("no record found")
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
	SetOrganizationEmailAddress(ctx context.Context, organization *domain.Organization, emailAddress string) error

	ListServices(ctx context.Context, organizationSerialNumber string) ([]*domain.Service, error)
	ListOrganizations(ctx context.Context) ([]*domain.Organization, error)
	ListVersionStatistics(ctx context.Context) ([]*domain.VersionStatistics, error)

	RegisterOutway(*domain.Outway) error
	GetOutway(name, organizationSerialNumber string) (*domain.Outway, error)

	ListParticipants(ctx context.Context) ([]*domain.Participant, error)

	Shutdown() error
}
