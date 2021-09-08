// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package scheduler

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
	"go.nlx.io/nlx/management-api/pkg/util/clock"
)

type schedulerMocks struct {
	db         *mock_database.MockConfigDatabase
	directory  *mock_directory.MockClient
	management *mock_management.MockClient
	ctrl       *gomock.Controller
}

func newTestScheduler(t *testing.T) (schedulerMocks, *scheduler) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	requests := make(chan *database.OutgoingAccessRequest, 10)
	mocks := schedulerMocks{
		ctrl:       ctrl,
		db:         mock_database.NewMockConfigDatabase(ctrl),
		directory:  mock_directory.NewMockClient(ctrl),
		management: mock_management.NewMockClient(ctrl),
	}
	scheduler := &scheduler{
		clock:           &clock.FakeClock{},
		logger:          zap.NewNop(),
		directoryClient: mocks.directory,
		configDatabase:  mocks.db,
		orgCert:         nil,
		requests:        requests,
		createManagementClientFunc: func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error) {
			return mocks.management, nil
		},
	}

	return mocks, scheduler
}

func TestSchedulePendingRequests(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		setupMock func(schedulerMocks)
		wantErr   bool
	}{
		"returns_an_error_when_list_outgoing_access_requests_errors": {
			wantErr: true,
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(ctx).
					Return(nil, errors.New("random error"))
			},
		},

		"schedules_received_requests": {
			wantErr: true,
			setupMock: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(ctx).
					Return(&database.OutgoingAccessRequest{
						// Give it an error state so that it won't actually be scheduled
						State: database.OutgoingAccessRequestFailed,
					}, nil)
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mocks, scheduler := newTestScheduler(t)

			tt.setupMock(mocks)

			err := scheduler.schedulePendingRequest(ctx)
			if tt.wantErr {
				assert.Error(t, err, "schedulePendingRequests should error")
			} else {
				assert.NoError(t, err, "schedulePendingRequests shouldn't error")
			}
		})
	}
}

//nolint:funlen // covers all test-cases of access-proofs
func TestSyncAccessProof(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		setupMocks func(schedulerMocks)
		request    *database.OutgoingAccessRequest
		wantErr    bool
	}{
		"returns_an_error_when_get_organization_inway_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("", errors.New("arbitrary error"))
			},
		},

		"returns_an_error_when_external_get_access_proof_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
				ServiceName:      "service",
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(nil, errors.New("random error"))

				mocks.management.
					EXPECT().
					Close()
			},
		},

		"returns_an_error_when_parsing_access_proof_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
				ServiceName:      "service",
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						CreatedAt: &timestamppb.Timestamp{
							// Trigger an invalid timestamp by setting the year to > 10.000
							Seconds: 553371149436,
						},
					}, nil)

				mocks.management.
					EXPECT().
					Close()
			},
		},

		"returns_an_error_when_database_getting_access_proof_for_outgoing_access_request_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
				ServiceName:      "service",
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						CreatedAt: timestamppb.Now(),
						RevokedAt: nil,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
					Return(nil, errors.New("random error"))

				mocks.management.
					EXPECT().
					Close()
			},
		},

		"returns_an_error_when_database_create_access_proof_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
				ServiceName:      "service",
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						RevokedAt:        nil,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
					Return(nil, database.ErrNotFound)

				mocks.db.
					EXPECT().
					CreateAccessProof(ctx, &database.OutgoingAccessRequest{
						ID:               1,
						OrganizationName: "organization-b",
						ServiceName:      "service",
					}).
					Return(nil, errors.New("random error"))

				mocks.management.
					EXPECT().
					Close()
			},
		},

		"returns_an_error_when_database_revoke_access_proof_errors": {
			wantErr: true,
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
				ServiceName:      "service",
			},
			setupMocks: func(mocks schedulerMocks) {
				ts := timestamppb.Now()
				t := timestamppb.New(ts.AsTime())

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						CreatedAt:        ts,
						RevokedAt:        ts,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
					Return(&database.AccessProof{
						ID: 2,
						OutgoingAccessRequest: &database.OutgoingAccessRequest{
							ID:               1,
							OrganizationName: "organization-b",
							ServiceName:      "service",
						},
						CreatedAt: t.AsTime(),
					}, nil)

				mocks.db.
					EXPECT().
					RevokeAccessProof(
						ctx,
						uint(2),
						t.AsTime(),
					).
					Return(nil, errors.New("random error"))

				mocks.management.
					EXPECT().
					Close()
			},
		},

		"successfully_revokes_an_access_grant_when_its_revoked": {
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
				ServiceName:      "service",
			},
			setupMocks: func(mocks schedulerMocks) {
				ts := timestamppb.Now()
				t := timestamppb.New(ts.AsTime())

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						CreatedAt:        ts,
						RevokedAt:        ts,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
					Return(&database.AccessProof{
						ID: 2,
						OutgoingAccessRequest: &database.OutgoingAccessRequest{
							ID:               1,
							OrganizationName: "organization-b",
							ServiceName:      "service",
						},
						CreatedAt: t.AsTime(),
					}, nil)

				mocks.db.
					EXPECT().
					RevokeAccessProof(
						ctx,
						uint(2),
						t.AsTime(),
					).
					Return(nil, nil)

				mocks.management.
					EXPECT().
					Close()
			},
		},

		"successfully_creates_an_access_proof_when_its_found": {
			wantErr: false,
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
				ServiceName:      "service",
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						RevokedAt:        nil,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
					Return(nil, database.ErrNotFound)

				mocks.db.
					EXPECT().
					CreateAccessProof(ctx, &database.OutgoingAccessRequest{
						ID:               1,
						OrganizationName: "organization-b",
						ServiceName:      "service",
					}).
					Return(nil, nil)

				mocks.management.
					EXPECT().
					Close()
			},
		},

		"successfully_delete_an_access_proof_when_the_corresponding_service_no_longer_exists": {
			wantErr: false,
			request: &database.OutgoingAccessRequest{
				ID:               1,
				OrganizationName: "organization-b",
				ServiceName:      "service",
			},
			setupMocks: func(mocks schedulerMocks) {
				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(ctx, "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(ctx, &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(nil, fmt.Errorf("mock grpc wrapper: %w", server.ErrServiceDoesNotExist))

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(ctx, "organization-b", "service").
					Return(nil)

				mocks.management.
					EXPECT().
					Close()
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mocks, scheduler := newTestScheduler(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tt.setupMocks(mocks)

			err := scheduler.syncAccessProof(ctx, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
