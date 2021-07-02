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

	selectInwayByAddressStatement   *sqlx.NamedStmt
	setOrganizationInwayStatement   *sqlx.NamedStmt
	clearOrganizationInwayStatement *sqlx.Stmt
}

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
		selectInwayByAddressStatement:   selectInwayByAddressStatement,
		setOrganizationInwayStatement:   setOrganizationInwayStatement,
		clearOrganizationInwayStatement: clearOrganizationInwayStatement,
	}, nil
}
