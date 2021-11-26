// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package storage_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pgadapter_test_setup "go.nlx.io/nlx/directory-api/adapters/storage/postgres/test_setup"
	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func new(t *testing.T, enableFixtures bool) (storage.Repository, func() error) {
	repo, close := pgadapter_test_setup.New(t)
	if enableFixtures {
		err := loadFixtures(repo)
		require.NoError(t, err)
	}

	return repo, close
}

func loadFixtures(repo storage.Repository) error {
	newOrganizationsArgs := []struct {
		name         string
		serialNumber string
	}{
		{
			name:         "fixture-organization-name",
			serialNumber: "01234567890123456789",
		},
		{
			name:         "fixture-second-organization-name",
			serialNumber: "01234567890123456781",
		},
		{
			name:         "duplicate-org-name",
			serialNumber: "11111111111111111111",
		},
		{
			name:         "duplicate-org-name",
			serialNumber: "22222222222222222222",
		},
	}

	organizationsModels := make([]*domain.Organization, len(newOrganizationsArgs))

	for i, args := range newOrganizationsArgs {
		organization, err := domain.NewOrganization(args.name, args.serialNumber)
		if err != nil {
			return err
		}

		organizationsModels[i] = organization
	}

	newInwaysArgs := []*domain.NewInwayArgs{
		{
			Name:                "fixture-inway-name-one",
			Address:             "fixture-inway-address-one.com",
			IsOrganizationInway: true,
			Organization:        organizationsModels[0],
			NlxVersion:          "1.0.0",
			CreatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:                "fixture-inway-name-two",
			Address:             "fixture-inway-address-two.com",
			IsOrganizationInway: false,
			Organization:        organizationsModels[0],
			NlxVersion:          "1.0.0",
			CreatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:                "fixture-inway-name-three",
			Address:             "fixture-inway-address-three.com",
			IsOrganizationInway: false,
			Organization:        organizationsModels[1],
			NlxVersion:          "1.0.0",
			CreatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:                "fixture-inway-name-four",
			Address:             "fixture-inway-address-four.com",
			IsOrganizationInway: false,
			Organization:        organizationsModels[2],
			NlxVersion:          "1.0.0",
			CreatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:                "fixture-inway-name-five",
			Address:             "fixture-inway-address-five.com",
			IsOrganizationInway: false,
			Organization:        organizationsModels[3],
			NlxVersion:          "1.0.0",
			CreatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
	}

	for _, args := range newInwaysArgs {
		inway, err := domain.NewInway(args)
		if err != nil {
			return err
		}

		err = repo.RegisterInway(inway)
		if err != nil {
			return err
		}

		if inway.IsOrganizationInway() {
			repo.SetOrganizationInway(context.Background(), inway.Organization().SerialNumber(), inway.Address())
		}
	}

	newServicesArgs := []*domain.NewServiceArgs{
		{
			Name:                 "fixture-service-name",
			Organization:         organizationsModels[0],
			DocumentationURL:     "https://fixture-documentation-url.com",
			APISpecificationType: "OpenAPI3",
			Internal:             false,
			TechSupportContact:   "fixture@tech-support-contact.com",
			PublicSupportContact: "fixture@public-support-contact.com",
			Costs: &domain.NewServiceCostsArgs{
				OneTime: 1,
				Monthly: 2,
				Request: 3,
			},
		},
	}

	for _, args := range newServicesArgs {
		service, err := domain.NewService(args)
		if err != nil {
			return err
		}

		err = repo.RegisterService(service)
		if err != nil {
			return err
		}
	}

	newOutwaysArgs := []*domain.NewOutwayArgs{
		{
			Name:         "fixture-outway-name-one",
			Organization: organizationsModels[0],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:         "fixture-outway-name-two",
			Organization: organizationsModels[1],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:         "fixture-outway-name-three",
			Organization: organizationsModels[1],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:         "fixture-outway-name-four",
			Organization: organizationsModels[1],
			NlxVersion:   "2.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:         "fixture-outway-name-five",
			Organization: organizationsModels[3],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
	}

	for _, args := range newOutwaysArgs {
		outway, err := domain.NewOutway(args)
		if err != nil {
			return err
		}

		err = repo.RegisterOutway(outway)
		if err != nil {
			return err
		}
	}

	newOutwaysArgs := []*domain.NewOutwayArgs{
		{
			Name:         "fixture-inway-name-one",
			Organization: organizationsModels[0],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:         "fixture-inway-name-two",
			Organization: organizationsModels[1],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:         "fixture-inway-name-three",
			Organization: organizationsModels[1],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:         "fixture-inway-name-four",
			Organization: organizationsModels[1],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:         "fixture-inway-name-five",
			Organization: organizationsModels[3],
			NlxVersion:   "1.0.0",
			CreatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:    time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
	}

	for _, args := range newOutwaysArgs {
		outway, err := domain.NewOutway(args)
		if err != nil {
			return err
		}

		err = repo.RegisterOutway(outway)
		if err != nil {
			return err
		}
	}

	return nil
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

func assertOutwayInRepository(t *testing.T, repo storage.Repository, ow *domain.Outway) {
	require.NotNil(t, ow)

	outwayFromRepo, err := repo.GetOutway(ow.Name(), ow.Organization().SerialNumber())
	require.NoError(t, err)

	assert.Equal(t, ow, outwayFromRepo)
}

func assertServiceInRepository(t *testing.T, repo storage.Repository, s *domain.Service) {
	require.NotNil(t, s)

	model, err := repo.GetService(s.ID())
	require.NoError(t, err)

	assert.EqualValues(t, s, model)

}
