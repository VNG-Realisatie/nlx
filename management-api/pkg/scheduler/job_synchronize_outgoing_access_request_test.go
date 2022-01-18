// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint funlen: these are tests
package scheduler_test

import (
	"context"
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
	request    *database.OutgoingAccessRequest
	wantErr    error
}

func TestSynchronizeOutgoingAccessRequest(t *testing.T) {
	testGroups := []map[string]testCase{
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
				err := job.Synchronize(context.Background(), tt.request)

				if tt.wantErr != nil {
					require.EqualError(t, err, tt.wantErr.Error())
				} else {
					require.Nil(t, err)
				}
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
