// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package pgadapter_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"go.nlx.io/nlx/directory-api/domain"
)

func TestRegisterService(t *testing.T) {
	t.Parallel()

	type organization struct {
		SerialNumber string
		Name         string
	}

	tests := map[string]struct {
		organization *organization
		services     []*domain.NewServiceArgs
	}{
		"new_service_offered_by_single_inway": {
			services: []*domain.NewServiceArgs{
				{
					Name:                 "my-service",
					Internal:             true,
					DocumentationURL:     "documentation-url",
					APISpecificationType: domain.OpenAPI3,
					PublicSupportContact: "public-support-contact",
					TechSupportContact:   "tech-support-contact",
					Costs: &domain.NewServiceCostsArgs{
						OneTime: 1,
						Monthly: 2,
						Request: 3,
					},
					Availabilities: []*domain.NewServiceAvailability{
						{InwayAddress: "fixture-inway-address-one.com:443", State: domain.InwayDOWN},
					},
				},
			},
			organization: &organization{
				SerialNumber: "01234567890123456789",
				Name:         "fixture-organization-name",
			},
		},
		"new_service_offered_by_multiple_inway": {
			services: []*domain.NewServiceArgs{
				{
					Name:                 "my-service",
					Internal:             true,
					DocumentationURL:     "documentation-url",
					APISpecificationType: domain.OpenAPI3,
					PublicSupportContact: "public-support-contact",
					TechSupportContact:   "tech-support-contact",
					Costs: &domain.NewServiceCostsArgs{
						OneTime: 1,
						Monthly: 2,
						Request: 3,
					},
					Availabilities: []*domain.NewServiceAvailability{
						{InwayAddress: "fixture-inway-address-one.com:443",
							State: domain.InwayDOWN},
					},
				},
			},
			organization: &organization{
				SerialNumber: "01234567890123456789",
				Name:         "fixture-organization-name",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage, close := new(t, true)
			defer close()

			organization, err := domain.NewOrganization(tt.organization.Name, tt.organization.SerialNumber)
			assert.NoError(t, err)

			for _, newServiceArgs := range tt.services {
				service, err := domain.NewService(&domain.NewServiceArgs{
					Name:                 newServiceArgs.Name,
					Organization:         organization,
					Internal:             newServiceArgs.Internal,
					DocumentationURL:     newServiceArgs.DocumentationURL,
					APISpecificationType: newServiceArgs.APISpecificationType,
					PublicSupportContact: newServiceArgs.PublicSupportContact,
					TechSupportContact:   newServiceArgs.TechSupportContact,
					Costs:                newServiceArgs.Costs,
					Availabilities:       newServiceArgs.Availabilities,
				})
				assert.NoError(t, err)

				err = storage.RegisterService(service)
				assert.NoError(t, err)
				assertServiceInRepository(t, storage, service)
			}

		})
	}
}
