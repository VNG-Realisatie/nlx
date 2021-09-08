// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package scheduler_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/scheduler"
	"go.nlx.io/nlx/management-api/pkg/server"
)

type testCase struct {
	setupMocks    func(schedulerMocks)
	accessRequest *database.OutgoingAccessRequest
	wantErr       error
}

func getCreatedAccessRequests() map[string]testCase {
	return map[string]testCase{
		"returns_an_error_when_getting_organization_management_client_fails": {
			accessRequest: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-a",
				ServiceName:      "service",
				State:            database.OutgoingAccessRequestCreated,
				ReferenceID:      2,
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("", errors.New("arbitrary error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any()).
					Return(nil)
			},
		},
		"delete_outgoing_access_request_when_service_has_been_deleted": {
			accessRequest: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-a",
				ServiceName:      "service",
				State:            database.OutgoingAccessRequestCreated,
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(nil, fmt.Errorf("mock grpc wrapper: %w", server.ErrServiceDoesNotExist))

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), "organization-a", "service").
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
		"scheduling_a_created_access_request_succeeds": {
			accessRequest: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-a",
				ServiceName:      "service",
				State:            database.OutgoingAccessRequestCreated,
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&external.RequestAccessResponse{
						ReferenceId: 2,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestReceived, uint(2), nil).
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
	}
}

func getReceivedAccessRequests() map[string]testCase {
	return map[string]testCase{
		"returns_an_error_when_update_access_request_state_returns_an_error": {
			accessRequest: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-a",
				ServiceName:      "service",
				State:            database.OutgoingAccessRequestReceived,
			},
			wantErr: errors.New("arbitrary error"),
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(gomock.Any(), &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_APPROVED,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestApproved, uint(0), nil).
					Return(errors.New("arbitrary error"))

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
		"scheduling_a_pending_access_request_returns_service_not_found": {
			accessRequest: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-a",
				ServiceName:      "service",
				State:            database.OutgoingAccessRequestReceived,
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(gomock.Any(), &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(nil, fmt.Errorf("mock grpc wrapper: %w", server.ErrServiceDoesNotExist))

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), "organization-a", "service").
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
		"happy_flow": {
			accessRequest: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-a",
				ServiceName:      "service",
				State:            database.OutgoingAccessRequestReceived,
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(gomock.Any(), &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_APPROVED,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestApproved, uint(0), nil).
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
	}
}

// nolint:funlen,dupl // this is a test
func TestSynchronizeOutgoingAccessRequest(t *testing.T) {
	testGroups := []map[string]testCase{
		getCreatedAccessRequests(),
		getReceivedAccessRequests(),
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
				err := job.Run(context.Background(), tt.accessRequest)
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
