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

func TestCreateAccessProof(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		accessRequestOutgoingID uint
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AccessProof
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accessRequestOutgoingID: 1,
			},
			want: &database.AccessProof{
				ID:                      fixturesStartID,
				AccessRequestOutgoingID: 1,
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

			got, err := configDb.CreateAccessProof(context.Background(), tt.args.accessRequestOutgoingID)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				// NOTE: we're testing individual properties, since we don't have control over the CreatedAt timestamp
				require.False(t, got.CreatedAt.IsZero())

				require.Equal(t, tt.want.ID, got.ID)
				require.Equal(t, tt.want.AccessRequestOutgoingID, got.AccessRequestOutgoingID)
			}
		})
	}
}
