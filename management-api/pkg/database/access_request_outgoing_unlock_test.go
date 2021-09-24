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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestUnlockOutgoingAccessRequest(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		outgoingAccessRequest *database.OutgoingAccessRequest
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		wantErr      error
	}{
		"happy_flow": {
			loadFixtures: true,
			args: args{
				outgoingAccessRequest: &database.OutgoingAccessRequest{
					ID: 1,
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

			err := configDb.UnlockOutgoingAccessRequest(context.Background(), tt.args.outgoingAccessRequest)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertOutgoingAccessRequestUnlocked(t, configDb, tt.args.outgoingAccessRequest.ID)
			}
		})
	}
}

func assertOutgoingAccessRequestUnlocked(t *testing.T, repo database.ConfigDatabase, accessRequestID uint) {
	accessRequest, err := repo.GetOutgoingAccessRequest(context.Background(), accessRequestID)
	require.NoError(t, err)
	require.NotNil(t, accessRequest)

	assert.Equal(t, sql.NullTime{}, accessRequest.LockExpiresAt)
}
