// Copyright © VNG Realisatie 2022
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

func TestListInways(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	tests := map[string]struct {
		loadFixtures bool
		want         []*database.Inway
	}{
		"no_inways_present": {
			loadFixtures: false,
			want:         []*database.Inway{},
		},
		"happy_flow": {
			loadFixtures: true,
			want: []*database.Inway{
				{
					ID:          1,
					Name:        "fixture-inway",
					Version:     "unknown",
					Hostname:    "fixture-server",
					IPAddress:   "127.0.0.1",
					SelfAddress: "fixture.local",
					Services: []*database.Service{
						{
							ID:                   1,
							Name:                 "fixture-service-name",
							EndpointURL:          "http://fixture-api:8000",
							DocumentationURL:     "https://fixture-documentation-url.com",
							APISpecificationURL:  "http://fixture-api:8000/schema?format=openapi-json",
							Internal:             false,
							TechSupportContact:   "fixture@tech-support-contact.com",
							PublicSupportContact: "fixture@public-support-contact.com",
							Inways:               nil,
							OneTimeCosts:         1,
							MonthlyCosts:         2,
							RequestCosts:         3,
							CreatedAt:            fixtureTime,
							UpdatedAt:            fixtureTime,
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
							Inways:               nil,
							OneTimeCosts:         1,
							MonthlyCosts:         2,
							RequestCosts:         3,
							CreatedAt:            fixtureTime,
							UpdatedAt:            fixtureTime,
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
							Inways:               nil,
							OneTimeCosts:         1,
							MonthlyCosts:         2,
							RequestCosts:         3,
							CreatedAt:            fixtureTime,
							UpdatedAt:            fixtureTime,
						},
					},
					CreatedAt: fixtureTime,
					UpdatedAt: fixtureTime,
				},
				{
					ID:          2,
					Version:     "unknown",
					Hostname:    "fixture-server",
					IPAddress:   "127.0.0.1",
					SelfAddress: "fixture.local",
					Services:    []*database.Service{},
					CreatedAt:   fixtureTime,
					UpdatedAt:   fixtureTime,
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			got, err := configDb.ListInways(context.Background())

			require.NoError(t, err)
			require.EqualValues(t, tt.want, got)
		})
	}
}
