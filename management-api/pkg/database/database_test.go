// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package database_test

import (
	"os"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"go.nlx.io/nlx/testing/testingutils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.nlx.io/nlx/management-api/pkg/database"
)

var setupOnce sync.Once
var fixtureMutex sync.Mutex

func setup(t *testing.T) {
	setupOnce.Do(func() {
		setupPostgreSQL(t)
	})
}

func setupPostgreSQL(t *testing.T) {
	dsnBase := os.Getenv("POSTGRES_DSN")
	dsn, err := testingutils.CreateTestDatabase(dsnBase, "test_management_api")
	if err != nil {
		t.Fatal(err)
	}

	// Necessary to prevent migration version collision with directory database migrations
	dsnForMigrations := testingutils.AddQueryParamToAddress(dsn, "x-migrations-table", "management_migrations")
	err = database.PostgresPerformMigrations(dsnForMigrations)
	if err != nil {
		t.Fatal(err)
	}

	txdb.Register("txdb", "postgres", dsn)
}

func newPostgresConfigDatabase(t *testing.T, id string, loadFixtures bool) (database.ConfigDatabase, func() error) {
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

	if loadFixtures {
		fixtures, err := testfixtures.New(
			testfixtures.Database(db),
			testfixtures.Dialect("postgres"),
			testfixtures.Directory("testdata/fixtures/postgres"),
			testfixtures.DangerousSkipTestDatabaseCheck(),
		)

		fixtureMutex.Lock()

		err = fixtures.Load()

		fixtureMutex.Unlock()

		require.NoError(t, err)
	}

	return &database.PostgresConfigDatabase{
		DB: gormDB,
	}, db.Close
}

func newConfigDatabase(t *testing.T, id string, loadFixtures bool) (database.ConfigDatabase, func() error) {
	return newPostgresConfigDatabase(t, id, loadFixtures)
}
