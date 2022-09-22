// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

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

func TestListOutgoingOrdersByOrganization(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		organizationSerialNumber string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         []*database.OutgoingOrder
		wantErr      error
	}{
		"happy_flow_when_no_orders": {
			loadFixtures: false,
			args: args{
				organizationSerialNumber: "arbitrary",
			},
			want:    []*database.OutgoingOrder{},
			wantErr: nil,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				organizationSerialNumber: "00000000000000000003",
			},
			want: []*database.OutgoingOrder{
				{
					ID:           4,
					Reference:    "fixture-reference-four",
					Description:  "fixture-description",
					Delegatee:    "00000000000000000003",
					PublicKeyPEM: fixturePublicKeyPEM,
					RevokedAt:    sql.NullTime{},
					ValidFrom:    fixtureTime,
					ValidUntil:   fixtureTime,
					CreatedAt:    fixtureTime,
					OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
						{
							AccessProofID:   1,
							OutgoingOrderID: 4,
							AccessProof: &database.AccessProof{
								ID:                      1,
								AccessRequestOutgoingID: 1,
								CreatedAt:               fixtureTime,
								OutgoingAccessRequest: &database.OutgoingAccessRequest{
									ID: 1,
									Organization: database.Organization{
										Name:         "fixture-organization-name",
										SerialNumber: "00000000000000000001",
									},
									ServiceName:          "fixture-service-name",
									ReferenceID:          1,
									State:                database.OutgoingAccessRequestReceived,
									PublicKeyFingerprint: "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=",
									ErrorCode:            0,
									ErrorCause:           "",
									ErrorStackTrace:      nil,
									CreatedAt:            fixtureTime,
									UpdatedAt:            fixtureTime,
								},
							},
						},
					},
				},
				{
					ID:           3,
					Reference:    "fixture-reference-three",
					Description:  "fixture-description",
					Delegatee:    "00000000000000000003",
					PublicKeyPEM: fixturePublicKeyPEM,
					RevokedAt: sql.NullTime{
						Time:  fixtureTime,
						Valid: true,
					},
					ValidFrom:  fixtureTime,
					ValidUntil: fixtureTime,
					CreatedAt:  fixtureTime,
					OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
						{
							AccessProofID:   1,
							OutgoingOrderID: 3,
							AccessProof: &database.AccessProof{
								ID:                      1,
								AccessRequestOutgoingID: 1,
								CreatedAt:               fixtureTime,
								OutgoingAccessRequest: &database.OutgoingAccessRequest{
									ID: 1,
									Organization: database.Organization{
										Name:         "fixture-organization-name",
										SerialNumber: "00000000000000000001",
									},
									ServiceName:          "fixture-service-name",
									ReferenceID:          1,
									State:                database.OutgoingAccessRequestReceived,
									PublicKeyFingerprint: "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=",
									ErrorCode:            0,
									ErrorCause:           "",
									ErrorStackTrace:      nil,
									CreatedAt:            fixtureTime,
									UpdatedAt:            fixtureTime,
								},
							},
						},
					},
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

			got, err := configDb.ListOutgoingOrdersByOrganization(context.Background(), tt.args.organizationSerialNumber)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}
