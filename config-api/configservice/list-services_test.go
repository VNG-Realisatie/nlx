// nolint:dupl
package configservice_test

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/config-api/configservice"
	mock "go.nlx.io/nlx/config-api/configservice/mock"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

func TestListServices(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock.NewMockConfigDatabase(mockCtrl)
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
