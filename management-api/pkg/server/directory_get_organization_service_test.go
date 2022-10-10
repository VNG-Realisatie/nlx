// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // its a unittest
func TestGetOrganizationService(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		req     *api.GetOrganizationServiceRequest
		setup   func(context.Context, directoryServiceMocks)
		want    *api.DirectoryService
		wantErr error
	}{
		"happy_flow": {
			req: &api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			setup: func(ctx context.Context, mocks directoryServiceMocks) {
				mocks.d.
					EXPECT().
					ListOrganizations(ctx, &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{
						Organizations: []*directoryapi.Organization{
							{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
						},
					}, nil)

				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "test-service").
					Return([]*database.OutgoingAccessRequest{{
						ID:                   1,
						ServiceName:          "test-service",
						PublicKeyFingerprint: "public-key-fingerprint",
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
						},
						State:     database.OutgoingAccessRequestReceived,
						CreatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
						UpdatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
					}}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(&database.AccessProof{
						ID:                      1,
						AccessRequestOutgoingID: 1,
						OutgoingAccessRequest: &database.OutgoingAccessRequest{
							ID:          1,
							ServiceName: "test-service",
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
							},
							State:     database.OutgoingAccessRequestReceived,
							CreatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
							UpdatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
						},
						CreatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
					}, nil)

				mocks.d.
					EXPECT().
					ListServices(gomock.Any(), gomock.Any()).
					Return(&directoryapi.ListServicesResponse{
						Services: []*directoryapi.ListServicesResponse_Service{{
							Name: "test-service",
							Organization: &directoryapi.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
						}},
					}, nil)
			},
			want: &api.DirectoryService{
				Organization: &api.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "Organization One",
				},
				ServiceName: "test-service",
				//nolint dupl: this is a test
				AccessStates: []*api.DirectoryService_AccessState{
					{
						AccessRequest: &api.OutgoingAccessRequest{
							Id: 1,
							Organization: &api.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
							ServiceName:          "test-service",
							PublicKeyFingerprint: "public-key-fingerprint",
							State:                external.AccessRequestState_ACCESS_REQUEST_STATE_RECEIVED,
							CreatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
							UpdatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
						},
						AccessProof: &api.AccessProof{
							Id: 1,
							Organization: &api.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "Organization One",
							},
							ServiceName:     "test-service",
							AccessRequestId: 1,
							CreatedAt:       timestamppb.New(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
						},
					},
				},
			},
			wantErr: nil,
		},
		"happy_flow_without_latest_access_request_and_grant": {
			req: &api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			setup: func(ctx context.Context, mocks directoryServiceMocks) {
				mocks.d.
					EXPECT().
					ListOrganizations(ctx, &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{}, nil)

				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "test-service").
					Return([]*database.OutgoingAccessRequest{}, nil)

				mocks.d.
					EXPECT().
					ListServices(gomock.Any(), gomock.Any()).
					Return(&directoryapi.ListServicesResponse{
						Services: []*directoryapi.ListServicesResponse_Service{{
							Name: "test-service",
							Organization: &directoryapi.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "test-organization",
							},
						}},
					}, nil)
			},
			want: &api.DirectoryService{
				Organization: &api.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName:  "test-service",
				AccessStates: []*api.DirectoryService_AccessState{},
			},
			wantErr: nil,
		},
		"directory_call_fails": {
			req: &api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			setup: func(ctx context.Context, mocks directoryServiceMocks) {
				mocks.d.
					EXPECT().
					ListServices(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "directory not available"),
		},
		"database_call_fail_list_latest_outgoing_access_requests": {
			req: &api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			setup: func(ctx context.Context, mocks directoryServiceMocks) {
				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "test-service").
					Return(nil, errors.New("arbitrary error"))

				mocks.d.
					EXPECT().
					ListServices(gomock.Any(), gomock.Any()).Return(&directoryapi.ListServicesResponse{
					Services: []*directoryapi.ListServicesResponse_Service{{
						Name: "test-service",
						Organization: &directoryapi.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization",
						},
					}},
				}, nil)
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"database_call_fail_get_latest_access_proof": {
			req: &api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			setup: func(ctx context.Context, mocks directoryServiceMocks) {
				mocks.d.
					EXPECT().
					ListOrganizations(ctx, &directoryapi.ListOrganizationsRequest{}).
					Return(&directoryapi.ListOrganizationsResponse{}, nil)

				mocks.db.
					EXPECT().
					ListLatestOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "test-service").
					Return([]*database.OutgoingAccessRequest{{
						ID: 1,
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization",
						},
						ServiceName: "test-service",
					}}, nil)

				mocks.db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(nil, errors.New("arbitrary error"))

				mocks.d.
					EXPECT().
					ListServices(gomock.Any(), gomock.Any()).
					Return(&directoryapi.ListServicesResponse{
						Services: []*directoryapi.ListServicesResponse_Service{{
							Name: "test-service",
							Organization: &directoryapi.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "test-organization",
							},
						}},
					}, nil)
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "database error"),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mocks := newDirectoryService(t)

			tt.setup(ctx, mocks)

			returnedService, err := service.GetOrganizationService(ctx, tt.req)

			assert.Equal(t, tt.want, returnedService)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
