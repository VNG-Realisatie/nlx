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
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test method
func TestRevokeOutgoingOrder(t *testing.T) {
	tests := map[string]struct {
		setup            func(context.Context, serviceMocks)
		ctx              context.Context
		req              *api.RevokeOutgoingOrderRequest
		expectedResponse *api.RevokeOutgoingOrderResponse
		expectedErr      error
	}{
		"missing_required_permission": {
			setup: func(ctx context.Context, mocks serviceMocks) {},
			ctx:   testCreateUserWithoutPermissionsContext(),
			req: &api.RevokeOutgoingOrderRequest{
				Delegatee: "00000000000000000001",
				Reference: "test-reference",
			},
			expectedResponse: nil,
			expectedErr:      status.New(codes.PermissionDenied, "user needs the permission \"permissions.outgoing_order.revoke\" to execute this request").Err(),
		},
		"when_revoking_order_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingRevoke(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"test-reference",
					)

				mocks.db.
					EXPECT().
					RevokeOutgoingOrderByReference(ctx, "00000000000000000001", "test-reference", gomock.Any()).
					Return(errors.New("arbitrary error"))
			},
			ctx: testCreateAdminUserContext(),
			req: &api.RevokeOutgoingOrderRequest{
				Delegatee: "00000000000000000001",
				Reference: "test-reference",
			},
			expectedResponse: nil,
			expectedErr:      status.Errorf(codes.Internal, "failed to revoke outgoing order"),
		},
		"when_writing_audit_logs_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingRevoke(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"test-reference",
					).
					Return(errors.New("arbitrary error"))
			},
			ctx: testCreateAdminUserContext(),
			req: &api.RevokeOutgoingOrderRequest{
				Delegatee: "00000000000000000001",
				Reference: "test-reference",
			},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.Internal, "failed to write to auditlog"),
		},
		"when_order_not_found": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingRevoke(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"test-reference",
					)

				mocks.db.
					EXPECT().
					RevokeOutgoingOrderByReference(ctx, "00000000000000000001", "test-reference", gomock.Any()).
					Return(database.ErrNotFound)
			},
			ctx: testCreateAdminUserContext(),
			req: &api.RevokeOutgoingOrderRequest{
				Delegatee: "00000000000000000001",
				Reference: "test-reference",
			},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.NotFound, "outgoing order with delegatee 00000000000000000001 and reference test-reference does not exist"),
		},
		"when_delegatee_missing": {
			setup: func(ctx context.Context, mocks serviceMocks) {},
			ctx:   testCreateAdminUserContext(),
			req: &api.RevokeOutgoingOrderRequest{
				Delegatee: "",
				Reference: "test-reference",
			},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.InvalidArgument, "delegatee is required"),
		},
		"when_reference_missing": {
			setup: func(ctx context.Context, mocks serviceMocks) {},
			ctx:   testCreateAdminUserContext(),
			req: &api.RevokeOutgoingOrderRequest{
				Delegatee: "00000000000000000001",
				Reference: "",
			},
			expectedResponse: nil,
			expectedErr:      status.Error(codes.InvalidArgument, "reference is required"),
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrderOutgoingRevoke(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
						"00000000000000000001",
						"test-reference",
					)

				mocks.db.
					EXPECT().
					RevokeOutgoingOrderByReference(ctx, "00000000000000000001", "test-reference", gomock.Any()).
					Return(nil)
			},
			ctx: testCreateAdminUserContext(),
			req: &api.RevokeOutgoingOrderRequest{
				Delegatee: "00000000000000000001",
				Reference: "test-reference",
			},
			expectedResponse: &api.RevokeOutgoingOrderResponse{},
			expectedErr:      nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t, nil)
			tt.setup(tt.ctx, mocks)

			actual, err := service.RevokeOutgoingOrder(tt.ctx, tt.req)

			assert.Equal(t, tt.expectedResponse, actual)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
