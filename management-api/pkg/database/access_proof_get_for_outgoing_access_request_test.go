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

func TestGetAccessProofForOutgoingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		accessRequestID uint
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.AccessProof
		wantErr      error
	}{
		"when_access_request_not_found": {
			loadFixtures: false,
			args: args{
				accessRequestID: 9999,
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accessRequestID: 1,
			},
			want: &database.AccessProof{
				ID:                      1,
				AccessRequestOutgoingID: 1,
				OutgoingAccessRequest: &database.OutgoingAccessRequest{
					ID:                   1,
					OrganizationName:     "fixture-organization-name",
					ServiceName:          "fixture-service-name",
					ReferenceID:          1,
					State:                database.OutgoingAccessRequestCreated,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					PublicKeyPEM:         fixturePublicKeyPEM,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
				},
				CreatedAt: fixtureTime,
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

			got, err := configDb.GetAccessProofForOutgoingAccessRequest(context.Background(), tt.args.accessRequestID)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
