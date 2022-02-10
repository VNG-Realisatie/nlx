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
)

func TestTakePendingOutgoingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	tests := map[string]struct {
		loadFixtures bool
		wantErr      error
		wantIDs      []uint64
	}{
		"happy_flow_no_pending_requests": {
			loadFixtures: false,
			wantErr:      nil,
		},
		"happy_flow": {
			loadFixtures: true,
			wantErr:      nil,
			wantIDs:      []uint64{1, 2, 3, 5},
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

			if tt.wantErr == nil {
				assert.Len(t, gots, len(tt.wantIDs))

				foundIDs := []uint{}
				for _, wantedID := range tt.wantIDs {
					for _, got := range gots {
						if got.ID == uint(wantedID) {
							foundIDs = append(foundIDs, got.ID)
						}
					}
				}

				assert.Len(t, foundIDs, len(tt.wantIDs))
			}
		})
	}
}
