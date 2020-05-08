//nolint:dupl // test package
package configservice_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/configapi"
	"go.nlx.io/nlx/management-api/configservice"

	mock_configservice "go.nlx.io/nlx/management-api/configservice/mock"
)

func TestCreateService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	requestService := &configapi.Service{
		Name:                  "my-service",
		EndpointURL:           "my-service.test",
		AuthorizationSettings: &configapi.Service_AuthorizationSettings{Mode: "none"},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().CreateService(ctx, requestService)

	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

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

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	getServiceRequest := &configapi.GetServiceRequest{
		Name: "my-service",
	}

	mockDatabase.EXPECT().GetService(ctx, "my-service")

	_, actualError := service.GetService(ctx, getServiceRequest)
	expectedError := status.Error(codes.NotFound, "service not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	mockServiceResponse := &configapi.Service{
		Name: "my-service",
	}

	mockDatabase.EXPECT().GetService(ctx, "my-service").Return(mockServiceResponse, nil)

	getServiceResponse, err := service.GetService(ctx, getServiceRequest)
	if err != nil {
		t.Fatal("could not get service", err)
	}

	assert.Equal(t, mockServiceResponse, getServiceResponse)
}

func TestUpdateService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockService := &configapi.Service{
		Name:                  "my-service",
		EndpointURL:           "my-service.test",
		AuthorizationSettings: &configapi.Service_AuthorizationSettings{Mode: "none"},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().UpdateService(ctx, "my-service", mockService)

	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	updateServiceRequest := &configapi.UpdateServiceRequest{
		Name:    "my-service",
		Service: mockService,
	}

	updateServiceResponse, err := service.UpdateService(ctx, updateServiceRequest)
	assert.Nil(t, err)

	assert.Equal(t, mockService, updateServiceResponse)

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

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().DeleteService(ctx, "my-service")

	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	deleteRequest := &configapi.DeleteServiceRequest{
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

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)
	myService := &configapi.Service{
		Name:   "my-service",
		Inways: []string{"inway.mock"},
	}
	anotherService := &configapi.Service{
		Name:   "another-service",
		Inways: []string{"another-inway.mock"},
	}
	thirdService := &configapi.Service{
		Name: "third-service",
	}

	mockListServices := []*configapi.Service{
		myService,
		anotherService,
		thirdService,
	}

	tests := []struct {
		request          *configapi.ListServicesRequest
		expectedResponse *configapi.ListServicesResponse
	}{
		{
			request: &configapi.ListServicesRequest{
				InwayName: "inway.mock",
			},
			expectedResponse: &configapi.ListServicesResponse{
				Services: []*configapi.Service{myService},
			},
		},
		{
			request: &configapi.ListServicesRequest{
				InwayName: "another-inway.mock",
			},
			expectedResponse: &configapi.ListServicesResponse{
				Services: []*configapi.Service{anotherService},
			},
		},
		{
			request: &configapi.ListServicesRequest{},
			expectedResponse: &configapi.ListServicesResponse{
				Services: mockListServices,
			},
		},
	}

	for _, test := range tests {
		mockDatabase.EXPECT().ListServices(ctx).Return(mockListServices, nil)
		actualResponse, err := service.ListServices(ctx, test.request)

		if err != nil {
			t.Fatal("could not get list of services", err)
		}

		assert.Equal(t, test.expectedResponse, actualResponse)
	}
}
