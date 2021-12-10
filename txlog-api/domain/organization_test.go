// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/txlog-api/domain"
)

//nolint:funlen // this is a test
func Test_NewOrganization(t *testing.T) {
	type organizationParams struct {
		serialNumber string
	}

	tests := map[string]struct {
		organization organizationParams
		expectedErr  string
	}{
		"when_invalid_organization_serial_number": {
			organization: organizationParams{
				serialNumber: "this_serial_number_is_too_long",
			},
			expectedErr: "error validating organization serial number: too long, max 20 bytes",
		},
		"happy_flow": {
			organization: organizationParams{
				serialNumber: "000000000000000001",
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewOrganization(
				tt.organization.serialNumber,
			)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.organization.serialNumber, result.SerialNumber())
			}
		})
	}
}
