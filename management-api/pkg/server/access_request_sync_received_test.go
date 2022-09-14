// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/server"
)

//nolint:funlen // this is a test
func getReceivedTestCases(t *testing.T) syncOutgoingAccessRequestTestCases {
	certBundle, err := newCertificateBundle()
	require.NoError(t, err)

	testPublicKeyFingerprint := certBundle.PublicKeyFingerprint()

	return syncOutgoingAccessRequestTestCases{
		"failed_to_get_access_request_state": {
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
							State: database.OutgoingAccessRequestReceived,
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					Close().
					Return(nil)

				mocks.mc.
					EXPECT().
					GetAccessRequestState(gomock.Any(),
						&external.GetAccessRequestStateRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(nil, errors.New("arbitrary error"))
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "error occurred while syncing at least one Outgoing Access Request").Err(),
		},
		"service_has_been_removed": {
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
							State: database.OutgoingAccessRequestReceived,
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					Close().
					Return(nil)

				mocks.mc.
					EXPECT().
					GetAccessRequestState(gomock.Any(),
						&external.GetAccessRequestStateRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(nil, server.ErrServiceDoesNotExist)

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "my-service").
					Return(nil)
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
		"service_has_been_removed_failed_to_delete_access_proof": {
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
							State: database.OutgoingAccessRequestReceived,
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("address", nil)

				mocks.mc.
					EXPECT().
					Close().
					Return(nil)

				mocks.mc.
					EXPECT().
					GetAccessRequestState(gomock.Any(),
						&external.GetAccessRequestStateRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(nil, server.ErrServiceDoesNotExist)

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "my-service").
					Return(errors.New("arbitrary error"))
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "error occurred while syncing at least one Outgoing Access Request").Err(),
		},
		"failed_to_update_state": {
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
							State: database.OutgoingAccessRequestReceived,
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
					Close().
					Return(nil)

				mocks.mc.
					EXPECT().
					GetAccessRequestState(gomock.Any(),
						&external.GetAccessRequestStateRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_RECEIVED,
					}, nil)

				week := time.Hour * 24 * 7

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(
						gomock.Any(),
						uint(42),
						database.OutgoingAccessRequestReceived,
						uint(0),
						nil,
						fixtureTime.Add(week),
					).
					Return(errors.New("arbitrary error"))
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "error occurred while syncing at least one Outgoing Access Request").Err(),
		},
		"state_went_from_received_to_approved": {
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
							State: database.OutgoingAccessRequestReceived,
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
					Close().
					Return(nil)

				mocks.mc.
					EXPECT().
					GetAccessRequestState(gomock.Any(),
						&external.GetAccessRequestStateRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_APPROVED,
					}, nil)

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
							State: database.OutgoingAccessRequestReceived,
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
					Close().
					Return(nil)

				mocks.mc.
					EXPECT().
					GetAccessRequestState(gomock.Any(),
						&external.GetAccessRequestStateRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_RECEIVED,
					}, nil)

				week := time.Hour * 24 * 7

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(
						gomock.Any(),
						uint(42),
						database.OutgoingAccessRequestReceived,
						uint(0),
						nil,
						fixtureTime.Add(week),
					).
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
