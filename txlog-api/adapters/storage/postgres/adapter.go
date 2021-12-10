// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"go.nlx.io/nlx/txlog-api/adapters/storage/postgres/queries"
	"go.uber.org/zap"
)

type PostgreSQLRepository struct {
	logger  *zap.Logger
	db      *sqlx.DB
	queries *queries.Queries
}

func New(logger *zap.Logger, db *sqlx.DB) (*PostgreSQLRepository, error) {
	if logger == nil {
		panic("missing logger")
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

	return db, nil
}

func (r *PostgreSQLRepository) Shutdown() error {
	err := r.queries.Close()
	if err != nil {
		return err
	}

	return r.db.Close()
}

func PostgreSQLPerformMigrations(dsn string) error {
	migrator, err := migrate.New("file://../../../../txlog-db/migrations", dsn)
	if err != nil {
		return fmt.Errorf("setup migrator: %v", err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("running migrations: %v", err)
	}

	return nil
}
