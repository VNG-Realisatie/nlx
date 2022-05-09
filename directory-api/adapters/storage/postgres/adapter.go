// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/directory-api/migrations"
)

const driverName = "embed"

var registerDriverOnce sync.Once

type PostgreSQLRepository struct {
	logger  *zap.Logger
	db      *sqlx.DB
	queries *queries.Queries
}

//nolint gocyclo: all checks in this function are necessary
func New(logger *zap.Logger, db *sqlx.DB) (*PostgreSQLRepository, error) {
	if logger == nil {
		return nil, errors.New("no logger provided")
	}

	if db == nil {
		panic("missing db")
	}

	querier, err := queries.Prepare(context.Background(), db)
	if err != nil {
		return nil, err
	}

	return &PostgreSQLRepository{
		logger:  logger.Named("postgres repository"),
		db:      db,
		queries: querier,
	}, nil
}

func NewPostgreSQLConnection(dsn string) (*sqlx.DB, error) {
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
	db.MapperFunc(xstrings.ToSnakeCase)

	return db, nil
}

func (r *PostgreSQLRepository) Shutdown() error {
	return r.db.Close()
}

func setupMigrator(dsn string) (*migrate.Migrate, error) {
	registerDriverOnce.Do(func() {
		migrations.RegisterDriver(driverName)
	})

	return migrate.New(fmt.Sprintf("%s://", driverName), dsn)
}

func PostgreSQLPerformMigrations(dsn string) error {
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

func PostgresMigrationStatus(dsn string) (version uint, dirty bool, err error) {
	migrator, err := setupMigrator(dsn)
	if err != nil {
		return 0, false, err
	}

	return migrator.Version()
}
