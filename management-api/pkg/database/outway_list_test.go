// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"net"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgtype"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func mockIP(t *testing.T, ip string) pgtype.Inet {
	_, ipNet, err := net.ParseCIDR(ip)
	require.NoError(t, err)

	return pgtype.Inet{
		Status: pgtype.Present,
		IPNet:  ipNet,
	}
}
func TestListOutways(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	tests := map[string]struct {
		loadFixtures bool
		want         []*database.Outway
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			want: []*database.Outway{
				{
					ID:                   1,
					Name:                 "fixture-outway-1",
					PublicKeyPEM:         fixturePublicKeyPEM,
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					IPAddress:            mockIP(t, "127.0.0.1/32"),
					SelfAddressAPI:       "self-address-api",
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
					SelfAddressAPI:       "self-address-api-2",
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

			got, err := configDb.ListOutways(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
