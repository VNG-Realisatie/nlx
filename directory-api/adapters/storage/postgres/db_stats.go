// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

type VersionStatistics struct {
	Type    VersionStatisticsType
	Version string
	Amount  uint32
}

type VersionStatisticsType string

const (
	TypeInway  VersionStatisticsType = "inway"
	TypeOutway VersionStatisticsType = "outway"
)

func (r *PostgreSQLRepository) ListVersionStatistics(_ context.Context) ([]*domain.VersionStatistics, error) {
	var result []*VersionStatistics

	err := r.selectVersionStatisticsStmt.Select(&result)
	if err != nil {
		return nil, err
	}

	statistics := make([]*domain.VersionStatistics, len(result))
	for i, s := range result {
		statistics[i], err = domain.NewVersionStatistics(&domain.NewVersionStatisticsArgs{
			GatewayType: domain.VersionStatisticsType(s.Type),
			Version:     s.Version,
			Amount:      s.Amount,
		})
		if err != nil {
			return nil, fmt.Errorf("invalid version statistics model in database: %v", err)
		}
	}

	return statistics, nil
}

func prepareSelectVersionStatisticsStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	// All the outways announcements for the last day (24 hours) are fetched and counted per version,
	// the inways are updated per organization so they have no time constraint.
	selectVersionStatisticsStatement, err := db.Preparex(`
		SELECT 'outway' AS type
		,      version
		,      COUNT(*) AS amount
		FROM   directory.outways
		WHERE  announced > now() - interval '1 days'
		GROUP BY version
		UNION
		SELECT 'inway' AS type
		,      version
		,      COUNT(*) AS amount
		FROM   directory.inways
		GROUP BY version
		ORDER BY type, version DESC
	`)

	if err != nil {
		return nil, err
	}

	return selectVersionStatisticsStatement, nil
}
