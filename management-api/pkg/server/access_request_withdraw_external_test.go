// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
)

//nolint:funlen // this is a test
func TestWithdrawAccessRequestExternal(t *testing.T) {
	tests := map[string]struct {
		setup   func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context
		req     *external.WithdrawAccessRequestRequest
		want    *external.WithdrawAccessRequestResponse
		wantErr error
	}{
		"when_the_peer_context_is_missing": {
			setup: func(*testing.T, *mock_database.MockConfigDatabase, *tls.CertificateBundle) context.Context {
				return context.Background()
			},
			wantErr: status.Error(codes.Internal, "missing metadata from the management proxy"),
		},
		"when_get_incoming_access_request_from_database_fails": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(nil, fmt.Errorf("arbitrary error"))

				return ctx
			},
			req: &external.WithdrawAccessRequestRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.NewInternal("internal", nil),
		},
		"when_update_incoming_access_request_state_in_database_fails": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.IncomingAccessRequest{ID: 1, State: database.IncomingAccessRequestReceived}, nil)

				db.EXPECT().
					UpdateIncomingAccessRequestState(ctx, uint(1), database.IncomingAccessRequestWithdrawn).
					Return(fmt.Errorf("arbitrary error"))

				return ctx
			},
			req: &external.WithdrawAccessRequestRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.NewInternal("internal", nil),
		},
		"incoming_access_request_not_found": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(nil, database.ErrNotFound)

				return ctx
			},
			req: &external.WithdrawAccessRequestRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.New(codes.NotFound, external.ErrorReason_ERROR_REASON_ACCESS_REQUEST_NOT_FOUND, "access request could not be found", nil),
		},
		"incoming_access_request_state_is_not_received": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.IncomingAccessRequest{ID: 1, State: database.IncomingAccessRequestApproved}, nil)

				return ctx
			},
			req: &external.WithdrawAccessRequestRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			wantErr: grpcerrors.New(codes.FailedPrecondition, external.ErrorReason_ERROR_REASON_ACCESS_REQUEST_INVALID_STATE, "expected state: received, actual state: approved", nil),
		},
		"happy_flow": {
			setup: func(t *testing.T, db *mock_database.MockConfigDatabase, certBundle *tls.CertificateBundle) context.Context {
				ctx := setProxyMetadataWithCertBundle(t, context.Background(), certBundle)

				db.
					EXPECT().
					GetLatestIncomingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.IncomingAccessRequest{ID: 1, State: database.IncomingAccessRequestReceived}, nil)

				db.EXPECT().
					UpdateIncomingAccessRequestState(ctx, uint(1), database.IncomingAccessRequestWithdrawn).
					Return(nil)

				return ctx
			},
			req: &external.WithdrawAccessRequestRequest{
				ServiceName:          "service-name",
				PublicKeyFingerprint: "fingerprint",
			},
			want: &external.WithdrawAccessRequestResponse{},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, certBundle, mocks := newService(t, nil)
			ctx := tt.setup(t, mocks.db, certBundle)

			actual, err := service.WithdrawAccessRequest(ctx, tt.req)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, actual)
		})
	}
}
