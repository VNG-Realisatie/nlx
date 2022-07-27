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
)

func TestGetSettings(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		loadFixtures bool
		settings     func() *domain.Settings
		expectedErr  error
	}{
		"no_settings_present": {
			loadFixtures: false,
			settings: func() *domain.Settings {
				settings, err := domain.NewSettings(
					"",
					"",
				)
				require.NoError(t, err)

				return settings
			},
			expectedErr: nil,
		},
		"happy_flow": {
			loadFixtures: true,
			settings: func() *domain.Settings {
				settings, err := domain.NewSettings(
					"fixture-inway",
					"fixture@example.com",
				)
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

			wantSettings := tt.settings()

			settingsInDB, err := configDb.GetSettings(context.Background())
			require.ErrorIs(t, err, tt.expectedErr)

			if tt.expectedErr == nil {
				require.NoError(t, err)
				require.Equal(t, wantSettings.OrganizationInwayName(), settingsInDB.OrganizationInwayName())
				require.Equal(t, wantSettings.OrganizationEmailAddress(), settingsInDB.OrganizationEmailAddress())
			}
		})
	}
}
