// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package adapters

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"go.uber.org/zap"
)

var (
	ErrDuplicateAddress     = errors.New("another inway is already registered with this address")
	ErrNoInwayWithAddress   = errors.New("no inway found for address")
	ErrOrganizationNotFound = errors.New("no organization found")
)

type PostgreSQLRepository struct {
	logger                             *zap.Logger
	db                                 *sqlx.DB
	registerInwayStmt                  *sqlx.NamedStmt
	getInwayStmt                       *sqlx.NamedStmt
	registerServiceStmt                *sqlx.NamedStmt
	getServiceStmt                     *sqlx.NamedStmt
	selectInwayByAddressStmt           *sqlx.NamedStmt
	setOrganizationInwayStmt           *sqlx.NamedStmt
	clearOrganizationInwayStmt         *sqlx.NamedStmt
	selectOrganizationInwayAddressStmt *sqlx.NamedStmt
}

//nolint gocyclo: all checks in this function are necessary
func New(logger *zap.Logger, db *sqlx.DB) (*PostgreSQLRepository, error) {
	if logger == nil {
		panic("missing logger")
	}

	if db == nil {
		panic("missing db")
	}

	registerInwayStmt, err := prepareRegisterInwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare register inway statement: %s", err)
	}

	getInwayStmt, err := prepareGetInwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get inway statement: %s", err)
	}

	registerServiceStmt, err := prepareRegisterServiceStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare register service statement: %s", err)
	}

	getServiceStmt, err := prepareGetServiceStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get service statement: %s", err)
	}

	selectInwayByAddressStmt, err := prepareSelectInwayByAddressStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select inway by address statement: %s", err)
	}

	setOrganizationInwayStmt, err := prepareSetOrganizationInwayStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare set organization inway statement: %s", err)
	}

	clearOrganizationInwayStmt, err := prepareClearOrganizationInwayStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare clear organization inway statement: %s", err)
	}

	selectOrganizationInwayAddressStmt, err := prepareSelectOrganizationInwayAddressStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select organization inway address statement: %s", err)
	}

	return &PostgreSQLRepository{
		logger:                             logger.Named("postgres repository"),
		db:                                 db,
		registerInwayStmt:                  registerInwayStmt,
		getInwayStmt:                       getInwayStmt,
		registerServiceStmt:                registerServiceStmt,
		getServiceStmt:                     getServiceStmt,
		selectInwayByAddressStmt:           selectInwayByAddressStmt,
		setOrganizationInwayStmt:           setOrganizationInwayStmt,
		clearOrganizationInwayStmt:         clearOrganizationInwayStmt,
		selectOrganizationInwayAddressStmt: selectOrganizationInwayAddressStmt,
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

func PostgreSQLPerformMigrations(dsn string) error {
	migrator, err := migrate.New("file://../../directory-db/migrations", dsn)
	if err != nil {
		return fmt.Errorf("setup migrator: %v", err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("running migrations: %v", err)
	}

	return nil
}
