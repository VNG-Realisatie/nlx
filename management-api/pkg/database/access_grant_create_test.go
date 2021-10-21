// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestCreateAccessGrant(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		accesRequestID uint
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AccessGrant
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accesRequestID: 1,
			},
			want: &database.AccessGrant{
				ID:                      fixturesStartID,
				IncomingAccessRequestID: 1,
				RevokedAt:               sql.NullTime{},
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

			got, err := configDb.CreateAccessGrant(context.Background(), &database.IncomingAccessRequest{
				ID: tt.args.accesRequestID,
			})
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				// NOTE: we're testing individual properties, since we don't have control over the CreatedAt timestamp
				require.False(t, got.CreatedAt.IsZero())

				require.Equal(t, tt.want.ID, got.ID, "id does not match")
				require.Equal(t, tt.want.IncomingAccessRequestID, got.IncomingAccessRequestID)
				require.Equal(t, tt.want.IncomingAccessRequest, got.IncomingAccessRequest)
				require.Equal(t, tt.want.RevokedAt, got.RevokedAt)
			}
		})
	}
}
