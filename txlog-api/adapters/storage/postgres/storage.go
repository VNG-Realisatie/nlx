// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package postgresadapter

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver

	"go.nlx.io/nlx/txlog-api/adapters/storage/postgres/migrations"
	"go.nlx.io/nlx/txlog-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/txlog-api/domain/record"
)

const driverName = "embed"

var registerDriverOnce sync.Once

type PostgreSQLRepository struct {
	db      *sqlx.DB
	queries *queries.Queries
}

func New(db *sqlx.DB) (record.Repository, error) {
	if db == nil {
		panic("missing db")
	}

	querier, err := queries.Prepare(context.Background(), db)
	if err != nil {
		return nil, err
	}

	return &PostgreSQLRepository{
		db:      db,
		queries: querier,
	}, nil
}

func NewConnection(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open connection to postgres: %s", err)
	}

	const (
		FiveMinutes        = 5 * time.Minute
		MaxIdleConnections = 2
	)

	db.SetConnMaxLifetime(FiveMinutes)
	db.SetMaxIdleConns(MaxIdleConnections)

	return db, nil
}

func (r *PostgreSQLRepository) Shutdown() error {
	err := r.queries.Close()
	if err != nil {
		return err
	}

	return r.db.Close()
}

func setupMigrator(dsn string) (*migrate.Migrate, error) {
	registerDriverOnce.Do(func() {
		migrations.RegisterDriver(driverName)
	})

	return migrate.New(fmt.Sprintf("%s://", driverName), dsn)
}

func PerformMigrations(dsn string) error {
	migrator, err := setupMigrator(dsn)
	if err != nil {
		return err
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("running migrations: %v", err)
	}

	return nil
}

func MigrationStatus(dsn string) (version uint, dirty bool, err error) {
	migrator, err := setupMigrator(dsn)
	if err != nil {
		return 0, false, err
	}

	return migrator.Version()
}
