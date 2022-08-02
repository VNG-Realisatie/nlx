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

	inways := make([]*domain.NewServiceAvailability, len(row.InwayAddresses))

	for j, inwayAddress := range row.InwayAddresses {
		inwayArgs := &domain.NewServiceAvailability{
			InwayAddress: inwayAddress,
			State:        domain.InwayDOWN,
		}

		if row.HealthyStatuses[j] {
			inwayArgs.State = domain.InwayUP
		}

		inways[j] = inwayArgs
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
			Availabilities: inways,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid service model in database: %v", err)
	}

	model.SetID(uint(row.ID))

	return model, nil
}

func (r *PostgreSQLRepository) RegisterService(model *domain.Service) error {
	tx, err := r.db.Begin()
	if err != nil {
		return nil
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return
			}
			r.logger.Error("cannot rollback database transaction for register service", zap.Error(err))
		}
	}()

	queriesWithTx := r.queries.WithTx(tx)

	for _, availability := range model.Availabilities() {
		id, err := queriesWithTx.RegisterService(context.Background(), &queries.RegisterServiceParams{
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
			InwayAddress:             availability.InwayAddress(),
		})
		if err != nil {
			return err
		}

		model.SetID(uint(id))
	}

	return tx.Commit()
}

func (r *PostgreSQLRepository) ListServices(ctx context.Context, organizationSerialNumber string) ([]*domain.Service, error) {
	rows, err := r.queries.SelectServices(ctx, organizationSerialNumber)
	if err != nil {
		return nil, err
	}

	return convertServiceRowsToModel(rows)
}

func convertServiceRowsToModel(rows []*queries.SelectServicesRow) ([]*domain.Service, error) {
	result := make([]*domain.Service, len(rows))

	for i, row := range rows {
		organization, err := domain.NewOrganization(row.OrganizationName, row.OrganizationSerialNumber)
		if err != nil {
			return nil, err
		}

		inways := make([]*domain.NewServiceAvailability, len(row.InwayAddresses))

		for j, inwayAddress := range row.InwayAddresses {
			inwayArgs := &domain.NewServiceAvailability{
				InwayAddress: inwayAddress,
				State:        domain.InwayDOWN,
			}

			if row.HealthyStatuses[j] {
				inwayArgs.State = domain.InwayUP
			}

			inways[j] = inwayArgs
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
			Availabilities: inways,
		})
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
