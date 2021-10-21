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

func TestGetLatestOutgoingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getCustomFixtureTime(t, "2021-01-03T01:02:03Z")

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		organizationSerialNumber string
		serviceName              string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.OutgoingAccessRequest
		wantErr      error
	}{
		"when_there_are_no_access_requests_present": {
			loadFixtures: false,
			args: args{
				organizationSerialNumber: "arbitrary",
				serviceName:              "arbitrary",
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationSerialNumber: "00000000000000000001",
				serviceName:              "fixture-service-name",
			},
			want: &database.OutgoingAccessRequest{
				ID: 2,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "fixture-organization-name",
				},
				ServiceName:          "fixture-service-name",
				ReferenceID:          1,
				State:                database.OutgoingAccessRequestApproved,
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

			got, err := configDb.GetLatestOutgoingAccessRequest(context.Background(), tt.args.organizationSerialNumber, tt.args.serviceName)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
