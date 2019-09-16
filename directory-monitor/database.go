package monitor

import (
	"time"

	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
)

func InitDatabase(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(2)
	db.MapperFunc(xstrings.ToSnakeCase)

	return db, nil
}
