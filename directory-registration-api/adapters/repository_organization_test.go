// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package adapters_test

import (
	"context"
	"go.nlx.io/nlx/directory-registration-api/domain"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-registration-api/adapters"
)

func TestSetOrganizationInway(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationName string
		inwayAddress     string
	}

	tests := map[string]struct {
		setup       func(*testing.T) *domain.NewInwayArgs
		input       inputParams
		expectedErr error
	}{
		"inway_address_not_found": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "TestSetOrganizationInwayinwayaddressnotfound"),
					Address:      "my-org-e.com",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationName: "TestSetOrganizationInwayinwayaddressnotfound",
				inwayAddress:     "does-not-exist.com",
			},
			expectedErr: adapters.ErrNoInwayWithAddress,
		},
		"happy_flow": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "TestSetOrganizationInwayhappyflow"),
					Address:      "my-org-e.com",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationName: "TestSetOrganizationInwayhappyflow",
				inwayAddress:     "my-org-e.com",
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, close := newRepo(t, t.Name())
			defer close()

			inwayArgs := tt.setup(t)

			inwayModel, err := domain.NewInway(inwayArgs)
			require.NoError(t, err)

			err = repo.RegisterInway(inwayModel)
			require.NoError(t, err)

			err = repo.SetOrganizationInway(context.Background(), tt.input.organizationName, tt.input.inwayAddress)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				assertOrganizationInwayAddress(t, repo, tt.input.organizationName, tt.input.inwayAddress)
			}
		})
	}
}

func TestClearOrganizationInway(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationName string
	}

	tests := map[string]struct {
		setup       func(*testing.T) *domain.NewInwayArgs
		input       inputParams
		expectedErr error
	}{
		"organization_not_found": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "my-organization-name"),
					Address:      "my-org-g.com",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationName: "organization-does-not-exist",
			},
			expectedErr: adapters.ErrOrganizationNotFound,
		},
		"happy_flow": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "my-organization"),
					Address:      "my-org-h.com",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationName: "my-organization",
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, close := newRepo(t, t.Name())
			defer close()

			inwayArgs := tt.setup(t)

			inwayModel, err := domain.NewInway(inwayArgs)
			require.NoError(t, err)

			err = repo.RegisterInway(inwayModel)
			require.NoError(t, err)

			err = repo.SetOrganizationInway(context.Background(), inwayModel.Organization().Name(), inwayModel.Address())
			require.Equal(t, nil, err)

			log.Println(tt.input.organizationName)

			err = repo.ClearOrganizationInway(context.Background(), tt.input.organizationName)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				assertOrganizationInwayAddress(t, repo, tt.input.organizationName, "")
			}
		})
	}
}

func TestGetOrganizationInwayAddress(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationName string
	}

	tests := map[string]struct {
		setup           func(*testing.T) *domain.NewInwayArgs
		input           inputParams
		expectedAddress string
		expectedErr     error
	}{
		"organization_not_found": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "my-organization-name"),
					Address:      "my-org-i.com",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationName: "organization-does-not-exist",
			},
			expectedAddress: "",
			expectedErr:     adapters.ErrOrganizationNotFound,
		},
		"happy_flow": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "my-organization"),
					Address:      "my-org-i.com",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationName: "my-organization",
			},
			expectedAddress: "my-org-i.com",
			expectedErr:     nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo, close := newRepo(t, t.Name())
			defer close()

			inwayArgs := tt.setup(t)

			inwayModel, err := domain.NewInway(inwayArgs)
			require.NoError(t, err)

			err = repo.RegisterInway(inwayModel)
			require.NoError(t, err)

			err = repo.SetOrganizationInway(context.Background(), inwayModel.Organization().Name(), inwayModel.Address())
			require.Equal(t, nil, err)

			address, err := repo.GetOrganizationInwayAddress(context.Background(), tt.input.organizationName)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				require.Equal(t, tt.expectedAddress, address)
			}
		})
	}
}
