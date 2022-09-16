// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package monitor

import (
	"time"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/common/strings"
)

type DBConnectionArgs struct {
	DSN                string
	MaxIdleConnections int
	MaxOpenConnections int
	ConnectionTimeout  time.Duration
}

func InitDatabase(d *DBConnectionArgs) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", d.DSN)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(d.ConnectionTimeout)
	db.SetMaxIdleConns(d.MaxIdleConnections)
	db.SetMaxOpenConns(d.MaxOpenConnections)
	db.MapperFunc(strings.ToSnakeCase)

	return db, nil
}
