// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

type recordMetadata struct {
	Delegatee            string `json:"delegatee,omitempty"`
	InwayName            string `json:"inwayName,omitempty"`
	Reference            string `json:"reference,omitempty"`
	PublicKeyFingerprint string `json:"publicKeyFingerprint,omitempty"`
}

func marshallMetadata(t *testing.T, metadata *recordMetadata) string {
	data, err := json.Marshal(metadata)

	require.NoError(t, err)

	return string(data)
}

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
			loadFixtures: false,
			args: args{
				auditLog: &database.AuditLog{
					UserName:   "test-username",
					ActionType: database.IncomingAccessRequestAccept,
					UserAgent:  "test-user-agent",
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
							CreatedAt: now,
						},
					},
					CreatedAt: now,
				},
			},
			want: &database.AuditLog{
				ID:         1,
				UserName:   "test-username",
				ActionType: database.IncomingAccessRequestAccept,
				UserAgent:  "test-user-agent",
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
						CreatedAt: stripNanosecondsFromTime(now).UTC(),
					},
				},
				CreatedAt: stripNanosecondsFromTime(now).UTC(),
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

			_, err := configDb.CreateAuditLogRecord(context.Background(), tt.args.auditLog)
			require.ErrorIs(t, err, tt.wantErr)

			auditLogRecords, err := configDb.ListAuditLogRecords(context.Background(), 1)
			require.NoError(t, err)
			require.NotEmpty(t, auditLogRecords)

			// explicitly added check for created at, since this failed a couple of
			// times in the past. Want to keep this for some time as it
			// provides extra information while debugging
			require.Equal(t, tt.want.CreatedAt, auditLogRecords[0].CreatedAt)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, auditLogRecords[0])
			}
		})
	}
}

func stripNanosecondsFromTime(input time.Time) time.Time {
	return input.Round(time.Microsecond)
}
