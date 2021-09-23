// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//go:build integration
// +build integration

package database_test

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"go.nlx.io/nlx/common/diagnostics"
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
		referenceID     uint
		schedulerErr    *diagnostics.ErrorDetails
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
		"happy_flow_with_scheduler_error": {
			loadFixtures: true,
			args: args{
				accessRequestID: 1,
				state:           database.OutgoingAccessRequestApproved,
				referenceID:     2,
				schedulerErr: &diagnostics.ErrorDetails{
					Cause:      "cause",
					Code:       diagnostics.InternalError,
					StackTrace: []string{"a", "b", "c"},
				},
			},
			wantErr: nil,
		},
		"happy_flow": {
			loadFixtures: true,
			args: args{
				accessRequestID: 1,
				state:           database.OutgoingAccessRequestApproved,
				referenceID:     2,
				schedulerErr:    nil,
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

			err := configDb.UpdateOutgoingAccessRequestState(context.Background(), tt.args.accessRequestID, tt.args.state, tt.args.referenceID, tt.args.schedulerErr)
			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assertOutgoingAccessRequest(t, configDb, tt.args.accessRequestID, tt.args.state, tt.args.referenceID, tt.args.schedulerErr)
			}
		})
	}
}

func assertOutgoingAccessRequest(t *testing.T, repo database.ConfigDatabase, accessRequestID uint, state database.OutgoingAccessRequestState, referenceID uint, schedulerErr *diagnostics.ErrorDetails) {
	accessRequest, err := repo.GetOutgoingAccessRequest(context.Background(), accessRequestID)
	require.NoError(t, err)
	require.NotNil(t, accessRequest)

	assert.Equal(t, state, accessRequest.State)
	assert.Equal(t, referenceID, accessRequest.ReferenceID)
	assert.Nil(t, accessRequest.LockID)
	assert.Equal(t, sql.NullTime{}, accessRequest.LockExpiresAt)

	if schedulerErr != nil {
		assert.Equal(t, int(schedulerErr.Code), accessRequest.ErrorCode)
		assert.Equal(t, schedulerErr.Cause, accessRequest.ErrorCause)
		assert.Equal(t, pq.StringArray(schedulerErr.StackTrace), accessRequest.ErrorStackTrace)
	} else {
		assert.Equal(t, 0, accessRequest.ErrorCode)
		assert.Equal(t, "", accessRequest.ErrorCause)
		assert.Nil(t, accessRequest.ErrorStackTrace)
	}
}
