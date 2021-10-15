// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestRevokeOutgoingOrderByReference(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)
	fixtureTimeNew := getCustomFixtureTime(t, "2021-01-04T01:02:03Z")

	type args struct {
		delegatee string
		reference string
		revokedAt time.Time
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.OutgoingOrder
		wantErr      error
	}{
		"when_order_does_not_exist": {
			loadFixtures: false,
			args: args{
				delegatee: "arbitrary",
				reference: "arbitrary",
				revokedAt: fixtureTimeNew,
			},
			wantErr: database.ErrNotFound,
		},
		"when_order_is_already_revoked": {
			loadFixtures: false,
			args: args{
				delegatee: "00000000000000000001",
				reference: "fixture-reference",
				revokedAt: fixtureTimeNew,
			},
			want: &database.OutgoingOrder{
				Reference:   "fixture-reference",
				Description: "fixture-description",
				Delegatee:   "00000000000000000001",
				RevokedAt: sql.NullTime{
					Valid: true,
					Time:  fixtureTimeNew,
				},
				ValidFrom:  fixtureTime,
				ValidUntil: fixtureTime,
			},
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				delegatee: "00000000000000000001",
				reference: "fixture-reference",
				revokedAt: fixtureTimeNew,
			},
			want: &database.OutgoingOrder{
				Reference:   "fixture-reference",
				Description: "fixture-description",
				Delegatee:   "00000000000000000001",
				RevokedAt: sql.NullTime{
					Valid: true,
					Time:  fixtureTimeNew,
				},
				ValidFrom:  fixtureTime,
				ValidUntil: fixtureTime,
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

			err := configDb.RevokeOutgoingOrderByReference(context.Background(), tt.args.delegatee, tt.args.reference, tt.args.revokedAt)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertOutgoingOrderRevoked(t, configDb, tt.args.delegatee, tt.args.reference, tt.want)
			}
		})
	}
}

func assertOutgoingOrderRevoked(t *testing.T, repo database.ConfigDatabase, delegatee string, reference string, want *database.OutgoingOrder) {
	outgoingOrders, err := repo.ListOutgoingOrders(context.Background())
	require.NoError(t, err)
	require.NotNil(t, outgoingOrders)

	var outgoingOrder *database.OutgoingOrder

	for _, o := range outgoingOrders {
		if o.Delegatee != delegatee || o.Reference != reference {
			continue
		}

		outgoingOrder = o
	}

	assert.NotNil(t, outgoingOrder)
	assert.Equal(t, want.RevokedAt, outgoingOrder.RevokedAt)
}
