// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"fmt"
	"time"

	"github.com/go-errors/errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"go.uber.org/zap"
)

// PostgreSQLDirectoryDatabase is the PostgreSQL implementation of the DirectoryDatabase
type PostgreSQLDirectoryDatabase struct {
	logger *zap.Logger
	db     *sqlx.DB

	selectServicesStatement                 *sqlx.Stmt
	registerOutwayStatement                 *sqlx.NamedStmt
	selectOrganizationsStatement            *sqlx.Stmt
	selectOrganizationInwayAddressStatement *sqlx.NamedStmt
	selectVersionStatisticsStatement        *sqlx.Stmt
}

// NewPostgreSQLDirectoryDatabase constructs a new PostgreSQLDirectoryDatabase
func NewPostgreSQLDirectoryDatabase(db *sqlx.DB) (DirectoryDatabase, error) {
	if db == nil {
		panic("missing db")
	}

	selectServicesStatement, err := prepareSelectServicesStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create select services prepared statement: %s", err)
	}

	registerOutwayStatement, err := prepareRegisterOutwayStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create register outway prepared statement: %s", err)
	}

	selectOrganizationsStatement, err := prepareSelectOrganizationsStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create select organizations prepared statement: %s", err)
	}

	selectOrganizationInwayAddressStatement, err := prepareSelectOrganizationInwayAddressStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create select organization inway prepared statement: %s", err)
	}

	selectVersionStatisticsStatement, err := prepareSelectVersionStatisticsStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create select version statistics prepared statement: %s", err)
	}

	return &PostgreSQLDirectoryDatabase{
		db:                                      db,
		selectServicesStatement:                 selectServicesStatement,
		registerOutwayStatement:                 registerOutwayStatement,
		selectOrganizationsStatement:            selectOrganizationsStatement,
		selectOrganizationInwayAddressStatement: selectOrganizationInwayAddressStatement,
		selectVersionStatisticsStatement:        selectVersionStatisticsStatement,
	}, nil
}

func (db *PostgreSQLDirectoryDatabase) Shutdown() error {
	return db.db.Close()
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

func PostgreSQLPerformMigrations(dsn string) error {
	migrator, err := migrate.New("file://../../../directory-db/migrations", dsn)
	if err != nil {
		return fmt.Errorf("setup migrator: %v", err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("running migrations: %v", err)
	}

	return nil
}
