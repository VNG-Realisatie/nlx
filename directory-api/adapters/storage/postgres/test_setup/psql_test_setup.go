// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package pgadapter_test_setup

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
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	pgadapter "go.nlx.io/nlx/directory-api/adapters/storage/postgres"
	"go.nlx.io/nlx/testing/testingutils"
)

const fixtureSuffix = "_fixtures"

var setupOnceFixtures sync.Once
var setupOnce sync.Once

func setupDatabase(t *testing.T, loadFixtures bool) {
	dbName := getDBName(loadFixtures)

	dsnBase := os.Getenv("POSTGRES_DSN")
	dsn, err := testingutils.CreateTestDatabase(dsnBase, dbName)
	if err != nil {
		t.Fatal(err)
	}

	dsnForMigrations := testingutils.AddQueryParamToAddress(dsn, "x-migrations-table", dbName)
	err = pgadapter.PostgreSQLPerformMigrations(dsnForMigrations)
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
			testfixtures.Directory("../../../adapters/storage/postgres/test_setup/testdata/fixtures"),
			testfixtures.DangerousSkipTestDatabaseCheck(),
		)

		err = fixtures.Load()
		require.NoError(t, err)
	}
}

func New(t *testing.T, loadFixtures bool) (*pgadapter.PostgreSQLRepository, func() error) {
	if loadFixtures {
		setupOnceFixtures.Do(func() {
			setupDatabase(t, true)
		})
	} else {
		setupOnce.Do(func() {
			setupDatabase(t, false)
		})
	}

	db, err := sqlx.Open(getDriverName(loadFixtures), t.Name())
	require.NoError(t, err)

	db.MapperFunc(xstrings.ToSnakeCase)

	repo, err := pgadapter.New(zap.NewNop(), db)
	require.NoError(t, err)

	return repo, db.Close
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

	return fmt.Sprintf("test_directory%s", suffix)
}
