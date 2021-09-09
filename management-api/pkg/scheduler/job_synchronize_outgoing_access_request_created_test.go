// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint funlen: these are tests
package scheduler_test

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"

	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/server"
)

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
