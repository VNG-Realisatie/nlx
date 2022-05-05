// Copyright © VNG Realisatie 2022
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

func TestClearOrganizationInway(t *testing.T) {
	t.Parallel()

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationSerialNumber string
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
					Organization: createNewOrganization(t, "my-organization-name", testOrganizationSerialNumber),
					Address:      "my-org-g.com:443",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationSerialNumber: "12345678900987654321",
			},
			expectedErr: storage.ErrNotFound,
		},
		"happy_flow": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "my-organization", testOrganizationSerialNumber),
					Address:      "my-org-h.com:443",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationSerialNumber: testOrganizationSerialNumber,
			},
			expectedErr: nil,
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
			require.NoError(t, err)

			t.Logf("tt.input.organizationSerialNumber = %s", tt.input.organizationSerialNumber)

			err = storage.ClearOrganizationInway(context.Background(), tt.input.organizationSerialNumber)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				assertOrganizationInwayAddress(t, storage, tt.input.organizationSerialNumber, "")
			}
		})
	}
}
