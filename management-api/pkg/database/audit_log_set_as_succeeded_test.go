// Copyright © VNG Realisatie 2022
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

func Test_SetAsSucceeded(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	type args struct {
		id uint64
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AuditLog
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args:         args{id: uint64(1)},
			want: &database.AuditLog{
				ID:         2,
				UserName:   "fixture-user-name",
				ActionType: database.OutgoingAccessRequestWithdraw,
				UserAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36",
				Data: sql.NullString{
					String: marshallMetadata(t, &recordMetadata{
						PublicKeyFingerprint: "WoiWRyIOVNa9ihaBciRSC7XHjliYS9VwUGOIud4PB18=",
					}),
					Valid: true,
				},
				Services:     []database.AuditLogService{},
				CreatedAt:    fixtureTime,
				HasSucceeded: true,
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

			numberOfAuditLogs := 100

			err := configDb.SetAuditLogAsSucceeded(context.Background(), int64(tt.want.ID))
			require.NoError(t, err)

			auditLogRecords, err := configDb.ListAuditLogRecords(context.Background(), numberOfAuditLogs)
			require.NoError(t, err)

			var got *database.AuditLog

			for _, record := range auditLogRecords {
				if record.ID == tt.want.ID {
					got = record
				}
			}

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
