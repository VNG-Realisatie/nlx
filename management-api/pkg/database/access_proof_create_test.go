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

func TestCreateAccessProof(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		accessRequest *database.OutgoingAccessRequest
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
				accessRequest: &database.OutgoingAccessRequest{
					ID: 1,
					Organization: database.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "fixture-organization-name",
					},
					ServiceName:          "fixture-service-name",
					ReferenceID:          1,
					State:                database.OutgoingAccessRequestCreated,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					PublicKeyPEM:         fixturePublicKeyPEM,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
				},
			},
			want: &database.AccessProof{
				ID:                      fixturesStartID,
				AccessRequestOutgoingID: 1,
				OutgoingAccessRequest: &database.OutgoingAccessRequest{
					ID: 1,
					Organization: database.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "fixture-organization-name",
					},
					ServiceName:          "fixture-service-name",
					ReferenceID:          1,
					State:                database.OutgoingAccessRequestCreated,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					PublicKeyPEM:         fixturePublicKeyPEM,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
				},
				RevokedAt: sql.NullTime{},
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

			got, err := configDb.CreateAccessProof(context.Background(), tt.args.accessRequest)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want.ID, got.ID)
				require.Equal(t, tt.want.AccessRequestOutgoingID, got.AccessRequestOutgoingID)
				require.EqualValues(t, tt.want.OutgoingAccessRequest, got.OutgoingAccessRequest)
				require.False(t, got.CreatedAt.IsZero())
				require.Equal(t, tt.want.RevokedAt, got.RevokedAt)
			}
		})
	}
}
