// nolint:dupl
package configservice_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/config-api/configservice"
	mock_configservice "go.nlx.io/nlx/config-api/configservice/mock"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

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
