// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//notlint:dupl // test function
package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestGetInsight(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := server.NewManagementService(
		logger,
		testProcess,
		mock_directory.NewMockClient(mockCtrl),
		nil,
		mockDatabase,
		nil,
		mock_auditlog.NewMockLogger(mockCtrl),
		management.NewClient,
	)

	emptyRequest := &emptypb.Empty{}

	mockDatabase.EXPECT().GetSettings(ctx).Return(nil, database.ErrNotFound)

	_, actualError := service.GetInsightConfiguration(ctx, emptyRequest)
	expectedError := status.Error(codes.NotFound, "insight configuration not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	mockInsightResponse := &database.Settings{
		InsightAPIURL: "http://insight-api-url",
		IrmaServerURL: "http://irma-server-url",
	}

	mockDatabase.EXPECT().GetSettings(ctx).Return(mockInsightResponse, nil)

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
	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		"username":               "Jane Doe",
		"grpcgateway-user-agent": "nlxctl",
	}))

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().PutInsightConfiguration(ctx, "http://irma-url.com", "http://insight-url.com")

	mockDirectoryClient := mock_directory.NewMockClient(mockCtrl)
	mockDirectoryClient.EXPECT().SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: "http://insight-url.com",
		IrmaServerURL: "http://irma-url.com",
	}).Return(&emptypb.Empty{}, nil)

	auditLogger := mock_auditlog.NewMockLogger(mockCtrl)
	auditLogger.EXPECT().OrganizationInsightConfigurationUpdate(ctx, "Jane Doe", "nlxctl")

	service := server.NewManagementService(
		logger,
		testProcess,
		mockDirectoryClient,
		nil,
		mockDatabase,
		nil,
		auditLogger,
		management.NewClient,
	)

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
