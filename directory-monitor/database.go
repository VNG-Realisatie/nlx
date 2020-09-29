package monitor

import (
	"time"

	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
)

const (
	dbConnectionMaxLifetime = 5 * time.Minute
	dbMaxIdleConnections    = 2
)

func InitDatabase(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(dbConnectionMaxLifetime)
	db.SetMaxIdleConns(dbMaxIdleConnections)
	db.MapperFunc(xstrings.ToSnakeCase)

	return db, nil
}
