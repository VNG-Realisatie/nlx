// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestUpdateOutgoingOrder(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	require.NoError(t, err)

	fixtureTime := getFixtureTime(t)

	type args struct {
		outgoingOrder *database.UpdateOutgoingOrder
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		updatedOrder *database.OutgoingOrder
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				outgoingOrder: &database.UpdateOutgoingOrder{
					ID:             1,
					Reference:      "fixture-reference",
					Description:    "description updated",
					ValidFrom:      now,
					ValidUntil:     now,
					PublicKeyPEM:   "PUBLIC_KEY_HERE",
					AccessProofIds: []uint64{1},
				},
			},
			updatedOrder: &database.OutgoingOrder{
				ID:           1,
				Delegatee:    "00000000000000000001",
				Reference:    "fixture-reference",
				Description:  "description updated",
				PublicKeyPEM: "PUBLIC_KEY_HERE",
				ValidFrom:    now,
				ValidUntil:   now,
				CreatedAt:    fixtureTime,
				OutgoingOrderAccessProofs: []*database.OutgoingOrderAccessProof{
					{
						OutgoingOrderID: 1,
						AccessProofID:   1,
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
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			err := configDb.UpdateOutgoingOrder(context.Background(), tt.args.outgoingOrder)
			require.NoError(t, err)

			assertOutgoingOrder(t, configDb, tt.updatedOrder)

		})
	}
}

func assertOutgoingOrder(t *testing.T, repo database.ConfigDatabase, want *database.OutgoingOrder) {
	got, err := repo.GetOutgoingOrderByReference(context.Background(), want.Reference)
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.EqualValues(t, want, got)
}
