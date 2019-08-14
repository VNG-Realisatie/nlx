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

func TestDeleteInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().DeleteInway(ctx, "inway42.test")

	service := configservice.New(logger, testProcess, mockDatabase)

	deleteRequest := &configapi.DeleteInwayRequest{
		Name: "inway42.test",
	}

	_, err := service.DeleteInway(ctx, deleteRequest)
	if err != nil {
		t.Error("could not delete inway", err)
	}
}
