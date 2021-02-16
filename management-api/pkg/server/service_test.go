// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestCreateService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	databaseService := &database.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		"username":               "Jane Doe",
		"grpcgateway-user-agent": "nlxctl",
	}))

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().CreateServiceWithInways(ctx, databaseService, []string{})

	auditLogger := mock_auditlog.NewMockLogger(mockCtrl)
	auditLogger.EXPECT().ServiceCreate(gomock.Any(), "Jane Doe", "nlxctl")

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, auditLogger)

	requestService := &api.CreateServiceRequest{
		Name:        "my-service",
		EndpointURL: "my-service.test",
		Inways:      []string{},
	}

	responseService, err := service.CreateService(ctx, requestService)
	if err != nil {
		t.Error("could not create service", err)
	}

	assert.Equal(t, &api.CreateServiceResponse{
		Name:        "my-service",
		EndpointURL: "my-service.test",
		Inways:      []string{},
	}, responseService)
}

func TestGetService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, mock_auditlog.NewMockLogger(mockCtrl))

	getServiceRequest := &api.GetServiceRequest{
		Name: "my-service",
	}

	mockDatabase.EXPECT().GetService(ctx, "my-service").Return(nil, database.ErrNotFound)

	_, actualError := service.GetService(ctx, getServiceRequest)
	expectedError := status.Error(codes.NotFound, "service not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	databaseService := &database.Service{
		Name:   "my-service",
		Inways: []*database.Inway{},
	}

	mockDatabase.EXPECT().GetService(ctx, "my-service").Return(databaseService, nil)

	getServiceResponse, err := service.GetService(ctx, getServiceRequest)
	assert.NoError(t, err)

	expectedResponse := &api.GetServiceResponse{
		Name:   "my-service",
		Inways: []string{},
	}

	assert.Equal(t, expectedResponse, getServiceResponse)
}

func TestUpdateService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	databaseService := &database.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		"username":               "Jane Doe",
		"grpcgateway-user-agent": "nlxctl",
	}))

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().GetService(ctx, "my-service").Return(databaseService, nil)
	mockDatabase.EXPECT().GetService(ctx, "other-service").Return(nil, database.ErrNotFound)
	mockDatabase.EXPECT().UpdateServiceWithInways(ctx, databaseService, []string{})

	auditLogger := mock_auditlog.NewMockLogger(mockCtrl)
	auditLogger.EXPECT().ServiceUpdate(gomock.Any(), "Jane Doe", "nlxctl")

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, auditLogger)

	updateServiceRequest := &api.UpdateServiceRequest{
		Name:        "my-service",
		EndpointURL: "my-service.test",
		Inways:      []string{},
	}

	updateServiceResponse, err := service.UpdateService(ctx, updateServiceRequest)
	assert.NoError(t, err)

	expectedResponse := &api.UpdateServiceResponse{
		Name:        "my-service",
		EndpointURL: "my-service.test",
		Inways:      []string{},
	}

	assert.Equal(t, expectedResponse, updateServiceResponse)

	updateServiceRequest.Name = "other-service"

	_, err = service.UpdateService(ctx, updateServiceRequest)
	assert.EqualError(t, err, "rpc error: code = NotFound desc = service not found")
}

func TestDeleteService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().DeleteService(ctx, "my-service")

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, mock_auditlog.NewMockLogger(mockCtrl))

	deleteRequest := &api.DeleteServiceRequest{
		Name: "my-service",
	}

	_, err := service.DeleteService(ctx, deleteRequest)
	if err != nil {
		t.Error("could not delete service", err)
	}
}

//nolint:funlen // alot of scenario's to test
func TestListServices(t *testing.T) {
	databaseServices := []*database.Service{
		{
			Name:   "my-service",
			Inways: []*database.Inway{{Name: "inway.mock"}},
		},
		{
			Name:   "another-service",
			Inways: []*database.Inway{{Name: "another-inway.mock"}},
		},
		{
			Name: "third-service",
		},
	}

	tests := []struct {
		name             string
		request          *api.ListServicesRequest
		db               func(ctrl *gomock.Controller) database.ConfigDatabase
		expectedResponse *api.ListServicesResponse
		expectedError    error
	}{
		{
			name: "happy flow for a specific inway",
			request: &api.ListServicesRequest{
				InwayName: "inway.mock",
			},
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)

				db.EXPECT().GetInway(gomock.Any(), "inway.mock").Return(&database.Inway{
					Name: "inway.mock",
					Services: []*database.Service{
						{
							Name: "my-service",
							Inways: []*database.Inway{
								{
									Name: "inway.mock",
								},
							},
						},
					},
				}, nil)

				db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{}, nil)

				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "my-service").Return([]*database.AccessGrant{{
					ID:                      1,
					IncomingAccessRequestID: 1,
					IncomingAccessRequest: &database.IncomingAccessRequest{
						ID:               1,
						OrganizationName: "mock-organization-name",
						ServiceID:        1,
						Service: &database.Service{
							ID:   1,
							Name: "my-service",
						},
						PublicKeyFingerprint: "mock-publickey-fingerprint",
					},
				}}, nil)
				return db
			},
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.ListServicesResponse_Service{
					{
						Name:   "my-service",
						Inways: []string{"inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode: "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{
								{
									OrganizationName: "mock-organization-name",
									PublicKeyHash:    "mock-publickey-fingerprint",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "happy flow for another specific inway",
			request: &api.ListServicesRequest{
				InwayName: "another-inway.mock",
			},
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().GetInway(gomock.Any(), "another-inway.mock").Return(&database.Inway{
					Name: "another-inway.mock",
					Services: []*database.Service{
						{
							Name: "another-service",
							Inways: []*database.Inway{
								{
									Name: "another-inway.mock",
								},
							},
						},
					},
				}, nil)

				db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{}, nil)

				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "another-service").Return([]*database.AccessGrant{}, nil)
				return db
			},
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.ListServicesResponse_Service{
					{
						Name:   "another-service",
						Inways: []string{"another-inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			name:    "happy flow without inway filter",
			request: &api.ListServicesRequest{},
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().ListServices(gomock.Any()).Return(databaseServices, nil)

				db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{}, nil)

				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "my-service").Return([]*database.AccessGrant{}, nil)
				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "another-service").Return([]*database.AccessGrant{}, nil)
				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "third-service").Return([]*database.AccessGrant{}, nil)
				return db
			},
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.ListServicesResponse_Service{
					{
						Name:   "my-service",
						Inways: []string{"inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
					},
					{
						Name:   "another-service",
						Inways: []string{"another-inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
					},
					{
						Name:   "third-service",
						Inways: []string{},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			name:    "happy flow with incoming access requests",
			request: &api.ListServicesRequest{},
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().ListServices(gomock.Any()).Return(databaseServices, nil)

				db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{
					"my-service":      2,
					"another-service": 0,
					"third-service":   0,
				}, nil)

				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "my-service").Return([]*database.AccessGrant{}, nil)
				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "another-service").Return([]*database.AccessGrant{}, nil)
				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "third-service").Return([]*database.AccessGrant{}, nil)
				return db
			},
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.ListServicesResponse_Service{
					{
						Name:   "my-service",
						Inways: []string{"inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
						IncomingAccessRequestCount: 2,
					},
					{
						Name:   "another-service",
						Inways: []string{"another-inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
						IncomingAccessRequestCount: 0,
					},
					{
						Name:   "third-service",
						Inways: []string{},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
						IncomingAccessRequestCount: 0,
					},
				},
			},
			expectedError: nil,
		},
		{
			name:    "when database call for service fails",
			request: &api.ListServicesRequest{},
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().ListServices(gomock.Any()).Return(nil, errors.New("arbitrary error"))
				return db
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		{
			name:    "when database for access grants fails",
			request: &api.ListServicesRequest{},
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().ListServices(gomock.Any()).Return(databaseServices, nil)

				db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{}, nil)

				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "my-service").Return(nil, errors.New("arbitrary error"))
				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "another-service").Return(nil, errors.New("arbitrary error"))
				db.EXPECT().ListAccessGrantsForService(gomock.Any(), "third-service").Return(nil, errors.New("arbitrary error"))
				return db
			},
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.ListServicesResponse_Service{},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			logger := zap.NewNop()
			testProcess := process.NewProcess(logger)
			ctrl := gomock.NewController(t)
			ctx := context.Background()

			service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(ctrl), nil, test.db(ctrl), mock_auditlog.NewMockLogger(ctrl))
			actualResponse, err := service.ListServices(ctx, test.request)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedResponse, actualResponse)
		})
	}
}

func TestGetStatisticsOfServices(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().GetIncomingAccessRequestCountByService(ctx).Return(map[string]int{
		"service-a": 3,
	}, nil)

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, mock_auditlog.NewMockLogger(mockCtrl))

	requestGetStatisticsOfServices := &api.GetStatisticsOfServicesRequest{}

	responseService, err := service.GetStatisticsOfServices(ctx, requestGetStatisticsOfServices)
	if err != nil {
		t.Error("could not get stats for services", err)
	}

	assert.Equal(t, &api.GetStatisticsOfServicesResponse{
		Services: []*api.ServiceStatistics{
			{
				Name:                       "service-a",
				IncomingAccessRequestCount: 3,
			},
		},
	}, responseService)

	mockDatabase.EXPECT().GetIncomingAccessRequestCountByService(ctx).Return(nil, errors.New("arbitrary error"))

	_, err = service.GetStatisticsOfServices(ctx, requestGetStatisticsOfServices)
	assert.Error(t, err)
}
