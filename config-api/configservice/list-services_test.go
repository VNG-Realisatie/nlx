// nolint:dupl
package configservice_test

import (
	"context"
	"testing"

	"go.nlx.io/nlx/config-api/configapi"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/config-api/configservice"
	mock "go.nlx.io/nlx/config-api/configservice/mock"

	"go.nlx.io/nlx/common/process"
	"go.uber.org/zap"
)

func TestListServices(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock.NewMockConfigDatabase(mockCtrl)
	service := configservice.New(logger, testProcess, mockDatabase)

	mockListServices := []*configapi.Service{
		{Name: "my-service"},
		{Name: "another-service"},
	}

	mockDatabase.EXPECT().ListServices(ctx).Return(mockListServices, nil)

	actualResponse, err := service.ListServices(ctx, &configapi.ListServicesRequest{})
	if err != nil {
		t.Fatal("could not get list of services", err)
	}

	expectedResponse := &configapi.ListServicesResponse{
		Services: mockListServices,
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
