// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) ListServices(_ context.Context, organizationSerialNumber string) ([]*domain.Service, error) {
	rows, err := r.selectServicesStmt.Queryx(organizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to execute stmtSelectServices: %v", err)
	}

	type dbService struct {
		ID                       uint           `db:"id"`
		Name                     string         `db:"name"`
		OrganizationSerialNumber string         `db:"organization_serial_number"`
		OrganizationName         string         `db:"organization_name"`
		DocumentationURL         string         `db:"documentation_url"`
		APISpecificationType     string         `db:"api_specification_type"`
		PublicSupportContact     string         `db:"public_support_contact"`
		TechSupportContact       string         `db:"tech_support_contact"`
		InwayAddresses           pq.StringArray `db:"inway_addresses"`
		HealthyStatuses          pq.BoolArray   `db:"healthy_statuses"`
		OneTimeCosts             int32          `db:"one_time_costs"`
		MonthlyCosts             int32          `db:"monthly_costs"`
		RequestCosts             int32          `db:"request_costs"`
		Internal                 bool           `db:"internal"`
	}

	var queryResult []*dbService

	err = sqlx.StructScan(rows, &queryResult)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Service, len(queryResult))

	for i, service := range queryResult {
		if len(service.InwayAddresses) != len(service.HealthyStatuses) {
			err := errors.New("length of the inwayadresses does not match healthchecks")
			r.logger.Error("failed to convert service to domain model", zap.Error(err))

			return nil, err
		}

		organization, err := domain.NewOrganization(service.OrganizationName, service.OrganizationSerialNumber)
		if err != nil {
			return nil, err
		}

		inways := make([]*domain.NewServiceInwayArgs, len(service.InwayAddresses))

		for i, inwayAddress := range service.InwayAddresses {
			inwayArgs := &domain.NewServiceInwayArgs{
				Address: inwayAddress,
				State:   domain.InwayDOWN,
			}

			if service.HealthyStatuses[i] {
				inwayArgs.State = domain.InwayUP
			}

			inways[i] = inwayArgs
		}

		result[i], err = domain.NewService(&domain.NewServiceArgs{
			Name:                 service.Name,
			Organization:         organization,
			Internal:             service.Internal,
			DocumentationURL:     service.DocumentationURL,
			APISpecificationType: domain.SpecificationType(service.APISpecificationType),
			PublicSupportContact: service.PublicSupportContact,
			TechSupportContact:   service.TechSupportContact,
			Costs: &domain.NewServiceCostsArgs{
				OneTime: uint(service.OneTimeCosts),
				Monthly: uint(service.MonthlyCosts),
				Request: uint(service.RequestCosts),
			},
			Inways: inways,
		})
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func prepareSelectServicesStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	selectServicesStatement, err := db.Preparex(`
		SELECT
			o.serial_number as organization_serial_number,
			o.name AS organization_name,
			s.name AS name,
			s.internal as internal,
			s.one_time_costs as one_time_costs,
			s.monthly_costs as monthly_costs,
			s.request_costs as request_costs,
			array_remove(array_agg(i.address), NULL) AS inway_addresses,
			COALESCE(s.documentation_url, '') AS documentation_url,
			COALESCE(s.api_specification_type, '') AS api_specification_type,
			COALESCE(s.public_support_contact, '') AS public_support_contact,
			array_remove(array_agg(a.healthy), NULL) as healthy_statuses
		FROM directory.services s
		INNER JOIN directory.availabilities a ON a.service_id = s.id
		INNER JOIN directory.organizations o ON o.id = s.organization_id
		INNER JOIN directory.inways i ON i.id = a.inway_id
		WHERE (
			internal = false
			OR (
				internal = true
				AND o.serial_number = $1
			)
		)
		GROUP BY s.id, o.id
		ORDER BY o.name, s.name
	`)
	if err != nil {
		return nil, err
	}

	return selectServicesStatement, nil
}
