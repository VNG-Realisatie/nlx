// nolint:dupl
package configservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/config-api/configservice"
	mock_configservice "go.nlx.io/nlx/config-api/configservice/mock"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	mock_registrationapi "go.nlx.io/nlx/directory-registration-api/registrationapi/mock"
)

func TestPutInsight(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockInsightConfig := &configapi.InsightConfiguration{
		IrmaServerURL: "http://irma-url.com",
		InsightAPIURL: "http://insight-url.com",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().PutInsightConfiguration(ctx, mockInsightConfig)

	mockCtrlDirectoryRegistrationAPI := gomock.NewController(t)
	defer mockCtrlDirectoryRegistrationAPI.Finish()

	mockDirectoryRegistrationClient := mock_registrationapi.NewMockDirectoryRegistrationClient(mockCtrlDirectoryRegistrationAPI)
	mockDirectoryRegistrationClient.EXPECT().SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: mockInsightConfig.InsightAPIURL,
		IrmaServerURL: mockInsightConfig.IrmaServerURL,
	}).Return(&registrationapi.Empty{}, nil)
	service := configservice.New(logger, testProcess, mockDirectoryRegistrationClient, mockDatabase)

	putInsightResponse, err := service.PutInsightConfiguration(ctx, mockInsightConfig)
	if err != nil {
		t.Error("could not update inway", err)
	}

	assert.Equal(t, mockInsightConfig, putInsightResponse)

	mockError := fmt.Errorf("mock error")
	mockDirectoryRegistrationClient.EXPECT().SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: mockInsightConfig.InsightAPIURL,
		IrmaServerURL: mockInsightConfig.IrmaServerURL,
	}).Return(nil, mockError)

	putInsightResponse, err = service.PutInsightConfiguration(ctx, mockInsightConfig)
	assert.NotNil(t, err)
	assert.Nil(t, putInsightResponse)
	assert.Equal(t, mockError, err)
}
