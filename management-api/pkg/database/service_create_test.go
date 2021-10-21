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

func TestCreateService(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	require.NoError(t, err)

	type args struct {
		service *database.Service
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		wantErr      error
	}{
		"when_service_already_exist": {
			loadFixtures: true,
			args: args{
				service: &database.Service{
					ID:                   1,
					Name:                 "service-name",
					EndpointURL:          "http://api:8000",
					DocumentationURL:     "https://documentation-url.com",
					APISpecificationURL:  "http://api:8000/schema?format=openapi-json",
					Internal:             false,
					TechSupportContact:   "test@tech-support-contact.com",
					PublicSupportContact: "test@public-support-contact.com",
					Inways:               []*database.Inway{},
					OneTimeCosts:         1,
					MonthlyCosts:         2,
					RequestCosts:         3,
					CreatedAt:            now,
					UpdatedAt:            now,
				},
			},
			wantErr: database.ErrServiceAlreadyExists,
		},
		"happy_flow": {
			loadFixtures: false,
			args: args{
				service: &database.Service{
					ID:                   1,
					Name:                 "service-name",
					EndpointURL:          "http://api:8000",
					DocumentationURL:     "https://documentation-url.com",
					APISpecificationURL:  "http://api:8000/schema?format=openapi-json",
					Internal:             false,
					TechSupportContact:   "test@tech-support-contact.com",
					PublicSupportContact: "test@public-support-contact.com",
					Inways:               []*database.Inway{},
					OneTimeCosts:         1,
					MonthlyCosts:         2,
					RequestCosts:         3,
					CreatedAt:            now,
					UpdatedAt:            now,
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

			err := configDb.CreateService(context.Background(), tt.args.service)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertService(t, configDb, tt.args.service)
			}
		})
	}
}

func assertService(t *testing.T, repo database.ConfigDatabase, want *database.Service) {
	got, err := repo.GetService(context.Background(), want.Name)
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.EqualValues(t, want, got)
}
