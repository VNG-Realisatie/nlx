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

func TestCreateAuditLogRecord(t *testing.T) {
	t.Parallel()

	setup(t)

	now := time.Now()

	type args struct {
		auditLog *database.AuditLog
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AuditLog
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				auditLog: &database.AuditLog{
					UserName:   "test-username",
					ActionType: database.IncomingAccessRequestAccept,
					UserAgent:  "test-user-agent",
					Delegatee:  "fixture-organization-name",
					Data:       sql.NullString{},
					Services: []database.AuditLogService{
						{
							AuditLogID: 1,
							Service:    "fixture-service-name",
							Organization: database.AuditLogServiceOrganization{
								SerialNumber: "00000000000000000001",
								Name:         "fixture-organization-name",
							},
							CreatedAt: now,
						},
					},
					CreatedAt: now,
				},
			},
			want: &database.AuditLog{
				ID:         fixturesStartID,
				UserName:   "test-username",
				ActionType: database.IncomingAccessRequestAccept,
				UserAgent:  "test-user-agent",
				Delegatee:  "fixture-organization-name",
				Data:       sql.NullString{},
				Services: []database.AuditLogService{
					{
						AuditLogID: 1,
						Service:    "fixture-service-name",
						Organization: database.AuditLogServiceOrganization{
							SerialNumber: "00000000000000000001",
							Name:         "fixture-organization-name",
						},
						CreatedAt: now,
					},
				},
				CreatedAt: now,
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

			got, err := configDb.CreateAuditLogRecord(context.Background(), tt.args.auditLog)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
