// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestCreateServiceWithInways(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	require.NoError(t, err)

	fixtureTime := getFixtureTime(t)

	type args struct {
		service    *database.Service
		inwayNames []string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.Service
		wantErr      error
	}{
		"when_inway_not_found": {
			loadFixtures: false,
			args: args{
				service: &database.Service{
					Name:                 "service-name",
					EndpointURL:          "http://api:8000",
					DocumentationURL:     "https://documentation-url.com",
					APISpecificationURL:  "http://api:8000/schema?format=openapi-json",
					Internal:             false,
					TechSupportContact:   "test@tech-support-contact.com",
					PublicSupportContact: "test@public-support-contact.com",
					OneTimeCosts:         1,
					MonthlyCosts:         2,
					RequestCosts:         3,
					CreatedAt:            now,
					UpdatedAt:            now,
				},
				inwayNames: []string{
					"arbitrary",
				},
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				service: &database.Service{
					Name:                 "service-name",
					EndpointURL:          "http://api:8000",
					DocumentationURL:     "https://documentation-url.com",
					APISpecificationURL:  "http://api:8000/schema?format=openapi-json",
					Internal:             false,
					TechSupportContact:   "test@tech-support-contact.com",
					PublicSupportContact: "test@public-support-contact.com",
					OneTimeCosts:         1,
					MonthlyCosts:         2,
					RequestCosts:         3,
					CreatedAt:            now,
					UpdatedAt:            now,
				},
				inwayNames: []string{
					"fixture-inway",
				},
			},
			want: &database.Service{
				ID:                   fixturesStartID,
				Name:                 "service-name",
				EndpointURL:          "http://api:8000",
				DocumentationURL:     "https://documentation-url.com",
				APISpecificationURL:  "http://api:8000/schema?format=openapi-json",
				Internal:             false,
				TechSupportContact:   "test@tech-support-contact.com",
				PublicSupportContact: "test@public-support-contact.com",
				Inways: []*database.Inway{
					{
						ID:          1,
						Name:        "fixture-inway",
						Version:     "unknown",
						Hostname:    "fixture-server",
						IPAddress:   "127.0.0.1",
						SelfAddress: "fixture.local",
						CreatedAt:   fixtureTime,
						UpdatedAt:   fixtureTime,
					},
				},
				OneTimeCosts: 1,
				MonthlyCosts: 2,
				RequestCosts: 3,
				CreatedAt:    now,
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

			err := configDb.CreateServiceWithInways(context.Background(), tt.args.service, tt.args.inwayNames)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertServiceWithInways(t, configDb, tt.args.service.Name, tt.want)
			}
		})
	}
}

func assertServiceWithInways(t *testing.T, repo database.ConfigDatabase, serviceName string, want *database.Service) {
	got, err := repo.GetService(context.Background(), serviceName)
	require.NoError(t, err)
	require.NotNil(t, got)

	// Overwrite updatedAt because we don't know what the value will be
	got.UpdatedAt = time.Time{}

	assert.EqualValues(t, want, got)
}
