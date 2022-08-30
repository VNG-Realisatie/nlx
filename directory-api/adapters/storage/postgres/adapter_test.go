// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package pgadapter_test

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/strings"
	pgadapter "go.nlx.io/nlx/directory-api/adapters/storage/postgres"
	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
	"go.nlx.io/nlx/testing/testingutils"
)

const dbName = "test_directory"
const dbDriver = "txdb"

var setupOnce sync.Once

func setupDatabase(t *testing.T) {
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

	txdb.Register(dbDriver, "postgres", dsn)

	// This is necessary because the default BindVars for txdb isn't correct
	sqlx.BindDriver(dbDriver, sqlx.DOLLAR)

}

func New(t *testing.T) (*pgadapter.PostgreSQLRepository, func() error) {
	setupOnce.Do(func() {
		setupDatabase(t)
	})

	db, err := sqlx.Open(dbDriver, t.Name())
	require.NoError(t, err)

	db.MapperFunc(strings.ToSnakeCase)

	repo, err := pgadapter.New(zap.NewNop(), db)
	require.NoError(t, err)

	return repo, db.Close
}

func new(t *testing.T, enableFixtures bool) (storage.Repository, func() error) {
	repo, close := New(t)
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
			Name:                      "fixture-inway-name-one",
			Address:                   "fixture-inway-address-one.com:443",
			ManagementAPIProxyAddress: "fixture-inway-proxy-address-one.com:8443",
			IsOrganizationInway:       true,
			Organization:              organizationsModels[0],
			NlxVersion:                "1.0.0",
			CreatedAt:                 time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:                 time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:                "fixture-inway-name-two",
			Address:             "fixture-inway-address-two.com:443",
			IsOrganizationInway: false,
			Organization:        organizationsModels[0],
			NlxVersion:          "1.0.0",
			CreatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:                "fixture-inway-name-three",
			Address:             "fixture-inway-address-three.com:443",
			IsOrganizationInway: false,
			Organization:        organizationsModels[1],
			NlxVersion:          "1.0.0",
			CreatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:                "fixture-inway-name-four",
			Address:             "fixture-inway-address-four.com:443",
			IsOrganizationInway: false,
			Organization:        organizationsModels[2],
			NlxVersion:          "1.0.0",
			CreatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
			UpdatedAt:           time.Date(2021, 1, 2, 1, 2, 3, 0, time.UTC),
		},
		{
			Name:                "fixture-inway-name-five",
			Address:             "fixture-inway-address-five.com:443",
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
			Availabilities: []*domain.NewServiceAvailability{
				{InwayAddress: "fixture-inway-address-one.com:443"},
				{InwayAddress: "fixture-inway-address-two.com:443"},
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
