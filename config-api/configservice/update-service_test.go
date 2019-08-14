// nolint:dupl
package configservice_test

import (
	"context"
	"testing"

	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/config-api/configservice"

	mock_configservice "go.nlx.io/nlx/config-api/configservice/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/process"
	"go.uber.org/zap"
)

func TestUpdateService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockService := &configapi.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().UpdateService(ctx, "my-service", mockService)

	service := configservice.New(logger, testProcess, mockDatabase)

	updateServiceRequest := &configapi.UpdateServiceRequest{
		Name:    "my-service",
		Service: mockService,
	}

	updateServiceResponse, err := service.UpdateService(ctx, updateServiceRequest)
	if err != nil {
		t.Error("could not update inway", err)
	}

	assert.Equal(t, mockService, updateServiceResponse)
}
