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
		organizationName string
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
				organizationName: "arbritrary organization name",
			},
			want:    nil,
			wantErr: nil,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationName: "fixture-name",
			},
			want: []*database.Service{
				{
					Name:                  "fixture-service-name",
					Organization:          "fixture-organization-name",
					EndpointURL:           "",
					DocumentationURL:      "https://fixture-documentation-url.com",
					APISpecificationURL:   "",
					APISpecificationType:  "OpenAPI3",
					Internal:              false,
					TechSupportContact:    "",
					PublicSupportContact:  "fixture@public-support-contact.com",
					OneTimeCosts:          1,
					MonthlyCosts:          2,
					RequestCosts:          3,
					AuthorizationSettings: nil,
					Inways: []*database.Inway{
						{
							Address: "https://fixture-address.com",
							State:   database.InwayUP,
						},
					},
					InwayAddresses: []string{
						"https://fixture-address.com",
					},
					HealthyStates: []bool{
						true,
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

			got, err := db.ListServices(context.Background(), tt.args.organizationName)
			require.Equal(t, tt.wantErr, err)

			if tt.wantErr == nil {
				assert.EqualValues(t, tt.want, got)
			}
		})
	}
}
