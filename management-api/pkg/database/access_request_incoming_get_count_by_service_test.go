// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestGetIncomingAccessRequestCountByService(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		loadFixtures bool
		want         map[string]int
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			want: map[string]int{
				"fixture-service-name":       1,
				"fixture-service-name-two":   0,
				"fixture-service-name-three": 0,
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

			got, err := configDb.GetIncomingAccessRequestCountByService(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
