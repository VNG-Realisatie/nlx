// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

// +build integration

package adapters_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-registration-api/adapters"
	"go.nlx.io/nlx/directory-registration-api/domain/directory"
	"go.nlx.io/nlx/directory-registration-api/domain/inway"
)

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
				inwayAddress:     "does-not-exist.com",
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
