// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/environment"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func newDirectoryService(t *testing.T) (s *server.DirectoryService, m *mock_directory.MockClient, db *mock_database.MockConfigDatabase) {
	logger := zaptest.Logger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	db = mock_database.NewMockConfigDatabase(ctrl)

	m = mock_directory.NewMockClient(ctrl)

	s = server.NewDirectoryService(logger, &environment.Environment{}, m, db)

	return
}

var directoryServiceStateTests = []struct {
	ExpectedState api.DirectoryService_State
	Inways        []*directoryapi.Inway
}{
	{
		api.DirectoryService_unknown,
		nil,
	},
	{
		api.DirectoryService_unknown,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_UNKNOWN},
		},
	},
	{
		api.DirectoryService_up,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_UP},
		},
	},
	{
		api.DirectoryService_up,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_UP},
			{State: directoryapi.Inway_UP},
			{State: directoryapi.Inway_UP},
		},
	},
	{
		api.DirectoryService_down,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_down,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_DOWN},
			{State: directoryapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_degraded,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_UP},
			{State: directoryapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_degraded,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_DOWN},
			{State: directoryapi.Inway_UNKNOWN},
		},
	},
}

func TestDirectoryServiceState(t *testing.T) {
	for i, test := range directoryServiceStateTests {
		name := strconv.Itoa(i + 1)
		test := test

		t.Run(name, func(t *testing.T) {
			state := server.DetermineDirectoryServiceState(test.Inways)
			assert.Equal(t, test.ExpectedState, state)
		})
	}
}

//nolint dupl: this is a test
func TestListDirectoryServices(t *testing.T) {
	logger := zap.NewNop()
	env := &environment.Environment{}
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	clientServices := []*directoryapi.ListServicesResponse_Service{
		{
			Name: "test-service-1",
			Organization: &directoryapi.Organization{
				SerialNumber: "00000000000000000001",
				Name:         "test-organization-a",
			},
			ApiSpecificationType: "OpenAPI3",
			DocumentationUrl:     "https://example.com",
			PublicSupportContact: "test@example.com",
			Costs: &directoryapi.ListServicesResponse_Costs{
				OneTime: 1,
				Monthly: 5,
				Request: 250,
			},
		},
	}

	client := mock_directory.NewMockClient(mockCtrl)
	client.EXPECT().
		ListServices(ctx, &emptypb.Empty{}).
		Return(&directoryapi.ListServicesResponse{Services: clientServices}, nil)

	db := mock_database.NewMockConfigDatabase(mockCtrl)
	service := clientServices[0]

	db.EXPECT().
		ListLatestOutgoingAccessRequests(ctx, service.Organization.SerialNumber, service.Name).
		Return([]*database.OutgoingAccessRequest{
			{
				ID: 1,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization-a",
				},
				ServiceName:          "test-service-1",
				State:                database.OutgoingAccessRequestCreated,
				PublicKeyFingerprint: "public-key-fingerprint",
				CreatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
				UpdatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
			},
		}, nil)

	db.EXPECT().
		GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
		Return(&database.AccessProof{
			ID:                      1,
			AccessRequestOutgoingID: 1,
			OutgoingAccessRequest: &database.OutgoingAccessRequest{
				ID: 1,
				Organization: database.Organization{
					SerialNumber: "00000000000000000001",
					Name:         "test-organization-a",
				},
				ServiceName:          "test-service-1",
				State:                database.OutgoingAccessRequestCreated,
				PublicKeyFingerprint: "public-key-fingerprint",
				CreatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
				UpdatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
			},

			CreatedAt: time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
		}, nil)

	directoryService := server.NewDirectoryService(logger, env, client, db)
	response, err := directoryService.ListServices(ctx, &emptypb.Empty{})
	assert.NoError(t, err)

	expected := []*api.DirectoryService{
		{
			ServiceName: "test-service-1",
			Organization: &api.Organization{
				SerialNumber: "00000000000000000001",
				Name:         "test-organization-a",
			},
			ApiSpecificationType: "OpenAPI3",
			DocumentationURL:     "https://example.com",
			PublicSupportContact: "test@example.com",
			State:                api.DirectoryService_unknown,
			OneTimeCosts:         1,
			MonthlyCosts:         5,
			RequestCosts:         250,
			//nolint dupl: this is a test
			AccessStates: []*api.DirectoryService_AccessState{
				{
					AccessRequest: &api.OutgoingAccessRequest{
						Id: 1,
						Organization: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization-a",
						},
						ServiceName:          "test-service-1",
						PublicKeyFingerprint: "public-key-fingerprint",
						State:                api.AccessRequestState_CREATED,
						CreatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
						UpdatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
					},
					AccessProof: &api.AccessProof{
						Id: 1,
						Organization: &api.Organization{
							SerialNumber: "00000000000000000001",
							Name:         "test-organization-a",
						},
						ServiceName:          "test-service-1",
						AccessRequestId:      1,
						PublicKeyFingerprint: "public-key-fingerprint",
						CreatedAt:            timestamppb.New(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
					},
				},
			},
		},
	}

	assert.EqualValues(t, expected, response.Services)
}

// nolint:funlen,dupl // this is a test method
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

func TestGetTermsOfService(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		directoryClient func(directoryClient *mock_directory.MockClient)
		wantResult      *api.GetTermsOfServiceResponse
		wantErr         error
	}{
		"failed_to_fetch_from_directory_client": {
			directoryClient: func(directoryClient *mock_directory.MockClient) {
				directoryClient.
					EXPECT().
					GetTermsOfService(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			wantResult: nil,
			wantErr:    status.Error(codes.Internal, "unable to get terms of service from directory"),
		},
		"happy_flow": {
			directoryClient: func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().GetTermsOfService(gomock.Any(), gomock.Any()).Return(&directoryapi.GetTermsOfServiceResponse{
					Enabled: true,
					Url:     "https://example.com",
				}, nil)
			},
			wantResult: &api.GetTermsOfServiceResponse{
				Enabled: true,
				Url:     "https://example.com",
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mockDirectoryClient, _ := newDirectoryService(t)

			tt.directoryClient(mockDirectoryClient)

			result, err := service.GetTermsOfService(ctx, &emptypb.Empty{})

			assert.Equal(t, tt.wantResult, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
