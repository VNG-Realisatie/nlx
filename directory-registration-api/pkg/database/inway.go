// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// nolint:gocritic // these are valid regex patterns
var organizationNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-._\s]{1,100}$`)

type RegisterInwayParams struct {
	OrganizationName string
	Address          string
	NlxVersion       string
}

func (params *RegisterInwayParams) Validate() error {
	return validation.ValidateStruct(
		params,
		validation.Field(&params.OrganizationName, validation.Required, validation.Match(organizationNameRegex)),
		validation.Field(
			&params.Address,
			validation.Required,
			validation.When(strings.Contains(params.Address, ":"), is.DialString),
			validation.When(!strings.Contains(params.Address, ":"), is.DNSName),
		),
		validation.Field(&params.NlxVersion, validation.When(params.NlxVersion != "unknown", is.Semver)),
	)
}

func (db PostgreSQLDirectoryDatabase) RegisterInway(params *RegisterInwayParams) error {
	type upsertInwayParams struct {
		OrganizationName string `db:"organization_name"`
		InwayAddress     string `db:"inway_address"`
		InwayVersion     string `db:"inway_version"`
	}

	_, err := db.upsertInwayStmt.Exec(&upsertInwayParams{
		OrganizationName: params.OrganizationName,
		InwayAddress:     params.Address,
		InwayVersion:     params.NlxVersion,
	})
	if err != nil {
		return errors.Wrap(err, "failed upsert inway and its organization")
	}

	return nil
}

func prepareUpsertInwayStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
			insert into directory.organizations 
			            (name)
			     values (:organization_name)
			on conflict 
			    		on constraint organizations_uq_name
			  			do update 
			      			set name = excluded.name -- no-op update to return id
						returning id
		) 
		insert into directory.inways 
		    		(organization_id, address, version)
			 select organization.id, :inway_address, nullif(:inway_version, '')
			   from organization
	   	on conflict 
	   	    		on constraint inways_uq_address
	      			do update set 
	      			              address = excluded.address, 
								  version = excluded.version;
	`

	return db.PrepareNamed(query)
}
