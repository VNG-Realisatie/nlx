// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package adapters_test

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-registration-api/adapters"
	"go.nlx.io/nlx/directory-registration-api/domain/directory"
	"go.nlx.io/nlx/directory-registration-api/domain/inway"
	"go.nlx.io/nlx/directory-registration-api/domain/service"
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

func newPostgreSQLRepository(t *testing.T, id string) (*adapters.PostgreSQLRepository, func() error) {
	db, err := sqlx.Open("txdb", id)
	require.NoError(t, err)

	db.MapperFunc(xstrings.ToSnakeCase)

	repo, err := adapters.NewPostgreSQLRepository(db)
	require.NoError(t, err)

	return repo, db.Close
}

func newRepo(t *testing.T, id string) (directory.Repository, func() error) {
	return newPostgreSQLRepository(t, id)
}

func assertOrganizationInwayAddress(t *testing.T, repo directory.Repository, organizationName, inwayAddress string) {
	result, err := repo.GetOrganizationInwayAddress(context.Background(), organizationName)
	require.NoError(t, err)

	assert.Equal(t, inwayAddress, result)
}

func assertInwayInRepository(t *testing.T, repo directory.Repository, iw *inway.Inway) {
	require.NotNil(t, iw)

	inwayFromRepo, err := repo.GetInway(iw.Name(), iw.OrganizationName())
	require.NoError(t, err)

	assert.Equal(t, iw, inwayFromRepo)
}

func assertServiceInRepository(t *testing.T, repo directory.Repository, s *service.Service) {
	require.NotNil(t, s)

	model, err := repo.GetService(s.ID())
	require.NoError(t, err)

	assert.EqualValues(t, s, model)

}
