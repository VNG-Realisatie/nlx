//notlint:dupl // test function
package configapi_test

import (
	"context"
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
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
)

func TestGetInsight(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := configapi.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	emptyRequest := &configapi.Empty{}

	mockDatabase.EXPECT().GetInsightConfiguration(ctx)

	_, actualError := service.GetInsightConfiguration(ctx, emptyRequest)
	expectedError := status.Error(codes.NotFound, "insight configuration not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	mockInsightResponse := &database.InsightConfiguration{
		InsightAPIURL: "http://insight-api-url",
		IrmaServerURL: "http://irma-server-url",
	}

	mockDatabase.EXPECT().GetInsightConfiguration(ctx).Return(mockInsightResponse, nil)

	actualResponse, err := service.GetInsightConfiguration(ctx, emptyRequest)
	assert.Nil(t, err)

	expectedResponse := &configapi.InsightConfiguration{
		InsightAPIURL: "http://insight-api-url",
		IrmaServerURL: "http://irma-server-url",
	}

	assert.Equal(t, expectedResponse, actualResponse)
}

func TestPutInsight(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockInsightConfig := &database.InsightConfiguration{
		IrmaServerURL: "http://irma-url.com",
		InsightAPIURL: "http://insight-url.com",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().PutInsightConfiguration(ctx, mockInsightConfig)

	mockCtrlDirectoryRegistrationAPI := gomock.NewController(t)
	defer mockCtrlDirectoryRegistrationAPI.Finish()

	mockDirectoryRegistrationClient := mock_registrationapi.NewMockDirectoryRegistrationClient(mockCtrlDirectoryRegistrationAPI)
	mockDirectoryRegistrationClient.EXPECT().SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: "http://insight-url.com",
		IrmaServerURL: "http://irma-url.com",
	}).Return(&registrationapi.Empty{}, nil)

	service := configapi.New(logger, testProcess, mockDirectoryRegistrationClient, mockDatabase)

	request := &configapi.InsightConfiguration{
		IrmaServerURL: "http://irma-url.com",
		InsightAPIURL: "http://insight-url.com",
	}

	putInsightResponse, err := service.PutInsightConfiguration(ctx, request)
	assert.NoError(t, err)

	expectedResponse := &configapi.InsightConfiguration{
		InsightAPIURL: "http://insight-url.com",
		IrmaServerURL: "http://irma-url.com",
	}

	assert.Equal(t, expectedResponse, putInsightResponse)
}
