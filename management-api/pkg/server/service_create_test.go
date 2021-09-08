// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestCreateService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

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
	mockDatabase.EXPECT().CreateServiceWithInways(ctx, databaseService, []string{})

	auditLogger := mock_auditlog.NewMockLogger(mockCtrl)
	auditLogger.EXPECT().ServiceCreate(gomock.Any(), "Jane Doe", "nlxctl", "my-service")

	service := server.NewManagementService(
		logger,
		testProcess,
		mock_directory.NewMockClient(mockCtrl),
		nil,
		mockDatabase,
		nil,
		auditLogger,
		management.NewClient,
	)

	requestService := &api.CreateServiceRequest{
		Name:        "my-service",
		EndpointURL: "my-service.test",
		Inways:      []string{},
	}

	responseService, err := service.CreateService(ctx, requestService)
	if err != nil {
		t.Error("could not create service", err)
	}

	assert.Equal(t, &api.CreateServiceResponse{
		Name:        "my-service",
		EndpointURL: "my-service.test",
		Inways:      []string{},
	}, responseService)
}
