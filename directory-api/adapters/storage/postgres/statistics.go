// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"fmt"

	"go.nlx.io/nlx/directory-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) ListVersionStatistics(ctx context.Context) ([]*domain.VersionStatistics, error) {
	rows, err := r.queries.SelectVersionStatistics(ctx)
	if err != nil {
		return nil, err
	}

	statistics := make([]*domain.VersionStatistics, len(rows))
	for i, row := range rows {
		statistics[i], err = domain.NewVersionStatistics(&domain.NewVersionStatisticsArgs{
			GatewayType: domain.VersionStatisticsType(fmt.Sprintf("%s", row.Type)),
			Version:     row.Version,
			Amount:      uint32(row.Amount),
		})
		if err != nil {
			return nil, fmt.Errorf("invalid version statistics model in database: %v", err)
		}
	}

	return statistics, nil
}

func (r *PostgreSQLRepository) ListParticipants(ctx context.Context) ([]*domain.Participant, error) {
	rows, err := r.queries.SelectParticipants(ctx)
	if err != nil {
		return nil, err
	}

	return convertParticipantRowsToModels(rows)
}

func convertParticipantRowsToModels(rows []*queries.SelectParticipantsRow) ([]*domain.Participant, error) {
	result := make([]*domain.Participant, len(rows))

	for i, row := range rows {
		org, err := domain.NewOrganization(row.Name, row.SerialNumber)
		if err != nil {
			return nil, err
		}

		p, err := domain.NewParticipant(&domain.NewParticipantArgs{
			Organization: org,
			Statistics: &domain.NewParticipantStatisticsArgs{
				Inways:   uint(row.Inways),
				Outways:  uint(row.Outways),
				Services: uint(row.Services),
			},
			CreatedAt: row.CreatedAt,
		})
		if err != nil {
			return nil, err
		}

		result[i] = p
	}

	return result, nil
}
