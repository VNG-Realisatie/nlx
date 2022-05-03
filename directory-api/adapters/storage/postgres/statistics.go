// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"fmt"

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
