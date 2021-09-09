// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint funlen: these are tests
package scheduler_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/scheduler"
)

type testCase struct {
	setupMocks func(schedulerMocks)
	wantErr    error
}

func getGenericTests() map[string]testCase {
	return map[string]testCase{
		"when_taking_a_pending_access_request_errors": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			wantErr: errors.New("arbitrary error"),
		},
		"when_there_is_no_pending_access_request_available": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(nil, nil)
			},
		},
		"when_the_status_of_the_access_request_is_unknown": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(&database.OutgoingAccessRequest{
						State: "unknown state",
					}, nil)
			},
			wantErr: errors.New("invalid state 'unknown state' for pending access request"),
		},
	}
}

func TestSynchronizeOutgoingAccessRequest(t *testing.T) {
	testGroups := []map[string]testCase{
		getGenericTests(),
		getCreatedAccessRequests(),
		getReceivedAccessRequests(),
		getApprovedAccessRequests(),
	}

	for _, tests := range testGroups {
		for name, tt := range tests {
			tt := tt

			t.Run(name, func(t *testing.T) {
				mocks := newMocks(t)

				tt.setupMocks(mocks)

				job := scheduler.NewSynchronizeOutgoingAccessRequestJob(
					context.Background(),
					mocks.directory,
					mocks.db,
					nil,
					func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error) {
						return mocks.management, nil
					},
				)
				err := job.Run(context.Background())
				require.Equal(t, tt.wantErr, err)
			})
		}
	}
}

type schedulerMocks struct {
	db         *mock_database.MockConfigDatabase
	directory  *mock_directory.MockClient
	management *mock_management.MockClient
	ctrl       *gomock.Controller
}

func newMocks(t *testing.T) schedulerMocks {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := schedulerMocks{
		ctrl:       ctrl,
		db:         mock_database.NewMockConfigDatabase(ctrl),
		directory:  mock_directory.NewMockClient(ctrl),
		management: mock_management.NewMockClient(ctrl),
	}

	return mocks
}
