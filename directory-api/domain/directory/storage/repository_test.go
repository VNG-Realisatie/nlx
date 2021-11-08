// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package storage_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pgadapter_test_setup "go.nlx.io/nlx/directory-api/adapters/storage/postgres/test_setup"
	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func new(t *testing.T, loadFixtures bool) (storage.Repository, func() error) {
	return pgadapter_test_setup.New(t, loadFixtures)
}

func assertOrganizationInwayAddress(t *testing.T, repo storage.Repository, serialNumber, inwayAddress string) {
	t.Logf("serial number in assertOrganizationInwayAddress: %s", serialNumber)
	result, err := repo.GetOrganizationInwayAddress(context.Background(), serialNumber)
	require.NoError(t, err)

	assert.Equal(t, inwayAddress, result)
}

func assertInwayInRepository(t *testing.T, repo storage.Repository, iw *domain.Inway) {
	require.NotNil(t, iw)

	inwayFromRepo, err := repo.GetInway(iw.Name(), iw.Organization().SerialNumber())
	require.NoError(t, err)

	assert.Equal(t, iw, inwayFromRepo)
}

func assertServiceInRepository(t *testing.T, repo storage.Repository, s *domain.Service) {
	require.NotNil(t, s)

	model, err := repo.GetService(s.ID())
	require.NoError(t, err)

	assert.EqualValues(t, s, model)

}
