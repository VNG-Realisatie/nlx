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
		ID: 1,
		Organization: database.Organization{
			SerialNumber: "00000000000000000001",
		},
		ServiceName: "service",
		State:       database.OutgoingAccessRequestReceived,
	}

	return map[string]testCase{
		"when_updating_the_access_request_state_returns_an_error": {
			wantErr: errors.New("arbitrary error"),
			request: accessRequest,
			setupMocks: func(mocks schedulerMocks) {

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
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
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestApproved, uint(0), nil, gomock.Any()).
					Return(fmt.Errorf("arbitrary error"))

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
		"when_the_service_has_been_deleted": {
			request: accessRequest,
			setupMocks: func(mocks schedulerMocks) {

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					GetAccessRequestState(gomock.Any(), &external.GetAccessRequestStateRequest{
						ServiceName: "service",
					}, gomock.Any()).
					Return(nil, fmt.Errorf("mock grpc wrapper: %w", server.ErrServiceDoesNotExist))

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "service").
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
		"happy_flow": {
			request: accessRequest,
			setupMocks: func(mocks schedulerMocks) {

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
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
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestApproved, uint(0), nil, gomock.Any()).
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
	}
}
