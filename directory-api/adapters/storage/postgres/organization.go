// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package pgadapter

import (
	"context"

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
