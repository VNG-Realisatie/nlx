// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-registration-api/domain/service"
)

func Test_NewService(t *testing.T) {
	type serviceParams struct {
		name                 string
		apiSpecificationType service.SpecificationType
	}

	tests := map[string]struct {
		service     serviceParams
		expectedErr string
	}{
		"invalid_name": {
			service: serviceParams{
				name:                 "#*%",
				apiSpecificationType: service.OpenAPI2,
			},
			expectedErr: "name: must be in a valid format",
		},
		"happy_flow": {
			service: serviceParams{
				name:                 "name",
				apiSpecificationType: service.OpenAPI2,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := service.NewService(
				tt.service.name,
				"",
				"",
				tt.service.apiSpecificationType,
				"",
				"",
				0,
				0,
				0,
				false,
			)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.service.name, result.Name())
				assert.Equal(t, tt.service.apiSpecificationType, result.APISpecificationType())
			}
		})
	}
}
