// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"fmt"

	"go.nlx.io/nlx/directory-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) GetInway(name, serialNumber string) (*domain.Inway, error) {
	inway, err := r.queries.GetInway(context.Background(), &queries.GetInwayParams{
		SerialNumber: serialNumber,
		Name:         name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get inway (name: %s, serialNumber: %s): %s", name, serialNumber, err)
	}

	organizationModel, err := domain.NewOrganization(inway.OrganizationName, inway.OrganizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid organization model in database: %v", err)
	}

	model, err := domain.NewInway(&domain.NewInwayArgs{
		Name:         inway.Name,
		Organization: organizationModel,
		Address:      inway.Address,
		NlxVersion:   inway.NlxVersion,
		CreatedAt:    inway.CreatedAt,
		UpdatedAt:    inway.UpdatedAt,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid inway model in database: %v", err)
	}

	return model, nil
}
