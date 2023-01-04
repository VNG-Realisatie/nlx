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

func TestCreateOutgoingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		accessRequest *database.OutgoingAccessRequest
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.OutgoingAccessRequest
		wantErr      error
	}{
		"when_an_outgoing_access_request_for_the_service_is_already_present": {
			loadFixtures: true,
			args: args{
				accessRequest: &database.OutgoingAccessRequest{
					Organization: database.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "fixture-organization-name",
					},
					ServiceName:          "fixture-service-name",
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					PublicKeyPEM:         fixturePEM,
					State:                database.OutgoingAccessRequestReceived,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
				},
			},
			wantErr: database.ErrActiveAccessRequest,
		},
		"happy_flow": {
			loadFixtures: false,
			args: args{
				accessRequest: &database.OutgoingAccessRequest{
					Organization: database.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "my-org",
					},
					ServiceName:          "my-service",
					PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
					PublicKeyPEM:         fixturePEM,
					State:                database.OutgoingAccessRequestReceived,
					CreatedAt:            fixtureTime,
					UpdatedAt:            fixtureTime,
					ReferenceID:          42,
				},
			},
			want: &database.OutgoingAccessRequest{
				ID: nonFixturesStartID,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "my-org",
				},
				ServiceName:          "my-service",
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
				PublicKeyPEM:         fixturePEM,
				State:                database.OutgoingAccessRequestReceived,
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
				ReferenceID:          42,
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

			got, err := configDb.CreateOutgoingAccessRequest(context.Background(), tt.args.accessRequest)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
