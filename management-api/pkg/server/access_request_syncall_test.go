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

	common_grpcerrors "go.nlx.io/nlx/common/grpcerrors"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/grpcerrors"
)

//nolint:funlen // this is a test
func Test_SyncAllOutgoingAccessRequests(t *testing.T) {
	type testCase struct {
		ctx     context.Context
		setup   func(mocks serviceMocks)
		want    *emptypb.Empty
		wantErr error
	}

	testCases := map[string]testCase{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(mocks serviceMocks) {},
			want:    nil,
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.outgoing_access_requests.sync\" to execute this request").Err(),
		},
		"db_fails_to_retrieve_all_outgoing_access_requests": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListAllLatestOutgoingAccessRequests(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			want:    nil,
			wantErr: status.New(codes.Internal, "internal error").Err(),
		},
		"failed_to_retrieve_inway_proxy_address": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListAllLatestOutgoingAccessRequests(gomock.Any()).
					Return([]*database.OutgoingAccessRequest{
						{ID: 42, Organization: database.Organization{SerialNumber: "00000000000000000001"}},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000001").
					Return("", errors.New("arbitrary error"))

				mocks.dc.
					EXPECT().
					ListOrganizations(gomock.Any(), &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "my-organization",
							},
						},
					}, nil)
			},
			want: nil,
			wantErr: grpcerrors.NewInternal("unreachable organizations", &common_grpcerrors.Metadata{
				Metadata: map[string]string{
					"organizations": "my-organization",
				},
			}),
		},
		"happy_flow_no_outgoing_access_requests": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListAllLatestOutgoingAccessRequests(gomock.Any()).
					Return([]*database.OutgoingAccessRequest{}, nil)
			},
			want:    &emptypb.Empty{},
			wantErr: nil,
		},
		"happy_flow_multiple_access_requests_same_organization": {
			ctx: testCreateAdminUserContext(),
			setup: func(mocks serviceMocks) {
				mocks.db.
					EXPECT().
					ListAllLatestOutgoingAccessRequests(gomock.Any()).
					Return([]*database.OutgoingAccessRequest{
						{ID: 42, Organization: database.Organization{SerialNumber: "00000000000000000002"}, State: database.OutgoingAccessRequestReceived},
						{ID: 43, Organization: database.Organization{SerialNumber: "00000000000000000002"}, State: database.OutgoingAccessRequestReceived},
						{ID: 44, Organization: database.Organization{SerialNumber: "00000000000000000003"}, State: database.OutgoingAccessRequestReceived},
						{ID: 45, Organization: database.Organization{SerialNumber: "00000000000000000003"}, State: database.OutgoingAccessRequestReceived},
					}, nil)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000002").
					Return("", nil).
					MaxTimes(1)

				mocks.dc.
					EXPECT().
					GetOrganizationInwayProxyAddress(gomock.Any(), "00000000000000000003").
					Return("", nil).
					MaxTimes(1)

				mocks.mc.EXPECT().
					GetAccessRequestState(gomock.Any(), gomock.Any()).
					Return(nil, status.Error(codes.NotFound, "service does not exist")).
					AnyTimes()

				mocks.db.
					EXPECT().
					DeleteOutgoingAccessRequests(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil).
					AnyTimes()

				mocks.mc.
					EXPECT().
					Close().
					Return(nil).
					Times(2)
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
			got, err := service.SynchronizeAllOutgoingAccessRequests(tc.ctx, &emptypb.Empty{})

			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
