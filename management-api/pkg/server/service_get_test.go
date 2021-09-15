// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestGetService(t *testing.T) {
	logger := zap.NewNop()
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := server.NewManagementService(
		logger,
		mock_directory.NewMockClient(mockCtrl),
		nil,
		mockDatabase,
		nil,
		mock_auditlog.NewMockLogger(mockCtrl),
		management.NewClient,
	)

	getServiceRequest := &api.GetServiceRequest{
		Name: "my-service",
	}

	mockDatabase.EXPECT().GetService(ctx, "my-service").Return(nil, database.ErrNotFound)

	_, actualError := service.GetService(ctx, getServiceRequest)
	expectedError := status.Error(codes.NotFound, "service not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	databaseService := &database.Service{
		Name:   "my-service",
		Inways: []*database.Inway{},
	}

	mockDatabase.EXPECT().GetService(ctx, "my-service").Return(databaseService, nil)

	getServiceResponse, err := service.GetService(ctx, getServiceRequest)
	assert.NoError(t, err)

	expectedResponse := &api.GetServiceResponse{
		Name:   "my-service",
		Inways: []string{},
	}

	assert.Equal(t, expectedResponse, getServiceResponse)
}
