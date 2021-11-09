// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/directory-api/domain"
)

func Test_NewService(t *testing.T) {
	type args struct {
		Name                     string
		APISpecificationType     domain.SpecificationType
		OrganizationSerialNumber string
		OrganizationName         string
		Costs                    *domain.NewServiceCostsArgs
	}

	tests := map[string]struct {
		args        *args
		expectedErr string
	}{
		"invalid_name": {
			args: &args{
				Name:                     "#*%",
				APISpecificationType:     domain.OpenAPI2,
				OrganizationSerialNumber: "00000000000000000001",
				OrganizationName:         "org",
				Costs: &domain.NewServiceCostsArgs{
					OneTime: 1,
					Monthly: 2,
					Request: 3,
				},
			},
			expectedErr: "Name: must be in a valid format.",
		},
		"happy_flow": {
			args: &args{
				Name:                     "name",
				APISpecificationType:     domain.OpenAPI2,
				OrganizationSerialNumber: "00000000000000000001",
				OrganizationName:         "org",
				Costs: &domain.NewServiceCostsArgs{
					OneTime: 1,
					Monthly: 2,
					Request: 3,
				},
			},
			expectedErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			org, err := domain.NewOrganization(tt.args.OrganizationName, tt.args.OrganizationSerialNumber)
			assert.NoError(t, err)

			result, err := domain.NewService(&domain.NewServiceArgs{
				Name:                 tt.args.Name,
				APISpecificationType: tt.args.APISpecificationType,
				Organization:         org,
				Costs:                tt.args.Costs,
			})

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
