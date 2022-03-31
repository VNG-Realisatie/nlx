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

func TestGetInway(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	type args struct {
		name string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.Inway
		wantErr      error
	}{
		"inway_does_not_exist": {
			loadFixtures: false,
			args:         args{name: "arbitrary-name"},
			want:         &database.Inway{},
			wantErr:      database.ErrNotFound,
		},
		"empty_inway_name": {
			loadFixtures: true,
			args:         args{name: ""},
			want: &database.Inway{
				ID:          2,
				Name:        "",
				Version:     "unknown",
				Hostname:    "fixture-server",
				IPAddress:   "127.0.0.1",
				SelfAddress: "fixture.local:443",
				Services:    []*database.Service{},
				CreatedAt:   fixtureTime,
				UpdatedAt:   fixtureTime,
			},
		},
		"happy_flow": {
			loadFixtures: true,
			args:         args{name: "fixture-inway"},
			want: &database.Inway{
				ID:          1,
				Name:        "fixture-inway",
				Version:     "unknown",
				Hostname:    "fixture-server",
				IPAddress:   "127.0.0.1",
				SelfAddress: "fixture.local:443",
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
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			got, err := configDb.GetInway(context.Background(), tt.args.name)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
