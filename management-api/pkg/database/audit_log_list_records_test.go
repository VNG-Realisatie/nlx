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
					Data: sql.NullString{
						String: marshallMetadata(t, &recordMetadata{
							Delegatee: "fixture-delegatee",
							InwayName: "fixture-inway-name",
							Reference: "fixture-reference",
						}),
						Valid: true,
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
					CreatedAt:    fixtureTime,
					HasSucceeded: true,
				},
				{
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
					HasSucceeded: false,
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

			numberOfAuditLogs := 100

			got, err := configDb.ListAuditLogRecords(context.Background(), numberOfAuditLogs)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
