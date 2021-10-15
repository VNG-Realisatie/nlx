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

func TestCreateOutgoingOrder(t *testing.T) {
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
		wantErr      error
	}{
		"when_an_order_with_the_same_reference_already_exists": {
			loadFixtures: true,
			args: args{
				outgoingOrder: &database.OutgoingOrder{
					Reference:   "fixture-reference",
					Description: "description",
					Delegatee:   "00000000000000000001",
					RevokedAt:   sql.NullTime{},
					ValidFrom:   now,
					ValidUntil:  now,
					CreatedAt:   now,
					Services: []database.OutgoingOrderService{
						{
							Service: "service",
							Organization: database.OutgoingOrderServiceOrganization{
								Name:         "organization-one",
								SerialNumber: "10000000000000000001",
							},
						},
					},
				},
			},
			wantErr: database.ErrDuplicateOutgoingOrder,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				outgoingOrder: &database.OutgoingOrder{
					ID:          fixturesStartID,
					Reference:   "reference-one",
					Description: "description",
					Delegatee:   "00000000000000000001",
					RevokedAt:   sql.NullTime{},
					ValidFrom:   now,
					ValidUntil:  now,
					CreatedAt:   now,
					Services: []database.OutgoingOrderService{
						{
							OutgoingOrderID: fixturesStartID,
							Service:         "fixture-service",
							Organization: database.OutgoingOrderServiceOrganization{
								Name:         "fixture-organization",
								SerialNumber: "10000000000000000001",
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

			err := configDb.CreateOutgoingOrder(context.Background(), tt.args.outgoingOrder)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertOutgoingOrder(t, configDb, tt.args.outgoingOrder)
			}
		})
	}
}

func assertOutgoingOrder(t *testing.T, repo database.ConfigDatabase, want *database.OutgoingOrder) {
	got, err := repo.GetOutgoingOrderByReference(context.Background(), want.Reference)
	require.NoError(t, err)
	require.NotNil(t, got)

	assert.EqualValues(t, want, got)
}
