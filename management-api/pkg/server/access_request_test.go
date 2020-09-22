// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func newService(t *testing.T) (s *server.ManagementService, ctrl *gomock.Controller, db *mock_database.MockConfigDatabase) {
	logger := zaptest.Logger(t)
	proc := process.NewProcess(logger)

	ctrl = gomock.NewController(t)

	db = mock_database.NewMockConfigDatabase(ctrl)
	s = server.NewManagementService(logger, proc, mock_directory.NewMockClient(ctrl), db)

	return
}

func TestCreateAccessRequest(t *testing.T) {
	service, ctrl, db := newService(t)
	ctrl.Finish()

	ctx := context.Background()

	createTimestamp := func(ti time.Time) *types.Timestamp {
		return &types.Timestamp{
			Seconds: ti.Unix(),
			Nanos:   int32(ti.Nanosecond()),
		}
	}

	tests := []struct {
		name        string
		req         *api.CreateAccessRequestRequest
		ar          *database.OutgoingAccessRequest
		returnReq   *database.OutgoingAccessRequest
		returnErr   error
		expectedReq *api.OutgoingAccessRequest
		expectedErr error
	}{
		{
			"without active access request",
			&api.CreateAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "test-organization",
					ServiceName:      "test-service",
				},
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					ID:               "12345abcde",
					OrganizationName: "test-organization",
					ServiceName:      "test-service",
					State:            database.AccessRequestCreated,
					CreatedAt:        time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
					UpdatedAt:        time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
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
		{
			"with active access request",
			&api.CreateAccessRequestRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			&database.OutgoingAccessRequest{
				AccessRequest: database.AccessRequest{
					OrganizationName: "test-organization",
					ServiceName:      "test-service",
				},
			},
			nil,
			database.ErrActiveAccessRequest,
			nil,
			status.New(codes.AlreadyExists, "there is already an active access request").Err(),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			db.EXPECT().CreateOutgoingAccessRequest(ctx, tt.ar).
				Return(tt.returnReq, tt.returnErr)
			actual, err := service.CreateAccessRequest(ctx, tt.req)

			assert.Equal(t, tt.expectedReq, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
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
			service, ctrl, db := newService(t)

			defer ctrl.Finish()

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
					GetLatestOutgoingAccessRequest(ctx, "organization-a", "service").
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
					GetLatestOutgoingAccessRequest(ctx, "organization-a", "service").
					Return(&database.OutgoingAccessRequest{
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
			service, ctrl, db := newService(t)
			defer ctrl.Finish()

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
