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

	type registration struct {
		inwayArgs   *inway.NewInwayArgs
		expectedErr error
	}

	type testArgs struct {
		registrations []registration
		expectedInway *inway.NewInwayArgs
	}

	tests := map[string]func(*testing.T) testArgs{
		"new_inway": func(t *testing.T) testArgs {
			return testArgs{
				registrations: []registration{
					{
						inwayArgs: &inway.NewInwayArgs{
							Name:             "my-inway-name",
							OrganizationName: uniqueOrganizationName(t),
							Address:          "localhost",
							NlxVersion:       inway.NlxVersionUnknown,
							CreatedAt:        now,
							UpdatedAt:        now,
						},
						expectedErr: nil,
					},
				},
				expectedInway: &inway.NewInwayArgs{
					Name:             "my-inway-name",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "localhost",
					NlxVersion:       inway.NlxVersionUnknown,
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			}
		},
		"inway_without_name": func(t *testing.T) testArgs {
			return testArgs{
				registrations: []registration{
					{
						inwayArgs: &inway.NewInwayArgs{
							Name:             "",
							OrganizationName: uniqueOrganizationName(t),
							Address:          "localhost",
							NlxVersion:       inway.NlxVersionUnknown,
							CreatedAt:        now,
							UpdatedAt:        now,
						},
						expectedErr: nil,
					},
				},
				expectedInway: &inway.NewInwayArgs{
					Name:             "",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "localhost",
					NlxVersion:       inway.NlxVersionUnknown,
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			}
		},
		"existing_inway_for_same_organization": func(t *testing.T) testArgs {
			return testArgs{
				registrations: []registration{
					{
						inwayArgs: &inway.NewInwayArgs{
							Name:             "my-inway",
							OrganizationName: uniqueOrganizationName(t),
							Address:          "localhost",
							NlxVersion:       inway.NlxVersionUnknown,
							CreatedAt:        now,
							UpdatedAt:        now,
						},
						expectedErr: nil,
					},
					{
						inwayArgs: &inway.NewInwayArgs{
							Name:             "my-inway",
							OrganizationName: uniqueOrganizationName(t),
							Address:          "nlx-inway.io",
							NlxVersion:       "0.0.1",
							CreatedAt:        now,
							UpdatedAt:        now,
						},
						expectedErr: nil,
					},
				},
				expectedInway: &inway.NewInwayArgs{
					Name:             "my-inway",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "nlx-inway.io",
					NlxVersion:       "0.0.1",
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			}
		},
		"inways_with_different_name_but_same_address": func(t *testing.T) testArgs {
			return testArgs{
				registrations: []registration{
					{
						inwayArgs: &inway.NewInwayArgs{
							Name:             "my-first-inway",
							OrganizationName: uniqueOrganizationName(t),
							Address:          "localhost",
							NlxVersion:       inway.NlxVersionUnknown,
							CreatedAt:        now,
							UpdatedAt:        now,
						},
						expectedErr: nil,
					},
					{
						inwayArgs: &inway.NewInwayArgs{
							Name:             "my-second-inway",
							OrganizationName: uniqueOrganizationName(t),
							Address:          "localhost",
							NlxVersion:       inway.NlxVersionUnknown,
							CreatedAt:        now,
							UpdatedAt:        now,
						},
						expectedErr: adapters.ErrDuplicateAddress,
					},
				},
				expectedInway: &inway.NewInwayArgs{
					Name:             "my-first-inway",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "localhost",
					NlxVersion:       inway.NlxVersionUnknown,
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			}
		},
		"created_at_should_not_update_when_registering_an_existing_inway": func(t *testing.T) testArgs {
			return testArgs{
				registrations: []registration{
					{
						inwayArgs: &inway.NewInwayArgs{
							Name:             "my-inway",
							OrganizationName: uniqueOrganizationName(t),
							Address:          "localhost",
							NlxVersion:       inway.NlxVersionUnknown,
							CreatedAt:        now.Add(-1 * time.Hour),
							UpdatedAt:        now.Add(-1 * time.Hour),
						},
						expectedErr: nil,
					},
					{
						inwayArgs: &inway.NewInwayArgs{
							Name:             "my-inway",
							OrganizationName: uniqueOrganizationName(t),
							Address:          "localhost",
							NlxVersion:       inway.NlxVersionUnknown,
							CreatedAt:        now,
							UpdatedAt:        now,
						},
						expectedErr: nil,
					},
				},
				expectedInway: &inway.NewInwayArgs{
					Name:             "my-inway",
					OrganizationName: uniqueOrganizationName(t),
					Address:          "localhost",
					NlxVersion:       inway.NlxVersionUnknown,
					CreatedAt:        now.Add(-1 * time.Hour),
					UpdatedAt:        now,
				},
			}
		},
	}

	for name, createTestInput := range tests {
		createTestInput := createTestInput

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tt := createTestInput(t)
			registrations := tt.registrations

			for _, registration := range registrations {
				iw := createNewInway(t, registration.inwayArgs)
				err = repo.RegisterInway(iw)
				require.Equal(t, registration.expectedErr, err)
			}

			expectedInway := createNewInway(t, tt.expectedInway)
			assertInwayInRepository(t, repo, expectedInway)
		})
	}
}

func createNewInway(t *testing.T, inwayArgs *inway.NewInwayArgs) *inway.Inway {
	iw, err := inway.NewInway(inwayArgs)
	require.NoError(t, err)

	return iw
}
