// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestListServices(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	tests := map[string]struct {
		loadFixtures bool
		want         []*database.Service
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			want: []*database.Service{
				{
					ID:                   1,
					Name:                 "fixture-service-name",
					EndpointURL:          "http://fixture-api:8000",
					DocumentationURL:     "https://fixture-documentation-url.com",
					APISpecificationURL:  "http://fixture-api:8000/schema?format=openapi-json",
					Internal:             false,
					TechSupportContact:   "fixture@tech-support-contact.com",
					PublicSupportContact: "fixture@public-support-contact.com",
					Inways: []*database.Inway{
						{
							ID:          1,
							Name:        "fixture-inway",
							Version:     "unknown",
							Hostname:    "fixture-server",
							IPAddress:   "127.0.0.1",
							SelfAddress: "fixture.local:443",
							CreatedAt:   fixtureTime,
							UpdatedAt:   fixtureTime,
						},
					},
					OneTimeCosts: 1,
					MonthlyCosts: 2,
					RequestCosts: 3,
					CreatedAt:    fixtureTime,
					UpdatedAt:    fixtureTime,
				},
				{
					ID:                   2,
					Name:                 "fixture-service-name-two",
					EndpointURL:          "http://fixture-api:8000",
					DocumentationURL:     "https://fixture-documentation-url.com",
					APISpecificationURL:  "http://fixture-api:8000/schema?format=openapi-json",
					Internal:             false,
					TechSupportContact:   "fixture@tech-support-contact.com",
					PublicSupportContact: "fixture@public-support-contact.com",
					Inways: []*database.Inway{
						{
							ID:          1,
							Name:        "fixture-inway",
							Version:     "unknown",
							Hostname:    "fixture-server",
							IPAddress:   "127.0.0.1",
							SelfAddress: "fixture.local:443",
							CreatedAt:   fixtureTime,
							UpdatedAt:   fixtureTime,
						},
					},
					OneTimeCosts: 1,
					MonthlyCosts: 2,
					RequestCosts: 3,
					CreatedAt:    fixtureTime,
					UpdatedAt:    fixtureTime,
				},
				{
					ID:                   3,
					Name:                 "fixture-service-name-three",
					EndpointURL:          "http://fixture-api:8000",
					DocumentationURL:     "https://fixture-documentation-url.com",
					APISpecificationURL:  "http://fixture-api:8000/schema?format=openapi-json",
					Internal:             false,
					TechSupportContact:   "fixture@tech-support-contact.com",
					PublicSupportContact: "fixture@public-support-contact.com",
					Inways: []*database.Inway{
						{
							ID:          1,
							Name:        "fixture-inway",
							Version:     "unknown",
							Hostname:    "fixture-server",
							IPAddress:   "127.0.0.1",
							SelfAddress: "fixture.local:443",
							CreatedAt:   fixtureTime,
							UpdatedAt:   fixtureTime,
						},
					},
					OneTimeCosts: 1,
					MonthlyCosts: 2,
					RequestCosts: 3,
					CreatedAt:    fixtureTime,
					UpdatedAt:    fixtureTime,
				},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			got, err := configDb.ListServices(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
