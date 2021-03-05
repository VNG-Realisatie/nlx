// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	context "context"
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

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
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
	Inways        []*inspectionapi.Inway
}{
	{
		api.DirectoryService_unknown,
		nil,
	},
	{
		api.DirectoryService_unknown,
		[]*inspectionapi.Inway{
			{State: inspectionapi.Inway_UNKNOWN},
		},
	},
	{
		api.DirectoryService_up,
		[]*inspectionapi.Inway{
			{State: inspectionapi.Inway_UP},
		},
	},
	{
		api.DirectoryService_up,
		[]*inspectionapi.Inway{
			{State: inspectionapi.Inway_UP},
			{State: inspectionapi.Inway_UP},
			{State: inspectionapi.Inway_UP},
		},
	},
	{
		api.DirectoryService_down,
		[]*inspectionapi.Inway{
			{State: inspectionapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_down,
		[]*inspectionapi.Inway{
			{State: inspectionapi.Inway_DOWN},
			{State: inspectionapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_degraded,
		[]*inspectionapi.Inway{
			{State: inspectionapi.Inway_UP},
			{State: inspectionapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_degraded,
		[]*inspectionapi.Inway{
			{State: inspectionapi.Inway_DOWN},
			{State: inspectionapi.Inway_UNKNOWN},
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

func TestListDirectoryServices(t *testing.T) {
	logger := zap.NewNop()
	env := &environment.Environment{}
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	clientServices := []*inspectionapi.ListServicesResponse_Service{
		{
			ServiceName:          "test-service-1",
			OrganizationName:     "test-organization-a",
			ApiSpecificationType: "OpenAPI3",
			DocumentationUrl:     "https://example.com",
			PublicSupportContact: "test@example.com",
			OneTimeCosts:         1,
			MonthlyCosts:         5,
			RequestCosts:         250,
		},
	}

	client := mock_directory.NewMockClient(mockCtrl)
	client.EXPECT().
		ListServices(ctx, &emptypb.Empty{}).
		Return(&inspectionapi.ListServicesResponse{Services: clientServices}, nil)

	db := mock_database.NewMockConfigDatabase(mockCtrl)
	service := clientServices[0]

	db.EXPECT().
		GetLatestOutgoingAccessRequest(ctx, service.OrganizationName, service.ServiceName).
		Return(&database.OutgoingAccessRequest{
			ID:                   1,
			OrganizationName:     "test-organization-a",
			ServiceName:          "test-service-1",
			State:                database.OutgoingAccessRequestCreated,
			PublicKeyFingerprint: "test-finger-print",
			CreatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
			UpdatedAt:            time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
		}, nil)

	db.EXPECT().
		GetAccessProofForOutgoingAccessRequest(ctx, uint(1)).
		Return(&database.AccessProof{
			ID:                      1,
			AccessRequestOutgoingID: 1,
			OutgoingAccessRequest: &database.OutgoingAccessRequest{
				ID:                   1,
				OrganizationName:     "test-organization-a",
				ServiceName:          "test-service-1",
				State:                database.OutgoingAccessRequestCreated,
				PublicKeyFingerprint: "test-finger-print",
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
			ServiceName:          "test-service-1",
			OrganizationName:     "test-organization-a",
			APISpecificationType: "OpenAPI3",
			DocumentationURL:     "https://example.com",
			PublicSupportContact: "test@example.com",
			State:                api.DirectoryService_unknown,
			OneTimeCosts:         1,
			MonthlyCosts:         5,
			RequestCosts:         250,
			LatestAccessProof: &api.AccessProof{
				Id:               1,
				OrganizationName: "test-organization-a",
				ServiceName:      "test-service-1",
				AccessRequestId:  1,
				CreatedAt:        timestampProto(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
			},
			LatestAccessRequest: &api.OutgoingAccessRequest{
				Id:               1,
				OrganizationName: "test-organization-a",
				ServiceName:      "test-service-1",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        timestampProto(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
				UpdatedAt:        timestampProto(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
			},
		},
	}

	assert.Equal(t, expected, response.Services)
}

func timestampProto(t time.Time) *types.Timestamp {
	tp, _ := types.TimestampProto(t)
	return tp
}

//nolint:funlen // this is a test method
func TestGetOrganizationService(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name            string
		req             *api.GetOrganizationServiceRequest
		db              func(db *mock_database.MockConfigDatabase)
		directoryClient func(directoryClient *mock_directory.MockClient)
		expectedReq     *api.DirectoryService
		expectedErr     error
	}{
		{
			"happy_flow",
			&api.GetOrganizationServiceRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
				db.EXPECT().
					GetLatestOutgoingAccessRequest(gomock.Any(), "test-organization", "test-service").
					Return(&database.OutgoingAccessRequest{
						ID:               1,
						ServiceName:      "test-service",
						OrganizationName: "test-organization",
						State:            database.OutgoingAccessRequestCreated,
						CreatedAt:        time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
						UpdatedAt:        time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
					}, nil)

				db.EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(&database.AccessProof{
						ID:                      1,
						AccessRequestOutgoingID: 1,
						OutgoingAccessRequest: &database.OutgoingAccessRequest{
							ID:               1,
							ServiceName:      "test-service",
							OrganizationName: "test-organization",
							State:            database.OutgoingAccessRequestCreated,
							CreatedAt:        time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
							UpdatedAt:        time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
						},
						CreatedAt: time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
					}, nil)
			},
			func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return(&inspectionapi.ListServicesResponse{
					Services: []*inspectionapi.ListServicesResponse_Service{{
						ServiceName:      "test-service",
						OrganizationName: "test-organization",
					}},
				}, nil)
			},
			&api.DirectoryService{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
				LatestAccessRequest: &api.OutgoingAccessRequest{
					Id:               1,
					OrganizationName: "test-organization",
					ServiceName:      "test-service",
					State:            api.AccessRequestState_CREATED,
					CreatedAt:        timestampProto(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
					UpdatedAt:        timestampProto(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
				},
				LatestAccessProof: &api.AccessProof{
					Id:               1,
					AccessRequestId:  1,
					ServiceName:      "test-service",
					OrganizationName: "test-organization",
					CreatedAt:        timestampProto(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
				},
			},
			nil,
		},
		{
			"happy_flow_without_latest_access_request_and_grant",
			&api.GetOrganizationServiceRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
				db.
					EXPECT().
					GetLatestOutgoingAccessRequest(gomock.Any(), "test-organization", "test-service").
					Return(nil, database.ErrNotFound)
			},
			func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return(&inspectionapi.ListServicesResponse{
					Services: []*inspectionapi.ListServicesResponse_Service{{
						ServiceName:      "test-service",
						OrganizationName: "test-organization",
					}},
				}, nil)
			},
			&api.DirectoryService{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			nil,
		},
		{
			"directory_call_fails",
			&api.GetOrganizationServiceRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
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
			"database_call_fail_get_latest_outgoing_access_request",
			&api.GetOrganizationServiceRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
				db.
					EXPECT().
					GetLatestOutgoingAccessRequest(gomock.Any(), "test-organization", "test-service").
					Return(nil, errors.New("arbitrary error"))
			},

			func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return(&inspectionapi.ListServicesResponse{
					Services: []*inspectionapi.ListServicesResponse_Service{{
						ServiceName:      "test-service",
						OrganizationName: "test-organization",
					}},
				}, nil)
			},
			nil,
			status.Error(codes.Internal, "database error"),
		},
		{
			"database_call_fail_get_latest_access_proof",
			&api.GetOrganizationServiceRequest{
				OrganizationName: "test-organization",
				ServiceName:      "test-service",
			},
			func(db *mock_database.MockConfigDatabase) {
				db.
					EXPECT().
					GetLatestOutgoingAccessRequest(gomock.Any(), "test-organization", "test-service").
					Return(&database.OutgoingAccessRequest{
						ID:               1,
						OrganizationName: "test-organization",
						ServiceName:      "test-service",
					}, nil)

				db.
					EXPECT().
					GetAccessProofForOutgoingAccessRequest(gomock.Any(), uint(1)).
					Return(nil, errors.New("arbitrary error"))
			},

			func(directoryClient *mock_directory.MockClient) {
				directoryClient.
					EXPECT().
					ListServices(gomock.Any(), gomock.Any()).
					Return(&inspectionapi.ListServicesResponse{
						Services: []*inspectionapi.ListServicesResponse_Service{{
							ServiceName:      "test-service",
							OrganizationName: "test-organization",
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

			assert.Equal(t, tt.expectedReq, returnedService)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
