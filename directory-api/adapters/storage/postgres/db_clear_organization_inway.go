package pgadapter

import (
	"context"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (r *PostgreSQLRepository) ClearOrganizationInway(ctx context.Context, organizationSerialNumber string) error {
	arg := map[string]interface{}{
		"organization_serial_number": organizationSerialNumber,
	}

	res, err := r.clearOrganizationInwayStmt.ExecContext(ctx, arg)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n != 1 {
		return storage.ErrOrganizationNotFound
	}

	return nil
}

func prepareClearOrganizationInwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		UPDATE directory.organizations
		SET inway_id = null
		WHERE serial_number = :organization_serial_number
	`

	return db.PrepareNamed(query)
}
