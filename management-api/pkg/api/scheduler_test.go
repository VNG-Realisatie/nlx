// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package api

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

type schedulerMocks struct {
	db         *mock_database.MockConfigDatabase
	directory  *mock_directory.MockClient
	management *mock_management.MockClient
	ctrl       *gomock.Controller
}

func newTestScheduler(t *testing.T) (schedulerMocks, *accessRequestScheduler) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	requests := make(chan *database.OutgoingAccessRequest, 10)
	mocks := schedulerMocks{
		ctrl:       ctrl,
		db:         mock_database.NewMockConfigDatabase(ctrl),
		directory:  mock_directory.NewMockClient(ctrl),
		management: mock_management.NewMockClient(ctrl),
	}
	scheduler := &accessRequestScheduler{
		clock:           &clock.FakeClock{},
		logger:          zap.NewNop(),
		directoryClient: mocks.directory,
		configDatabase:  mocks.db,
		orgCert:         nil,
		requests:        requests,
		createManagementClientFunc: func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error) {
			return mocks.management, nil
		},
	}

	return mocks, scheduler
}

func TestListCurrentAccessRequests(t *testing.T) {
	mocks, scheduler := newTestScheduler(t)
	ctx := context.Background()

	requests := []*database.OutgoingAccessRequest{
		{
			AccessRequest: database.AccessRequest{
				ID:               "id-1",
				OrganizationName: "organization-a",
				ServiceName:      "service",
				State:            database.AccessRequestCreated,
			},
		},
		{
			AccessRequest: database.AccessRequest{
				ID:               "id-2",
				OrganizationName: "organization-b",
				ServiceName:      "service",
				State:            database.AccessRequestFailed,
			},
		},
		{
			AccessRequest: database.AccessRequest{
				ID:               "id-3",
				OrganizationName: "organization-c",
				ServiceName:      "service",
				State:            database.AccessRequestCreated,
			},
		},
		{
			AccessRequest: database.AccessRequest{
				ID:               "id-4",
				OrganizationName: "organization-d",
				ServiceName:      "service",
				State:            database.AccessRequestCreated,
			},
		},
		{
			AccessRequest: database.AccessRequest{
				ID:               "id-5",
				OrganizationName: "organization-e",
				ServiceName:      "service",
				State:            database.AccessRequestCreated,
			},
		},
	}

	mocks.db.
		EXPECT().
		ListAllOutgoingAccessRequests(ctx).
		Return(requests, nil)

	err := scheduler.listCurrentAccessRequests(ctx)
	assert.NoError(t, err)

	actual := []*database.OutgoingAccessRequest{}

	for i := 0; i < 4; i++ {
		actual = append(actual, <-scheduler.requests)
	}

	expected := append(requests[:1], requests[2:]...)

	assert.Equal(t, expected, actual)
}

func TestGetOrganizationManagementClient(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		setupMock        func(*mock_directory.MockClient)
		organizationName string
		wantErr          bool
	}{
		"returns_an_error_when_get_organization_inway_returns_an_error": {
			organizationName: "organization-a",
			wantErr:          true,
			setupMock: func(directory *mock_directory.MockClient) {
				directory.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(nil, errors.New("error"))
			},
		},

		"returns_err_when_inway_address_is_invalid": {
			organizationName: "organization-a",
			wantErr:          true,
			setupMock: func(directory *mock_directory.MockClient) {
				directory.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "hostname",
					}, nil)
			},
		},

		"returns_the_client_when_everything_is_ok": {
			organizationName: "organization-a",
			wantErr:          false,
			setupMock: func(directory *mock_directory.MockClient) {
				directory.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "hostname:8000",
					}, nil)
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mocks, scheduler := newTestScheduler(t)

			tt.setupMock(mocks.directory)

			client, err := scheduler.getOrganizationManagementClient(ctx, tt.organizationName)
			if tt.wantErr {
				assert.Error(t, err, "getOrganizationManagementClient should error")
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err, "getOrganizationManagementClient shouldn't error")
				assert.NotNil(t, client)
			}
		})
	}
}

//nolint:funlen,dupl // unittest
func TestSchedule(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		setupMock func(schedulerMocks)
		request   *database.OutgoingAccessRequest
		wantErr   bool
		wantState database.AccessRequestState
	}{
		"handling_a_locked_request_returns_nil": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					LockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(database.ErrAccessRequestLocked)
			},
		},

		"returns_an_error_when_locking_fails": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			wantErr: true,
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					LockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(errors.New("error"))
			},
		},

		"returns_an_error_when_getting_organization_management_client_fails": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			wantErr: true,
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					LockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(nil, errors.New("random error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(ctx, gomock.Any(), database.AccessRequestFailed).
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(nil)
			},
		},

		"returns_an_error_when_update_access_request_state_returns_an_error": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestReceived,
				},
			},
			wantErr: true,
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					LockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "hostname:7200",
					}, nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(ctx, &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_APPROVED,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(ctx, gomock.Any(), database.AccessRequestApproved).
					Return(errors.New("error"))

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},

		"scheduling_a_created_access_request_succeeds": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					LockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "hostname:7200",
					}, nil)

				mocks.management.
					EXPECT().
					RequestAccess(ctx, &external.RequestAccessRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&types.Empty{}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(ctx, gomock.Any(), database.AccessRequestReceived).
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(ctx, gomock.Any())

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},

		"scheduling_a_pending_access_request_succeeds": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestReceived,
				},
			},
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					LockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(&inspectionapi.GetOrganizationInwayResponse{
						Address: "hostname:7200",
					}, nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(ctx, &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_APPROVED,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(ctx, gomock.Any(), database.AccessRequestApproved).
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(ctx, gomock.Any())

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mocks, scheduler := newTestScheduler(t)

			tt.setupMock(mocks)

			err := scheduler.schedule(ctx, tt.request)
			if tt.wantErr {
				assert.Error(t, err, "schedule should error")
			} else {
				assert.NoError(t, err, "schedule shouldn't error")
			}
		})
	}
}

func TestSchedulePendingRequests(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		setupMock    func(schedulerMocks)
		wantErr      bool
		wantRequests []*database.OutgoingAccessRequest
	}{
		"returns_an_error_when_list_outgoing_access_requests_errors": {
			wantErr: true,
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					ListAllOutgoingAccessRequests(ctx).
					Return(nil, errors.New("random error"))
			},
		},

		"schedules_received_requests": {
			wantErr: false,
			wantRequests: []*database.OutgoingAccessRequest{
				{
					AccessRequest: database.AccessRequest{
						State: database.AccessRequestReceived,
					},
				},
				{
					AccessRequest: database.AccessRequest{
						State: database.AccessRequestReceived,
					},
				},
				{
					AccessRequest: database.AccessRequest{
						State: database.AccessRequestReceived,
					},
				},
			},
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					ListAllOutgoingAccessRequests(ctx).
					Return([]*database.OutgoingAccessRequest{
						{
							AccessRequest: database.AccessRequest{
								State: database.AccessRequestReceived,
							},
						},
						{
							AccessRequest: database.AccessRequest{
								State: database.AccessRequestCreated,
							},
						},
						{
							AccessRequest: database.AccessRequest{
								State: database.AccessRequestReceived,
							},
						},
						{
							AccessRequest: database.AccessRequest{
								State: database.AccessRequestApproved,
							},
						},
						{
							AccessRequest: database.AccessRequest{
								State: database.AccessRequestReceived,
							},
						},
					}, nil)
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mocks, scheduler := newTestScheduler(t)

			tt.setupMock(mocks)

			err := scheduler.schedulePendingRequests(ctx)
			if tt.wantErr {
				assert.Error(t, err, "schedulePendingRequests should error")
			} else {
				assert.NoError(t, err, "schedulePendingRequests shouldn't error")

				requests := []*database.OutgoingAccessRequest{}

				for i := 0; i < 3; i++ {
					request := <-scheduler.requests
					requests = append(requests, request)
				}

				assert.Equal(t, tt.wantRequests, requests)
			}
		})
	}
}
