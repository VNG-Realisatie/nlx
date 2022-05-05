// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package pgadapter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-api/domain"
)

func TestRegisterService(t *testing.T) {
	t.Parallel()

	type organization struct {
		SerialNumber string
		Name         string
	}

	tests := map[string]struct {
		services     []*domain.NewServiceArgs
		organization *organization
		expectedErr  error
	}{
		"new_service": {
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
				},
			},
			organization: &organization{
				SerialNumber: testOrganizationSerialNumber,
				Name:         "organization-d",
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
		if err != nil {
			t.Error(err)
		}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage, close := new(t, false)
			defer close()

			organization, err := domain.NewOrganization(tt.organization.SerialNumber, tt.organization.Name)
			assert.NoError(t, err)

			services := make([]*domain.Service, len(tt.services))
			for i, s := range tt.services {
				services[i], err = domain.NewService(&domain.NewServiceArgs{
					Name:                 s.Name,
					Organization:         organization,
					Internal:             s.Internal,
					DocumentationURL:     s.DocumentationURL,
					APISpecificationType: s.APISpecificationType,
					PublicSupportContact: s.PublicSupportContact,
					TechSupportContact:   s.TechSupportContact,
					Costs:                s.Costs,
				})
				assert.NoError(t, err)
			}

			inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
				Name:         "inway-for-service",
				Organization: organization,
				Address:      "my-org.com:443",
				NlxVersion:   domain.NlxVersionUnknown,
				CreatedAt:    now,
				UpdatedAt:    now,
			})
			require.NoError(t, err)

			err = storage.RegisterInway(inwayModel)
			require.NoError(t, err)

			var lastErr error
			for _, s := range services {
				err := storage.RegisterService(s)
				lastErr = err
			}

			require.Equal(t, tt.expectedErr, lastErr)

			if tt.expectedErr == nil {
				lastRegistration := services[len(services)-1]
				assertServiceInRepository(t, storage, lastRegistration)
			}
		})
	}
}
