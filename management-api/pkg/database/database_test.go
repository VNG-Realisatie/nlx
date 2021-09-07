// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package database_test

import (
	"os"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.nlx.io/nlx/management-api/pkg/database"
)

var setupOnce sync.Once

func setup(t *testing.T) {
	setupOnce.Do(func() {
		setupPostgreSQL(t)
	})
}

func setupPostgreSQL(t *testing.T) {
	dsn := os.Getenv("POSTGRES_DSN")

	err := database.PostgresPerformMigrations(dsn)
	require.NoError(t, err)

	txdb.Register("txdb", "postgres", dsn)
}

func newPostgresConfigDatabase(t *testing.T, id string) (database.ConfigDatabase, func() error) {
	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			DriverName: "txdb",
			DSN:        id,
		}),
		&gorm.Config{},
	)
	require.NoError(t, err)

	db, err := gormDB.DB()
	require.NoError(t, err)

	return &database.PostgresConfigDatabase{
		DB: gormDB,
	}, db.Close
}

func newConfigDatabase(t *testing.T, id string) (database.ConfigDatabase, func() error) {
	return newPostgresConfigDatabase(t, id)
}
