// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (r *PostgreSQLRepository) GetOrganizationInwayAddress(ctx context.Context, organizationSerialNumber string) (string, error) {
	var address sql.NullString

	arg := map[string]interface{}{
		"organization_serial_number": organizationSerialNumber,
	}

	err := r.selectOrganizationInwayAddressStmt.GetContext(ctx, &address, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrNotFound
		}

		return "", err
	}

	return address.String, nil
}

func prepareSelectOrganizationInwayAddressStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		SELECT i.address
		FROM directory.organizations o
		LEFT JOIN directory.inways i ON o.inway_id = i.id
		WHERE o.serial_number = :organization_serial_number
	`

	return db.PrepareNamed(query)
}
