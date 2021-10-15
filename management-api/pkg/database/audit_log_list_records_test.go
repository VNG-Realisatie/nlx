// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

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

func TestListAuditLogRecords(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	tests := map[string]struct {
		loadFixtures bool
		want         []*database.AuditLog
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			want: []*database.AuditLog{
				{
					ID:         1,
					UserName:   "fixture-user-name",
					ActionType: database.LoginSuccess,
					UserAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36",
					Delegatee:  "fixture-delegatee",
					Data: sql.NullString{
						String: "{}",
						Valid:  true,
					},
					Services: []database.AuditLogService{
						{
							AuditLogID: 1,
							Service:    "fixture-service-name",
							Organization: database.AuditLogServiceOrganization{
								SerialNumber: "00000000000000000001",
								Name:         "fixture-organization-name",
							},
							CreatedAt: fixtureTime,
						},
					},
					CreatedAt: fixtureTime,
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

			got, err := configDb.ListAuditLogRecords(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
