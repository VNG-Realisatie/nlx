// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"go.uber.org/zap"
)

// WaitForLatestDBVersion blocks until the database is migrated to the latest version.
// This function can exit the program when failing to query postgres for the version.
func WaitForLatestDBVersion(logger *zap.Logger, db *sql.DB, requiredVersion int) {
	dbMigrateDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatal("could not open driver to postgres connection", zap.Error(err))
	}
	for {
		currentVersion, dirty, err := dbMigrateDriver.Version()
		if err != nil && err != migrate.ErrNilVersion {
			logger.Fatal("could not obtain db version", zap.Error(err))
		}
		if dirty {
			logger.Info(fmt.Sprintf("db migration is dirty at version %d, waiting...", currentVersion))
			time.Sleep(1 * time.Second)
			continue
		}
		if currentVersion != requiredVersion {
			logger.Info(fmt.Sprintf("db is at version %d, require version %d", currentVersion, requiredVersion))
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
}
