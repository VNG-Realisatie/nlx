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
		name string
	}

	tests := map[string]struct {
		organization organizationParams
		expectedErr  string
	}{
		"invalid_name": {
			organization: organizationParams{
				name: "#*%",
			},
			expectedErr: "organization name: must be in a valid format",
		},
		"empty_name": {
			organization: organizationParams{
				name: "",
			},
			expectedErr: "organization name: cannot be blank",
		},
		"happy_flow": {
			organization: organizationParams{
				name: "name",
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewOrganization(
				tt.organization.name,
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
