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

func TestListLatestOutgoingAccessRequests(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)
	fixtureCustomTime := getCustomFixtureTime(t, "2021-01-03T01:02:03Z")

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	type args struct {
		serviceName              string
		organizationSerialNumber string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         []*database.OutgoingAccessRequest
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				serviceName:              "fixture-service-name",
				organizationSerialNumber: "00000000000000000001",
			},
			want: []*database.OutgoingAccessRequest{&database.OutgoingAccessRequest{
				ID: 2,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "fixture-organization-name",
				},
				ServiceName:          "fixture-service-name",
				ReferenceID:          1,
				State:                database.OutgoingAccessRequestApproved,
				CreatedAt:            fixtureCustomTime,
				UpdatedAt:            fixtureCustomTime,
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
			}, {
				ID: 6,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "fixture-organization-name-two",
				},
				ServiceName:          "fixture-service-name",
				ReferenceID:          1,
				State:                database.OutgoingAccessRequestReceived,
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
				PublicKeyFingerprint: "h+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=",
			}},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			got, err := configDb.ListLatestOutgoingAccessRequests(context.Background(), tt.args.organizationSerialNumber, tt.args.serviceName)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
