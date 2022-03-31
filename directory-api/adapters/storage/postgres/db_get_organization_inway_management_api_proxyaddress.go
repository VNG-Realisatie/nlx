// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint:dupl // duplication should be resolved when moving to sqlc
package pgadapter

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (r *PostgreSQLRepository) GetOrganizationInwayManagementAPIProxyAddress(ctx context.Context, organizationSerialNumber string) (string, error) {
	var address sql.NullString

	arg := map[string]interface{}{
		"organization_serial_number": organizationSerialNumber,
	}

	err := r.selectOrganizationInwayManagementAPIProxyAddressStmt.GetContext(ctx, &address, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrNotFound
		}

		return "", err
	}

	return address.String, nil
}

func prepareSelectOrganizationInwayManagementAPIProxyAddressStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		SELECT i.management_api_proxy_address
		FROM directory.organizations o
		LEFT JOIN directory.inways i ON o.inway_id = i.id
		WHERE o.serial_number = :organization_serial_number
	`

	return db.PrepareNamed(query)
}
