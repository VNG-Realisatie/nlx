package pgadapter

import (
	"time"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (r *PostgreSQLRepository) RegisterInway(model *domain.Inway) error {
	type registerParams struct {
		OrganizationName         string    `db:"organization_name"`
		OrganizationSerialNumber string    `db:"organization_serial_number"`
		Name                     string    `db:"inway_name"`
		Address                  string    `db:"inway_address"`
		NlxVersion               string    `db:"inway_version"`
		CreatedAt                time.Time `db:"inway_created_at"`
		UpdatedAt                time.Time `db:"inway_updated_at"`
	}

	organization := model.Organization()

	_, err := r.registerInwayStmt.Exec(&registerParams{
		OrganizationName:         organization.Name(),
		OrganizationSerialNumber: organization.SerialNumber(),
		Name:                     model.Name(),
		Address:                  model.Address(),
		NlxVersion:               model.NlxVersion(),
		CreatedAt:                model.CreatedAt(),
		UpdatedAt:                model.UpdatedAt(),
	})

	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"inways_uq_address\"" {
		return storage.ErrDuplicateAddress
	}

	return err
}

func prepareRegisterInwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
			insert into directory.organizations
			            (serial_number, name)
			     values (:organization_serial_number, :organization_name)
			on conflict
			    		on constraint organizations_uq_serial_number
			  			do update
			      			set serial_number = excluded.serial_number, -- no-op update to return id
			      			    name 		  = excluded.name
						returning id
		)
		insert into directory.inways
		    		(name, organization_id, address, version, created_at, updated_at)
			 select :inway_name, organization.id, :inway_address, nullif(:inway_version, ''), :inway_created_at, :inway_updated_at
			   from organization
	   	on conflict (name, organization_id) do update set
	      			              name = 	excluded.name,
	      			              address = excluded.address,
								  version = excluded.version,
								  updated_at = excluded.updated_at;
	`

	return db.PrepareNamed(query)
}
