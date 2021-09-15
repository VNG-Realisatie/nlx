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
	"google.golang.org/grpc/metadata"

	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestUpdateService(t *testing.T) {
	logger := zap.NewNop()

	databaseService := &database.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		"username":               "Jane Doe",
		"grpcgateway-user-agent": "nlxctl",
	}))

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().GetService(ctx, "my-service").Return(databaseService, nil)
	mockDatabase.EXPECT().GetService(ctx, "other-service").Return(nil, database.ErrNotFound)
	mockDatabase.EXPECT().UpdateServiceWithInways(ctx, databaseService, []string{})

	auditLogger := mock_auditlog.NewMockLogger(mockCtrl)
	auditLogger.EXPECT().ServiceUpdate(gomock.Any(), "Jane Doe", "nlxctl", "my-service")

	service := server.NewManagementService(
		logger,
		mock_directory.NewMockClient(mockCtrl),
		nil,
		mockDatabase,
		nil,
		auditLogger,
		management.NewClient,
	)

	updateServiceRequest := &api.UpdateServiceRequest{
		Name:        "my-service",
		EndpointURL: "my-service.test",
		Inways:      []string{},
	}

	updateServiceResponse, err := service.UpdateService(ctx, updateServiceRequest)
	assert.NoError(t, err)

	expectedResponse := &api.UpdateServiceResponse{
		Name:        "my-service",
		EndpointURL: "my-service.test",
		Inways:      []string{},
	}

	assert.Equal(t, expectedResponse, updateServiceResponse)

	updateServiceRequest.Name = "other-service"

	_, err = service.UpdateService(ctx, updateServiceRequest)
	assert.EqualError(t, err, "rpc error: code = NotFound desc = service not found")
}
