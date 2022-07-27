// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/management-api/migrations"
)

const driverName = "embed"

var ErrNotFound = errors.New("database: value not found")

var registerDriverOnce sync.Once

type PostgresConfigDatabase struct {
	*gorm.DB
	queries *queries.Queries
}

func New(connectionString string) (ConfigDatabase, error) {
	gormDB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	querier, err := queries.Prepare(context.Background(), db)
	if err != nil {
		return nil, err
	}

	return &PostgresConfigDatabase{
		DB:      gormDB,
		queries: querier,
	}, nil
}

func NewWithGorm(gormDB *gorm.DB) (ConfigDatabase, error) {
	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	querier, err := queries.Prepare(context.Background(), db)
	if err != nil {
		return nil, err
	}

	return &PostgresConfigDatabase{
		DB:      gormDB,
		queries: querier,
	}, nil
}

func setupMigrator(dsn string) (*migrate.Migrate, error) {
	registerDriverOnce.Do(func() {
		migrations.RegisterDriver(driverName)
	})

	return migrate.New(fmt.Sprintf("%s://", driverName), dsn)
}

func PostgresPerformMigrations(dsn string) error {
	migrator, err := setupMigrator(dsn)
	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("migrations are up-to-date")
			return nil
		}

		return err
	}

	return nil
}

func PostgresMigrationStatus(dsn string) (version uint, dirty bool, err error) {
	migrator, err := setupMigrator(dsn)
	if err != nil {
		return 0, false, err
	}

	return migrator.Version()
}
