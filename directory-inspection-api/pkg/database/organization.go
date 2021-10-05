// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var ErrNoOrganization = errors.New("no organization found")

func (db PostgreSQLDirectoryDatabase) ListOrganizations(_ context.Context) ([]*Organization, error) {
	var result []*Organization

	rows, err := db.selectOrganizationsStatement.Queryx()
	if err != nil {
		return nil, fmt.Errorf("failed to execute stmtSelectOrganizations: %v", err)
	}

	for rows.Next() {
		var organization = &Organization{}
		err = rows.Scan(
			&organization.SerialNumber,
			&organization.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan into struct: %v", err)
		}

		result = append(result, organization)
	}

	return result, nil
}

func (db PostgreSQLDirectoryDatabase) GetOrganizationInwayAddress(ctx context.Context, organizationSerialNumber string) (string, error) {
	var address string

	arg := map[string]interface{}{
		"organization_serial_number": organizationSerialNumber,
	}

	err := db.selectOrganizationInwayAddressStatement.GetContext(ctx, &address, arg)

	if errors.Is(err, sql.ErrNoRows) {
		return address, ErrNoOrganization
	}

	return address, err
}

func prepareSelectOrganizationInwayAddressStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	s := `
		SELECT i.address
		FROM directory.inways i
		INNER JOIN directory.organizations o ON o.inway_id = i.id
		WHERE o.serial_number = :organization_serial_number
	`

	return db.PrepareNamed(s)
}
