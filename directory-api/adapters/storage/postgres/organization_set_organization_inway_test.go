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

func TestSetOrganizationInway(t *testing.T) {
	t.Parallel()

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		t.Error(err)
	}

	type inputParams struct {
		organizationSerialNumber string
		inwayAddress             string
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
					Organization: createNewOrganization(t, "TestSetOrganizationInwayinwayaddressnotfound", testOrganizationSerialNumber),
					Address:      "my-org-e.com:443",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationSerialNumber: testOrganizationSerialNumber,
				inwayAddress:             "does-not-exist.com",
			},
			expectedErr: storage.ErrNoInwayWithAddress,
		},
		"happy_flow": {
			setup: func(t *testing.T) *domain.NewInwayArgs {
				return &domain.NewInwayArgs{
					Name:         "inway-for-service",
					Organization: createNewOrganization(t, "TestSetOrganizationInwayhappyflow", testOrganizationSerialNumber),
					Address:      "my-org-e.com:443",
					NlxVersion:   domain.NlxVersionUnknown,
					CreatedAt:    now,
					UpdatedAt:    now,
				}
			},
			input: inputParams{
				organizationSerialNumber: testOrganizationSerialNumber,
				inwayAddress:             "my-org-e.com:443",
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

			err = storage.SetOrganizationInway(context.Background(), tt.input.organizationSerialNumber, tt.input.inwayAddress)
			require.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				assertOrganizationInwayAddress(t, storage, tt.input.organizationSerialNumber, tt.input.inwayAddress)
			}
		})
	}
}
