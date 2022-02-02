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

func getCreatedAccessRequests() map[string]testCase {
	accessRequest := &database.OutgoingAccessRequest{
		ID: 1,
		Organization: database.Organization{
			SerialNumber: "00000000000000000001",
		},
		PublicKeyFingerprint: "public-key-fingerprint",
		ReferenceID:          2,
		ServiceName:          "service",
		State:                database.OutgoingAccessRequestCreated,
	}

	return map[string]testCase{
		"when_getting_the_organization_management_client_fails": {
			request: accessRequest,
			setupMocks: func(mocks schedulerMocks) {

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("", errors.New("arbitrary error"))

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestFailed, uint(0), gomock.Any(), gomock.Any()).
					Return(nil)
			},
		},
		"when_service_has_been_deleted": {
			request: accessRequest,
			setupMocks: func(mocks schedulerMocks) {

				mocks.directory.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("hostname:7200", nil)

				mocks.management.
					EXPECT().
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName:          "service",
						PublicKeyFingerprint: "public-key-fingerprint",
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
					RequestAccess(gomock.Any(), &external.RequestAccessRequest{
						ServiceName:          "service",
						PublicKeyFingerprint: "public-key-fingerprint",
					}, gomock.Any()).
					Return(&external.RequestAccessResponse{
						ReferenceId:        2,
						AccessRequestState: api.AccessRequestState_RECEIVED,
					}, nil)

				mocks.db.
					EXPECT().
					UpdateOutgoingAccessRequestState(gomock.Any(), uint(1), database.OutgoingAccessRequestReceived, uint(2), nil, gomock.Any()).
					Return(nil)

				mocks.management.
					EXPECT().
					Close().
					Return(nil)
			},
		},
	}
}
