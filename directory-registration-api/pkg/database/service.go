// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// nolint:gocritic // these are valid regex patterns
var serviceNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-.\s]{1,100}$`)

const (
	OpenAPI2 string = "OpenAPI2"
	OpenAPI3 string = "OpenAPI3"
)

type RegisterServiceParams struct {
	OrganizationName     string
	Name                 string
	Internal             bool
	DocumentationURL     string
	APISpecificationType string
	PublicSupportContact string
	TechSupportContact   string
	OneTimeCosts         int32
	MonthlyCosts         int32
	RequestCosts         int32
}

func (params *RegisterServiceParams) Validate() error {
	return validation.ValidateStruct(
		params,
		validation.Field(&params.Name, validation.Required, validation.Match(serviceNameRegex)),
		validation.Field(&params.APISpecificationType, validation.In("", OpenAPI2, OpenAPI3)),
	)
}

func (db PostgreSQLDirectoryDatabase) RegisterService(params *RegisterServiceParams) error {
	type upsertServiceParams struct {
		OrganizationName     string `db:"organization_name"`
		Name                 string `db:"name"`
		Internal             bool   `db:"internal"`
		DocumentationURL     string `db:"documentation_url"`
		APISpecificationType string `db:"api_specification_type"`
		PublicSupportContact string `db:"public_support_contact"`
		TechSupportContact   string `db:"tech_support_contact"`
		OneTimeCosts         int32  `db:"one_time_costs"`
		MonthlyCosts         int32  `db:"monthly_costs"`
		RequestCosts         int32  `db:"request_costs"`
	}

	_, err := db.upsertServiceStmt.ExecContext(context.Background(), &upsertServiceParams{
		OrganizationName:     params.OrganizationName,
		Name:                 params.Name,
		Internal:             params.Internal,
		DocumentationURL:     params.DocumentationURL,
		APISpecificationType: params.APISpecificationType,
		PublicSupportContact: params.PublicSupportContact,
		TechSupportContact:   params.TechSupportContact,
		OneTimeCosts:         params.OneTimeCosts,
		MonthlyCosts:         params.MonthlyCosts,
		RequestCosts:         params.RequestCosts,
	})
	if err != nil {
		return errors.Wrap(err, "failed upsert service")
	}

	return nil
}

func prepareUpsertServiceStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
		    select id from directory.organizations where name = :organization_name
		),
	    inway as (
		    select directory.inways.id from directory.inways, organization where organization_id = organization.id
		),
		service as (
			insert into directory.services 
			    		(organization_id, name, internal, documentation_url, api_specification_type, public_support_contact, tech_support_contact, request_costs, monthly_costs, one_time_costs)
				 select organization.id, :name, :internal, nullif(:documentation_url, ''), nullif(:api_specification_type, ''), nullif(:public_support_contact, ''), nullif(:tech_support_contact, ''), :request_costs, :monthly_costs, :one_time_costs
				   from organization
			on conflict on constraint services_uq_name 
		  do update set internal = excluded.internal,
		  				documentation_url = excluded.documentation_url,
	 					api_specification_type = excluded.api_specification_type,
	 					public_support_contact = excluded.public_support_contact,
	   					tech_support_contact = excluded.tech_support_contact,
						request_costs = excluded.request_costs,
	          			monthly_costs = excluded.monthly_costs,
	         			one_time_costs = excluded.one_time_costs
		      returning id
		)
		insert into directory.availabilities 
		    		(inway_id, service_id, last_announced)
			 select inway.id, service.id, now()
			   from inway, service
		on conflict on constraint availabilities_uq_inway_service
	  do update set last_announced = now(),
				    active = true
	`

	return db.PrepareNamed(query)
}
