// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package adapters_test

import (
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-registration-api/adapters"
	"go.nlx.io/nlx/directory-registration-api/domain/directory"
	"go.nlx.io/nlx/directory-registration-api/domain/inway"
)

func TestRepository(t *testing.T) {
	t.Run("register_inway", func(t *testing.T) {
		repo := newPostgreSQLRepository(t)
		testRegisterInway(t, repo)
	})
}

func testRegisterInway(t *testing.T, repo directory.Repository) {
	t.Helper()

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
					"organization-b",
					"localhost",
					inway.NlxVersionUnknown,
				)
				require.NoError(t, err)

				second, err := inway.NewInway(
					"my-inway",
					"organization-b",
					"nlx-inway.io",
					"0.0.1",
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
					"organization-c",
					"localhost",
					inway.NlxVersionUnknown,
				)
				require.NoError(t, err)

				second, err := inway.NewInway(
					"my-second-inway",
					"organization-c",
					"localhost",
					inway.NlxVersionUnknown,
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

func assertInwayInRepository(t *testing.T, repo directory.Repository, iw *inway.Inway) {
	require.NotNil(t, iw)

	inwayFromRepo, err := repo.GetInway(iw.Name(), iw.OrganizationName())
	require.NoError(t, err)

	assert.Equal(t, iw, inwayFromRepo)
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
