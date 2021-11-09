package pgadapter

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) RegisterService(model *domain.Service) error {
	type registerParams struct {
		Name                 string `db:"name"`
		SerialNumber         string `db:"organization_serial_number"`
		DocumentationURL     string `db:"documentation_url"`
		APISpecificationType string `db:"api_specification_type"`
		PublicSupportContact string `db:"public_support_contact"`
		TechSupportContact   string `db:"tech_support_contact"`
		OneTimeCosts         int32  `db:"one_time_costs"`
		MonthlyCosts         int32  `db:"monthly_costs"`
		RequestCosts         int32  `db:"request_costs"`
		Internal             bool   `db:"internal"`
	}

	type dbResult struct {
		ID uint `db:"id"`
	}

	result := dbResult{}

	err := r.registerServiceStmt.Get(
		&result,
		&registerParams{
			Name:                 model.Name(),
			SerialNumber:         model.Organization().SerialNumber(),
			Internal:             model.Internal(),
			DocumentationURL:     model.DocumentationURL(),
			APISpecificationType: string(model.APISpecificationType()),
			PublicSupportContact: model.PublicSupportContact(),
			TechSupportContact:   model.TechSupportContact(),
			OneTimeCosts:         int32(model.Costs().OneTime()),
			MonthlyCosts:         int32(model.Costs().Monthly()),
			RequestCosts:         int32(model.Costs().Request()),
		})

	if err != nil {
		return fmt.Errorf("failed to register service: %v", err)
	}

	model.SetID(result.ID)

	return err
}

func prepareRegisterServiceStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		with organization as (
		    select id from directory.organizations where directory.organizations.serial_number = :organization_serial_number
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
		),
		availabilities as (
					insert into directory.availabilities
								(inway_id, service_id, last_announced)
						 select inway.id, service.id, now()
						   from inway, service
					on conflict on constraint availabilities_uq_inway_service
				  do update set last_announced = now(),
								active = true
		) select id from service;

	`

	return db.PrepareNamed(query)
}
