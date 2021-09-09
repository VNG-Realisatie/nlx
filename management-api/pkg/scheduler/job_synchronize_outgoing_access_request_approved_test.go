// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint funlen: these are tests
package scheduler_test

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/server"
)

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
