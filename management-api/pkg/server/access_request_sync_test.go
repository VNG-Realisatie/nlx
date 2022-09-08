// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

type syncOutgoingAccessRequestTestCases map[string]struct {
	ctx        context.Context
	setupMocks func(mocks serviceMocks)
	req        *api.SynchronizeOutgoingAccessRequestsRequest
	want       *emptypb.Empty
	wantErr    error
}

func Test_SyncOutgoingAccessRequests(t *testing.T) {
	testGroups := map[string]syncOutgoingAccessRequestTestCases{
		"generic":  getGenericTestCases(),
		"received": getReceivedTestCases(t),
		"approved": getApprovedTestCases(t),
	}

	for groupName, testGroup := range testGroups {
		testGroup := testGroup

		for name, test := range testGroup {
			test := test

			t.Run(groupName+" "+name, func(t *testing.T) {
				service, _, mocks := newService(t)

				test.setupMocks(mocks)
				got, err := service.SynchronizeOutgoingAccessRequests(test.ctx, test.req)

				assert.Equal(t, test.want, got)
				assert.Equal(t, test.wantErr, err)
			})
		}
	}
}

func getGenericTestCases() syncOutgoingAccessRequestTestCases {
	return syncOutgoingAccessRequestTestCases{
		"missing_required_permission": {
			ctx:        testCreateUserWithoutPermissionsContext(),
			setupMocks: func(mocks serviceMocks) {},
			req:        &api.SynchronizeOutgoingAccessRequestsRequest{},
			want:       nil,
			wantErr:    status.New(codes.PermissionDenied, "user needs the permission \"permissions.outgoing_access_requests.sync\" to execute this request").Err(),
		},
		"db_fails_to_retrieve_latest_outgoing_access_requests": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(
						gomock.Any(),
						"00000000000000000001",
						"my-service",
					).
					Return(nil, errors.New("arbitrary error"))
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "internal error").Err(),
		},
		"failed_to_retrieve_inway_proxy_address": {
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
						{ID: 42},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("", errors.New("arbitrary error"))
			},
			req: &api.SynchronizeOutgoingAccessRequestsRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "my-service",
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "internal error").Err(),
		},
		"happy_flow_no_outgoing_access_requests": {
			ctx: testCreateAdminUserContext(),
			setupMocks: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(
						gomock.Any(),
						"00000000000000000001",
						"my-service",
					).Return([]*database.OutgoingAccessRequest{}, nil)
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
