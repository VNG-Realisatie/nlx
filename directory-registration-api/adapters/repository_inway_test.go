// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package adapters_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-registration-api/adapters"
	"go.nlx.io/nlx/directory-registration-api/domain/directory"
	"go.nlx.io/nlx/directory-registration-api/domain/inway"
)

func testRegisterInway(t *testing.T, repo directory.Repository) {
	t.Helper()

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	tests := map[string]struct {
		createRegistrations func(*testing.T) []*inway.Inway
		expectedErr         error
		expectedInway       func(*testing.T) *inway.Inway
	}{
		"new_inway": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				iw, err := inway.NewInway(
					"my-inway-name",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{iw}
			},
			expectedInway: func(t *testing.T) *inway.Inway {
				iw, err := inway.NewInway(
					"my-inway-name",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return iw
			},
			expectedErr: nil,
		},
		"inway_without_name": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				iw, err := inway.NewInway(
					"",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{iw}
			},
			expectedInway: func(t *testing.T) *inway.Inway {
				iw, err := inway.NewInway(
					"",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return iw
			},
			expectedErr: nil,
		},
		"existing_inway_for_same_organization": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				first, err := inway.NewInway(
					"my-inway",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				second, err := inway.NewInway(
					"my-inway",
					uniqueOrganizationName(t),
					"nlx-inway.io",
					"0.0.1",
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{first, second}
			},
			expectedInway: func(t *testing.T) *inway.Inway {
				iw, err := inway.NewInway(
					"my-inway",
					uniqueOrganizationName(t),
					"nlx-inway.io",
					"0.0.1",
					now,
					now,
				)
				require.NoError(t, err)

				return iw
			},
			expectedErr: nil,
		},
		"inways_with_different_name_but_same_address": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				first, err := inway.NewInway(
					"my-first-inway",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				second, err := inway.NewInway(
					"my-second-inway",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{first, second}
			},
			expectedInway: nil,
			expectedErr:   adapters.ErrDuplicateAddress,
		},
		"created_at_should_not_update_when_registering_an_existing_inway": {
			createRegistrations: func(t *testing.T) []*inway.Inway {
				first, err := inway.NewInway(
					"my-inway",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now.Add(-1*time.Hour),
					now.Add(-1*time.Hour),
				)
				require.NoError(t, err)

				second, err := inway.NewInway(
					"my-inway",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now,
					now,
				)
				require.NoError(t, err)

				return []*inway.Inway{first, second}
			},
			expectedInway: func(t *testing.T) *inway.Inway {
				iw, err := inway.NewInway(
					"my-inway",
					uniqueOrganizationName(t),
					"localhost",
					inway.NlxVersionUnknown,
					now.Add(-1*time.Hour),
					now,
				)
				require.NoError(t, err)

				return iw
			},
			expectedErr: nil,
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
				expectedInway := tt.expectedInway(t)
				assertInwayInRepository(t, repo, expectedInway)
			}
		})
	}
}
