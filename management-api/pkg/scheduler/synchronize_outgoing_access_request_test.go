// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint funlen: these are tests
package scheduler_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/scheduler"
	"go.nlx.io/nlx/management-api/pkg/server"
)

type testCase struct {
	setupMocks func(schedulerMocks)
	wantErr    error
}

func getGenericTests() map[string]testCase {
	return map[string]testCase{
		"when_taking_a_pending_access_request_errors": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			wantErr: errors.New("arbitrary error"),
		},
		"when_there_is_no_pending_access_request_available": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(nil, nil)
			},
		},
		"when_the_status_of_the_access_request_is_unknown": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(&database.OutgoingAccessRequest{
						State: "unknown state",
					}, nil)
			},
			wantErr: errors.New("invalid state 'unknown state' for pending access request"),
		},
	}
}

func getCreatedAccessRequests() map[string]testCase {
	accessRequest := &database.OutgoingAccessRequest{
		ID:               1,
		OrganizationName: "organization-a",
		ServiceName:      "service",
		State:            database.OutgoingAccessRequestCreated,
		ReferenceID:      2,
	}

	return map[string]testCase{
		"when_getting_the_organization_management_client_fails": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("", errors.New("arbitrary error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any()).
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"when_service_has_been_deleted": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(nil, fmt.Errorf("mock grpc wrapper: %w", server.ErrServiceDoesNotExist))

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), "organization-a", "service").
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"happy_flow": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&external.RequestAccessResponse{
						ReferenceId: 2,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestReceived, uint(2), nil).
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
	}
}

func getReceivedAccessRequests() map[string]testCase {
	accessRequest := &database.OutgoingAccessRequest{
		ID:               1,
		OrganizationName: "organization-a",
		ServiceName:      "service",
		State:            database.OutgoingAccessRequestReceived,
	}

	return map[string]testCase{
		"when_updating_the_access_request_state_returns_an_error": {
			wantErr: errors.New("arbitrary error"),
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(gomock.Any(), &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_APPROVED,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestApproved, uint(0), nil).
					Return(errors.New("arbitrary error"))

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
		"when_the_service_has_been_deleted": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(gomock.Any(), &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(nil, fmt.Errorf("mock grpc wrapper: %w", server.ErrServiceDoesNotExist))

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), "organization-a", "service").
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"happy_flow": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-a").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(gomock.Any(), &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(&external.GetAccessRequestStateResponse{
						State: api.AccessRequestState_APPROVED,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestApproved, uint(0), nil).
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
	}
}

func getApprovedAccessRequests() map[string]testCase {
	accessRequest := &database.OutgoingAccessRequest{
		ID:               1,
		OrganizationName: "organization-b",
		ServiceName:      "service",
		State:            database.OutgoingAccessRequestApproved,
	}

	return map[string]testCase{
		"when_getting_the_organization_inway_proxy_address_fails": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("", errors.New("arbitrary error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any()).
					Return(nil)

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"when_getting_the_access_proof_fails": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(gomock.Any(), &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(nil, errors.New("random error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any()).
					Return(nil)

				mocks.management.
					EXPECT().
					Close()

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"when_parsing_the_access_proof_fails": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(gomock.Any(), &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						CreatedAt: &timestamppb.Timestamp{
							// Trigger an invalid timestamp by setting the year to > 10.000
							Seconds: 553371149436,
						},
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any()).
					Return(nil)

				mocks.management.
					EXPECT().
					Close()

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},

		"when_database_getting_access_proof_errors": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(gomock.Any(), &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						CreatedAt: timestamppb.Now(),
						RevokedAt: nil,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(nil, errors.New("random error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any()).
					Return(nil)

				mocks.management.
					EXPECT().
					Close()

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"when_database_create_access_proof_errors": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(gomock.Any(), &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						RevokedAt:        nil,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(nil, database.ErrNotFound)

				mocks.db.
					EXPECT().
					CreateAccessProof(gomock.Any(), accessRequest).
					Return(nil, errors.New("random error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any()).
					Return(nil)

				mocks.management.
					EXPECT().
					Close()

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"when_database_revoke_access_proof_errors": {
			setupMocks: func(mocks schedulerMocks) {
				ts := timestamppb.Now()
				t := timestamppb.New(ts.AsTime())

				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(gomock.Any(), &external.GetAccessProofRequest{
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
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(&database.AccessProof{
						ID:                    2,
						OutgoingAccessRequest: accessRequest,
						CreatedAt:             t.AsTime(),
					}, nil)

				mocks.db.
					EXPECT().
					RevokeAccessProof(gomock.Any(), uint(2), t.AsTime()).
					Return(nil, errors.New("arbitrary error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any()).
					Return(nil)

				mocks.management.
					EXPECT().
					Close()

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"successfully_revokes_an_access_grant_when_its_revoked": {
			setupMocks: func(mocks schedulerMocks) {
				ts := timestamppb.Now()
				t := timestamppb.New(ts.AsTime())

				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(gomock.Any(), &external.GetAccessProofRequest{
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
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(&database.AccessProof{
						ID:                    2,
						OutgoingAccessRequest: accessRequest,
						CreatedAt:             t.AsTime(),
					}, nil)

				mocks.db.
					EXPECT().
					RevokeAccessProof(gomock.Any(), uint(2), t.AsTime()).
					Return(nil, nil)

				mocks.management.
					EXPECT().
					Close()

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"successfully_creates_an_access_proof_when_its_found": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(gomock.Any(), &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(&api.AccessProof{
						OrganizationName: "organization-a",
						ServiceName:      "service",
						RevokedAt:        nil,
					}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(nil, database.ErrNotFound)

				mocks.db.
					EXPECT().
					CreateAccessProof(gomock.Any(), &database.OutgoingAccessRequest{
						ID:               1,
						OrganizationName: "organization-b",
						ServiceName:      "service",
						State:            database.OutgoingAccessRequestApproved,
					}).
					Return(nil, nil)

				mocks.management.
					EXPECT().
					Close()

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
		"successfully_delete_an_access_proof_when_the_corresponding_service_no_longer_exists": {
			setupMocks: func(mocks schedulerMocks) {
				mocks.db.
					EXPECT().
					TakePendingOutgoingAccessRequest(gomock.Any()).
					Return(accessRequest, nil)

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "organization-b").
					Return("localhost:8000", nil)

				mocks.management.
					EXPECT().
					GetAccessProof(gomock.Any(), &external.GetAccessProofRequest{
						ServiceName: "service",
					}).
					Return(nil, fmt.Errorf("mock grpc wrapper: %w", server.ErrServiceDoesNotExist))

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), "organization-b", "service").
					Return(nil)

				mocks.management.
					EXPECT().
					Close()

				mocks.db.
					EXPECT().
					UnlockOutgoingAccessRequest(gomock.Any(), accessRequest)
			},
		},
	}
}

// nolint:funlen,dupl // this is a test
func TestSynchronizeOutgoingAccessRequest(t *testing.T) {
	testGroups := []map[string]testCase{
		getGenericTests(),
		getCreatedAccessRequests(),
		getReceivedAccessRequests(),
		getApprovedAccessRequests(),
	}

	for _, tests := range testGroups {
		for name, tt := range tests {
			tt := tt

			t.Run(name, func(t *testing.T) {
				mocks := newMocks(t)

				tt.setupMocks(mocks)

				job := scheduler.NewSynchronizeOutgoingAccessRequestJob(
					context.Background(),
					mocks.directory,
					mocks.db,
					nil,
					func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error) {
						return mocks.management, nil
					},
				)
				err := job.Run(context.Background())
				require.Equal(t, tt.wantErr, err)
			})
		}
	}
}

type schedulerMocks struct {
	db         *mock_database.MockConfigDatabase
	directory  *mock_directory.MockClient
	management *mock_management.MockClient
	ctrl       *gomock.Controller
}

func newMocks(t *testing.T) schedulerMocks {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := schedulerMocks{
		ctrl:       ctrl,
		db:         mock_database.NewMockConfigDatabase(ctrl),
		directory:  mock_directory.NewMockClient(ctrl),
		management: mock_management.NewMockClient(ctrl),
	}

	return mocks
}
