// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package adapters_test

import (
	"context"
	"os"
	"testing"
	"time"

	"go.nlx.io/nlx/directory-registration-api/domain/service"

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

func testRegisterInway(t *testing.T, repo directory.Repository) {
	t.Helper()

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	tests := map[string]struct {
		createRegistrations func(*testing.T) []*inway.Inway
		expectedErr         error
	}{
		"new_inway": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				iw, err := inway.NewInway(
					"my-new-inway",
					"organization-a",
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{iw}
			},
			expectedErr: nil,
		},
		"inway_without_name": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				iw, err := inway.NewInway(
					"",
					"organization-b",
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{iw}
			},
			expectedErr: nil,
		},
		"existing_inway_for_same_organization": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				first, err := inway.NewInway(
					"my-inway",
					"organization-c",
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				second, err := inway.NewInway(
					"my-inway",
					"organization-c",
					"nlx-inway.io",
					"0.0.1",
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{first, second}
			},
			expectedErr: nil,
		},
		"inways_with_different_name_but_same_address": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				first, err := inway.NewInway(
					"my-first-inway",
					"organization-d",
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				second, err := inway.NewInway(
					"my-second-inway",
					"organization-d",
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{first, second}
			},
			expectedErr: adapters.ErrDuplicateAddress,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			inways := tt.createRegistrations(t)

			var lastErr error
			for _, inwayToRegister := range inways {
				err := repo.RegisterInway(inwayToRegister)
				lastErr = err
			}

			require.Equal(t, tt.expectedErr, lastErr)

			if tt.expectedErr == nil {
				lastRegistration := inways[len(inways)-1]
				assertInwayInRepository(t, repo, lastRegistration)
			}
		})
	}
}

func testRegisterService(t *testing.T, repo directory.Repository) {
	t.Helper()

	tests := map[string]struct {
		createRegistrations func(*testing.T) []*service.Service
		expectedErr         error
	}{
		"new_service": {
			createRegistrations: func(t *testing.T) []*service.Service {
				s, err := service.NewService(
					"my-service",
					"organization-d",
					"documentation-url",
					service.OpenAPI3,
					"public-support-contact",
					"tech-support-contact",
					1,
					2,
					3,
					true,
				)
				require.NoError(t, err)

				return []*service.Service{s}
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
		if err != nil {
			t.Error(err)
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			models := tt.createRegistrations(t)

			inwayModel, err := inway.NewInway(
				"inway-for-service",
				"organization-d",
				"my-org.com",
				inway.NlxVersionUnknown,
				now,
				now,
			)
			require.NoError(t, err)

			err = repo.RegisterInway(inwayModel)
			require.NoError(t, err)

			var lastErr error
			for _, model := range models {
				err := repo.RegisterService(model)
				lastErr = err
			}

			require.Equal(t, tt.expectedErr, lastErr)

			if tt.expectedErr == nil {
				lastRegistration := models[len(models)-1]
				assertServiceInRepository(t, repo, lastRegistration)
			}
		})
	}
}

func testSetOrganizationInway(t *testing.T, repo directory.Repository) {
	t.Helper()

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationName string
		inwayAddress     string
	}

	tests := map[string]struct {
		setup       func(*testing.T) *inway.Inway
		input       inputParams
		expectedErr error
	}{
		"inway_address_not_found": {
			setup: func(t *testing.T) *inway.Inway {
				inwayModel, err := inway.NewInway(
					"inway-for-service",
					"organization-e",
					"my-org-e.com",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)
				return inwayModel
			},
			input: inputParams{
				organizationName: "organization-e",
				inwayAddress:     "doesn-exist.com",
			},
			expectedErr: adapters.ErrNoInwayWithAddress,
		},
		"happy_flow": {
			setup: func(t *testing.T) *inway.Inway {
				inwayModel, err := inway.NewInway(
					"inway-for-service",
					"organization-e",
					"my-org-e.com",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)
				return inwayModel
			},
			input: inputParams{
				organizationName: "organization-e",
				inwayAddress:     "my-org-e.com",
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			inwayModel := tt.setup(t)

			err = repo.RegisterInway(inwayModel)
			require.NoError(t, err)

			err := repo.SetOrganizationInway(context.Background(), tt.input.organizationName, tt.input.inwayAddress)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				assertOrganizationInwayAddress(t, repo, tt.input.organizationName, tt.input.inwayAddress)
			}
		})
	}
}

func testGetOrganizationInwayAddress(t *testing.T, repo directory.Repository) {
	t.Helper()

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationName string
	}

	tests := map[string]struct {
		setup           func(*testing.T) *inway.Inway
		input           inputParams
		expectedAddress string
		expectedErr     error
	}{
		"organization_not_found": {
			setup: func(t *testing.T) *inway.Inway {
				inwayModel, err := inway.NewInway(
					"inway-for-service",
					"organization-i",
					"my-org-i.com",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)
				return inwayModel
			},
			input: inputParams{
				organizationName: "organization-does-not-exist",
			},
			expectedAddress: "",
			expectedErr:     adapters.ErrOrganizationNotFound,
		},
		"happy_flow": {
			setup: func(t *testing.T) *inway.Inway {
				inwayModel, err := inway.NewInway(
					"inway-for-service",
					"organization-i",
					"my-org-i.com",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)
				return inwayModel
			},
			input: inputParams{
				organizationName: "organization-i",
			},
			expectedAddress: "my-org-i.com",
			expectedErr:     nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			inwayModel := tt.setup(t)

			err = repo.RegisterInway(inwayModel)
			require.NoError(t, err)

			err := repo.SetOrganizationInway(context.Background(), inwayModel.OrganizationName(), inwayModel.Address())
			require.Equal(t, nil, err)

			address, err := repo.GetOrganizationInwayAddress(context.Background(), tt.input.organizationName)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				require.Equal(t, tt.expectedAddress, address)
			}
		})
	}
}

func testClearOrganizationInway(t *testing.T, repo directory.Repository) {
	t.Helper()

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationName string
	}

	tests := map[string]struct {
		setup       func(*testing.T) *inway.Inway
		input       inputParams
		expectedErr error
	}{
		"organization_not_found": {
			setup: func(t *testing.T) *inway.Inway {
				inwayModel, err := inway.NewInway(
					"inway-for-service",
					"organization-g",
					"my-org-g.com",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)
				return inwayModel
			},
			input: inputParams{
				organizationName: "organization-does-not-exist",
			},
			expectedErr: adapters.ErrOrganizationNotFound,
		},
		"happy_flow": {
			setup: func(t *testing.T) *inway.Inway {
				inwayModel, err := inway.NewInway(
					"inway-for-service",
					"organization-h",
					"my-org-h.com",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)
				return inwayModel
			},
			input: inputParams{
				organizationName: "organization-h",
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			inwayModel := tt.setup(t)

			err = repo.RegisterInway(inwayModel)
			require.NoError(t, err)

			err := repo.SetOrganizationInway(context.Background(), inwayModel.OrganizationName(), inwayModel.Address())
			require.Equal(t, nil, err)

			err = repo.ClearOrganizationInway(context.Background(), tt.input.organizationName)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				assertOrganizationInwayAddress(t, repo, tt.input.organizationName, "")
			}
		})
	}
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
