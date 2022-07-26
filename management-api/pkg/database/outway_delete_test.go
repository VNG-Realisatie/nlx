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

func TestDeleteOutway(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		loadFixtures bool
		outwayName   string
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			outwayName:   "fixture-outway-1",
			wantErr:      nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			err := configDb.DeleteOutway(context.Background(), tt.outwayName)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertOutwayDeleted(t, configDb, tt.outwayName)
			}
		})
	}
}

func assertOutwayDeleted(t *testing.T, repo database.ConfigDatabase, outwayName string) {
	_, err := repo.GetOutway(context.Background(), outwayName)
	require.Equal(t, database.ErrNotFound, err)

	outways, err := repo.ListOutways(context.Background())
	require.NoError(t, err)

	for _, o := range outways {
		if o.Name == outwayName {
			t.Errorf("outway %q is not deleted", outwayName)
		}
	}

}
