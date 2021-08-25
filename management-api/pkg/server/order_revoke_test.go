// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test method
func TestRevokeOutgoingOrder(t *testing.T) {
	tests := map[string]struct {
		setup            func(context.Context, serviceMocks)
		ctx              context.Context
		req              *api.RevokeOutgoingOrderRequest
		expectedResponse *emptypb.Empty
		expectedErr      error
	}{
		"when_revoking_order_fails": {
			func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingRevoke(
						gomock.Any(),
						"Jane Doe",
						"nlxctl",
						"test-delegatee",
						"test-reference",
					)

				mocks.db.
					EXPECT().
					RevokeOutgoingOrderByReference(ctx, "test-delegatee", "test-reference", gomock.Any()).
					Return(errors.New("arbitrary error"))
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RevokeOutgoingOrderRequest{
				Delegatee: "test-delegatee",
				Reference: "test-reference",
			},
			nil,
			status.Errorf(codes.Internal, "failed to revoke outgoing order"),
		},
		"when_writing_audit_logs_fails": {
			func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingRevoke(
						gomock.Any(),
						"Jane Doe",
						"nlxctl",
						"test-delegatee",
						"test-reference",
					).
					Return(errors.New("arbitrary error"))
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RevokeOutgoingOrderRequest{
				Delegatee: "test-delegatee",
				Reference: "test-reference",
			},
			nil,
			status.Error(codes.Internal, "failed to write to auditlog"),
		},
		"when_order_not_found": {
			func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingRevoke(
						gomock.Any(),
						"Jane Doe",
						"nlxctl",
						"test-delegatee",
						"test-reference",
					)

				mocks.db.
					EXPECT().
					RevokeOutgoingOrderByReference(ctx, "test-delegatee", "test-reference", gomock.Any()).
					Return(database.ErrNotFound)
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RevokeOutgoingOrderRequest{
				Delegatee: "test-delegatee",
				Reference: "test-reference",
			},
			nil,
			status.Error(codes.NotFound, "outgoing order with delegatee test-delegatee and reference test-reference does not exist"),
		},
		"when_delegatee_missing": {
			func(ctx context.Context, mocks serviceMocks) {},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RevokeOutgoingOrderRequest{
				Delegatee: "",
				Reference: "test-reference",
			},
			nil,
			status.Error(codes.InvalidArgument, "delegatee is required"),
		},
		"when_reference_missing": {
			func(ctx context.Context, mocks serviceMocks) {},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RevokeOutgoingOrderRequest{
				Delegatee: "test-delegatee",
				Reference: "",
			},
			nil,
			status.Error(codes.InvalidArgument, "reference is required"),
		},
		"happy_flow": {
			func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingRevoke(
						gomock.Any(),
						"Jane Doe",
						"nlxctl",
						"test-delegatee",
						"test-reference",
					)

				mocks.db.
					EXPECT().
					RevokeOutgoingOrderByReference(ctx, "test-delegatee", "test-reference", gomock.Any()).
					Return(nil)
			},
			metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			&api.RevokeOutgoingOrderRequest{
				Delegatee: "test-delegatee",
				Reference: "test-reference",
			},
			&emptypb.Empty{},
			nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setup(tt.ctx, mocks)

			actual, err := service.RevokeOutgoingOrder(tt.ctx, tt.req)

			assert.Equal(t, tt.expectedResponse, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
