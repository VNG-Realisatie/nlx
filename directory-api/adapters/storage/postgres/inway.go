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

func (r *PostgreSQLRepository) GetInway(name, serialNumber string) (*domain.Inway, error) {
	inway, err := r.queries.GetInway(context.Background(), &queries.GetInwayParams{
		SerialNumber: serialNumber,
		Name:         name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get inway (name: %s, serialNumber: %s): %s", name, serialNumber, err)
	}

	return convertInwayRowToModel(inway)
}

func convertInwayRowToModel(row *queries.GetInwayRow) (*domain.Inway, error) {
	organizationModel, err := domain.NewOrganization(row.OrganizationName, row.OrganizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid organization model in database: %v", err)
	}

	model, err := domain.NewInway(&domain.NewInwayArgs{
		Name:         row.Name,
		Organization: organizationModel,
		Address:      row.Address,
		NlxVersion:   row.NlxVersion,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid inway model in database: %v", err)
	}

	return model, nil
}

func (r *PostgreSQLRepository) GetOrganizationInwayAddress(ctx context.Context, organizationSerialNumber string) (string, error) {
	inwayAddress, err := r.queries.SelectOrganizationInwayAddress(ctx, organizationSerialNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrNotFound
		}

		return "", err
	}

	return inwayAddress.String, nil
}

func (r *PostgreSQLRepository) GetOrganizationInwayManagementAPIProxyAddress(ctx context.Context, organizationSerialNumber string) (string, error) {
	inwayProxyAddress, err := r.queries.SelectOrganizationInwayManagementAPIProxyAddress(ctx, organizationSerialNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrNotFound
		}

		return "", err
	}

	return inwayProxyAddress.String, nil
}

func (r *PostgreSQLRepository) RegisterInway(model *domain.Inway) error {
	err := r.queries.RegisterInway(context.Background(), &queries.RegisterInwayParams{
		OrganizationSerialNumber:  model.Organization().SerialNumber(),
		OrganizationName:          model.Organization().Name(),
		Name:                      model.Name(),
		Address:                   model.Address(),
		ManagementApiProxyAddress: model.ManagementAPIProxyAddress(),
		Version:                   model.NlxVersion(),
		CreatedAt:                 model.CreatedAt(),
		UpdatedAt:                 model.UpdatedAt(),
	})
	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"inways_uq_address\"" {
		return storage.ErrDuplicateAddress
	}

	return err
}
