// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"
	"go.nlx.io/nlx/directory-registration-api/domain"
)

type Repository interface {
	RegisterInway(*domain.Inway) error
	GetInway(name, organization string) (*domain.Inway, error)

	RegisterService(*domain.Service) error
	GetService(id uint) (*domain.Service, error)

	SetOrganizationInway(ctx context.Context, organizationName, inwayAddress string) error
	ClearOrganizationInway(ctx context.Context, organizationName string) error
	GetOrganizationInwayAddress(ctx context.Context, organizationName string) (string, error)
}
