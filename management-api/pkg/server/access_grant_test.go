// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	context "context"
	"errors"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
)

//nolint:funlen // this is a test method
func TestListAccessGrantsForService(t *testing.T) {
	createTimestamp := func(ti time.Time) *types.Timestamp {
		return &types.Timestamp{
			Seconds: ti.Unix(),
			Nanos:   int32(ti.Nanosecond()),
		}
	}

	tests := []struct {
		name             string
		req              *api.ListAccessGrantsForServiceRequest
		db               func(ctx context.Context, db *mock_database.MockConfigDatabase)
		expectedResponse *api.ListAccessGrantsForServiceResponse
		expectedErr      error
	}{
		{
			"happy_flow",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			func(ctx context.Context, db *mock_database.MockConfigDatabase) {
				db.EXPECT().GetService(ctx, "test-service").Return(&database.Service{
					Name: "test-service",
				}, nil)

				db.EXPECT().ListAccessGrantsForService(ctx, "test-service").Return([]*database.AccessGrant{
					{
						ID:                      1,
						IncomingAccessRequestID: 1,
						IncomingAccessRequest: &database.IncomingAccessRequest{
							ID:                   1,
							ServiceID:            1,
							OrganizationName:     "test-organization",
							State:                database.IncomingAccessRequestReceived,
							PublicKeyFingerprint: "test-finger-print",
							Service: &database.Service{
								ID:   1,
								Name: "test-service",
							},
						},
						CreatedAt: time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC),
					},
				}, nil)
			},
			&api.ListAccessGrantsForServiceResponse{
				AccessGrants: []*api.AccessGrant{
					{
						Id:                   1,
						OrganizationName:     "test-organization",
						ServiceName:          "test-service",
						PublicKeyFingerprint: "test-finger-print",
						CreatedAt:            createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					},
				},
			},

			nil,
		},
		{
			"service_not_found",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			func(ctx context.Context, db *mock_database.MockConfigDatabase) {
				db.EXPECT().GetService(ctx, "test-service").Return(nil, database.ErrNotFound)
			},
			nil,
			status.Error(codes.NotFound, "service not found"),
		},
		{
			"convert_access_grant_error",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			func(ctx context.Context, db *mock_database.MockConfigDatabase) {
				db.EXPECT().GetService(ctx, "test-service").Return(&database.Service{
					Name: "test-service",
				}, nil)

				db.EXPECT().ListAccessGrantsForService(ctx, "test-service").Return([]*database.AccessGrant{
					{
						ID:                      1,
						IncomingAccessRequestID: 1,
						IncomingAccessRequest: &database.IncomingAccessRequest{
							ID:                   1,
							ServiceID:            1,
							OrganizationName:     "test-org",
							State:                database.IncomingAccessRequestReceived,
							PublicKeyFingerprint: "test-finger-print",
							Service: &database.Service{
								ID:   1,
								Name: "test-service",
							}},
						CreatedAt: time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC),
					},
				}, nil)
			},
			nil,
			status.Error(codes.Internal, "error converting access grant"),
		},
		{
			"list_grants_database_error",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			func(ctx context.Context, db *mock_database.MockConfigDatabase) {
				db.EXPECT().GetService(ctx, "test-service").Return(&database.Service{
					Name: "test-service",
				}, nil)

				db.EXPECT().ListAccessGrantsForService(ctx, "test-service").Return(nil, errors.New("arbitrary error"))
			},
			nil,
			status.Error(codes.Internal, "database error"),
		},
		{
			"get_service_database_error",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			func(ctx context.Context, db *mock_database.MockConfigDatabase) {
				db.EXPECT().GetService(ctx, "test-service").Return(nil, errors.New("arbitrary error"))
			},
			nil,
			status.Error(codes.Internal, "database error"),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			service, db, _ := newService(t)

			ctx := context.Background()
			tt.db(ctx, db)

			actual, err := service.ListAccessGrantsForService(ctx, tt.req)

			assert.Equal(t, tt.expectedResponse, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
