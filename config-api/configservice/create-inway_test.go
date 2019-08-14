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

func TestCreateInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	requestInway := &configapi.Inway{
		Name: "inway42.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().CreateInway(ctx, requestInway)

	service := configservice.New(logger, testProcess, mockDatabase)

	responseInway, err := service.CreateInway(ctx, requestInway)
	if err != nil {
		t.Error("could not create inway", err)
	}

	assert.Equal(t, requestInway, responseInway)
}
