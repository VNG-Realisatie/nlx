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
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
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

	tests := []struct {
		name             string
		req              *api.ListAccessGrantsForServiceRequest
		setup            func(context.Context, serviceMocks)
		expectedResponse *api.ListAccessGrantsForServiceResponse
		expectedErr      error
	}{
		{
			"happy_flow",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetService(ctx, "test-service").
					Return(&database.Service{
						Name: "test-service",
					}, nil)

				mocks.db.
					EXPECT().
					ListAccessGrantsForService(ctx, "test-service").
					Return([]*database.AccessGrant{
						createDummyAccessGrant(time.Date(2020, time.July, 9, 14, 45, 5, 0, time.UTC)),
					}, nil)
			},
			&api.ListAccessGrantsForServiceResponse{
				AccessGrants: []*api.AccessGrant{
					{
						Id: 1,
						Organization: &api.AccessGrant_Organization{
							Name:         "test-organization",
							SerialNumber: "00000000000000000001",
						},
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
			func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetService(ctx, "test-service").
					Return(nil, database.ErrNotFound)
			},
			nil,
			status.Error(codes.NotFound, "service not found"),
		},
		{
			"list_grants_database_error",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetService(ctx, "test-service").Return(&database.Service{
					Name: "test-service",
				}, nil)

				mocks.db.
					EXPECT().
					ListAccessGrantsForService(ctx, "test-service").
					Return(nil, errors.New("arbitrary error"))
			},
			nil,
			status.Error(codes.Internal, "database error"),
		},
		{
			"get_service_database_error",
			&api.ListAccessGrantsForServiceRequest{
				ServiceName: "test-service",
			},
			func(ctx context.Context, mocks serviceMocks) {
				mocks.db.EXPECT().
					GetService(ctx, "test-service").
					Return(nil, errors.New("arbitrary error"))
			},
			nil,
			status.Error(codes.Internal, "database error"),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			service, _, mocks := newService(t)

			ctx := context.Background()
			tt.setup(ctx, mocks)

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
		setup            func(context.Context, serviceMocks)
		ctx              context.Context
		req              *api.RevokeAccessGrantRequest
		expectedResponse *api.AccessGrant
		expectedErr      error
	}{
		{
			name: "happy_flow",
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
						"Jane Doe",
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
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			req: &api.RevokeAccessGrantRequest{
				AccessGrantID: 42,
			},
			expectedResponse: &api.AccessGrant{
				Organization: &api.AccessGrant_Organization{},
				CreatedAt:    createTimestamp(createdAt),
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
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
