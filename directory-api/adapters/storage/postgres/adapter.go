// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-api/adapters/storage/postgres/queries"
)

type PostgreSQLRepository struct {
	logger                                               *zap.Logger
	db                                                   *sqlx.DB
	queries                                              *queries.Queries
	registerInwayStmt                                    *sqlx.NamedStmt
	registerServiceStmt                                  *sqlx.NamedStmt
	getServiceStmt                                       *sqlx.NamedStmt
	selectInwayByAddressStmt                             *sqlx.NamedStmt
	setOrganizationInwayStmt                             *sqlx.NamedStmt
	selectOrganizationInwayAddressStmt                   *sqlx.NamedStmt
	selectOrganizationInwayManagementAPIProxyAddressStmt *sqlx.NamedStmt
	setOrganizationEmailAddressStmt                      *sqlx.NamedStmt
	selectVersionStatisticsStmt                          *sqlx.Stmt
	selectServicesStmt                                   *sqlx.Stmt
	selectOrganizationsStmt                              *sqlx.Stmt
	registerOutwayStmt                                   *sqlx.NamedStmt
	getOutwayStmt                                        *sqlx.NamedStmt
	selectParticipantsStmt                               *sqlx.Stmt
}

//nolint gocyclo: all checks in this function are necessary
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

	registerInwayStmt, err := prepareRegisterInwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare register inway statement: %s", err)
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

	selectOrganizationInwayAddressStmt, err := prepareSelectOrganizationInwayAddressStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select organization inway address statement: %s", err)
	}

	selectOrganizationInwayManagementAPIProxyAddressStmt, err := prepareSelectOrganizationInwayManagementAPIProxyAddressStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select organization inway management api proxy address statement: %s", err)
	}

	selectVersionStatisticsStmt, err := prepareSelectVersionStatisticsStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create select version statistics prepared statement: %s", err)
	}

	selectServicesStatement, err := prepareSelectServicesStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select services statement: %s", err)
	}

	selectOrganizationsStmt, err := prepareSelectOrganizationsStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select organizations statement: %s", err)
	}

	registerOutwayStmt, err := prepareRegisterOutwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare register outway statement: %s", err)
	}

	getOutwayStmt, err := prepareGetOutwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get outway statement: %s", err)
	}

	selectParticipantsStmt, err := prepareSelectParticipantsStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select participants statement: %s", err)
	}

	setOrganizationEmailAddressStmt, err := prepareSetOrganizationEmailStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare set organization email statement: %s", err)
	}

	return &PostgreSQLRepository{
		logger:                             logger.Named("postgres repository"),
		db:                                 db,
		queries:                            querier,
		registerInwayStmt:                  registerInwayStmt,
		registerServiceStmt:                registerServiceStmt,
		getServiceStmt:                     getServiceStmt,
		selectInwayByAddressStmt:           selectInwayByAddressStmt,
		setOrganizationInwayStmt:           setOrganizationInwayStmt,
		setOrganizationEmailAddressStmt:    setOrganizationEmailAddressStmt,
		selectOrganizationInwayAddressStmt: selectOrganizationInwayAddressStmt,
		selectOrganizationInwayManagementAPIProxyAddressStmt: selectOrganizationInwayManagementAPIProxyAddressStmt,
		selectVersionStatisticsStmt:                          selectVersionStatisticsStmt,
		selectServicesStmt:                                   selectServicesStatement,
		selectOrganizationsStmt:                              selectOrganizationsStmt,
		registerOutwayStmt:                                   registerOutwayStmt,
		getOutwayStmt:                                        getOutwayStmt,
		selectParticipantsStmt:                               selectParticipantsStmt,
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

func PostgreSQLPerformMigrations(dsn string) error {
	migrator, err := migrate.New("file://../../../../directory-db/migrations", dsn)
	if err != nil {
		return fmt.Errorf("setup migrator: %v", err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("running migrations: %v", err)
	}

	return nil
}
