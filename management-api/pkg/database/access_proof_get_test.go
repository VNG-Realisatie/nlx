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

func TestGetAccessProofs(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	type args struct {
		accessProofIDs []uint64
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         []*database.AccessProof
	}{
		"when_access_request_not_found": {
			loadFixtures: false,
			args: args{
				accessProofIDs: []uint64{9999},
			},
			want: []*database.AccessProof{},
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accessProofIDs: []uint64{1, 2},
			},
			want: []*database.AccessProof{{
				ID:                      1,
				AccessRequestOutgoingID: 1,
				OutgoingAccessRequest: &database.OutgoingAccessRequest{
					ID: 1,
					Organization: database.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "fixture-organization-name",
					},
					ServiceName:          "fixture-service-name",
					ReferenceID:          1,
					State:                database.OutgoingAccessRequestReceived,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
				},
				CreatedAt: fixtureTime,
				RevokedAt: sql.NullTime{},
			}, {
				ID:                      2,
				AccessRequestOutgoingID: 1,
				OutgoingAccessRequest: &database.OutgoingAccessRequest{
					ID: 1,
					Organization: database.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "fixture-organization-name",
					},
					ServiceName:          "fixture-service-name",
					ReferenceID:          1,
					State:                database.OutgoingAccessRequestReceived,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
				},
				CreatedAt: fixtureTime,
				RevokedAt: sql.NullTime{
					Valid: true,
					Time:  fixtureTime,
				},
			}},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			got, err := configDb.GetAccessProofs(context.Background(), tt.args.accessProofIDs)
			require.NoError(t, err)

			require.EqualValues(t, tt.want, got)

		})
	}
}
