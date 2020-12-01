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
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-db/dbversion"
)

// PostgreSQLDirectoryDatabase is the PostgreSQL implementation of the DirectoryDatabase
type PostgreSQLDirectoryDatabase struct {
	logger *zap.Logger
	db     *sqlx.DB

	setInsightConfigurationStatement *sqlx.Stmt
	insertAvailabilityStatement      *sqlx.Stmt

	selectInwayByAddressStatement   *sqlx.NamedStmt
	setOrganizationInwayStatement   *sqlx.NamedStmt
	clearOrganizationInwayStatement *sqlx.Stmt
}

// NewPostgreSQLDirectoryDatabase constructs a new PostgreSQLDirectoryDatabase
func NewPostgreSQLDirectoryDatabase(dsn string, p *process.Process, logger *zap.Logger) (DirectoryDatabase, error) {
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

	p.CloseGracefully(db.Close)

	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

	setInsightConfigurationStatement, err := prepareSetInsightConfigurationStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare set insight configuration statement: %s", err)
	}

	insertAvailabilityStatement, err := prepareInsertAvailabilityStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare insert availability statement: %s", err)
	}

	selectInwayByAddressStatement, err := prepareSelectInwayByAddressStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select inway statement: %s", err)
	}

	setOrganizationInwayStatement, err := prepareSetOrganizationInwayStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select inway statement: %s", err)
	}

	clearOrganizationInwayStatement, err := prepareClearOrganizationInwayStatement(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select inway statement: %s", err)
	}

	return &PostgreSQLDirectoryDatabase{
		logger:                           logger,
		db:                               db,
		setInsightConfigurationStatement: setInsightConfigurationStatement,
		insertAvailabilityStatement:      insertAvailabilityStatement,
		selectInwayByAddressStatement:    selectInwayByAddressStatement,
		setOrganizationInwayStatement:    setOrganizationInwayStatement,
		clearOrganizationInwayStatement:  clearOrganizationInwayStatement,
	}, nil
}

func prepareSetInsightConfigurationStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	// NOTE: We do not have an endpoint yet to create services separately, therefore insert on demand.
	prepareSetInsightConfigurationStatement, err := db.Preparex(`
		INSERT INTO directory.organizations (name, insight_log_endpoint, insight_irma_endpoint)
			VALUES ($1, $2, $3)
			ON CONFLICT ON CONSTRAINT organizations_uq_name
				DO UPDATE SET
					insight_log_endpoint = NULLIF(EXCLUDED.insight_log_endpoint, ''),
					insight_irma_endpoint = NULLIF(EXCLUDED.insight_irma_endpoint, '')
			RETURNING id
	`)
	if err != nil {
		return nil, err
	}

	return prepareSetInsightConfigurationStatement, nil
}

func prepareInsertAvailabilityStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	// NOTE: We do not have an endpoint yet to create services separately, therefore insert on demand.
	prepareInsertAvailabilityStatement, err := db.Preparex(`
		WITH org AS (
			INSERT INTO directory.organizations (name, insight_log_endpoint, insight_irma_endpoint)
				VALUES ($1, $7, $8)
				ON CONFLICT ON CONSTRAINT organizations_uq_name
					DO UPDATE SET
						insight_log_endpoint = NULLIF(EXCLUDED.insight_log_endpoint, ''),
						insight_irma_endpoint = NULLIF(EXCLUDED.insight_irma_endpoint, '')
				RETURNING id
		), service AS (
			INSERT INTO directory.services (organization_id, name, internal, documentation_url, api_specification_type, public_support_contact, tech_support_contact)
				SELECT org.id, $2, $3, NULLIF($4, ''), NULLIF($5, ''), NULLIF($9, ''), NULLIF($10, '')
					FROM org
				ON CONFLICT ON CONSTRAINT services_uq_name
					DO UPDATE SET
						internal = EXCLUDED.internal,
						documentation_url = EXCLUDED.documentation_url,-- (possibly) no-op update to return id
						api_specification_type = EXCLUDED.api_specification_type,
						public_support_contact = EXCLUDED.public_support_contact,
						tech_support_contact = EXCLUDED.tech_support_contact
					RETURNING id
		), inway AS (
			INSERT INTO directory.inways (organization_id, address, version)
				SELECT org.id, $6, NULLIF($11, '')
					FROM org
				ON CONFLICT ON CONSTRAINT inways_uq_address
					DO UPDATE SET address = EXCLUDED.address -- no-op update to return id
				RETURNING id
		)
		INSERT INTO directory.availabilities (inway_id, service_id, last_announced)
			SELECT inway.id, service.id, NOW()
				FROM inway, service
			ON CONFLICT ON CONSTRAINT availabilities_uq_inway_service DO UPDATE
				SET last_announced = NOW(), active = true
	`)
	if err != nil {
		return nil, err
	}

	return prepareInsertAvailabilityStatement, nil
}
