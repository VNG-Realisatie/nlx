// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestGetStatisticsOfServices(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().GetIncomingAccessRequestCountByService(ctx).Return(map[string]int{
		"service-a": 3,
	}, nil)

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

	requestGetStatisticsOfServices := &api.GetStatisticsOfServicesRequest{}

	responseService, err := service.GetStatisticsOfServices(ctx, requestGetStatisticsOfServices)
	if err != nil {
		t.Error("could not get stats for services", err)
	}

	assert.Equal(t, &api.GetStatisticsOfServicesResponse{
		Services: []*api.ServiceStatistics{
			{
				Name:                       "service-a",
				IncomingAccessRequestCount: 3,
			},
		},
	}, responseService)

	mockDatabase.EXPECT().GetIncomingAccessRequestCountByService(ctx).Return(nil, errors.New("arbitrary error"))

	_, err = service.GetStatisticsOfServices(ctx, requestGetStatisticsOfServices)
	assert.Error(t, err)
}
