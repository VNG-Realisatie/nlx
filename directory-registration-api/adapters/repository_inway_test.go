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
		getRegistrations func(*testing.T) []*inway.NewInwayArgs
		expectedErr      error
		expectedInway    func(*testing.T) *inway.NewInwayArgs
	}{
		"new_inway": {
			getRegistrations: func(t *testing.T) []*inway.NewInwayArgs {
				return []*inway.NewInwayArgs{
					{
						Name:             "my-inway-name",
						OrganizationName: uniqueOrganizationName(t),
						Address:          "localhost",
						NlxVersion:       inway.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				}
			},
			expectedInway: func(t *testing.T) *inway.NewInwayArgs {
				return &inway.NewInwayArgs{
					Name:             "my-inway-name",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "localhost",
					NlxVersion:       inway.NlxVersionUnknown,
					CreatedAt:        now,
					UpdatedAt:        now,
				}
			},
			expectedErr: nil,
		},
		"inway_without_name": {
			getRegistrations: func(t *testing.T) []*inway.NewInwayArgs {
				return []*inway.NewInwayArgs{
					{
						Name:             "",
						OrganizationName: uniqueOrganizationName(t),
						Address:          "localhost",
						NlxVersion:       inway.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				}
			},
			expectedInway: func(t *testing.T) *inway.NewInwayArgs {
				return &inway.NewInwayArgs{
					Name:             "",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "localhost",
					NlxVersion:       inway.NlxVersionUnknown,
					CreatedAt:        now,
					UpdatedAt:        now,
				}
			},
			expectedErr: nil,
		},
		"existing_inway_for_same_organization": {
			getRegistrations: func(t *testing.T) []*inway.NewInwayArgs {
				return []*inway.NewInwayArgs{
					{
						Name:             "my-inway",
						OrganizationName: uniqueOrganizationName(t),
						Address:          "localhost",
						NlxVersion:       inway.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					{
						Name:             "my-inway",
						OrganizationName: uniqueOrganizationName(t),
						Address:          "nlx-inway.io",
						NlxVersion:       "0.0.1",
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				}
			},
			expectedInway: func(t *testing.T) *inway.NewInwayArgs {
				return &inway.NewInwayArgs{
					Name:             "my-inway",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "nlx-inway.io",
					NlxVersion:       "0.0.1",
					CreatedAt:        now,
					UpdatedAt:        now,
				}
			},
			expectedErr: nil,
		},
		"inways_with_different_name_but_same_address": {
			getRegistrations: func(t *testing.T) []*inway.NewInwayArgs {
				return []*inway.NewInwayArgs{
					{
						Name:             "my-first-inway",
						OrganizationName: uniqueOrganizationName(t),
						Address:          "localhost",
						NlxVersion:       inway.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					{
						Name:             "my-second-inway",
						OrganizationName: uniqueOrganizationName(t),
						Address:          "localhost",
						NlxVersion:       inway.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				}
			},
			expectedInway: nil,
			expectedErr:   adapters.ErrDuplicateAddress,
		},
		"created_at_should_not_update_when_registering_an_existing_inway": {
			getRegistrations: func(t *testing.T) []*inway.NewInwayArgs {
				return []*inway.NewInwayArgs{
					{
						Name:             "my-inway",
						OrganizationName: uniqueOrganizationName(t),
						Address:          "localhost",
						NlxVersion:       inway.NlxVersionUnknown,
						CreatedAt:        now.Add(-1 * time.Hour),
						UpdatedAt:        now.Add(-1 * time.Hour),
					},
					{
						Name:             "my-inway",
						OrganizationName: uniqueOrganizationName(t),
						Address:          "localhost",
						NlxVersion:       inway.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				}
			},
			expectedInway: func(t *testing.T) *inway.NewInwayArgs {
				return &inway.NewInwayArgs{
					Name:             "my-inway",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "localhost",
					NlxVersion:       inway.NlxVersionUnknown,
					CreatedAt:        now.Add(-1 * time.Hour),
					UpdatedAt:        now,
				}
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			inways := tt.getRegistrations(t)

			var lastErr error
			for _, inwayArgs := range inways {
				iw, err := inway.NewInway(inwayArgs)
				require.NoError(t, err)

				err = repo.RegisterInway(iw)
				lastErr = err
			}

			require.Equal(t, tt.expectedErr, lastErr)

			if tt.expectedErr == nil {
				inwayArgs := tt.expectedInway(t)

				expectedInway, err := inway.NewInway(inwayArgs)
				require.NoError(t, err)

				assertInwayInRepository(t, repo, expectedInway)
			}
		})
	}
}
