// nolint:dupl
package configservice_test

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/config-api/configservice"
	mock "go.nlx.io/nlx/config-api/configservice/mock"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.uber.org/zap"
)

func TestGetInsight(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock.NewMockConfigDatabase(mockCtrl)
	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	emptyRequest := &configapi.Empty{}

	mockDatabase.EXPECT().GetInsightConfiguration(ctx)

	_, actualError := service.GetInsightConfiguration(ctx, emptyRequest)
	expectedError := status.Error(codes.NotFound, "insight configuration not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	mockInsightResponse := &configapi.InsightConfiguration{
		InsightAPIURL: "http://insight-api-url",
		IrmaServerURL: "http://irma-server-url",
	}

	mockDatabase.EXPECT().GetInsightConfiguration(ctx).Return(mockInsightResponse, nil)

	getInsightConfigurationResponse, err := service.GetInsightConfiguration(ctx, emptyRequest)
	assert.Nil(t, err)

	assert.Equal(t, getInsightConfigurationResponse, getInsightConfigurationResponse)
}
