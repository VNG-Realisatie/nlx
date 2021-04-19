// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package database

import (
	"context"
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
