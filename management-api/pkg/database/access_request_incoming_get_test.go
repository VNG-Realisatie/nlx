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

func TestGetIncomingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		id uint
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.IncomingAccessRequest
		wantErr      error
	}{
		"when_the_access_request_does_not_exist": {
			loadFixtures: false,
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				id: 1,
			},
			want: &database.IncomingAccessRequest{
				ID:        1,
				ServiceID: 1,
				Service:   nil,
				Organization: database.IncomingAccessRequestOrganization{
					Name:         "fixture-organization-name",
					SerialNumber: "00000000000000000001",
				},
				State:                database.IncomingAccessRequestReceived,
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
				PublicKeyPEM:         fixturePEM,
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

			got, err := configDb.GetIncomingAccessRequest(context.Background(), tt.args.id)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
