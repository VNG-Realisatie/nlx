// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestRegisterInway(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	tests := map[string]struct {
		loadFixtures    bool
		inwayToRegister *database.Inway
		want            *database.Inway
	}{
		"happy_flow_insert": {
			loadFixtures: false,
			inwayToRegister: &database.Inway{
				Name:        "Inway-1",
				Version:     "1",
				Hostname:    "localhost",
				IPAddress:   "127.0.0.1",
				SelfAddress: "inway.local:443",
				CreatedAt:   fixtureTime,
				UpdatedAt:   fixtureTime,
			},
			want: &database.Inway{
				ID:          1,
				Name:        "Inway-1",
				Version:     "1",
				Hostname:    "localhost",
				IPAddress:   "127.0.0.1",
				SelfAddress: "inway.local:443",
				Services:    []*database.Service{},
				CreatedAt:   fixtureTime,
				UpdatedAt:   fixtureTime,
			},
		},
		"happy_flow_update": {
			loadFixtures: true,
			inwayToRegister: &database.Inway{
				Name:        "fixture-inway-3",
				Version:     "1",
				Hostname:    "localhost",
				IPAddress:   "127.0.0.2",
				SelfAddress: "inway.local:443",
				CreatedAt:   time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
			},
			want: &database.Inway{
				ID:          3,
				Name:        "fixture-inway-3",
				Version:     "1",
				Hostname:    "localhost",
				IPAddress:   "127.0.0.2",
				SelfAddress: "inway.local:443",
				Services:    []*database.Service{},
				CreatedAt:   fixtureTime,
				UpdatedAt:   time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			err := configDb.RegisterInway(context.Background(), tt.inwayToRegister)
			require.NoError(t, err)

			got, err := configDb.GetInway(context.Background(), tt.inwayToRegister.Name)

			require.NoError(t, err)
			require.EqualValues(t, tt.want, got)
		})
	}
}
