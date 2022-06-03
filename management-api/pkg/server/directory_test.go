// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

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

func TestGetTermsOfService(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		directoryClient func(directoryClient *mock_directory.MockClient)
		wantResult      *api.GetTermsOfServiceResponse
		wantErr         error
	}{
		"failed_to_fetch_from_directory_client": {
			directoryClient: func(directoryClient *mock_directory.MockClient) {
				directoryClient.
					EXPECT().
					GetTermsOfService(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			wantResult: nil,
			wantErr:    status.Error(codes.Internal, "unable to get terms of service from directory"),
		},
		"happy_flow": {
			directoryClient: func(directoryClient *mock_directory.MockClient) {
				directoryClient.EXPECT().GetTermsOfService(gomock.Any(), gomock.Any()).Return(&directoryapi.GetTermsOfServiceResponse{
					Enabled: true,
					Url:     "https://example.com",
				}, nil)
			},
			wantResult: &api.GetTermsOfServiceResponse{
				Enabled: true,
				Url:     "https://example.com",
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, mockDirectoryClient, _ := newDirectoryService(t)

			tt.directoryClient(mockDirectoryClient)

			result, err := service.GetTermsOfService(ctx, &emptypb.Empty{})

			assert.Equal(t, tt.wantResult, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
