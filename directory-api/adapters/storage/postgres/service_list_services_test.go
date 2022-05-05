// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package pgadapter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-api/domain"
)

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
							Address: "fixture-inway-address-one.com:443",
							State:   domain.InwayDOWN,
						},
						{
							Address: "fixture-inway-address-two.com:443",
							State:   domain.InwayDOWN,
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

			repo, close := new(t, tt.loadFixtures)
			defer close()

			want := make([]*domain.Service, len(tt.want))

			for i, s := range tt.want {
				var err error
				want[i], err = domain.NewService(s)
				require.NoError(t, err)
			}

			got, err := repo.ListServices(context.Background(), tt.args.organizationSerialNumber)
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, want, got)
			}
		})
	}
}
