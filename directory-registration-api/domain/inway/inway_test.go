// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package inway_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-registration-api/domain/inway"
)

//nolint:funlen // this is a test
func Test_NewInway(t *testing.T) {
	type inwayParams struct {
		name             string
		organizationName string
		address          string
		nlxVersion       string
	}

	tests := map[string]struct {
		inway       inwayParams
		expectedErr string
	}{
		"invalid_name": {
			inway: inwayParams{
				name:             "#*%",
				organizationName: "organization-name",
				address:          "address",
				nlxVersion:       "0.0.0",
			},
			expectedErr: "name: must be in a valid format",
		},
		"without_address": {
			inway: inwayParams{
				name:             "name",
				organizationName: "organization-name",
				address:          "",
				nlxVersion:       "0.0.0",
			},
			expectedErr: "address: cannot be blank",
		},
		"invalid_address": {
			inway: inwayParams{
				name:             "name",
				organizationName: "organization-name",
				address:          "foo::bar",
				nlxVersion:       "0.0.0",
			},
			expectedErr: "address: must be a valid dial string",
		},
		"invalid_version": {
			inway: inwayParams{
				name:             "name",
				organizationName: "organization-name",
				address:          "address",
				nlxVersion:       "invalid",
			},
			expectedErr: "nlx version: must be a valid semantic version",
		},
		"without_organization_name": {
			inway: inwayParams{
				name:             "name",
				organizationName: "",
				address:          "address",
				nlxVersion:       "0.0.0",
			},
			expectedErr: "organization name: cannot be blank",
		},
		"empty_name": {
			inway: inwayParams{
				name:             "",
				organizationName: "organization-name",
				address:          "address",
				nlxVersion:       "0.0.0",
			},
			expectedErr: "",
		},
		"happy_flow": {
			inway: inwayParams{
				name:             "name",
				organizationName: "organization-name",
				address:          "address",
				nlxVersion:       "0.0.0",
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := inway.NewInway(
				tt.inway.name,
				tt.inway.organizationName,
				tt.inway.address,
				tt.inway.nlxVersion,
			)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.inway.name, result.Name())
				assert.Equal(t, tt.inway.organizationName, result.OrganizationName())
				assert.Equal(t, tt.inway.address, result.Address())
				assert.Equal(t, tt.inway.nlxVersion, result.NlxVersion())
			}
		})
	}
}
