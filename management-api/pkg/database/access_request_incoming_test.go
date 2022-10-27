// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"go.nlx.io/nlx/management-api/pkg/database"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestDeleteIncomingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		id uint
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				id: 1,
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configDb, close := newConfigDatabase(t, t.Name(), tt.loadFixtures)
			defer close()

			err := configDb.DeleteIncomingAccessRequest(context.Background(), tt.args.id)
			require.NoError(t, err)

			assertIncomingAccessRequestWithIDDeleted(t, configDb, tt.args.id)
		})
	}
}

func assertIncomingAccessRequestWithIDDeleted(t *testing.T, repo database.ConfigDatabase, outgoingAccessRequestID uint) {
	_, err := repo.GetIncomingAccessRequest(context.Background(), outgoingAccessRequestID)
	require.Equal(t, err, database.ErrNotFound)
}
