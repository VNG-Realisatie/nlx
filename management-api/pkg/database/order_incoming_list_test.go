// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/domain"
)

func TestListIncomingOrders(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	tests := map[string]struct {
		loadFixtures bool
		want         []*domain.IncomingOrder
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			want: []*domain.IncomingOrder{
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference",
					Description: "fixture-description",
					Delegator:   "fixture-delegator",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("fixture-service", "fixture-organization"),
					},
				}),
				newIncomingOrder(t, &domain.NewIncomingOrderArgs{
					Reference:   "fixture-reference-two",
					Description: "fixture-description",
					Delegator:   "fixture-delegator-two",
					RevokedAt:   nil,
					ValidFrom:   fixtureTime,
					ValidUntil:  fixtureTime,
					Services: []domain.IncomingOrderService{
						domain.NewIncomingOrderService("fixture-service-two", "fixture-organization-two"),
					},
				}),
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

			got, err := configDb.ListIncomingOrders(context.Background())
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}

func newIncomingOrder(t *testing.T, args *domain.NewIncomingOrderArgs) *domain.IncomingOrder {
	model, err := domain.NewIncomingOrder(args)
	require.NoError(t, err)

	return model
}
