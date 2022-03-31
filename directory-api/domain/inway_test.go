// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-api/domain"
)

//nolint:funlen // this is a test
func TestNewInway(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		inwayArgs   *domain.NewInwayArgs
		expectedErr string
	}{
		"invalid_name": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "#*%",
				Organization: createNewOrganization(t),
				Address:      "address:443",
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "Name: must be in a valid format.",
		},
		"without_address": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				Address:      "",
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "Address: cannot be blank.",
		},
		"invalid_address": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				Address:      "foo::bar",
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "Address: must be a valid dial string.",
		},
		"invalid_version": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				Address:      "address:443",
				NlxVersion:   "invalid",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "NlxVersion: must be a valid semantic version.",
		},
		"without_organization": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "name",
				Organization: nil,
				Address:      "address:443",
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "Organization: is required.",
		},
		"empty_name": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "",
				Organization: createNewOrganization(t),
				Address:      "address:443",
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "",
		},
		"created_at_in_future": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				Address:      "address:443",
				NlxVersion:   "0.0.0",
				CreatedAt:    now.Add(1 * time.Hour),
				UpdatedAt:    now,
			},
			expectedErr: "CreatedAt: must not be in the future.",
		},
		"updated_at_in_future": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				Address:      "address:443",
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now.Add(1 * time.Hour),
			},
			expectedErr: "UpdatedAt: must not be in the future.",
		},
		"happy_flow": {
			inwayArgs: &domain.NewInwayArgs{
				Name:         "name",
				Organization: createNewOrganization(t),
				Address:      "address:443",
				NlxVersion:   "0.0.0",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewInway(tt.inwayArgs)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.inwayArgs.Name, result.Name())
				assert.Equal(t, tt.inwayArgs.Organization.Name(), result.Organization().Name())
				assert.Equal(t, tt.inwayArgs.Address, result.Address())
				assert.Equal(t, tt.inwayArgs.NlxVersion, result.NlxVersion())
			}
		})
	}
}

const testOrganizationName = "my-organization-name"

const testOrganizationSerialNumber = "01234567890123456789"

func createNewOrganization(t *testing.T) *domain.Organization {
	model, err := domain.NewOrganization(testOrganizationName, testOrganizationSerialNumber)
	require.NoError(t, err)

	return model
}
