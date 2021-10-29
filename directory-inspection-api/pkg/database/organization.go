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

type Organization struct {
	Name string
}

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
			&organization.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan into struct: %v", err)
		}

		result = append(result, organization)
	}

	return result, nil
}

func (db PostgreSQLDirectoryDatabase) GetOrganizationInwayAddress(ctx context.Context, organizationName string) (string, error) {
	var address string

	arg := map[string]interface{}{
		"organization_name": organizationName, // @TODO serial
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
		WHERE o.name = :organization_name
	`

	return db.PrepareNamed(s)
}
