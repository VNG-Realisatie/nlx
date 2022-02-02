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

func TestGetOutway(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		name string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.Outway
		wantErr      error
	}{
		"when_not_found": {
			loadFixtures: false,
			args: args{
				name: "arbitrary-name",
			},
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				name: "fixture-outway-1",
			},
			want: &database.Outway{
				ID:                   1,
				Name:                 "fixture-outway-1",
				PublicKeyPEM:         fixturePublicKeyPEM,
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
				IPAddress:            mockIP(t, "127.0.0.1/32"),
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

			got, err := configDb.GetOutway(context.Background(), tt.args.name)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
