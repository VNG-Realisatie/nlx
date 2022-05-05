// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (r *PostgreSQLRepository) GetService(id uint) (*domain.Service, error) {
	service, err := r.queries.GetService(context.Background(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed to get service with id %v: %s", id, err)
	}

	return convertServiceRowToModel(service)
}

func convertServiceRowToModel(row *queries.GetServiceRow) (*domain.Service, error) {
	organization, err := domain.NewOrganization(row.OrganizationName, row.OrganizationSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid organization model in database: %v", err)
	}

	model, err := domain.NewService(
		&domain.NewServiceArgs{
			Name:                 row.Name,
			Organization:         organization,
			Internal:             row.Internal,
			DocumentationURL:     row.DocumentationUrl.String,
			APISpecificationType: domain.SpecificationType(row.ApiSpecificationType.String),
			PublicSupportContact: row.PublicSupportContact.String,
			TechSupportContact:   row.TechSupportContact.String,
			Costs: &domain.NewServiceCostsArgs{
				OneTime: uint(row.OneTimeCosts),
				Monthly: uint(row.MonthlyCosts),
				Request: uint(row.RequestCosts),
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid service model in database: %v", err)
	}

	model.SetID(uint(row.ID))

	return model, nil
}

func (r *PostgreSQLRepository) RegisterService(model *domain.Service) error {
	id, err := r.queries.RegisterService(context.Background(), &queries.RegisterServiceParams{
		OrganizationSerialNumber: model.Organization().SerialNumber(),
		Name:                     model.Name(),
		Internal:                 model.Internal(),
		DocumentationUrl:         model.DocumentationURL(),
		ApiSpecificationType:     string(model.APISpecificationType()),
		PublicSupportContact:     model.PublicSupportContact(),
		TechSupportContact:       model.TechSupportContact(),
		RequestCosts:             int32(model.Costs().Request()),
		MonthlyCosts:             int32(model.Costs().Monthly()),
		OneTimeCosts:             int32(model.Costs().OneTime()),
	})
	if err != nil {
		return err
	}

	model.SetID(uint(id))

	return err
}

func (r *PostgreSQLRepository) ListServices(ctx context.Context, organizationSerialNumber string) ([]*domain.Service, error) {
	rows, err := r.queries.SelectServices(ctx, organizationSerialNumber)
	if err != nil {
		return nil, err
	}

	return convertServiceRowsToModel(r.logger, rows)
}

func convertServiceRowsToModel(logger *zap.Logger, rows []*queries.SelectServicesRow) ([]*domain.Service, error) {
	result := make([]*domain.Service, len(rows))

	for i, row := range rows {
		if len(row.InwayAddresses) != len(row.HealthyStatuses) {
			err := errors.New("length of the inwayadresses does not match healthchecks")
			logger.Error("failed to convert service to domain model", zap.Error(err))

			return nil, err
		}

		organization, err := domain.NewOrganization(row.OrganizationName, row.OrganizationSerialNumber)
		if err != nil {
			return nil, err
		}

		inways := make([]*domain.NewServiceInwayArgs, len(row.InwayAddresses))

		for i, inwayAddress := range row.InwayAddresses {
			inwayArgs := &domain.NewServiceInwayArgs{
				Address: inwayAddress,
				State:   domain.InwayDOWN,
			}

			if row.HealthyStatuses[i] {
				inwayArgs.State = domain.InwayUP
			}

			inways[i] = inwayArgs
		}

		result[i], err = domain.NewService(&domain.NewServiceArgs{
			Name:                 row.Name,
			Organization:         organization,
			Internal:             row.Internal,
			DocumentationURL:     row.DocumentationUrl,
			APISpecificationType: domain.SpecificationType(row.ApiSpecificationType),
			PublicSupportContact: row.PublicSupportContact,
			TechSupportContact:   "",
			Costs: &domain.NewServiceCostsArgs{
				OneTime: uint(row.OneTimeCosts),
				Monthly: uint(row.MonthlyCosts),
				Request: uint(row.RequestCosts),
			},
			Inways: inways,
		})
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
