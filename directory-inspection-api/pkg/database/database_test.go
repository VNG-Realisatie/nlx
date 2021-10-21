// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
	"go.nlx.io/nlx/testing/testingutils"
)

var setupOnce sync.Once

const fixtureSuffix = "_fixtures"

func setup(t *testing.T) {
	setupOnce.Do(func() {
		setupPostgreSQL(t)
	})
}

func setupPostgreSQL(t *testing.T) {
	setupDatabase(t, false) // Without fixtures
	setupDatabase(t, true)  // With fixtures
}

func setupDatabase(t *testing.T, loadFixtures bool) {
	dbName := getDBName(loadFixtures)

	dsnBase := os.Getenv("POSTGRES_DSN")
	dsn, err := testingutils.CreateTestDatabase(dsnBase, dbName)
	if err != nil {
		t.Fatal(err)
	}

	dsnForMigrations := testingutils.AddQueryParamToAddress(dsn, "x-migrations-table", dbName)
	err = database.PostgreSQLPerformMigrations(dsnForMigrations)
	if err != nil {
		t.Fatal(err)
	}

	dbDriver := getDriverName(loadFixtures)
	txdb.Register(dbDriver, "postgres", dsn)

	// This is necessary because the default BindVars for txdb isn't correct
	sqlx.BindDriver(dbDriver, sqlx.DOLLAR)

	if loadFixtures {
		db, err := sqlx.Open("postgres", dsn)
		require.NoError(t, err)

		fixtures, err := testfixtures.New(
			testfixtures.Database(db.DB),
			testfixtures.Dialect("postgres"),
			testfixtures.Directory("testdata/fixtures/postgres"),
			testfixtures.DangerousSkipTestDatabaseCheck(),
		)

		err = fixtures.Load()
		require.NoError(t, err)
	}
}

func newPostgresDirectoryDatabase(t *testing.T, id string, loadFixtures bool) (database.DirectoryDatabase, func() error) {

	db, err := sqlx.Open(getDriverName(loadFixtures), id)
	require.NoError(t, err)

	db.MapperFunc(xstrings.ToSnakeCase)

	repo, err := database.NewPostgreSQLDirectoryDatabase(db)
	require.NoError(t, err)

	return repo, db.Close
}

func newDirectoryDatabase(t *testing.T, id string, loadFixtures bool) (database.DirectoryDatabase, func() error) {
	return newPostgresDirectoryDatabase(t, id, loadFixtures)
}

func getDriverName(loadFixtures bool) string {
	var suffix string
	if loadFixtures {
		suffix = fixtureSuffix
	}

	return fmt.Sprintf("txdb%s", suffix)
}

func getDBName(loadFixtures bool) string {
	var suffix string
	if loadFixtures {
		suffix = fixtureSuffix
	}

	return fmt.Sprintf("test_direction_inspection%s", suffix)
}
