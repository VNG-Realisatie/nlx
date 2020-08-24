// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	context "context"
	"strconv"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/directory"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/environment"
	"go.nlx.io/nlx/management-api/pkg/server"
)

var directoryServiceStateTests = []struct {
	ExpectedState api.DirectoryService_State
	Inways        []*directory.Inway
}{
	{
		api.DirectoryService_unknown,
		nil,
	},
	{
		api.DirectoryService_unknown,
		[]*directory.Inway{
			{State: directory.InwayStateUnknown},
		},
	},
	{
		api.DirectoryService_up,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
		},
	},
	{
		api.DirectoryService_up,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
			{State: directory.InwayStateUp},
			{State: directory.InwayStateUp},
		},
	},
	{
		api.DirectoryService_down,
		[]*directory.Inway{
			{State: directory.InwayStateDown},
		},
	},
	{
		api.DirectoryService_down,
		[]*directory.Inway{
			{State: directory.InwayStateDown},
			{State: directory.InwayStateDown},
		},
	},
	{
		api.DirectoryService_degraded,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
			{State: directory.InwayStateDown},
		},
	},
	{
		api.DirectoryService_degraded,
		[]*directory.Inway{
			{State: directory.InwayStateDown},
			{State: directory.InwayStateUnknown},
		},
	},
}

func TestDirectoryServiceState(t *testing.T) {
	for i, test := range directoryServiceStateTests {
		name := strconv.Itoa(i + 1)
		test := test

		t.Run(name, func(t *testing.T) {
			state := server.DetermineDirectoryServiceState(test.Inways)
			assert.Equal(t, test.ExpectedState, state)
		})
	}
}

func TestListDirectoryServices(t *testing.T) {
	logger := zap.NewNop()
	env := &environment.Environment{}
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	clientServices := []*directory.InspectionAPIService{
		{
			Name:                 "test-service-1",
			OrganizationName:     "test-organization-a",
			APISpecificationType: "OpenAPI3",
		},
		{
			Name:                 "test-service-2",
			OrganizationName:     "test-organization-a",
			APISpecificationType: "OpenAPI3",
		},
		{
			Name:                 "test-service-3",
			OrganizationName:     "test-organization-b",
			APISpecificationType: "",
		},
	}

	databaseAccessRequests := map[string]*database.AccessRequest{
		"test-organization-b/test-service-3": {
			ID:               "161c188cfcea1939",
			OrganizationName: "test-organization-b",
			ServiceName:      "test-service-3",
			State:            database.AccessRequestCreated,
			CreatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
			UpdatedAt:        time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC),
		},
		"test-organization-a/test-service-1": {
			ID:               "161c1bd32da2b400",
			OrganizationName: "test-organization-a",
			ServiceName:      "test-service-1",
			State:            database.AccessRequestCreated,
			CreatedAt:        time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
			UpdatedAt:        time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC),
		},
	}

	client := mock_directory.NewMockClient(mockCtrl)
	client.EXPECT().ListServices().Return(clientServices, nil)

	db := mock_database.NewMockConfigDatabase(mockCtrl)
	db.EXPECT().ListAllLatestOutgoingAccessRequests(ctx).Return(databaseAccessRequests, nil)

	service := server.NewDirectoryService(logger, env, client, db)
	response, err := service.ListServices(ctx, &api.Empty{})
	assert.NoError(t, err)

	expected := []*api.DirectoryService{
		{
			ServiceName:          "test-service-1",
			OrganizationName:     "test-organization-a",
			APISpecificationType: "OpenAPI3",
			State:                api.DirectoryService_unknown,
			LatestAccessRequest: &api.AccessRequest{
				Id:               "161c1bd32da2b400",
				OrganizationName: "test-organization-a",
				ServiceName:      "test-service-1",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        timestampProto(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
				UpdatedAt:        timestampProto(time.Date(2020, time.June, 26, 13, 42, 42, 0, time.UTC)),
			},
		},
		{
			ServiceName:          "test-service-2",
			OrganizationName:     "test-organization-a",
			APISpecificationType: "OpenAPI3",
			State:                api.DirectoryService_unknown,
		},
		{
			ServiceName:          "test-service-3",
			OrganizationName:     "test-organization-b",
			APISpecificationType: "",
			State:                api.DirectoryService_unknown,
			LatestAccessRequest: &api.AccessRequest{
				Id:               "161c188cfcea1939",
				OrganizationName: "test-organization-b",
				ServiceName:      "test-service-3",
				State:            api.AccessRequestState_CREATED,
				CreatedAt:        timestampProto(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
				UpdatedAt:        timestampProto(time.Date(2020, time.June, 26, 12, 42, 42, 1337, time.UTC)),
			},
		},
	}

	assert.Equal(t, expected, response.Services)
}

func timestampProto(t time.Time) *types.Timestamp {
	tp, _ := types.TimestampProto(t)
	return tp
}
