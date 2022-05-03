// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"database/sql"
	"errors"

	"go.nlx.io/nlx/directory-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (r *PostgreSQLRepository) ClearOrganizationInway(ctx context.Context, organizationSerialNumber string) error {
	rowsAffected, err := r.queries.ClearOrganizationInway(ctx, organizationSerialNumber)
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return storage.ErrNotFound
	}

	return nil
}

func (r *PostgreSQLRepository) SetOrganizationInway(ctx context.Context, organizationSerialNumber, inwayAddress string) error {
	inway, err := r.queries.SelectInwayByAddress(ctx, &queries.SelectInwayByAddressParams{
		Address:      inwayAddress,
		SerialNumber: organizationSerialNumber,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrNoInwayWithAddress
		}

		return err
	}

	return r.queries.SetOrganizationInway(ctx, &queries.SetOrganizationInwayParams{
		InwayID:      sql.NullInt32{Int32: inway.InwayID, Valid: true},
		SerialNumber: organizationSerialNumber,
	})
}