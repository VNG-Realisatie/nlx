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

func Test_SyncOutgoingAccessRequests(t *testing.T) {
	type testCase struct {
		ctx     context.Context
		setup   func(mocks serviceMocks)
		req     *api.SynchronizeOutgoingAccessRequestsRequest
		want    *emptypb.Empty
		wantErr error
	}

	testCases := map[string]testCase{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(mocks serviceMocks) {},
			req:     &api.SynchronizeOutgoingAccessRequestsRequest{},
			want:    nil,
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.outgoing_access_requests.sync\" to execute this request").Err(),
		},
		"db_fails_to_retrieve_latest_outgoing_access_requests": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
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
			setup: func(mocks serviceMocks) {
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
			setup: func(mocks serviceMocks) {
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

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)

			tc.setup(mocks)
			got, err := service.SynchronizeOutgoingAccessRequests(tc.ctx, tc.req)

			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
