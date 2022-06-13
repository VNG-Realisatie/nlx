// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/environment"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func newDirectoryService(t *testing.T) (s *server.DirectoryService, m *mock_directory.MockClient, db *mock_database.MockConfigDatabase) {
	logger := zaptest.Logger(t)

	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	db = mock_database.NewMockConfigDatabase(ctrl)

	m = mock_directory.NewMockClient(ctrl)

	s = server.NewDirectoryService(logger, &environment.Environment{}, m, db)

	return
}

var directoryServiceStateTests = []struct {
	ExpectedState api.DirectoryService_State
	Inways        []*directoryapi.Inway
}{
	{
		api.DirectoryService_unknown,
		nil,
	},
	{
		api.DirectoryService_unknown,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_UNKNOWN},
		},
	},
	{
		api.DirectoryService_up,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_UP},
		},
	},
	{
		api.DirectoryService_up,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_UP},
			{State: directoryapi.Inway_UP},
			{State: directoryapi.Inway_UP},
		},
	},
	{
		api.DirectoryService_down,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_down,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_DOWN},
			{State: directoryapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_degraded,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_UP},
			{State: directoryapi.Inway_DOWN},
		},
	},
	{
		api.DirectoryService_degraded,
		[]*directoryapi.Inway{
			{State: directoryapi.Inway_DOWN},
			{State: directoryapi.Inway_UNKNOWN},
		},
	},
}
