// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package adapters_test

import (
	"context"
	"go.nlx.io/nlx/directory-registration-api/domain/service"
	"os"
	"regexp"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-registration-api/adapters"
	"go.nlx.io/nlx/directory-registration-api/domain/directory"
	"go.nlx.io/nlx/directory-registration-api/domain/inway"
)

func TestRepository(t *testing.T) {
	repo := newPostgreSQLRepository(t)

	t.Run("register_inway", func(t *testing.T) {
		t.Parallel()
		testRegisterInway(t, repo)
	})

	t.Run("register_service", func(t *testing.T) {
		t.Parallel()
		testRegisterService(t, repo)
	})

	t.Run("set_organization_inway", func(t *testing.T) {
		t.Parallel()
		testSetOrganizationInway(t, repo)
	})

	t.Run("clear_organization_inway", func(t *testing.T) {
		t.Parallel()
		testClearOrganizationInway(t, repo)
	})

	t.Run("get_organization_inway_address", func(t *testing.T) {
		t.Parallel()
		testGetOrganizationInwayAddress(t, repo)
	})
}

var alphanumericRegex = regexp.MustCompile("[^a-zA-Z0-9]+")

func alphanum(input string, maxLen int) string {
	result := alphanumericRegex.ReplaceAllString(input, "")

	if len(result) > maxLen {
		return result[0:maxLen]
	} else {
		return result
	}
}

func uniqueOrganizationName(t *testing.T) string {
	return alphanum(t.Name(), 100)
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

func newPostgreSQLRepository(t *testing.T) *adapters.PostgreSQLRepository {
	dsn := os.Getenv("POSTGRES_DSN")

	db, err := adapters.NewPostgreSQLConnection(dsn)
	require.NoError(t, err)

	err = adapters.PostgreSQLPerformMigrations(dsn)
	require.NoError(t, err)

	repo, err := adapters.NewPostgreSQLRepository(db)
	require.NoError(t, err)

	return repo
}
