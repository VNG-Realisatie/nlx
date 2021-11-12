package pgadapter

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) ListOrganizations(ctx context.Context) ([]*domain.Organization, error) {
	rows, err := r.selectOrganizationsStmt.Queryx()
	if err != nil {
		return nil, fmt.Errorf("failed to execute stmtSelectOrganizations: %v", err)
	}

	type dbOrganization struct {
		SerialNumber string `db:"serial_number"`
		Name         string `db:"name"`
	}

	var queryResult []*dbOrganization

	err = sqlx.StructScan(rows, &queryResult)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Organization, len(queryResult))

	for i, organizationInDB := range queryResult {
		org, err := domain.NewOrganization(organizationInDB.Name, organizationInDB.SerialNumber)
		if err != nil {
			return nil, err
		}

		result[i] = org
	}

	return result, nil
}

func prepareSelectOrganizationsStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	listOrganizationsStatement, err := db.Preparex(`
		SELECT
			serial_number,
			name
		FROM directory.organizations
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}

	return listOrganizationsStatement, nil
}
