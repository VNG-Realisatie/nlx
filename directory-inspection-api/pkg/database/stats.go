// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"

	"github.com/jmoiron/sqlx"
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

func (db PostgreSQLDirectoryDatabase) ListVersionStatistics(_ context.Context) ([]*VersionStatistics, error) {
	var result []*VersionStatistics

	err := db.selectVersionStatisticsStatement.Select(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
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
