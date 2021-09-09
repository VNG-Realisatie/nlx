// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint funlen: these are tests
package scheduler_test

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/server"
)

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
