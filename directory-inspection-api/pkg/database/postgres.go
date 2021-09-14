// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"fmt"
	"time"

	"github.com/go-errors/errors"
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/directory-db/dbversion"
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
func NewPostgreSQLDirectoryDatabase(dsn string, logger *zap.Logger) (DirectoryDatabase, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Errorf("could not open connection to postgres: %s", err)
	}

	const (
		FiveMinutes        = 5 * time.Minute
		MaxIdleConnections = 2
	)

	db.SetConnMaxLifetime(FiveMinutes)
	db.SetMaxIdleConns(MaxIdleConnections)
	db.MapperFunc(xstrings.ToSnakeCase)

	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

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
		logger:                                  logger,
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

func prepareSelectServicesStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	selectServicesStatement, err := db.Preparex(`
		SELECT
			o.name AS organization_name,
			s.name AS service_name,
			s.internal as service_internal,
			s.one_time_costs as one_time_costs,
			s.monthly_costs as monthly_costs, 
			s.request_costs as request_costs,
			array_remove(array_agg(i.address), NULL) AS inway_addresses,
			COALESCE(s.documentation_url, '') AS documentation_url,
			COALESCE(s.api_specification_type, '') AS api_specification_type,
			COALESCE(s.public_support_contact, '') AS public_support_contact,
			array_remove(array_agg(a.healthy), NULL) as healthy_statuses
		FROM directory.services s
		INNER JOIN directory.availabilities a ON a.service_id = s.id
		INNER JOIN directory.organizations o ON o.id = s.organization_id
		INNER JOIN directory.inways i ON i.id = a.inway_id
		WHERE (
			internal = false
			OR (
				internal = true
				AND o.name = $1
			)
		)
		GROUP BY s.id, o.id
		ORDER BY o.name, s.name
	`)
	if err != nil {
		return nil, err
	}

	return selectServicesStatement, nil
}

func prepareRegisterOutwayStatement(db *sqlx.DB) (*sqlx.NamedStmt, error) {
	registerOutwayStatement, err := db.PrepareNamed(`
		INSERT INTO directory.outways (version)
		VALUES (:version)
	`)
	if err != nil {
		return nil, err
	}

	return registerOutwayStatement, nil
}

func prepareSelectOrganizationsStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	listOrganizationsStatement, err := db.Preparex(`
		SELECT
			name
		FROM directory.organizations
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}

	return listOrganizationsStatement, nil
}

func prepareSelectVersionStatisticsStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	// All the outways announcements for the last day (24 hours) are fetched and counted per version,
	// the inways are updated per organization so they have no time constraint.
	selectVersionStatisticsStatement, err := db.Preparex(`
		SELECT 'outway' AS type
		,      version
		,      COUNT(*) AS amount
		FROM   directory.outways
		WHERE  announced > now() - interval '1 days'
		GROUP BY version
		UNION
		SELECT 'inway' AS type
		,      version
		,      COUNT(*) AS amount
		FROM   directory.inways
		GROUP BY version
		ORDER BY type, version DESC
	`)

	if err != nil {
		return nil, err
	}

	return selectVersionStatisticsStatement, nil
}
