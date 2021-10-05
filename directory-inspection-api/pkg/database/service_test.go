// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

func TestListServices(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		organizationSerialNumber string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         []*database.Service
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
			want: []*database.Service{
				{
					Name: "fixture-service-name",
					Organization: &database.Organization{
						Name:         "fixture-organization-name",
						SerialNumber: "01234567890123456789",
					},
					EndpointURL:          "",
					DocumentationURL:     "https://fixture-documentation-url.com",
					APISpecificationURL:  "",
					APISpecificationType: "OpenAPI3",
					Internal:             false,
					TechSupportContact:   "",
					PublicSupportContact: "fixture@public-support-contact.com",
					Costs: &database.ServiceCosts{
						OneTime: 1,
						Monthly: 2,
						Request: 3,
					},
					Inways: []*database.Inway{
						{
							Address: "https://fixture-inway-address.com",
							State:   database.InwayUP,
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

			db, close := newDirectoryDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			got, err := db.ListServices(context.Background(), tt.args.organizationSerialNumber)
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, tt.want, got)
			}
		})
	}
}
