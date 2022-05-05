// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package pgadapter_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func TestGetOrganizationInwayAddress(t *testing.T) {
	t.Parallel()

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationSerialNumber string
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
					Organization: createNewOrganization(t, "my-organization-name", testOrganizationSerialNumber),
					Address:      "my-org-i.com:443",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationSerialNumber: "010203040506070809",
			},
			expectedAddress: "",
			expectedErr:     storage.ErrNotFound,
		},
		"happy_flow": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "my-organization", testOrganizationSerialNumber),
					Address:      "my-org-i.com:443",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationSerialNumber: testOrganizationSerialNumber,
			},
			expectedAddress: "my-org-i.com:443",
			expectedErr:     nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage, close := new(t, false)
			defer close()

			inwayArgs := tt.setup(t)

			inwayModel, err := domain.NewInway(inwayArgs)
			require.NoError(t, err)

			err = storage.RegisterInway(inwayModel)
			require.NoError(t, err)

			err = storage.SetOrganizationInway(context.Background(), inwayModel.Organization().SerialNumber(), inwayModel.Address())
			require.Equal(t, nil, err)

			address, err := storage.GetOrganizationInwayAddress(context.Background(), tt.input.organizationSerialNumber)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				require.Equal(t, tt.expectedAddress, address)
			}
		})
	}
}
