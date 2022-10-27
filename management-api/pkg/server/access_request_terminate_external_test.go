// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
)

//nolint:funlen // this is a test
func TestTerminateAccessRequest(t *testing.T) {
	tests := map[string]struct {
		setup   func(*testing.T, serviceMocks) context.Context
		req     *external.TerminateAccessRequest
		want    *external.TerminateAccessResponse
		wantErr error
	}{
		"when_incoming_access_request_is_not_found_in_database": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				mocks.db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(nil, database.ErrNotFound)

				return ctx
			},
			req: &external.TerminateAccessRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.New(codes.NotFound, external.ErrorReason_ERROR_REASON_ACCESS_REQUEST_NOT_FOUND, "access request could not be found", nil),
		},
		"when_get_incoming_access_request_from_database_fails": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				mocks.db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(nil, fmt.Errorf("arbitrary error"))

				return ctx
			},
			req: &external.TerminateAccessRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.NewInternal("internal", nil),
		},
		"when_incoming_access_request_is_not_approved": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				mocks.db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.IncomingAccessRequest{
						State: database.IncomingAccessRequestReceived,
					}, nil)

				return ctx
			},
			req: &external.TerminateAccessRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.New(codes.FailedPrecondition, external.ErrorReason_ERROR_REASON_ACCESS_REQUEST_INVALID_STATE, "expected state: approved, actual state: received", nil),
		},
		"when_get_access_grant_from_db_fails": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				mocks.db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.IncomingAccessRequest{
						ID:    42,
						State: database.IncomingAccessRequestApproved,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessGrantIDForIncomingAccessRequest(ctx, uint(42)).
					Return(uint(0), errors.New("arbitrary error"))

				return ctx
			},
			req: &external.TerminateAccessRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.NewInternal("internal", nil),
		},
		"when_terminating_access_grant_fails": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				mocks.db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.IncomingAccessRequest{
						ID:    42,
						State: database.IncomingAccessRequestApproved,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessGrantIDForIncomingAccessRequest(ctx, uint(42)).
					Return(uint(1), nil)

				mocks.db.
					EXPECT().
					TerminateAccessGrant(ctx, uint(1), mocks.cl.Now()).
					Return(errors.New("arbitrary error"))

				return ctx
			},
			req: &external.TerminateAccessRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.NewInternal("internal", nil),
		},
		"when_access_grant_is_already_terminated": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				mocks.db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.IncomingAccessRequest{
						ID:    42,
						State: database.IncomingAccessRequestApproved,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessGrantIDForIncomingAccessRequest(ctx, uint(42)).
					Return(uint(1), nil)

				mocks.db.
					EXPECT().
					TerminateAccessGrant(ctx, uint(1), mocks.cl.Now()).
					Return(database.ErrAccessGrantAlreadyTerminated)

				return ctx
			},
			req: &external.TerminateAccessRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.New(codes.FailedPrecondition, external.ErrorReason_ERROR_REASON_ACCESS_GRANT_ALREADY_TERMINATED, "access grant already terminated", nil),
		},
		"happy_flow": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := setProxyMetadata(t, context.Background())

				mocks.db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.IncomingAccessRequest{
						ID:    42,
						State: database.IncomingAccessRequestApproved,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessGrantIDForIncomingAccessRequest(ctx, uint(42)).
					Return(uint(1), nil)

				mocks.db.
					EXPECT().
					TerminateAccessGrant(ctx, uint(1), mocks.cl.Now()).
					Return(nil)

				return ctx
			},
			req: &external.TerminateAccessRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: nil,
			want: &external.TerminateAccessResponse{
				TerminatedAt: timestamppb.New(fixtureTime),
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			ctx := tt.setup(t, mocks)

			actual, err := service.TerminateAccess(ctx, tt.req)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, actual)
		})
	}
}
