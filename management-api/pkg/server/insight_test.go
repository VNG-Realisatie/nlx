// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//notlint:dupl // test function
package server_test

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestGetInsight(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase)

	emptyRequest := &types.Empty{}

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

	expectedResponse := &api.InsightConfiguration{
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

	mockDirectoryClient := mock_directory.NewMockClient(mockCtrl)
	mockDirectoryClient.EXPECT().SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: "http://insight-url.com",
		IrmaServerURL: "http://irma-url.com",
	}).Return(&types.Empty{}, nil)

	service := server.NewManagementService(logger, testProcess, mockDirectoryClient, nil, mockDatabase)

	request := &api.InsightConfiguration{
		IrmaServerURL: "http://irma-url.com",
		InsightAPIURL: "http://insight-url.com",
	}

	putInsightResponse, err := service.PutInsightConfiguration(ctx, request)
	assert.NoError(t, err)

	expectedResponse := &api.InsightConfiguration{
		InsightAPIURL: "http://insight-url.com",
		IrmaServerURL: "http://irma-url.com",
	}

	assert.Equal(t, expectedResponse, putInsightResponse)
}
