//nolint:dupl // test package
package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestCreateService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	databaseService := &database.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().CreateService(ctx, databaseService)

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), mockDatabase)

	requestService := &api.Service{
		Name:                  "my-service",
		EndpointURL:           "my-service.test",
		AuthorizationSettings: &api.Service_AuthorizationSettings{Mode: "none"},
	}

	responseService, err := service.CreateService(ctx, requestService)
	if err != nil {
		t.Error("could not create service", err)
	}

	assert.Equal(t, requestService, responseService)
}

func TestGetService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), mockDatabase)

	getServiceRequest := &api.GetServiceRequest{
		Name: "my-service",
	}

	mockDatabase.EXPECT().GetService(ctx, "my-service")

	_, actualError := service.GetService(ctx, getServiceRequest)
	expectedError := status.Error(codes.NotFound, "service not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	databaseService := &database.Service{
		Name: "my-service",
	}

	mockDatabase.EXPECT().GetService(ctx, "my-service").Return(databaseService, nil)

	getServiceResponse, err := service.GetService(ctx, getServiceRequest)
	assert.NoError(t, err)

	expectedResponse := &api.Service{
		Name: "my-service",
	}

	assert.Equal(t, expectedResponse, getServiceResponse)
}

func TestUpdateService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	databseService := &database.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().UpdateService(ctx, "my-service", databseService)

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), mockDatabase)

	updateServiceRequest := &api.UpdateServiceRequest{
		Name: "my-service",
		Service: &api.Service{
			Name:                  "my-service",
			EndpointURL:           "my-service.test",
			AuthorizationSettings: &api.Service_AuthorizationSettings{Mode: "none"},
		},
	}

	updateServiceResponse, err := service.UpdateService(ctx, updateServiceRequest)
	assert.NoError(t, err)

	expectedResponse := &api.Service{
		Name:                  "my-service",
		EndpointURL:           "my-service.test",
		AuthorizationSettings: &api.Service_AuthorizationSettings{Mode: "none"},
	}

	assert.Equal(t, expectedResponse, updateServiceResponse)

	updateServiceRequest.Name = "other-name"

	_, err = service.UpdateService(ctx, updateServiceRequest)
	assert.Errorf(t, err, "changing the service name is not allowed")
}

func TestDeleteService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().DeleteService(ctx, "my-service")

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), mockDatabase)

	deleteRequest := &api.DeleteServiceRequest{
		Name: "my-service",
	}

	_, err := service.DeleteService(ctx, deleteRequest)
	if err != nil {
		t.Error("could not delete service", err)
	}
}

func TestListServices(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), mockDatabase)

	databaseServices := []*database.Service{
		{
			Name:   "my-service",
			Inways: []string{"inway.mock"},
		},
		{
			Name:   "another-service",
			Inways: []string{"another-inway.mock"},
		},
		{
			Name: "third-service",
		},
	}

	tests := []struct {
		request          *api.ListServicesRequest
		expectedResponse *api.ListServicesResponse
	}{
		{
			request: &api.ListServicesRequest{
				InwayName: "inway.mock",
			},
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.Service{
					{
						Name:   "my-service",
						Inways: []string{"inway.mock"},
					},
				},
			},
		},
		{
			request: &api.ListServicesRequest{
				InwayName: "another-inway.mock",
			},
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.Service{
					{
						Name:   "another-service",
						Inways: []string{"another-inway.mock"},
					},
				},
			},
		},
		{
			request: &api.ListServicesRequest{},
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.Service{
					{
						Name:   "my-service",
						Inways: []string{"inway.mock"},
					},
					{
						Name:   "another-service",
						Inways: []string{"another-inway.mock"},
					},
					{
						Name: "third-service",
					},
				},
			},
		},
	}

	for _, test := range tests {
		mockDatabase.EXPECT().ListServices(ctx).Return(databaseServices, nil)
		actualResponse, err := service.ListServices(ctx, test.request)

		if err != nil {
			t.Fatal("could not get list of services", err)
		}

		assert.Equal(t, test.expectedResponse, actualResponse)
	}
}
