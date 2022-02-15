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

func TestCreateOutgoingOrder(t *testing.T) {
	t.Parallel()

	setup(t)

	now, err := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
	require.NoError(t, err)

	type args struct {
		outgoingOrder *database.CreateOutgoingOrder
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		wantErr      error
	}{
		"when_an_order_with_the_same_reference_already_exists": {
			loadFixtures: true,
			args: args{
				outgoingOrder: &database.CreateOutgoingOrder{
					Reference:      "fixture-reference",
					Description:    "description",
					Delegatee:      "00000000000000000001",
					ValidFrom:      now,
					ValidUntil:     now,
					CreatedAt:      now,
					AccessProofIds: []uint64{1},
				},
			},
			wantErr: database.ErrDuplicateOutgoingOrder,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				outgoingOrder: &database.CreateOutgoingOrder{
					ID:             fixturesStartID,
					Reference:      "reference-one",
					Description:    "description",
					Delegatee:      "00000000000000000001",
					ValidFrom:      now,
					ValidUntil:     now,
					CreatedAt:      now,
					AccessProofIds: []uint64{1},
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

			err := configDb.CreateOutgoingOrder(context.Background(), tt.args.outgoingOrder)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertCreateOutgoingOrder(t, configDb, tt.args.outgoingOrder)
			}
		})
	}
}

func assertCreateOutgoingOrder(t *testing.T, repo database.ConfigDatabase, want *database.CreateOutgoingOrder) {
	got, err := repo.GetOutgoingOrderByReference(context.Background(), want.Reference)
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.Equal(t, want.PublicKeyPEM, got.PublicKeyPEM)
	assert.Equal(t, want.Reference, got.Reference)
	assert.Equal(t, want.CreatedAt, got.CreatedAt)
	assert.Equal(t, want.Delegatee, got.Delegatee)
	assert.Equal(t, want.Description, got.Description)
	assert.Equal(t, want.ValidFrom, got.ValidFrom)
	assert.Equal(t, want.ValidUntil, got.ValidUntil)

	// TODO: check if access proofs are equal
}
