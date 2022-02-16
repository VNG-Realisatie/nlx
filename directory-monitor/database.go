package monitor

import (
	"time"

	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
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
	db.MapperFunc(xstrings.ToSnakeCase)

	return db, nil
}
