// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestUpdateOutgoingAccessRequestState(t *testing.T) {
	t.Parallel()

	setup(t)

	type args struct {
		accessRequestID uint
		state           database.OutgoingAccessRequestState
	}

	tests := map[string]struct {
		loadFixtures bool
		args         args
		wantErr      error
	}{
		"when_access_request_does_not_exist": {
			loadFixtures: false,
			args: args{
				accessRequestID: 1,
			},
			wantErr: database.ErrNotFound,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accessRequestID: 1,
				state:           database.OutgoingAccessRequestApproved,
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

			err := configDb.UpdateOutgoingAccessRequestState(context.Background(), tt.args.accessRequestID, tt.args.state)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertOutgoingAccessRequest(t, configDb, tt.args.accessRequestID, tt.args.state)
			}
		})
	}
}

func assertOutgoingAccessRequest(t *testing.T, repo database.ConfigDatabase, accessRequestID uint, state database.OutgoingAccessRequestState) {
	accessRequest, err := repo.GetOutgoingAccessRequest(context.Background(), accessRequestID)
	require.NoError(t, err)
	require.NotNil(t, accessRequest)

	assert.Equal(t, state, accessRequest.State)
}
