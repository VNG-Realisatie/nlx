// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test method
func TestListAccessGrantsForService(t *testing.T) {
	createTimestamp := func(ti time.Time) *timestamppb.Timestamp {
		return &timestamppb.Timestamp{
			Seconds: ti.Unix(),
			Nanos:   int32(ti.Nanosecond()),
		}
	}

	tests := map[string]struct {
		ctx              context.Context
		req              *api.ListAccessGrantsForServiceRequest
		setup            func(context.Context, serviceMocks)
		expectedResponse *api.ListAccessGrantsForServiceResponse
		expectedErr      error
	}{
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			req: &api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetService(gomock.Any(), "test-service").
					Return(&database.Service{
						Name: "test-service",
					}, nil)

				mocks.db.
					EXPECT().
					ListAccessGrantsForService(gomock.Any(), "test-service").
					Return([]*database.AccessGrant{
						createDummyAccessGrant(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					}, nil)
			},
			expectedResponse: &api.ListAccessGrantsForServiceResponse{
				AccessGrants: []*api.AccessGrant{
					{
						Id: 1,
						Organization: &external.Organization{
							Name:         "test-organization",
							SerialNumber: "00000000000000000001",
						},
						ServiceName:          "test-service",
						PublicKeyFingerprint: "test-finger-print",
						CreatedAt:            createTimestamp(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					},
				},
			},
			expectedErr: nil,
		},
		"service_not_found": {
			ctx: testCreateAdminUserContext(),
			req: &api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetService(gomock.Any(), "test-service").
					Return(nil, database.ErrNotFound)
			},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.NotFound, "service not found"),
		},
		"list_grants_database_error": {
			ctx: testCreateAdminUserContext(),
			req: &api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetService(gomock.Any(), "test-service").Return(&database.Service{
					Name: "test-service",
				}, nil)

				mocks.db.
					EXPECT().
					ListAccessGrantsForService(gomock.Any(), "test-service").
					Return(nil, errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.Internal, "database error"),
		},
		"get_service_database_error": {
			ctx: testCreateAdminUserContext(),
			req: &api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.EXPECT().
					GetService(gomock.Any(), "test-service").
					Return(nil, errors.New("arbitrary error"))
			},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.Internal, "database error"),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			ctx := context.Background()
			tt.setup(ctx, mocks)

			actual, err := service.ListAccessGrantsForService(tt.ctx, tt.req)

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

	tests := map[string]struct {
		setup            func(context.Context, serviceMocks)
		ctx              context.Context
		req              *api.RevokeAccessGrantRequest
		expectedResponse *api.RevokeAccessGrantResponse
		expectedErr      error
	}{
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.EXPECT().
					GetAccessGrant(gomock.Any(), uint(42)).Return(
					&database.AccessGrant{
						CreatedAt: createdAt,
						IncomingAccessRequest: &database.IncomingAccessRequest{
							Service: &database.Service{
								Name: "test-service",
							},
							Organization: database.IncomingAccessRequestOrganization{
								SerialNumber: "00000000000000000001",
								Name:         "test-organization",
							},
						}}, nil,
				)

				mocks.al.
					EXPECT().
					AccessGrantRevoke(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"test-organization",
						"test-service",
					)

				mocks.db.
					EXPECT().
					RevokeAccessGrant(ctx, uint(42), gomock.Any()).
					Return(&database.AccessGrant{
						CreatedAt: createdAt,
						IncomingAccessRequest: &database.IncomingAccessRequest{
							Service:      &database.Service{},
							Organization: database.IncomingAccessRequestOrganization{},
						},
					}, nil)
			},
			req: &api.RevokeAccessGrantRequest{
				AccessGrantId: 42,
			},
			expectedResponse: &api.RevokeAccessGrantResponse{
				AccessGrant: &api.AccessGrant{
					Organization: &external.Organization{},
					CreatedAt:    createTimestamp(createdAt),
				},
			},
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setup(tt.ctx, mocks)

			actual, err := service.RevokeAccessGrant(tt.ctx, tt.req)

			assert.Equal(t, tt.expectedResponse, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func createDummyAccessGrant(createdAt time.Time) *database.AccessGrant {
	return &database.AccessGrant{
		ID:                      1,
		IncomingAccessRequestID: 1,
		IncomingAccessRequest: &database.IncomingAccessRequest{
			ID:        1,
			ServiceID: 1,
			Organization: database.IncomingAccessRequestOrganization{
				SerialNumber: "00000000000000000001",
				Name:         "test-organization",
			},
			State:                database.IncomingAccessRequestReceived,
			PublicKeyFingerprint: "test-finger-print",
			Service: &database.Service{
				ID:   1,
				Name: "test-service",
			},
		},
		CreatedAt: createdAt,
	}
}
