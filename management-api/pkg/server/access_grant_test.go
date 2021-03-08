// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
)

//nolint:funlen // this is a test method
func TestListAccessGrantsForService(t *testing.T) {
	createTimestamp := func(ti time.Time) *timestamppb.Timestamp {
		return &timestamppb.Timestamp{
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
							},
						},
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

//nolint:funlen // this is a test method
func TestRevokeAccessGrant(t *testing.T) {
	createTimestamp := func(ti time.Time) *timestamppb.Timestamp {
		return &timestamppb.Timestamp{
			Seconds: ti.Unix(),
			Nanos:   int32(ti.Nanosecond()),
		}
	}

	createdAt := time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)

	tests := []struct {
		name             string
		auditLog         func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger
		ctx              context.Context
		req              *api.RevokeAccessGrantRequest
		db               func(ctx context.Context, db *mock_database.MockConfigDatabase)
		expectedResponse *api.AccessGrant
		expectedErr      error
	}{
		{
			"happy_flow",
			func(auditLogger mock_auditlog.MockLogger) mock_auditlog.MockLogger {
				auditLogger.EXPECT().AccessGrantRevoke(gomock.Any(), "Jane Doe", "nlxctl", "test-organization", "test-service")
				return auditLogger
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RevokeAccessGrantRequest{
				AccessGrantID:    42,
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			func(ctx context.Context, db *mock_database.MockConfigDatabase) {
				db.EXPECT().RevokeAccessGrant(ctx, uint(42), gomock.Any()).Return(&database.AccessGrant{
					CreatedAt: createdAt,
					IncomingAccessRequest: &database.IncomingAccessRequest{
						Service: &database.Service{},
					},
				}, nil)
			},
			&api.AccessGrant{
				CreatedAt: createTimestamp(createdAt),
			},

			nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			service, db, auditLogger := newService(t)

			tt.auditLog(*auditLogger)

			tt.db(tt.ctx, db)

			actual, err := service.RevokeAccessGrant(tt.ctx, tt.req)

			assert.Equal(t, tt.expectedResponse, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
