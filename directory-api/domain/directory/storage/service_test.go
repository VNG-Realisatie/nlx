// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package storage_test

import (
	"context"
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
				Address:      "my-org.com",
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

func TestListServices(t *testing.T) {
	t.Parallel()

	type args struct {
		organizationSerialNumber string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         []*domain.NewServiceArgs
		wantErr      error
	}{
		"when_organization_not_found": {
			loadFixtures: false,
			args: args{
				organizationSerialNumber: "99999999999999999999",
			},
			want:    nil,
			wantErr: nil,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationSerialNumber: "01234567890123456789",
			},
			want: []*domain.NewServiceArgs{
				{
					Name:                 "fixture-service-name",
					Organization:         createNewOrganization(t, "fixture-organization-name", testOrganizationSerialNumber),
					DocumentationURL:     "https://fixture-documentation-url.com",
					APISpecificationType: "OpenAPI3",
					Internal:             false,
					TechSupportContact:   "",
					PublicSupportContact: "fixture@public-support-contact.com",
					Costs: &domain.NewServiceCostsArgs{
						OneTime: 1,
						Monthly: 2,
						Request: 3,
					},
					Inways: []*domain.NewServiceInwayArgs{
						{
							Address: "https://fixture-inway-address-one.com",
							State:   domain.InwayUP,
						},
					},
				},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, close := new(t, tt.loadFixtures)
			defer close()

			want := make([]*domain.Service, len(tt.want))

			for i, s := range tt.want {
				var err error
				want[i], err = domain.NewService(s)
				require.NoError(t, err)
			}

			got, err := db.ListServices(context.Background(), tt.args.organizationSerialNumber)
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, want, got)
			}
		})
	}
}
