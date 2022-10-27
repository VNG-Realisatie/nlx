// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//nolint:dupl // looks the same as CancelOutgoingAccessRequest but writes a different audit-log
package server_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test
func TestTerminateAccessProof(t *testing.T) {
	tests := map[string]struct {
		setup   func(*testing.T, serviceMocks) context.Context
		req     *api.TerminateAccessProofRequest
		want    *api.TerminateAccessProofResponse
		wantErr error
	}{
		"when_missing_required_permission": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := testCreateUserWithoutPermissionsContext()

				return ctx
			},
			req: &api.TerminateAccessProofRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "service-name",
				PublicKeyFingerprint:     "fingerprint",
			},
			wantErr: status.Error(codes.PermissionDenied, "user needs the permission \"permissions.access.terminate\" to execute this request"),
		},
		"when_get_outgoing_access_request_from_database_fails": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := testCreateAdminUserContext()

				mocks.db.
					EXPECT().
					GetLatestOutgoingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(nil, fmt.Errorf("arbitrary error"))

				return ctx
			},
			req: &api.TerminateAccessProofRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "service-name",
				PublicKeyFingerprint:     "fingerprint",
			},
			wantErr: status.Error(codes.Internal, "internal"),
		},
		"when_get_organization_inway_proxy_address_fails": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := testCreateAdminUserContext()

				mocks.db.
					EXPECT().
					GetLatestOutgoingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.OutgoingAccessRequest{ID: 1}, nil)

				mocks.al.EXPECT().
					AccessTerminate(ctx, "admin@example.com", "nlxctl", "00000000000000000001", "service-name", "fingerprint").
					Return(int64(1), nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "00000000000000000001").
					Return("", fmt.Errorf("arbitrary error"))

				return ctx
			},
			req: &api.TerminateAccessProofRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "service-name",
				PublicKeyFingerprint:     "fingerprint",
			},
			wantErr: status.Error(codes.Internal, "internal"),
		},
		"when_writing_audit_log_fails": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := testCreateAdminUserContext()

				mocks.db.
					EXPECT().
					GetLatestOutgoingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.OutgoingAccessRequest{ID: 1}, nil)

				mocks.al.EXPECT().
					AccessTerminate(ctx, "admin@example.com", "nlxctl", "00000000000000000001", "service-name", "fingerprint").
					Return(int64(0), fmt.Errorf("arbitrary error"))

				return ctx
			},
			req: &api.TerminateAccessProofRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "service-name",
				PublicKeyFingerprint:     "fingerprint",
			},
			wantErr: status.Error(codes.Internal, "could not create audit log"),
		},
		"when_external_terminate_access_grant_fails": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := testCreateAdminUserContext()

				mocks.db.
					EXPECT().
					GetLatestOutgoingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.OutgoingAccessRequest{ID: 1}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "00000000000000000001").
					Return("proxy.address", nil)

				mocks.al.EXPECT().
					AccessTerminate(ctx, "admin@example.com", "nlxctl", "00000000000000000001", "service-name", "fingerprint").
					Return(int64(1), nil)

				mocks.mc.EXPECT().
					TerminateAccess(ctx, &external.TerminateAccessRequest{
						ServiceName:          "service-name",
						PublicKeyFingerprint: "fingerprint",
					}).
					Return(nil, fmt.Errorf("arbitrary error"))

				return ctx
			},
			req: &api.TerminateAccessProofRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "service-name",
				PublicKeyFingerprint:     "fingerprint",
			},
			wantErr: status.Error(codes.Internal, "internal"),
		},
		"when_set_terminated_at_of_access_proof_in_database_fails": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := testCreateAdminUserContext()
				terminatedAt := timestamppb.New(time.Now())

				mocks.db.
					EXPECT().
					GetLatestOutgoingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.OutgoingAccessRequest{ID: 1}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "00000000000000000001").
					Return("proxy.address", nil)

				mocks.al.EXPECT().
					AccessTerminate(ctx, "admin@example.com", "nlxctl", "00000000000000000001", "service-name", "fingerprint").
					Return(int64(1), nil)

				mocks.mc.EXPECT().
					TerminateAccess(ctx, &external.TerminateAccessRequest{
						ServiceName:          "service-name",
						PublicKeyFingerprint: "fingerprint",
					}).
					Return(&external.TerminateAccessResponse{
						TerminatedAt: terminatedAt,
					}, nil)

				mocks.db.EXPECT().
					GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
					Return(&database.AccessProof{ID: uint(2)}, nil)

				mocks.db.EXPECT().
					TerminateAccessProof(ctx, uint(2), terminatedAt.AsTime()).
					Return(fmt.Errorf("arbitrary error"))

				return ctx
			},
			req: &api.TerminateAccessProofRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "service-name",
				PublicKeyFingerprint:     "fingerprint",
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "internal"),
		},
		"happy_flow_outgoing_access_request_not_found": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := testCreateAdminUserContext()

				mocks.db.
					EXPECT().
					GetLatestOutgoingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(nil, database.ErrNotFound)

				mocks.al.EXPECT().
					AccessTerminate(ctx, "admin@example.com", "nlxctl", "00000000000000000001", "service-name", "fingerprint").
					Return(int64(1), nil)

				mocks.al.EXPECT().
					SetAsSucceeded(ctx, int64(1)).
					Return(nil)

				return ctx
			},
			req: &api.TerminateAccessProofRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "service-name",
				PublicKeyFingerprint:     "fingerprint",
			},
			want: &api.TerminateAccessProofResponse{},
		},
		"happy_flow": {
			setup: func(t *testing.T, mocks serviceMocks) context.Context {
				ctx := testCreateAdminUserContext()
				terminatedAt := timestamppb.New(time.Now())

				mocks.db.
					EXPECT().
					GetLatestOutgoingAccessRequest(ctx, "00000000000000000001", "service-name", "fingerprint").
					Return(&database.OutgoingAccessRequest{ID: 1, State: database.OutgoingAccessRequestReceived}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "00000000000000000001").
					Return("proxy.address", nil)

				mocks.al.EXPECT().
					AccessTerminate(ctx, "admin@example.com", "nlxctl", "00000000000000000001", "service-name", "fingerprint").
					Return(int64(1), nil)

				mocks.mc.EXPECT().
					TerminateAccess(ctx, &external.TerminateAccessRequest{
						ServiceName:          "service-name",
						PublicKeyFingerprint: "fingerprint",
					}).
					Return(&external.TerminateAccessResponse{
						TerminatedAt: terminatedAt,
					}, nil)

				mocks.db.EXPECT().
					GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
					Return(&database.AccessProof{ID: uint(2)}, nil)

				mocks.db.EXPECT().
					TerminateAccessProof(ctx, uint(2), terminatedAt.AsTime()).
					Return(nil)

				mocks.al.EXPECT().
					SetAsSucceeded(ctx, int64(1)).
					Return(nil)

				return ctx
			},
			req: &api.TerminateAccessProofRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "service-name",
				PublicKeyFingerprint:     "fingerprint",
			},
			want: &api.TerminateAccessProofResponse{},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			ctx := tt.setup(t, mocks)

			actual, err := service.TerminateAccessProof(ctx, tt.req)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, actual)
		})
	}
}
