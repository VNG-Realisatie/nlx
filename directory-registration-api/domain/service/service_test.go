// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-registration-api/domain/service"
)

func Test_NewService(t *testing.T) {
	tests := map[string]struct {
		args        *service.NewServiceArgs
		expectedErr string
	}{
		"invalid_name": {
			args: &service.NewServiceArgs{
				Name:                 "#*%",
				APISpecificationType: service.OpenAPI2,
			},
			expectedErr: "Name: must be in a valid format.",
		},
		"happy_flow": {
			args: &service.NewServiceArgs{
				Name:                 "name",
				APISpecificationType: service.OpenAPI2,
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			result, err := service.NewService(tt.args)

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
