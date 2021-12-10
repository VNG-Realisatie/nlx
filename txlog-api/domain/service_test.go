// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/txlog-api/domain"
)

//nolint:funlen // this is a test
func Test_NewService(t *testing.T) {
	type serviceParams struct {
		name string
	}

	tests := map[string]struct {
		service     serviceParams
		expectedErr string
	}{
		"when_no_name": {
			service: serviceParams{
				name: "",
			},
			expectedErr: "cannot be blank",
		},
		"happy_flow": {
			service: serviceParams{
				name: "test-service",
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewService(
				tt.service.name,
			)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				if tt.expectedErr == "" {
					assert.Equal(t, tt.service.name, result.Name())
				}
			}
		})
	}
}
