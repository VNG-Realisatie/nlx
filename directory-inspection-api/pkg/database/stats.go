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

// ListVersionStatistics returns the statistics for every outway version
func (db PostgreSQLDirectoryDatabase) ListVersionStatistics(ctx context.Context) ([]*VersionStatistics, error) {
	var result []*VersionStatistics

	err := db.selectVersionStatisticsStatement.Select(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
