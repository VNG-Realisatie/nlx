// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package adapters_test

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"go.nlx.io/nlx/directory-registration-api/domain"
	"go.uber.org/zap"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-registration-api/adapters"
	"go.nlx.io/nlx/directory-registration-api/domain/directory"
)

var setupOnce sync.Once

func setup(t *testing.T) {
	setupOnce.Do(func() {
		setupPostgreSQLRepository(t)
	})
}

func setupPostgreSQLRepository(t *testing.T) {
	dsn := os.Getenv("POSTGRES_DSN")

	err := adapters.PostgreSQLPerformMigrations(dsn)
	require.NoError(t, err)

	txdb.Register("txdb", "postgres", dsn)

	// This is necessary because the default BindVars for txdb isn't correct
	sqlx.BindDriver("txdb", sqlx.DOLLAR)
}

func newPostgreSQLRepository(t *testing.T, id string, loadFixtures bool) (*adapters.PostgreSQLRepository, func() error) {
	db, err := sqlx.Open("txdb", id)
	require.NoError(t, err)

	if loadFixtures {
		fixtures, err := testfixtures.New(
			testfixtures.Database(db.DB),
			testfixtures.Dialect("postgres"),
			testfixtures.Directory("testdata/fixtures/postgres"),
			testfixtures.DangerousSkipTestDatabaseCheck(),
		)
		require.NoError(t, err)
		err = fixtures.Load()
		require.NoError(t, err)
	}

	db.MapperFunc(xstrings.ToSnakeCase)

	repo, err := adapters.New(zap.NewNop(), db)
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
