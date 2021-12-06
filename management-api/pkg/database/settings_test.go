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

	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestUpdateSettings(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		loadFixtures bool
		settings     func() *domain.Settings
		expectedErr  error
	}{
		"non_existing_inway": {
			loadFixtures: true,
			settings: func() *domain.Settings {
				settings, err := domain.NewSettings("does-not-exist", "mock@email.com")
				require.NoError(t, err)

				return settings
			},
			expectedErr: database.ErrInwayNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			settings: func() *domain.Settings {
				settings, err := domain.NewSettings("fixture-inway", "mock@email.com")
				require.NoError(t, err)

				return settings
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			settings := tt.settings()
			err := configDb.UpdateSettings(context.Background(), settings)
			require.ErrorIs(t, err, tt.expectedErr)

			if tt.expectedErr == nil {
				settingsInDB, err := configDb.GetSettings(context.Background())
				require.NoError(t, err)
				require.Equal(t, settings.OrganizationInwayName(), settingsInDB.OrganizationInwayName())
				require.Equal(t, settings.OrganizationEmailAddress(), settingsInDB.OrganizationEmailAddress())
			}
		})
	}
}

func newUint(x uint) *uint {
	return &x
}
