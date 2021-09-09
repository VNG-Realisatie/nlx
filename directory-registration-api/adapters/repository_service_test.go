// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package adapters_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.nlx.io/nlx/directory-registration-api/domain/inway"
	"go.nlx.io/nlx/directory-registration-api/domain/service"
)

func TestRegisterService(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		createRegistrations func(*testing.T) []*service.Service
		expectedErr         error
	}{
		"new_service": {
			createRegistrations: func(t *testing.T) []*service.Service {
				s, err := service.NewService(
					&service.NewServiceArgs{
						Name:                 "my-service",
						OrganizationName:     "organization-d",
						Internal:             true,
						DocumentationURL:     "documentation-url",
						APISpecificationType: service.OpenAPI3,
						PublicSupportContact: "public-support-contact",
						TechSupportContact:   "tech-support-contact",
						OneTimeCosts:         1,
						MonthlyCosts:         2,
						RequestCosts:         3,
					},
				)
				require.NoError(t, err)

				return []*service.Service{s}
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

			repo, close := newRepo(t, t.Name())
			defer close()

			models := tt.createRegistrations(t)

			inwayModel, err := inway.NewInway(&inway.NewInwayArgs{
				Name:             "inway-for-service",
				OrganizationName: "organization-d",
				Address:          "my-org.com",
				NlxVersion:       inway.NlxVersionUnknown,
				CreatedAt:        now,
				UpdatedAt:        now,
			})
			require.NoError(t, err)

			err = repo.RegisterInway(inwayModel)
			require.NoError(t, err)

			var lastErr error
			for _, model := range models {
				err := repo.RegisterService(model)
				lastErr = err
			}

			require.Equal(t, tt.expectedErr, lastErr)

			if tt.expectedErr == nil {
				lastRegistration := models[len(models)-1]
				assertServiceInRepository(t, repo, lastRegistration)
			}
		})
	}
}
