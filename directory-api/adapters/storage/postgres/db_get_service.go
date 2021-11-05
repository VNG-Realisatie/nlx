package pgadapter

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) GetService(id uint) (*domain.Service, error) {
	type params struct {
		ID uint `db:"id"`
	}

	type dbService struct {
		ID                       uint   `db:"id"`
		Name                     string `db:"name"`
		OrganizationSerialNumber string `db:"organization_serial_number"`
		OrganizationName         string `db:"organization_name"`
		DocumentationURL         string `db:"documentation_url"`
		APISpecificationType     string `db:"api_specification_type"`
		PublicSupportContact     string `db:"public_support_contact"`
		TechSupportContact       string `db:"tech_support_contact"`
		OneTimeCosts             int32  `db:"one_time_costs"`
		MonthlyCosts             int32  `db:"monthly_costs"`
		RequestCosts             int32  `db:"request_costs"`
		Internal                 bool   `db:"internal"`
	}

	result := dbService{}

	err := r.getServiceStmt.Get(&result, &params{ID: id})
	if err != nil {
		return nil, fmt.Errorf("failed to get service with id %v: %s", id, err)
	}

	organization, err := domain.NewOrganization(result.OrganizationName, result.OrganizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid organization model in database: %v", err)
	}

	model, err := domain.NewService(
		&domain.NewServiceArgs{
			Name:                 result.Name,
			Organization:         organization,
			Internal:             result.Internal,
			DocumentationURL:     result.DocumentationURL,
			APISpecificationType: domain.SpecificationType(result.APISpecificationType),
			PublicSupportContact: result.PublicSupportContact,
			TechSupportContact:   result.TechSupportContact,
			Costs: &domain.ServiceCosts{
				OneTime: uint(result.OneTimeCosts),
				Monthly: uint(result.MonthlyCosts),
				Request: uint(result.RequestCosts),
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid service model in database: %v", err)
	}

	model.SetID(result.ID)

	return model, nil
}

func prepareGetServiceStmt(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	query := `
		select directory.services.id as id, directory.services.name as name, documentation_url, api_specification_type, internal, tech_support_contact, public_support_contact, directory.organizations.serial_number as organization_serial_number, directory.organizations.name as organization_name, one_time_costs, monthly_costs, request_costs
		from directory.services
		join directory.organizations
		    on directory.services.organization_id = directory.organizations.id
		where
		      directory.services.id = :id;
	`

	return db.PrepareNamed(query)
}
