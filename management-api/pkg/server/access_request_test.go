// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/diagnostics"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"

	common_tls "go.nlx.io/nlx/common/tls"
)

func newService(t *testing.T) (s *server.ManagementService, db *mock_database.MockConfigDatabase) {
	logger := zaptest.Logger(t)
	proc := process.NewProcess(logger)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	db = mock_database.NewMockConfigDatabase(ctrl)
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")
	bundle, err := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	assert.NoError(t, err)

	s = server.NewManagementService(logger, proc, mock_directory.NewMockClient(ctrl), bundle, db)

	return
}

func createTimestamp(ti time.Time) *types.Timestamp {
	return &types.Timestamp{
		Seconds: ti.Unix(),
		Nanos:   int32(ti.Nanosecond()),
	}
}

func TestCreateAccessRequest(t *testing.T) {
	tests := map[string]struct {
		req         *api.CreateAccessRequestRequest
		ar          *database.OutgoingAccessRequest
		returnReq   *database.OutgoingAccessRequest
		returnErr   error
		expectedReq *api.OutgoingAccessRequest
		expectedErr error
	}{
		"without_an_active_access_request": {
			&api.CreateAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName:     "test-organization",
					ServiceName:          "test-service",
					PublicKeyFingerprint: "60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=",
				},
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:                   "12345abcde",
					OrganizationName:     "test-organization",
					ServiceName:          "test-service",
					PublicKeyFingerprint: "60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=",
					State:                database.AccessRequestCreated,
					CreatedAt:            time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
					UpdatedAt:            time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
				},
			},
			nil,
			&api.OutgoingAccessRequest{
				Id:               "12345abcde",
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
				UpdatedAt:        createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
			},
			nil,
		},

		"with_an_activeaccess_request": {
			&api.CreateAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName:     "test-organization",
					ServiceName:          "test-service",
					PublicKeyFingerprint: "60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4=",
				},
			},
			nil,
			database.ErrActiveAccessRequest,
			nil,
			status.New(codes.AlreadyExists, "there is already an active access request").Err(),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, db := newService(t)
			ctx := context.Background()

			db.EXPECT().CreateOutgoingAccessRequest(ctx, tt.ar).
				Return(tt.returnReq, tt.returnErr)
			actual, err := service.CreateAccessRequest(ctx, tt.req)

			assert.Equal(t, tt.expectedReq, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

//nolint:funlen // this is a test
func TestSendAccessRequest(t *testing.T) {
	tests := []struct {
		name             string
		request          *api.SendAccessRequestRequest
		accessRequest    *database.OutgoingAccessRequest
		accessRequestErr error
		updateMock       func(mock *gomock.Call)
		response         *api.OutgoingAccessRequest
		responseErr      error
	}{
		{
			"non_existing",
			&api.SendAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  "1",
			},
			nil,
			database.ErrNotFound,
			func(mock *gomock.Call) {
				mock.MaxTimes(0)
			},
			nil,
			status.New(codes.NotFound, "access request not found").Err(),
		},
		{
			"database_error",
			&api.SendAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  "1",
			},
			nil,
			errors.New("an error"),
			func(mock *gomock.Call) {
				mock.MaxTimes(0)
			},
			nil,
			status.New(codes.Internal, "database error").Err(),
		},
		{
			"organiation_mismatch",
			&api.SendAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  "1",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "1",
					OrganizationName: "other-organization",
					ServiceName:      "test-service",
				},
			},
			nil,
			func(mock *gomock.Call) {
				mock.MaxTimes(0)
			},
			nil,
			status.New(codes.NotFound, "organization not found").Err(),
		},
		{
			"service_mismatch",
			&api.SendAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  "1",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "1",
					OrganizationName: "test-organization",
					ServiceName:      "other-service",
				},
			},
			nil,
			func(mock *gomock.Call) {
				mock.MaxTimes(0)
			},
			nil,
			status.New(codes.NotFound, "service not found").Err(),
		},
		{
			"update_failed",
			&api.SendAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  "1",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "1",
					OrganizationName: "test-organization",
					ServiceName:      "test-service",
					State:            database.AccessRequestCreated,
				},
			},
			nil,
			func(mock *gomock.Call) {
				mock.Return(errors.New("an error"))
			},
			nil,
			status.New(codes.Internal, "database error").Err(),
		},
		{
			"created_state",
			&api.SendAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  "1",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "1",
					OrganizationName: "test-organization",
					ServiceName:      "test-service",
					State:            database.AccessRequestCreated,
				},
			},
			nil,
			func(mock *gomock.Call) {
				mock.Return(nil)
			},
			&api.OutgoingAccessRequest{
				Id:               "1",
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        createTimestamp(time.Time{}),
				UpdatedAt:        createTimestamp(time.Time{}),
			},
			nil,
		},
		{
			"failed_state",
			&api.SendAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				AccessRequestID:  "1",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "1",
					OrganizationName: "test-organization",
					ServiceName:      "test-service",
					State:            database.AccessRequestFailed,
				},
			},
			nil,
			func(mock *gomock.Call) {
				mock.Return(nil)
			},
			&api.OutgoingAccessRequest{
				Id:               "1",
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        createTimestamp(time.Time{}),
				UpdatedAt:        createTimestamp(time.Time{}),
			},
			nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			service, db := newService(t)
			ctx := context.Background()

			db.EXPECT().GetOutgoingAccessRequest(ctx, test.request.AccessRequestID).
				Return(test.accessRequest, test.accessRequestErr)

			updateMock := db.EXPECT().UpdateOutgoingAccessRequestState(ctx, test.accessRequest, database.AccessRequestCreated, "", nil).
				Do(func(_ context.Context, accessRequest *database.OutgoingAccessRequest, state database.AccessRequestState, _ string, errorDetails *diagnostics.ErrorDetails) error {
					accessRequest.State = state
					return nil
				})
			test.updateMock(updateMock)

			response, err := service.SendAccessRequest(ctx, test.request)

			assert.Equal(t, test.response, response)
			assert.Equal(t, test.responseErr, err)
		})
	}
}

//nolint:funlen // this is a test method
func TestApproveIncomingAccessRequest(t *testing.T) {
	tests := []struct {
		name             string
		request          *api.ApproveIncomingAccessRequestRequest
		service          *database.Service
		serviceErr       error
		accessRequest    *database.IncomingAccessRequest
		accessRequestErr error
		response         *types.Empty
		err              error
	}{
		{
			"unknown_service",
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			nil,
			database.ErrNotFound,
			nil,
			nil,
			nil,
			status.Error(codes.NotFound, "service not found"),
		},
		{
			"unknown_access_request",
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			&database.Service{
				Name: "test-service",
			},
			nil,
			nil,
			database.ErrNotFound,
			nil,
			status.Error(codes.NotFound, "access request not found"),
		},
		{
			"service_mismatch_access_request",
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			&database.Service{
				Name: "test-service",
			},
			nil,
			&database.IncomingAccessRequest{
				AccessRequest: database.AccessRequest{
					ServiceName: "other-service",
				},
			},
			nil,
			nil,
			status.Error(codes.InvalidArgument, "service name does not match the one from access request"),
		},
		{
			"access_request_already_approved",
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			&database.Service{
				Name: "test-service",
			},
			nil,
			&database.IncomingAccessRequest{
				AccessRequest: database.AccessRequest{
					ServiceName: "test-service",
					State:       database.AccessRequestApproved,
				},
			},
			nil,
			nil,
			status.Error(codes.AlreadyExists, "access request is already approved"),
		},
		{
			"happy_flow",
			&api.ApproveIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			&database.Service{
				Name: "test-service",
			},
			nil,
			&database.IncomingAccessRequest{
				AccessRequest: database.AccessRequest{
					ServiceName: "test-service",
				},
			},
			nil,
			&types.Empty{},
			nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			service, db := newService(t)
			ctx := context.Background()

			db.EXPECT().GetService(ctx, test.request.ServiceName).Return(test.service, test.serviceErr)

			if test.service != nil {
				db.EXPECT().GetIncomingAccessRequest(ctx, test.request.AccessRequestID).Return(test.accessRequest, test.accessRequestErr)
			}

			if test.response != nil {
				db.EXPECT().CreateAccessGrant(ctx, test.accessRequest)
			}

			actual, err := service.ApproveIncomingAccessRequest(ctx, test.request)

			assert.Equal(t, test.response, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

//nolint:funlen // this is a test
func TestRejectIncomingAccessRequest(t *testing.T) {
	tests := []struct {
		name             string
		request          *api.RejectIncomingAccessRequestRequest
		service          *database.Service
		serviceErr       error
		accessRequest    *database.IncomingAccessRequest
		accessRequestErr error
		expectUpdateCall bool
		updateErr        error
		response         *types.Empty
		err              error
	}{
		{
			"unknown_service",
			&api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			nil,
			database.ErrNotFound,
			nil,
			nil,
			false,
			nil,
			nil,
			status.Error(codes.NotFound, "service not found"),
		},
		{
			"unknown_access_request",
			&api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			&database.Service{
				Name: "test-service",
			},
			nil,
			nil,
			database.ErrNotFound,
			false,
			nil,
			nil,
			status.Error(codes.NotFound, "access request not found"),
		},
		{
			"service_mismatch_access_request",
			&api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			&database.Service{
				Name: "test-service",
			},
			nil,
			&database.IncomingAccessRequest{
				AccessRequest: database.AccessRequest{
					ServiceName: "other-service",
				},
			},
			nil,
			false,
			nil,
			nil,
			status.Error(codes.InvalidArgument, "service name does not match the one from access request"),
		},
		{
			"update_state_fails",
			&api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			&database.Service{
				Name: "test-service",
			},
			nil,
			&database.IncomingAccessRequest{
				AccessRequest: database.AccessRequest{
					ServiceName: "test-service",
				},
			},
			nil,
			true,
			errors.New("arbitrary error"),
			nil,
			status.Error(codes.Internal, "database error"),
		},
		{
			"happy_flow",
			&api.RejectIncomingAccessRequestRequest{
				ServiceName:     "test-service",
				AccessRequestID: "1",
			},
			&database.Service{
				Name: "test-service",
			},
			nil,
			&database.IncomingAccessRequest{
				AccessRequest: database.AccessRequest{
					ServiceName: "test-service",
				},
			},
			nil,
			true,
			nil,
			&types.Empty{},
			nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			service, db := newService(t)
			ctx := context.Background()

			db.EXPECT().GetService(ctx, test.request.ServiceName).Return(test.service, test.serviceErr)

			if test.service != nil {
				db.EXPECT().GetIncomingAccessRequest(ctx, test.request.AccessRequestID).Return(test.accessRequest, test.accessRequestErr)
			}

			if test.expectUpdateCall {
				db.EXPECT().UpdateIncomingAccessRequestState(ctx, test.accessRequest, database.AccessRequestRejected).Return(test.updateErr)
			}

			actual, err := service.RejectIncomingAccessRequest(ctx, test.request)

			assert.Equal(t, test.response, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestApproveIncomingAccessRequestModified(t *testing.T) {
	serverService, db := newService(t)
	ctx := context.Background()

	service := &database.Service{
		Name: "test-service",
	}

	accessRequest := &database.IncomingAccessRequest{
		AccessRequest: database.AccessRequest{
			ServiceName: "test-service",
		},
	}

	db.EXPECT().GetService(ctx, "test-service").Return(service, nil)
	db.EXPECT().GetIncomingAccessRequest(ctx, "1").Return(accessRequest, nil)
	db.EXPECT().CreateAccessGrant(ctx, accessRequest).Return(nil, database.ErrAccessRequestModified)

	request := &api.ApproveIncomingAccessRequestRequest{
		ServiceName:     "test-service",
		AccessRequestID: "1",
	}

	accessGrant, err := serverService.ApproveIncomingAccessRequest(ctx, request)

	assert.Nil(t, accessGrant)
	assert.Equal(t, status.Error(codes.Aborted, "access request modified"), err)
}

func setProxyMetadata(ctx context.Context) context.Context {
	md := metadata.Pairs(
		"nlx-organization", "organization-a",
		"nlx-public-key-fingerprint", "1655A0AB68576280",
	)

	return metadata.NewIncomingContext(ctx, md)
}

func TestExternalRequestAccess(t *testing.T) {
	tests := map[string]struct {
		wantCode codes.Code
		setup    func(*mock_database.MockConfigDatabase) context.Context
	}{
		"errors_when_peer_context_is_missing": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return context.Background()
			},
		},

		"returns_error_when_creating_access_request_errors": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					CreateIncomingAccessRequest(ctx, gomock.Any()).
					Return(nil, errors.New("error"))

				return ctx
			},
		},

		"returns_empty_when_creating_the_access_request_succeeds": {
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					CreateIncomingAccessRequest(ctx, gomock.Any()).
					Return(&database.IncomingAccessRequest{}, nil)

				return ctx
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, db := newService(t)
			ctx := tt.setup(db)

			_, err := service.RequestAccess(ctx, &external.RequestAccessRequest{
				ServiceName: "service",
			})

			if tt.wantCode > 0 {
				assert.Error(t, err)

				st, ok := status.FromError(err)

				assert.True(t, ok)
				assert.Equal(t, tt.wantCode, st.Code())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExternalGetAccessRequestState(t *testing.T) {
	tests := map[string]struct {
		want     api.AccessRequestState
		wantCode codes.Code
		setup    func(*mock_database.MockConfigDatabase) context.Context
	}{
		"errors_when_peer_context_is_missing": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				return context.Background()
			},
		},

		"returns_error_when_retrieving_state_errors": {
			wantCode: codes.Internal,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "organization-a", "service").
					Return(nil, errors.New("error"))

				return ctx
			},
		},

		"returns_the_right_state_when_the_request_exists": {
			want: api.AccessRequestState_RECEIVED,
			setup: func(db *mock_database.MockConfigDatabase) context.Context {
				ctx := setProxyMetadata(context.Background())

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "organization-a", "service").
					Return(&database.IncomingAccessRequest{
						AccessRequest: database.AccessRequest{
							State: database.AccessRequestReceived,
						},
					}, nil)

				return ctx
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, db := newService(t)
			ctx := tt.setup(db)

			response, err := service.GetAccessRequestState(ctx, &external.GetAccessRequestStateRequest{
				ServiceName: "service",
			})

			if tt.wantCode > 0 {
				assert.Error(t, err)

				st, ok := status.FromError(err)

				assert.True(t, ok)
				assert.Equal(t, tt.wantCode, st.Code())
			} else {
				assert.NoError(t, err)

				if assert.NotNil(t, response) {
					assert.Equal(t, tt.want, response.State)
				}
			}
		})
	}
}
