// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test
func getApprovedTestCases(t *testing.T) syncOutgoingAccessRequestTestCases {
	certBundle, err := newCertificateBundle()
	require.NoError(t, err)

	testPublicKeyFingerprint := certBundle.PublicKeyFingerprint()

	return syncOutgoingAccessRequestTestCases{
		"failed_to_get_access_proof": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(
						gomock.Any(),
						"00000000000000000001",
						"my-service",
					).
					Return([]*database.OutgoingAccessRequest{
						{
							ID:    42,
							State: database.OutgoingAccessRequestApproved,
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
							ReferenceID:          1,
						},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					GetAccessProof(gomock.Any(),
						&external.GetAccessProofRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(nil, errors.New("arbitrary error"))

				mocks.mc.
					EXPECT().
					Close().
					Return(nil)
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "error occurred while syncing at least one Outgoing Access Request").Err(),
		},
		"access_proof_is_linked_to_different_access_request": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(
						gomock.Any(),
						"00000000000000000001",
						"my-service",
					).
					Return([]*database.OutgoingAccessRequest{
						{
							ID:    42,
							State: database.OutgoingAccessRequestApproved,
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
							ReferenceID:          1,
						},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					GetAccessProof(gomock.Any(),
						&external.GetAccessProofRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&api.AccessProof{
						AccessRequestId: 999,
						Organization: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "my-organization",
						},
						ServiceName: "my-service",
						RevokedAt:   nil,
					}, nil)

				week := time.Hour * 24 * 7

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(
						gomock.Any(),
						uint(42),
						database.OutgoingAccessRequestApproved,
						uint(0),
						nil,
						fixtureTime.Add(week),
					).
					Return(nil)

				mocks.mc.
					EXPECT().
					Close().
					Return(nil)
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
		"remote_acces_proof_is_revoked": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(
						gomock.Any(),
						"00000000000000000001",
						"my-service",
					).
					Return([]*database.OutgoingAccessRequest{
						{
							ID:    42,
							State: database.OutgoingAccessRequestApproved,
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
							ReferenceID:          1,
						},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				revokedAt := timestamppb.New(time.Now())

				mocks.mc.
					EXPECT().
					GetAccessProof(gomock.Any(),
						&external.GetAccessProofRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&api.AccessProof{
						AccessRequestId: 1,
						Organization: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "my-organization",
						},
						ServiceName: "my-service",
						RevokedAt:   revokedAt,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(42)).
					Return(&database.AccessProof{
						ID: 2,
						RevokedAt: sql.NullTime{
							Valid: false,
						},
					}, nil)

				mocks.db.
					EXPECT().
					RevokeAccessProof(
						gomock.Any(),
						uint(2),
						revokedAt.AsTime(),
					).
					Return(nil, nil)

				week := time.Hour * 24 * 7

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(
						gomock.Any(),
						uint(42),
						database.OutgoingAccessRequestApproved,
						uint(0),
						nil,
						fixtureTime.Add(week),
					).
					Return(nil)

				mocks.mc.
					EXPECT().
					Close().
					Return(nil)
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(
						gomock.Any(),
						"00000000000000000001",
						"my-service",
					).
					Return([]*database.OutgoingAccessRequest{
						{
							ID:    42,
							State: database.OutgoingAccessRequestApproved,
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
							ReferenceID:          1,
						},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					GetAccessProof(gomock.Any(),
						&external.GetAccessProofRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&api.AccessProof{
						AccessRequestId: 1,
						Organization: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "my-organization",
						},
						ServiceName: "my-service",
						RevokedAt:   nil,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(42)).
					Return(nil, database.ErrNotFound)

				mocks.db.
					EXPECT().
					CreateAccessProof(gomock.Any(), uint(42)).
					Return(nil, nil)

				week := time.Hour * 24 * 7

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(
						gomock.Any(),
						uint(42),
						database.OutgoingAccessRequestApproved,
						uint(0),
						nil,
						fixtureTime.Add(week),
					).
					Return(nil)

				mocks.mc.
					EXPECT().
					Close().
					Return(nil)
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
	}
}
