// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package syncer_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/server"
	"go.nlx.io/nlx/management-api/pkg/syncer"
)

//nolint:funlen // this is a test
func getReceivedTestCases(t *testing.T) syncOutgoingAccessRequestTestCases {
	certBundle, err := newCertificateBundle()
	require.NoError(t, err)

	testPublicKeyFingerprint := certBundle.PublicKeyFingerprint()

	return syncOutgoingAccessRequestTestCases{
		"failed_to_get_access_request_state": {
			setup: func(mocks syncMocks) {
				mocks.mc.
					EXPECT().
					GetAccessRequestState(gomock.Any(),
						&external.GetAccessRequestStateRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(nil, errors.New("arbitrary error"))
			},
			createArgs: func(mocks syncMocks) *syncer.SyncArgs {
				return &syncer.SyncArgs{
					Ctx:    context.Background(),
					Logger: zap.NewNop(),
					DB:     mocks.db,
					Client: mocks.mc,
					Requests: []*database.OutgoingAccessRequest{
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
					},
				}
			},
			want: errors.New("error occurred while syncing at least one Outgoing Access Request"),
		},
		"service_has_been_removed": {
			setup: func(mocks syncMocks) {
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
			createArgs: func(mocks syncMocks) *syncer.SyncArgs {
				return &syncer.SyncArgs{
					Ctx:    context.Background(),
					Logger: zap.NewNop(),
					DB:     mocks.db,
					Client: mocks.mc,
					Requests: []*database.OutgoingAccessRequest{
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
					},
				}
			},
			want: nil,
		},
		"service_has_been_removed_failed_to_delete_access_proof": {
			setup: func(mocks syncMocks) {
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
			createArgs: func(mocks syncMocks) *syncer.SyncArgs {
				return &syncer.SyncArgs{
					Ctx:    context.Background(),
					Logger: zap.NewNop(),
					DB:     mocks.db,
					Client: mocks.mc,
					Requests: []*database.OutgoingAccessRequest{
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
					},
				}
			},
			want: errors.New("error occurred while syncing at least one Outgoing Access Request"),
		},
		"failed_to_update_state": {
			setup: func(mocks syncMocks) {
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

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(
						gomock.Any(),
						uint(42),
						database.OutgoingAccessRequestReceived,
						uint(0),
						nil,
					).
					Return(errors.New("arbitrary error"))
			},
			createArgs: func(mocks syncMocks) *syncer.SyncArgs {
				return &syncer.SyncArgs{
					Ctx:    context.Background(),
					Logger: zap.NewNop(),
					DB:     mocks.db,
					Client: mocks.mc,
					Requests: []*database.OutgoingAccessRequest{
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
					},
				}
			},
			want: errors.New("error occurred while syncing at least one Outgoing Access Request"),
		},
		"state_went_from_received_to_approved": {
			setup: func(mocks syncMocks) {
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

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(
						gomock.Any(),
						uint(42),
						database.OutgoingAccessRequestApproved,
						uint(0),
						nil,
					).
					Return(nil)
			},
			createArgs: func(mocks syncMocks) *syncer.SyncArgs {
				return &syncer.SyncArgs{
					Ctx:    context.Background(),
					Logger: zap.NewNop(),
					DB:     mocks.db,
					Client: mocks.mc,
					Requests: []*database.OutgoingAccessRequest{
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
					},
				}
			},
			want: nil,
		},
		"happy_flow": {
			setup: func(mocks syncMocks) {
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

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(
						gomock.Any(),
						uint(42),
						database.OutgoingAccessRequestReceived,
						uint(0),
						nil,
					).
					Return(nil)
			},
			createArgs: func(mocks syncMocks) *syncer.SyncArgs {
				return &syncer.SyncArgs{
					Ctx:    context.Background(),
					Logger: zap.NewNop(),
					DB:     mocks.db,
					Client: mocks.mc,
					Requests: []*database.OutgoingAccessRequest{
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
					},
				}
			},
			want: nil,
		},
	}
}
