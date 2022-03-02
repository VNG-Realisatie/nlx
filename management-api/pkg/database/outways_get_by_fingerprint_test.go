// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestGetOutwaysByPublicKeyFingerprint(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	tests := map[string]struct {
		loadFixtures         bool
		publicKeyFingerprint string
		want                 []*database.Outway
		wantErr              error
	}{
		"when_not_found": {
			loadFixtures:         false,
			publicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
			want:                 nil,
			wantErr:              database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures:         true,
			publicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
			want: []*database.Outway{
				{
					ID:                   1,
					Name:                 "fixture-outway-1",
					PublicKeyPEM:         fixturePublicKeyPEM,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					IPAddress:            mockIP(t, "127.0.0.1/32"),
					SelfAddress:          "self-address",
					Version:              "unknown",
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
				},
				{
					ID:                   2,
					Name:                 "fixture-outway-2",
					IPAddress:            mockIP(t, "127.0.0.2/32"),
					PublicKeyPEM:         fixturePublicKeyPEM,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					SelfAddress:          "self-address-2",
					Version:              "1.2.3",
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
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

			got, err := configDb.GetOutwaysByPublicKeyFingerprint(context.Background(), tt.publicKeyFingerprint)

			require.Equal(t, err, tt.wantErr)
			require.Equal(t, tt.want, got)
		})
	}
}
