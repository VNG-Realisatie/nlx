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

func TestTakePendingOutgoingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getCustomFixtureTime(t, "2021-01-03T01:02:03Z")

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		id uint
	}

	tests := map[string]struct {
		loadFixtures bool
		want         *database.OutgoingAccessRequest
		wantErr      error
		expectedIDs  []uint64
	}{
		"happy_flow_no_pending_requests": {
			loadFixtures: false,
			want:         nil,
			wantErr:      nil,
		},
		"happy_flow": {
			loadFixtures: true,
			want: &database.OutgoingAccessRequest{
				ID: 5,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "fixture-organization-name",
				},
				ServiceName:          "fixture-service-name-b",
				ReferenceID:          1,
				State:                database.OutgoingAccessRequestCreated,
				CreatedAt:            fixtureTime,
				UpdatedAt:            fixtureTime,
				PublicKeyPEM:         fixturePublicKeyPEM,
				PublicKeyFingerprint: fixtureCertBundle.PublicKeyFingerprint(),
			},
			wantErr:     nil,
			expectedIDs: []uint64{1, 2, 3, 4, 5},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			gots, err := configDb.TakePendingOutgoingAccessRequests(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil && tt.want != nil {
				assert.Len(t, gots, len(tt.expectedIDs))

				foundIDs := []uint{}
				for _, wantedID := range tt.expectedIDs {
					for _, got := range gots {
						if got.ID == uint(wantedID) {
							foundIDs = append(foundIDs, got.ID)
						}
					}
				}

				assert.Len(t, tt.expectedIDs, len(foundIDs))
			}
		})
	}
}
