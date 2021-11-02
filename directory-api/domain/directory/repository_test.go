// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package directory_test

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	pgadapter "go.nlx.io/nlx/directory-api/adapters/postgres"
	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory"
	"go.nlx.io/nlx/testing/testingutils"
)

var setupOnce sync.Once

const fixtureSuffix = "_fixtures"

func setup(t *testing.T) {
	setupOnce.Do(func() {
		setupPostgreSQLRepository(t)
	})
}

func setupPostgreSQLRepository(t *testing.T) {
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
			testfixtures.Directory("../../adapters/postgres/testdata/fixtures"),
			testfixtures.DangerousSkipTestDatabaseCheck(),
		)

		err = fixtures.Load()
		require.NoError(t, err)
	}
}

func newPostgreSQLRepository(t *testing.T, id string, loadFixtures bool) (*pgadapter.PostgreSQLRepository, func() error) {
	db, err := sqlx.Open(getDriverName(loadFixtures), id)
	require.NoError(t, err)

	db.MapperFunc(xstrings.ToSnakeCase)

	repo, err := pgadapter.New(zap.NewNop(), db)
	require.NoError(t, err)

	return repo, db.Close
}

func newRepo(t *testing.T, id string, loadFixtures bool) (directory.Repository, func() error) {
	return newPostgreSQLRepository(t, id, loadFixtures)
}

func assertOrganizationInwayAddress(t *testing.T, repo directory.Repository, serialNumber, inwayAddress string) {
	t.Logf("serial number in assertOrganizationInwayAddress: %s", serialNumber)
	result, err := repo.GetOrganizationInwayAddress(context.Background(), serialNumber)
	require.NoError(t, err)

	assert.Equal(t, inwayAddress, result)
}

func assertInwayInRepository(t *testing.T, repo directory.Repository, iw *domain.Inway) {
	require.NotNil(t, iw)

	inwayFromRepo, err := repo.GetInway(iw.Name(), iw.Organization().SerialNumber())
	require.NoError(t, err)

	assert.Equal(t, iw, inwayFromRepo)
}

func assertServiceInRepository(t *testing.T, repo directory.Repository, s *domain.Service) {
	require.NotNil(t, s)

	model, err := repo.GetService(s.ID())
	require.NoError(t, err)

	assert.EqualValues(t, s, model)

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

	return fmt.Sprintf("test_direction%s", suffix)
}
