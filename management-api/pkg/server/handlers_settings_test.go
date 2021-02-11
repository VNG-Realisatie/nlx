// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/directory"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestManagementService_GetSettings(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.ConfigDatabase
		expectedResponse *api.Settings
		expectedError    error
	}{
		{
			name: "when the database call fails",
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().GetSettings(gomock.Any()).Return(nil, errors.New("arbitrary error"))

				return db
			},

			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		{
			name: "happy flow",
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().GetSettings(gomock.Any()).Return(&database.Settings{
					Inway: &database.Inway{
						ID:   1,
						Name: "inway-name",
					},
				}, nil)

				return db
			},

			expectedResponse: &api.Settings{
				OrganizationInway: "inway-name",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l := zap.NewNop()
			p := process.NewProcess(l)
			d := mock_directory.NewMockClient(ctrl)

			h := server.NewManagementService(l, p, d, nil, tt.db(ctrl), mock_auditlog.NewMockLogger(ctrl))
			got, err := h.GetSettings(context.Background(), &types.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

//nolint:funlen // alot of scenario's to test
func TestManagementService_UpdateSettings(t *testing.T) {
	tests := []struct {
		name             string
		db               func(ctrl *gomock.Controller) database.ConfigDatabase
		directoryClient  func(ctrl *gomock.Controller) directory.Client
		req              *api.UpdateSettingsRequest
		expectedResponse *types.Empty
		expectedError    error
	}{
		{
			name: "when_the_getinway_database_call_fails",
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)

				db.EXPECT().
					GetInway(gomock.Any(), "inway-name").
					Return(nil, errors.New("random error"))

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)

				return directoryClient
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway: "inway-name",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		{
			name: "when_the_directory_call_fails",
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)

				db.EXPECT().
					GetInway(gomock.Any(), "inway-name").
					Return(&database.Inway{
						SelfAddress: "inway.localhost",
					}, nil)

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)
				directoryClient.EXPECT().SetOrganizationInway(gomock.Any(), &registrationapi.SetOrganizationInwayRequest{
					Address: "inway.localhost",
				}).Return(nil, errors.New("arbitrary error"))

				return directoryClient
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway: "inway-name",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		{
			name: "when_the_inway_is_empty_and_the_directory_call_fails",
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)
				directoryClient.EXPECT().
					ClearOrganizationInway(gomock.Any(), &types.Empty{}).
					Return(nil, errors.New("arbitrary error"))

				return directoryClient
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway: "",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		{
			name: "when_the_updatesettings_database_call_fails",
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)

				db.EXPECT().
					GetInway(gomock.Any(), "inway-name").
					Return(&database.Inway{
						SelfAddress: "inway.localhost",
					}, nil)

				db.EXPECT().PutOrganizationInway(
					gomock.Any(), gomock.Any(),
				).Return(nil, errors.New("arbitrary error"))

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)

				directoryClient.EXPECT().
					SetOrganizationInway(gomock.Any(), &registrationapi.SetOrganizationInwayRequest{
						Address: "inway.localhost",
					}).
					Return(&types.Empty{}, nil)

				return directoryClient
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway: "inway-name",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		{
			name: "happy_flow",
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)

				db.EXPECT().
					GetInway(gomock.Any(), "inway-name").
					Return(&database.Inway{
						ID:          1,
						SelfAddress: "inway.localhost",
					}, nil)

				db.EXPECT().PutOrganizationInway(gomock.Any(), createUintPointer(1)).Return(&database.Settings{}, nil)

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)
				directoryClient.EXPECT().SetOrganizationInway(gomock.Any(), &registrationapi.SetOrganizationInwayRequest{
					Address: "inway.localhost",
				}).Return(&types.Empty{}, nil)

				return directoryClient
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway: "inway-name",
			},
			expectedResponse: &types.Empty{},
			expectedError:    nil,
		},
		{
			name: "happy flow with empty inway name",
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().PutOrganizationInway(gomock.Any(), nil).Return(&database.Settings{}, nil)

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)
				directoryClient.EXPECT().ClearOrganizationInway(gomock.Any(), gomock.Any()).Return(&types.Empty{}, nil)
				return directoryClient
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway: "",
			},
			expectedResponse: &types.Empty{},
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l := zap.NewNop()
			p := process.NewProcess(l)

			h := server.NewManagementService(l, p, tt.directoryClient(ctrl), nil, tt.db(ctrl), mock_auditlog.NewMockLogger(ctrl))
			got, err := h.UpdateSettings(context.Background(), tt.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func createUintPointer(x uint) *uint {
	return &x
}
