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

func TestUpdateInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockInway := &configapi.Inway{
		Name: "inway42.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().UpdateInway(ctx, "inway42.test", mockInway)

	service := configservice.New(logger, testProcess, mockDatabase)

	updateInwayRequest := &configapi.UpdateInwayRequest{
		Name:  "inway42.test",
		Inway: mockInway,
	}

	updateInwayResponse, err := service.UpdateInway(ctx, updateInwayRequest)
	if err != nil {
		t.Error("could not update inway", err)
	}

	assert.Equal(t, mockInway, updateInwayResponse)
}
