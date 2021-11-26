// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) GetOutway(name, serialNumber string) (*domain.Outway, error) {
	type params struct {
		Name                     string `db:"name"`
		OrganizationSerialNumber string `db:"organization_serial_number"`
	}

	type dbOutway struct {
		Name                     string    `db:"name"`
		NlxVersion               string    `db:"nlx_version"`
		OrganizationName         string    `db:"organization_name"`
		OrganizationSerialNumber string    `db:"organization_serial_number"`
		CreatedAt                time.Time `db:"created_at"`
		UpdatedAt                time.Time `db:"updated_at"`
	}

	result := dbOutway{}

	err := r.getOutwayStmt.Get(&result, &params{
		Name:                     name,
		OrganizationSerialNumber: serialNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get outway (name: %s, serialNumber: %s): %s", name, serialNumber, err)
	}

	organizationModel, err := domain.NewOrganization(result.OrganizationName, result.OrganizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid organization model in database: %v", err)
	}

	model, err := domain.NewOutway(&domain.NewOutwayArgs{
		Name:         result.Name,
		Organization: organizationModel,
		NlxVersion:   result.NlxVersion,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid outway model in database: %v", err)
	}

	return model, nil
}

func prepareGetOutwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		select directory.outways.name as name, version as nlx_version, directory.outways.created_at as created_at, updated_at, directory.organizations.serial_number as organization_serial_number, directory.organizations.name as organization_name
		from directory.outways
		join directory.organizations
		    on directory.outways.organization_id = directory.organizations.id
		where
		      directory.organizations.serial_number = :organization_serial_number
		  and directory.outways.name = :name;
	`

	return db.PrepareNamed(query)
}
