// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package inway_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-registration-api/domain/inway"
)

//nolint:funlen // this is a test
func Test_NewInway(t *testing.T) {
	now := time.Now()

	tests := map[string]struct {
		inwayArgs   *inway.NewInwayArgs
		expectedErr string
	}{
		"invalid_name": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "#*%",
				OrganizationName: "organization-name",
				Address:          "address",
				NlxVersion:       "0.0.0",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectedErr: "Name: must be in a valid format.",
		},
		"without_address": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "name",
				OrganizationName: "organization-name",
				Address:          "",
				NlxVersion:       "0.0.0",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectedErr: "Address: cannot be blank.",
		},
		"invalid_address": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "name",
				OrganizationName: "organization-name",
				Address:          "foo::bar",
				NlxVersion:       "0.0.0",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectedErr: "Address: must be a valid dial string.",
		},
		"invalid_version": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "name",
				OrganizationName: "organization-name",
				Address:          "address",
				NlxVersion:       "invalid",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectedErr: "NlxVersion: must be a valid semantic version.",
		},
		"without_organization_name": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "name",
				OrganizationName: "",
				Address:          "address",
				NlxVersion:       "0.0.0",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectedErr: "OrganizationName: cannot be blank.",
		},
		"empty_name": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "",
				OrganizationName: "organization-name",
				Address:          "address",
				NlxVersion:       "0.0.0",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectedErr: "",
		},
		"created_at_in_future": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "name",
				OrganizationName: "organization-name",
				Address:          "address",
				NlxVersion:       "0.0.0",
				CreatedAt:        now.Add(1 * time.Hour),
				UpdatedAt:        now,
			},
			expectedErr: "CreatedAt: must not be in the future.",
		},
		"updated_at_in_future": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "name",
				OrganizationName: "organization-name",
				Address:          "address",
				NlxVersion:       "0.0.0",
				CreatedAt:        now,
				UpdatedAt:        now.Add(1 * time.Hour),
			},
			expectedErr: "UpdatedAt: must not be in the future.",
		},
		"happy_flow": {
			inwayArgs: &inway.NewInwayArgs{
				Name:             "name",
				OrganizationName: "organization-name",
				Address:          "address",
				NlxVersion:       "0.0.0",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := inway.NewInway(tt.inwayArgs)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.inwayArgs.Name, result.Name())
				assert.Equal(t, tt.inwayArgs.OrganizationName, result.OrganizationName())
				assert.Equal(t, tt.inwayArgs.Address, result.Address())
				assert.Equal(t, tt.inwayArgs.NlxVersion, result.NlxVersion())
			}
		})
	}
}
