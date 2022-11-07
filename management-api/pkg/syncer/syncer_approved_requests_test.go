// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package syncer_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/server"
	"go.nlx.io/nlx/management-api/pkg/syncer"
)

//nolint:funlen // this is a test
func getApprovedTestCases(t *testing.T) syncOutgoingAccessRequestTestCases {
	certBundle, err := newCertificateBundle()
	require.NoError(t, err)

	testPublicKeyFingerprint := certBundle.PublicKeyFingerprint()

	return syncOutgoingAccessRequestTestCases{
		"failed_to_get_access_proof": {
			setup: func(mocks syncMocks) {
				mocks.mc.
					EXPECT().
					GetAccessGrant(gomock.Any(),
						&external.GetAccessGrantRequest{
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
							State: database.OutgoingAccessRequestApproved,
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
		"service_of_access_proof_was_removed": {
			setup: func(mocks syncMocks) {
				mocks.mc.
					EXPECT().
					GetAccessGrant(gomock.Any(),
						&external.GetAccessGrantRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(nil, server.ErrServiceDoesNotExist)

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(
						gomock.Any(),
						"00000000000000000001",
						"my-service",
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
							State: database.OutgoingAccessRequestApproved,
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
		"access_proof_is_linked_to_different_access_request": {
			setup: func(mocks syncMocks) {
				mocks.mc.
					EXPECT().
					GetAccessGrant(gomock.Any(),
						&external.GetAccessGrantRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&external.GetAccessGrantResponse{
						AccessGrant: &external.AccessGrant{
							AccessRequestId: 999,
							Organization: &external.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName: "my-service",
							RevokedAt:   nil,
						},
					}, nil)
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
							State: database.OutgoingAccessRequestApproved,
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
		},
		"remote_acces_proof_is_revoked": {
			setup: func(mocks syncMocks) {
				revokedAt := timestamppb.New(time.Now())

				mocks.mc.
					EXPECT().
					GetAccessGrant(gomock.Any(),
						&external.GetAccessGrantRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&external.GetAccessGrantResponse{
						AccessGrant: &external.AccessGrant{
							AccessRequestId: 1,
							Organization: &external.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName: "my-service",
							RevokedAt:   revokedAt,
						},
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
							State: database.OutgoingAccessRequestApproved,
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
		"remote_acces_proof_is_terminated": {
			setup: func(mocks syncMocks) {
				terminatedAt := timestamppb.New(time.Now())

				mocks.mc.
					EXPECT().
					GetAccessGrant(gomock.Any(),
						&external.GetAccessGrantRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&external.GetAccessGrantResponse{
						AccessGrant: &external.AccessGrant{
							AccessRequestId: 1,
							Organization: &external.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName:  "my-service",
							TerminatedAt: terminatedAt,
						},
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(42)).
					Return(&database.AccessProof{
						ID: 2,
						TerminatedAt: sql.NullTime{
							Valid: false,
						},
					}, nil)

				mocks.db.
					EXPECT().
					TerminateAccessProof(
						gomock.Any(),
						uint(2),
						terminatedAt.AsTime(),
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
							State: database.OutgoingAccessRequestApproved,
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
		"remote_acces_proof_has_been_removed": {
			setup: func(mocks syncMocks) {
				mocks.mc.
					EXPECT().
					GetAccessGrant(gomock.Any(),
						&external.GetAccessGrantRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(nil, status.Error(codes.NotFound, "access grant not found"))

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(
						gomock.Any(),
						"00000000000000000001",
						"my-service",
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
							State: database.OutgoingAccessRequestApproved,
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
					GetAccessGrant(gomock.Any(),
						&external.GetAccessGrantRequest{
							ServiceName:          "my-service",
							PublicKeyFingerprint: testPublicKeyFingerprint,
						}).
					Return(&external.GetAccessGrantResponse{
						AccessGrant: &external.AccessGrant{
							AccessRequestId: 1,
							Organization: &external.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
							ServiceName: "my-service",
							RevokedAt:   nil,
						},
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(42)).
					Return(nil, database.ErrNotFound)

				mocks.db.
					EXPECT().
					CreateAccessProof(gomock.Any(), uint(42)).
					Return(nil, nil)
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
							State: database.OutgoingAccessRequestApproved,
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
