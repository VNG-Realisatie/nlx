// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package inway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInway(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			inway, err := NewInway(
				tt.inway.name,
				tt.inway.organizationName,
				tt.inway.address,
				tt.inway.nlxVersion,
			)

			if tt.expectedErr != "" {
				assert.Nil(t, inway)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, inway)
				assert.Nil(t, err)
			}
		})
	}
}
