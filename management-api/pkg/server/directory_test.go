// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zaptest"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/environment"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func newDirectoryService(t *testing.T) (*server.DirectoryService, directoryServiceMocks) {
	logger := zaptest.NewLogger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	mocks := directoryServiceMocks{
		d:  mock_directory.NewMockClient(ctrl),
		db: mock_database.NewMockConfigDatabase(ctrl),
	}

	s := server.NewDirectoryService(
		logger,
		&environment.Environment{},
		mocks.d,
		mocks.db,
	)

	return s, mocks
}

type directoryServiceMocks struct {
	d  *mock_directory.MockClient
	db *mock_database.MockConfigDatabase
}

var directoryServiceStateTests = []struct {
	ExpectedState api.DirectoryService_State
	Inways        []*directoryapi.Inway
}{
	{
		api.DirectoryService_STATE_UNSPECIFIED,
		nil,
	},
	{
		api.DirectoryService_STATE_UNSPECIFIED,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_STATE_UNSPECIFIED},
		},
	},
	{
		api.DirectoryService_STATE_UP,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_STATE_UP},
		},
	},
	{
		api.DirectoryService_STATE_UP,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_STATE_UP},
			{State: directoryapi.Inway_STATE_UP},
			{State: directoryapi.Inway_STATE_UP},
		},
	},
	{
		api.DirectoryService_STATE_DOWN,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_STATE_DOWN},
		},
	},
	{
		api.DirectoryService_STATE_DOWN,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_STATE_DOWN},
			{State: directoryapi.Inway_STATE_DOWN},
		},
	},
	{
		api.DirectoryService_STATE_DEGRADED,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_STATE_UP},
			{State: directoryapi.Inway_STATE_DOWN},
		},
	},
	{
		api.DirectoryService_STATE_DEGRADED,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_STATE_DOWN},
			{State: directoryapi.Inway_STATE_UNSPECIFIED},
		},
	},
}
