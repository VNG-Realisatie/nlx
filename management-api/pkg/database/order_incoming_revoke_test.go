// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"go.nlx.io/nlx/management-api/domain"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestRevokeIncomingOrder(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)
	fixtureTimeNew := getCustomFixtureTime(t, "2021-01-04T01:02:03Z")

	type args struct {
		delegator string
		reference string
		revokedAt time.Time
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *domain.IncomingOrder
		wantErr      error
	}{
		"when_order_does_not_exist": {
			loadFixtures: false,
			args: args{
				delegator: "arbitrary",
				reference: "arbitrary",
				revokedAt: fixtureTimeNew,
			},
			wantErr: database.ErrNotFound,
		},
		"when_order_is_already_revoked": {
			loadFixtures: false,
			args: args{
				delegator: "fixture-delegator",
				reference: "fixture-reference",
				revokedAt: fixtureTimeNew,
			},
			want: newIncomingOrder(t, &domain.NewIncomingOrderArgs{
				Reference:   "fixture-reference",
				Description: "fixture-description",
				Delegator:   "fixture-delegator",
				RevokedAt:   &fixtureTime,
				ValidFrom:   fixtureTime,
				ValidUntil:  fixtureTime,
			}),
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				delegator: "fixture-delegator",
				reference: "fixture-reference",
				revokedAt: fixtureTimeNew,
			},
			want: newIncomingOrder(t, &domain.NewIncomingOrderArgs{
				Reference:   "fixture-reference",
				Description: "fixture-description",
				Delegator:   "fixture-delegator",
				RevokedAt:   &fixtureTimeNew,
				ValidFrom:   fixtureTime,
				ValidUntil:  fixtureTime,
			}),
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			err := configDb.RevokeIncomingOrderByReference(context.Background(), tt.args.delegator, tt.args.reference, tt.args.revokedAt)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertIncomingOrder(t, configDb, tt.args.delegator, tt.args.reference, tt.want)
			}
		})
	}
}

func assertIncomingOrder(t *testing.T, repo database.ConfigDatabase, delegator string, reference string, want *domain.IncomingOrder) {
	incomingOrders, err := repo.ListIncomingOrders(context.Background())
	require.NoError(t, err)
	require.NotNil(t, incomingOrders)

	var incomingOrder *domain.IncomingOrder

	for _, model := range incomingOrders {
		if model.Delegator() != delegator || model.Reference() != reference {
			continue
		}

		incomingOrder = model
	}

	assert.NotNil(t, incomingOrder)
	assert.Equal(t, want.RevokedAt(), incomingOrder.RevokedAt())
}
