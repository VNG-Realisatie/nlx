// nolint:dupl
package configservice_test

import (
	"context"
	"testing"

	"go.nlx.io/nlx/config-api/configapi"

	"go.nlx.io/nlx/config-api/configservice"

	mock_configservice "go.nlx.io/nlx/config-api/configservice/mock"

	"github.com/golang/mock/gomock"

	"go.nlx.io/nlx/common/process"
	"go.uber.org/zap"
)

func TestDeleteService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().DeleteService(ctx, "my-service")

	service := configservice.New(logger, testProcess, mockDatabase)

	deleteRequest := &configapi.DeleteServiceRequest{
		Name: "my-service",
	}

	_, err := service.DeleteService(ctx, deleteRequest)
	if err != nil {
		t.Error("could not delete service", err)
	}
}
