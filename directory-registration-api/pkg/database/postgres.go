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

	upsertInwayStmt                 *sqlx.NamedStmt
	upsertServiceStmt               *sqlx.NamedStmt
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

	upsertInwayStmt, err := prepareUpsertInwayStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare upsert inway statement: %s", err)
	}

	upsertServiceStmt, err := prepareUpsertServiceStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare upsert service statement: %s", err)
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
		logger:                          logger,
		db:                              db,
		upsertInwayStmt:                 upsertInwayStmt,
		upsertServiceStmt:               upsertServiceStmt,
		selectInwayByAddressStatement:   selectInwayByAddressStatement,
		setOrganizationInwayStatement:   setOrganizationInwayStatement,
		clearOrganizationInwayStatement: clearOrganizationInwayStatement,
	}, nil
}
