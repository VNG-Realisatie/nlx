// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-registration-api/domain"
)

func Test_NewService(t *testing.T) {
	tests := map[string]struct {
		args        *domain.NewServiceArgs
		expectedErr string
	}{
		"invalid_name": {
			args: &domain.NewServiceArgs{
				Name:                 "#*%",
				APISpecificationType: domain.OpenAPI2,
			},
			expectedErr: "Name: must be in a valid format.",
		},
		"happy_flow": {
			args: &domain.NewServiceArgs{
				Name:                 "name",
				APISpecificationType: domain.OpenAPI2,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := domain.NewService(tt.args)

			if tt.expectedErr != "" {
				assert.Nil(t, result)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, result)
				assert.Nil(t, err)

				assert.Equal(t, tt.args.Name, result.Name())
				assert.Equal(t, tt.args.APISpecificationType, result.APISpecificationType())
			}
		})
	}
}
