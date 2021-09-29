// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestGetOutgoingOrderByReference(t *testing.T) {
	t.Parallel()

	setup(t)

	fixtureTime := getFixtureTime(t)

	fixtureCertBundle, err := newFixtureCertificateBundle()
	require.NoError(t, err)

	fixturePublicKeyPEM, err := fixtureCertBundle.PublicKeyPEM()
	require.NoError(t, err)

	type args struct {
		reference string
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		want         *database.OutgoingOrder
		wantErr      error
	}{
		"when_not_found": {
			loadFixtures: false,
			args: args{
				reference: "arbitrary",
			},
			want:    nil,
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				reference: "fixture-reference",
			},
			want: &database.OutgoingOrder{
				ID:           1,
				Reference:    "fixture-reference",
				Description:  "fixture-description",
				Delegatee:    "fixture-delegatee",
				PublicKeyPEM: fixturePublicKeyPEM,
				RevokedAt:    sql.NullTime{},
				ValidFrom:    fixtureTime,
				ValidUntil:   fixtureTime,
				CreatedAt:    fixtureTime,
				Services: []database.OutgoingOrderService{
					{
						OutgoingOrderID: 1,
						Service:         "fixture-service",
						Organization:    "fixture-organization",
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

			got, err := configDb.GetOutgoingOrderByReference(context.Background(), tt.args.reference)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				require.EqualValues(t, tt.want, got)
			}
		})
	}
}