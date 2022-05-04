// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"fmt"

	"go.nlx.io/nlx/directory-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) GetOutway(name, serialNumber string) (*domain.Outway, error) {
	row, err := r.queries.GetOutway(context.Background(), &queries.GetOutwayParams{
		SerialNumber: serialNumber,
		Name:         name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get outway (name: %s, serialNumber: %s): %s", name, serialNumber, err)
	}

	organizationModel, err := domain.NewOrganization(row.OrganizationName, row.OrganizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid organization model in database: %v", err)
	}

	model, err := domain.NewOutway(&domain.NewOutwayArgs{
		Name:         row.Name,
		Organization: organizationModel,
		NlxVersion:   row.NlxVersion,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid outway model in database: %v", err)
	}

	return model, nil
}

func (r *PostgreSQLRepository) RegisterOutway(model *domain.Outway) error {
	return r.queries.RegisterOutway(context.Background(), &queries.RegisterOutwayParams{
		SerialNumber: model.Organization().SerialNumber(),
		Name:         model.Organization().Name(),
		Name_2:       model.Name(),
		Column4:      model.NlxVersion(),
		CreatedAt:    model.CreatedAt(),
		UpdatedAt:    model.UpdatedAt(),
	})
}
