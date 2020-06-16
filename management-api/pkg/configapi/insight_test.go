//notlint:dupl // test function
package configapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	mock_registrationapi "go.nlx.io/nlx/directory-registration-api/registrationapi/mock"
	"go.nlx.io/nlx/management-api/pkg/configapi"
	mock_configapi "go.nlx.io/nlx/management-api/pkg/configapi/mock"
)

func TestGetInsight(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configapi.NewMockConfigDatabase(mockCtrl)
	service := configapi.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

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

	mockDatabase := mock_configapi.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().PutInsightConfiguration(ctx, mockInsightConfig)

	mockCtrlDirectoryRegistrationAPI := gomock.NewController(t)
	defer mockCtrlDirectoryRegistrationAPI.Finish()

	mockDirectoryRegistrationClient := mock_registrationapi.NewMockDirectoryRegistrationClient(mockCtrlDirectoryRegistrationAPI)
	mockDirectoryRegistrationClient.EXPECT().SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: mockInsightConfig.InsightAPIURL,
		IrmaServerURL: mockInsightConfig.IrmaServerURL,
	}).Return(&registrationapi.Empty{}, nil)

	service := configapi.New(logger, testProcess, mockDirectoryRegistrationClient, mockDatabase)

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
