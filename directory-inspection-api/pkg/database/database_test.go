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
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

var setupOnce sync.Once

func setup(t *testing.T) {
	setupOnce.Do(func() {
		setupPostgreSQL(t)
	})
}

func setupPostgreSQL(t *testing.T) {
	dsn := os.Getenv("POSTGRES_DSN")

	err := database.PostgreSQLPerformMigrations(dsn)
	require.NoError(t, err)

	txdb.Register("txdb", "postgres", dsn)

	// This is necessary because the default BindVars for txdb isn't correct
	sqlx.BindDriver("txdb", sqlx.DOLLAR)
}

func newPostgresDirectoryDatabase(t *testing.T, id string, loadFixtures bool) (database.DirectoryDatabase, func() error) {
	db, err := sqlx.Open("txdb", id)
	require.NoError(t, err)

	if loadFixtures {
		fixtures, err := testfixtures.New(
			testfixtures.Database(db.DB),
			testfixtures.Dialect("postgres"),
			testfixtures.Directory("testdata/fixtures/postgres"),
			testfixtures.DangerousSkipTestDatabaseCheck(),
		)

		err = fixtures.Load()
		require.NoError(t, err)
	}

	db.MapperFunc(xstrings.ToSnakeCase)

	repo, err := database.NewPostgreSQLDirectoryDatabase(db)
	require.NoError(t, err)

	return repo, db.Close
}

func newDirectoryDatabase(t *testing.T, id string, loadFixtures bool) (database.DirectoryDatabase, func() error) {
	return newPostgresDirectoryDatabase(t, id, loadFixtures)
}
