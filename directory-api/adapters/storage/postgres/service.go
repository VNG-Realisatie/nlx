// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.nlx.io/nlx/directory-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (r *PostgreSQLRepository) GetService(id uint) (*domain.Service, error) {
	service, err := r.queries.GetService(context.Background(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed to get service with id %v: %s", id, err)
	}

	return convertServiceRowToModel(service)
}

func convertServiceRowToModel(row *queries.GetServiceRow) (*domain.Service, error) {
	organization, err := domain.NewOrganization(row.OrganizationName, row.OrganizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid organization model in database: %v", err)
	}

	model, err := domain.NewService(
		&domain.NewServiceArgs{
			Name:                 row.Name,
			Organization:         organization,
			Internal:             row.Internal,
			DocumentationURL:     row.DocumentationUrl.String,
			APISpecificationType: domain.SpecificationType(row.ApiSpecificationType.String),
			PublicSupportContact: row.PublicSupportContact.String,
			TechSupportContact:   row.TechSupportContact.String,
			Costs: &domain.NewServiceCostsArgs{
				OneTime: uint(row.OneTimeCosts),
				Monthly: uint(row.MonthlyCosts),
				Request: uint(row.RequestCosts),
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid service model in database: %v", err)
	}

	model.SetID(uint(row.ID))

	return model, nil
}
