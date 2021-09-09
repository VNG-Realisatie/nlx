// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package adapters_test

import (
	"go.nlx.io/nlx/directory-registration-api/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-registration-api/adapters"
)

func TestRegisterInway(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type registration struct {
		inwayArgs   *domain.NewInwayArgs
		expectedErr error
	}

	tests := map[string]struct {
		registrations []registration
		expectedInway *domain.NewInwayArgs
	}{
		"new_inway": {
			registrations: []registration{
				{
					inwayArgs: &domain.NewInwayArgs{
						Name:             "my-inway-name",
						OrganizationName: "my-organization-name",
						Address:          "localhost",
						NlxVersion:       domain.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					expectedErr: nil,
				},
			},
			expectedInway: &domain.NewInwayArgs{
				Name:             "my-inway-name",
				OrganizationName: "my-organization-name",
				Address:          "localhost",
				NlxVersion:       domain.NlxVersionUnknown,
				CreatedAt:        now,
				UpdatedAt:        now,
			},
		},
		"inway_without_name": {
			registrations: []registration{
				{
					inwayArgs: &domain.NewInwayArgs{
						Name:             "",
						OrganizationName: "my-organization-name",
						Address:          "localhost",
						NlxVersion:       domain.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					expectedErr: nil,
				},
			},
			expectedInway: &domain.NewInwayArgs{
				Name:             "",
				OrganizationName: "my-organization-name",
				Address:          "localhost",
				NlxVersion:       domain.NlxVersionUnknown,
				CreatedAt:        now,
				UpdatedAt:        now,
			},
		},
		"existing_inway_for_same_organization": {
			registrations: []registration{
				{
					inwayArgs: &domain.NewInwayArgs{
						Name:             "my-inway",
						OrganizationName: "my-organization-name",
						Address:          "localhost",
						NlxVersion:       domain.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					expectedErr: nil,
				},
				{
					inwayArgs: &domain.NewInwayArgs{
						Name:             "my-inway",
						OrganizationName: "my-organization-name",
						Address:          "nlx-inway.io",
						NlxVersion:       "0.0.1",
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					expectedErr: nil,
				},
			},
			expectedInway: &domain.NewInwayArgs{
				Name:             "my-inway",
				OrganizationName: "my-organization-name",
				Address:          "nlx-inway.io",
				NlxVersion:       "0.0.1",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
		},
		"inways_with_different_name_but_same_address": {
			registrations: []registration{
				{
					inwayArgs: &domain.NewInwayArgs{
						Name:             "my-first-inway",
						OrganizationName: "my-organization-name",
						Address:          "localhost",
						NlxVersion:       domain.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					expectedErr: nil,
				},
				{
					inwayArgs: &domain.NewInwayArgs{
						Name:             "my-second-inway",
						OrganizationName: "my-organization-name",
						Address:          "localhost",
						NlxVersion:       domain.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					expectedErr: adapters.ErrDuplicateAddress,
				},
			},
			expectedInway: nil,
		},
		"created_at_should_not_update_when_registering_an_existing_inway": {
			registrations: []registration{
				{
					inwayArgs: &domain.NewInwayArgs{
						Name:             "my-inway",
						OrganizationName: "my-organization-name",
						Address:          "localhost",
						NlxVersion:       domain.NlxVersionUnknown,
						CreatedAt:        now.Add(-1 * time.Hour),
						UpdatedAt:        now.Add(-1 * time.Hour),
					},
					expectedErr: nil,
				},
				{
					inwayArgs: &domain.NewInwayArgs{
						Name:             "my-inway",
						OrganizationName: "my-organization-name",
						Address:          "localhost",
						NlxVersion:       domain.NlxVersionUnknown,
						CreatedAt:        now,
						UpdatedAt:        now,
					},
					expectedErr: nil,
				},
			},
			expectedInway: &domain.NewInwayArgs{
				Name:             "my-inway",
				OrganizationName: "my-organization-name",
				Address:          "localhost",
				NlxVersion:       domain.NlxVersionUnknown,
				CreatedAt:        now.Add(-1 * time.Hour),
				UpdatedAt:        now,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, close := newRepo(t, t.Name())
			defer close()

			registrations := tt.registrations

			for _, registration := range registrations {
				iw := createNewInway(t, registration.inwayArgs)
				err = repo.RegisterInway(iw)
				require.Equal(t, registration.expectedErr, err)
			}

			if tt.expectedInway != nil {
				expectedInway := createNewInway(t, tt.expectedInway)
				assertInwayInRepository(t, repo, expectedInway)
			}
		})
	}
}

func createNewInway(t *testing.T, inwayArgs *domain.NewInwayArgs) *domain.Inway {
	iw, err := domain.NewInway(inwayArgs)
	require.NoError(t, err)

	return iw
}
