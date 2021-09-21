// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-registration-api/domain"
)

//nolint:funlen // this is a test
func Test_NewOrganization(t *testing.T) {
	type organizationParams struct {
		name         string
		serialNumber string
	}

	tests := map[string]struct {
		organization organizationParams
		expectedErr  string
	}{
		"invalid_name": {
			organization: organizationParams{
				name:         "#*%",
				serialNumber: testOrganizationSerialNumber,
			},
			expectedErr: "error validating organization name: must be in a valid format",
		},
		"empty_name": {
			organization: organizationParams{
				name:         "",
				serialNumber: testOrganizationSerialNumber,
			},
			expectedErr: "error validating organization name: cannot be blank",
		},
		"short_serial_number": {
			organization: organizationParams{
				name:         "name",
				serialNumber: "0123456789",
			},
			expectedErr: "error validating organization serial number: must be in a valid format",
		},
		"long_serial_number": {
			organization: organizationParams{
				name:         "name",
				serialNumber: "012345678901234567890123456789",
			},
			expectedErr: "error validating organization serial number: must be in a valid format",
		},
		"invalid_serial_number": {
			organization: organizationParams{
				name:         "name",
				serialNumber: "0123456789abc456789",
			},
			expectedErr: "error validating organization serial number: must be in a valid format",
		},
		"empty_serial_number": {
			organization: organizationParams{
				name:         "name",
				serialNumber: "",
			},
			expectedErr: "error validating organization serial number: cannot be blank",
		},
		"happy_flow": {
			organization: organizationParams{
				name:         "name",
				serialNumber: testOrganizationSerialNumber,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewOrganization(
				tt.organization.name,
				tt.organization.serialNumber,
			)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.organization.name, result.Name())
			}
		})
	}
}
