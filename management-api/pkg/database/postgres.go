// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("database: value not found")

// PostgresConfigDatabase is the postgres implementation of ConfigDatabase
type PostgresConfigDatabase struct {
	*gorm.DB
}

// NewPostgresConfigDatabase constructs a new PostgresDatabase
func NewPostgresConfigDatabase(connectionString string) (ConfigDatabase, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresConfigDatabase{
		DB: db,
	}, nil
}
