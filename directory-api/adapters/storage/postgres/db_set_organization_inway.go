package pgadapter

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (r *PostgreSQLRepository) SetOrganizationInway(ctx context.Context, organizationSerialNumber, inwayAddress string) error {
	arg := map[string]interface{}{
		"inway_address":              inwayAddress,
		"organization_serial_number": organizationSerialNumber,
	}

	var ioID struct {
		InwayID        int `db:"inway_id"`
		OrganizationID int `db:"organization_id"`
	}

	err := r.selectInwayByAddressStmt.GetContext(ctx, &ioID, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrNoInwayWithAddress
		}

		return err
	}

	setOrgInwayArgs := map[string]interface{}{
		"inway_id":                   ioID.InwayID,
		"organization_serial_number": organizationSerialNumber,
	}

	_, err = r.setOrganizationInwayStmt.ExecContext(ctx, setOrgInwayArgs)
	if err != nil {
		return err
	}

	return nil
}

func prepareSetOrganizationInwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		UPDATE directory.organizations
		SET inway_id = :inway_id
		WHERE serial_number = :organization_serial_number
	`

	return db.PrepareNamed(query)
}

func prepareSelectInwayByAddressStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		SELECT i.id AS inway_id, i.organization_id
		FROM directory.inways i
		INNER JOIN directory.organizations o ON o.id = i.organization_id
		WHERE i.address = :inway_address
		AND o.serial_number = :organization_serial_number
	`

	return db.PrepareNamed(query)
}
