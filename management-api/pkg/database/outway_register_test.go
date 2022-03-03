// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestRegisterOutway(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	tests := map[string]struct {
		loadFixtures bool
		arg          *database.Outway
		want         *database.Outway
		wantErr      error
	}{
		"happy_flow_update_existing": {
			loadFixtures: true,
			arg: &database.Outway{
				Name:                 "fixture-outway-1",
				PublicKeyPEM:         "foobar",
				PublicKeyFingerprint: "fingerprint",
				SelfAddressAPI:       "new-self-address",
				IPAddress:            mockIP(t, "127.0.0.2/32"),
				Version:              "1.0.0",
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
			},
			want: &database.Outway{
				ID:                   1,
				Name:                 "fixture-outway-1",
				PublicKeyPEM:         "foobar",
				PublicKeyFingerprint: "fingerprint",
				SelfAddressAPI:       "new-self-address",
				IPAddress:            mockIP(t, "127.0.0.2/32"),
				Version:              "1.0.0",
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
			},
			wantErr: nil,
		},
		"happy_flow": {
			loadFixtures: false,
			arg: &database.Outway{
				Name:                 "register-outway-1",
				PublicKeyPEM:         fixturePublicKeyPEM,
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
				IPAddress:            mockIP(t, "127.0.0.1/32"),
				SelfAddressAPI:       "self-address",
				Version:              "unknown",
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
			},
			want: &database.Outway{
				ID:                   1,
				Name:                 "register-outway-1",
				PublicKeyPEM:         fixturePublicKeyPEM,
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
				IPAddress:            mockIP(t, "127.0.0.1/32"),
				SelfAddressAPI:       "self-address",
				Version:              "unknown",
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
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

			err := configDb.RegisterOutway(context.Background(), tt.arg)
			require.ErrorIs(t, err, tt.wantErr)

			assertOutway(t, configDb, tt.want)

		})
	}
}

func assertOutway(t *testing.T, repo database.ConfigDatabase, want *database.Outway) {
	got, err := repo.GetOutway(context.Background(), want.Name)
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.EqualValues(t, want, got)
}
