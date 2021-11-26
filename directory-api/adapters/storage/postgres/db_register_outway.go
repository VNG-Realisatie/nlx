// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"time"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) RegisterOutway(model *domain.Outway) error {
	type registerParams struct {
		OrganizationName         string    `db:"organization_name"`
		OrganizationSerialNumber string    `db:"organization_serial_number"`
		Name                     string    `db:"outway_name"`
		NlxVersion               string    `db:"outway_version"`
		CreatedAt                time.Time `db:"outway_created_at"`
		UpdatedAt                time.Time `db:"outway_updated_at"`
	}

	organization := model.Organization()

	_, err := r.registerOutwayStmt.Exec(&registerParams{
		OrganizationName:         organization.Name(),
		OrganizationSerialNumber: organization.SerialNumber(),
		Name:                     model.Name(),
		NlxVersion:               model.NlxVersion(),
		CreatedAt:                model.CreatedAt(),
		UpdatedAt:                model.UpdatedAt(),
	})

	return err
}

func prepareRegisterOutwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
			insert into directory.organizations
			            (serial_number, name)
			     values (:organization_serial_number, :organization_name)
			on conflict
			    		on constraint organizations_uq_serial_number
			  			do update
			      			set serial_number = excluded.serial_number,
			      			    name 		  = excluded.name
						returning id
		)
		insert into directory.outways
		    		(name, organization_id, version, created_at, updated_at)
			 select :outway_name, organization.id, nullif(:outway_version, ''), :outway_created_at, :outway_updated_at
			   from organization
	   	on conflict (name, organization_id) do update set
	      			              name = 	excluded.name,
								  version = excluded.version,
								  updated_at = excluded.updated_at;
	`

	return db.PrepareNamed(query)
}
