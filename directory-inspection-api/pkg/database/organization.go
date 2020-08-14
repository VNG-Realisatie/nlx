// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"fmt"
)

type Organization struct {
	Name                string
	InsightIrmaEndpoint string
	InsightLogEndpoint  string
}

// ListOrganizations returns a list of organizations
func (db PostgreSQLDirectoryDatabase) ListOrganizations(ctx context.Context) ([]*Organization, error) {
	var result []*Organization

	rows, err := db.selectOrganizationsStatement.Queryx()
	if err != nil {
		return nil, fmt.Errorf("failed to execute stmtSelectOrganizations: %v", err)
	}

	for rows.Next() {
		// TODO: rename respOrganization to organization
		var respOrganization = &Organization{}
		err = rows.Scan(
			&respOrganization.Name,
			&respOrganization.InsightIrmaEndpoint,
			&respOrganization.InsightLogEndpoint,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan into struct: %v", err)
		}

		result = append(result, respOrganization)
	}

	return result, nil
}
