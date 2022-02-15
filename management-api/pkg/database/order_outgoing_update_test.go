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
		outgoingOrder *database.UpdateOutgoingOrder
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		updatedOrder *database.UpdateOutgoingOrder
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				outgoingOrder: &database.UpdateOutgoingOrder{
					ID:           1,
					Description:  "description updated",
					ValidFrom:    now,
					ValidUntil:   now,
					PublicKeyPEM: "PUBLIC_KEY_HERE",
				},
			},
			updatedOrder: &database.UpdateOutgoingOrder{
				ID:           1,
				Description:  "description updated",
				PublicKeyPEM: "PUBLIC_KEY_HERE",
				ValidFrom:    now,
				ValidUntil:   now,
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
