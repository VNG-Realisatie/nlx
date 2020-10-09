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
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

type statusLoopMocks struct {
	db         *mock_database.MockConfigDatabase
	directory  *mock_directory.MockClient
	management *mock_management.MockClient
	ctrl       *gomock.Controller
}

func newTestAccessRequestStatusLoop(t *testing.T) (statusLoopMocks, *accessRequestStatusLoop) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	requests := make(chan *database.OutgoingAccessRequest, 10)
	mocks := statusLoopMocks{
		ctrl:       ctrl,
		db:         mock_database.NewMockConfigDatabase(ctrl),
		directory:  mock_directory.NewMockClient(ctrl),
		management: mock_management.NewMockClient(ctrl),
	}
	statusLoop := &accessRequestStatusLoop{
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

	return mocks, statusLoop
}

func TestListCurrentAccessRequests(t *testing.T) {
	mocks, statusLoop := newTestAccessRequestStatusLoop(t)
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

	err := statusLoop.listCurrentAccessRequests(ctx)
	assert.NoError(t, err)

	actual := []*database.OutgoingAccessRequest{}

	for i := 0; i < 4; i++ {
		actual = append(actual, <-statusLoop.requests)
	}

	expected := append(requests[:1], requests[2:]...)

	assert.Equal(t, expected, actual)
}

//nolint:funlen // lot's of mocks
func TestHandleRequest(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		setupMock func(statusLoopMocks)
		request   *database.OutgoingAccessRequest
		wantErr   bool
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
			setupMock: func(mocks statusLoopMocks) {
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
			setupMock: func(mocks statusLoopMocks) {
				mocks.db.
					EXPECT().
					LockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(errors.New("error"))
			},
		},

		"returns_an_error_when_get_organization_inway_returns_an_error": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			wantErr: true,
			setupMock: func(mocks statusLoopMocks) {
				mocks.db.
					EXPECT().
					LockOutgoingAccessRequest(ctx, gomock.Any()).
					Return(nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
						OrganizationName: "organization-a",
					}).
					Return(nil, errors.New("error"))

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

		"returns_err_when_inway_address_is_invalid": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			wantErr: true,
			setupMock: func(mocks statusLoopMocks) {
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
						Address: "hostname",
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(ctx, gomock.Any(), database.AccessRequestFailed).
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(ctx, gomock.Any())
			},
		},

		"returns_an_error_when_requestaccess_returns_an_error": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			wantErr: true,
			setupMock: func(mocks statusLoopMocks) {
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
					Return(nil, errors.New("error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(ctx, gomock.Any(), database.AccessRequestFailed).
					Return(nil)

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

		"returns_an_error_when_update_access_request_returns_an_error": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			wantErr: true,
			setupMock: func(mocks statusLoopMocks) {
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
					Return(nil, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(ctx, gomock.Any(), database.AccessRequestReceived).
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

		"handling_a_valid_access_requests_succeeds": {
			request: &database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "id-1",
					OrganizationName: "organization-a",
					ServiceName:      "service",
					State:            database.AccessRequestCreated,
				},
			},
			setupMock: func(mocks statusLoopMocks) {
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
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mocks, statusLoop := newTestAccessRequestStatusLoop(t)

			tt.setupMock(mocks)

			err := statusLoop.handleRequest(ctx, tt.request)
			if tt.wantErr {
				assert.Error(t, err, "handleRequest should error")
			} else {
				assert.NoError(t, err, "handleRequest shouldn't error")
			}
		})
	}
}
