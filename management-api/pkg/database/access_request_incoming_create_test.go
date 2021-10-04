// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestCreateIncomingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		accessRequest *database.IncomingAccessRequest
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.IncomingAccessRequest
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accessRequest: &database.IncomingAccessRequest{
					ServiceID: 1,
					Organization: database.IncomingAccessRequestOrganization{
						Name:         "organization-name",
						SerialNumber: "00000000000000000001",
					},
					State:                database.IncomingAccessRequestReceived,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
					PublicKeyPEM:         fixturePublicKeyPEM,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
				},
			},
			want: &database.IncomingAccessRequest{
				ID:        fixturesStartID,
				ServiceID: 1,
				Service:   nil,
				Organization: database.IncomingAccessRequestOrganization{
					Name:         "organization-name",
					SerialNumber: "00000000000000000001",
				},
				State:                database.IncomingAccessRequestReceived,
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
				PublicKeyPEM:         fixturePublicKeyPEM,
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
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

			got, err := configDb.CreateIncomingAccessRequest(context.Background(), tt.args.accessRequest)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
