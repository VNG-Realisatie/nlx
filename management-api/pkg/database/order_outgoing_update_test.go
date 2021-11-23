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
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestUpdateOutgoingOrder(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	require.NoError(t, err)

	type args struct {
		outgoingOrder *database.OutgoingOrder
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		updatedOrder *database.OutgoingOrder
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				outgoingOrder: &database.OutgoingOrder{
					ID:           1,
					Description:  "description updated",
					ValidFrom:    now,
					ValidUntil:   now,
					PublicKeyPEM: "PUBLIC_KEY_HERE",
					Services: []database.OutgoingOrderService{
						{
							OutgoingOrderID: 1,
							Service:         "fixture-service-two",
							Organization: database.OutgoingOrderServiceOrganization{
								Name:         "fixture-organization-two",
								SerialNumber: "10000000000000000002",
							},
						},
					},
				},
			},
			updatedOrder: &database.OutgoingOrder{
				ID:           1,
				Reference:    "fixture-reference",
				Description:  "description updated",
				PublicKeyPEM: "PUBLIC_KEY_HERE",
				Delegatee:    "00000000000000000001",
				ValidFrom:    now,
				ValidUntil:   now,
				CreatedAt:    getFixtureTime(t),
				Services: []database.OutgoingOrderService{
					{
						OutgoingOrderID: 1,
						Service:         "fixture-service-two",
						Organization: database.OutgoingOrderServiceOrganization{
							Name:         "fixture-organization-two",
							SerialNumber: "10000000000000000002",
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

			err := configDb.UpdateOutgoingOrder(context.Background(), tt.args.outgoingOrder)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertOutgoingOrder(t, configDb, tt.updatedOrder)
			}
		})
	}
}
