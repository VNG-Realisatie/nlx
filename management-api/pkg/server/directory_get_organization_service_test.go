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
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
)

//nolint:funlen // its a unittest
func TestGetOrganizationService(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name            string
		req             *api.GetOrganizationServiceRequest
		db              func(db *mock_database.MockConfigDatabase)
		directoryClient func(directoryClient *mock_directory.MockClient)
		want            *api.DirectoryService
		wantErr         error
	}{
		{
			"happy_flow",
			&api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
				db.EXPECT().
					ListLatestOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "test-service").
					Return([]*database.OutgoingAccessRequest{{
						ID:                   1,
						ServiceName:          "test-service",
						PublicKeyFingerprint: "public-key-fingerprint",
						Organization: database.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization",
						},
						State:     database.OutgoingAccessRequestCreated,
						CreatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
						UpdatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
					}}, nil)

				db.EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(&database.AccessProof{
						ID:                      1,
						AccessRequestOutgoingID: 1,
						OutgoingAccessRequest: &database.OutgoingAccessRequest{
							ID:          1,
							ServiceName: "test-service",
							Organization: database.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "test-organization",
							},
							State:     database.OutgoingAccessRequestCreated,
							CreatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
							UpdatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
						},
						CreatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
					}, nil)
			},
			func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return(&directoryapi.ListServicesResponse{
					Services: []*directoryapi.ListServicesResponse_Service{{
						Name: "test-service",
						Organization: &directoryapi.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization",
						},
					}},
				}, nil)
			},
			&api.DirectoryService{
				Organization: &api.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName: "test-service",
				//nolint dupl: this is a test
				AccessStates: []*api.DirectoryService_AccessState{
					{
						AccessRequest: &api.OutgoingAccessRequest{
							Id: 1,
							Organization: &api.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "test-organization",
							},
							ServiceName:          "test-service",
							PublicKeyFingerprint: "public-key-fingerprint",
							State:                api.AccessRequestState_CREATED,
							CreatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
							UpdatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
						},
						AccessProof: &api.AccessProof{
							Id: 1,
							Organization: &api.Organization{
								SerialNumber: "00000000000000000001",
								Name:         "test-organization",
							},
							ServiceName:     "test-service",
							AccessRequestId: 1,
							CreatedAt:       timestamppb.New(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
						},
					},
				},
			},
			nil,
		},
		{
			"happy_flow_without_latest_access_request_and_grant",
			&api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
				db.
					EXPECT().
					ListLatestOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "test-service").
					Return([]*database.OutgoingAccessRequest{}, nil)
			},
			func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return(&directoryapi.ListServicesResponse{
					Services: []*directoryapi.ListServicesResponse_Service{{
						Name: "test-service",
						Organization: &directoryapi.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization",
						},
					}},
				}, nil)
			},
			&api.DirectoryService{
				Organization: &api.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization",
				},
				ServiceName:  "test-service",
				AccessStates: []*api.DirectoryService_AccessState{},
			},
			nil,
		},
		{
			"directory_call_fails",
			&api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
			},
			func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return(nil, errors.New("arbitrary error"))
			},
			nil,
			status.Error(codes.Internal, "directory not available"),
		},
		{
			"database_call_fail_list_latest_outgoing_access_requests",
			&api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
				db.
					EXPECT().
					ListLatestOutgoingAccessRequests(gomock.Any(), "00000000000000000001", "test-service").
					Return(nil, errors.New("arbitrary error"))
			},

			func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return(&directoryapi.ListServicesResponse{
					Services: []*directoryapi.ListServicesResponse_Service{{
						Name: "test-service",
						Organization: &directoryapi.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization",
						},
					}},
				}, nil)
			},
			nil,
			status.Error(codes.Internal, "database error"),
		},
		{
			"database_call_fail_get_latest_access_proof",
			&api.GetOrganizationServiceRequest{
				OrganizationSerialNumber: "00000000000000000001",
				ServiceName:              "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
				db.
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

				db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(nil, errors.New("arbitrary error"))
			},

			func(directoryClient *mock_directory.MockClient) {
				directoryClient.
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
			nil,
			status.Error(codes.Internal, "database error"),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			service, mockDirectoryClient, db := newDirectoryService(t)

			tt.db(db)
			tt.directoryClient(mockDirectoryClient)

			returnedService, err := service.GetOrganizationService(ctx, tt.req)

			assert.Equal(t, tt.want, returnedService)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
