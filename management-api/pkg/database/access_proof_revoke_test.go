// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestRevokeAccessProof(t *testing.T) {
	t.Parallel()

	setup(t)

	now := time.Now()

	fixtureTime := getFixtureTime(t)

	type args struct {
		accessProofID uint
		revokedAt     time.Time
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AccessProof
		wantErr      error
	}{
		"when_access_grant_does_not_exist": {
			loadFixtures: false,
			args: args{
				accessProofID: 9999,
				revokedAt:     now,
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accessProofID: 1,
				revokedAt:     now,
			},
			want: &database.AccessProof{
				ID:                      1,
				AccessRequestOutgoingID: 1,
				CreatedAt:               fixtureTime,
				RevokedAt:               sql.NullTime{Time: now, Valid: true},
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

			got, err := configDb.RevokeAccessProof(context.Background(), tt.args.accessProofID, tt.args.revokedAt)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
