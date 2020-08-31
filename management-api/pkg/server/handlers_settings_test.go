// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/directory"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func generateMockConfigDatabase(t *testing.T) *mock_database.MockConfigDatabase {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return mock_database.NewMockConfigDatabase(mockCtrl)
}

func generateMockDirectoryClient(t *testing.T) directory.Client {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return mock_directory.NewMockClient(mockCtrl)
}

func TestManagementService_GetSettings(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		configDatabase  database.ConfigDatabase
		mainProcess     *process.Process
		directoryClient directory.Client
	}

	type args struct {
		ctx context.Context
		req *types.Empty
	}

	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse *api.Settings
		expectedError    error
	}{
		{
			name: "when the database call fails",
			fields: func() fields {
				logger := zap.NewNop()

				configDatabase := generateMockConfigDatabase(t)
				configDatabase.EXPECT().GetSettings(gomock.Any()).Return(nil, errors.New("arbitrary error")).AnyTimes()

				return fields{
					logger:          logger,
					mainProcess:     process.NewProcess(logger),
					configDatabase:  configDatabase,
					directoryClient: generateMockDirectoryClient(t),
				}
			}(),
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		{
			name: "happy flow",
			fields: func() fields {
				logger := zap.NewNop()

				configDatabase := generateMockConfigDatabase(t)
				configDatabase.EXPECT().GetSettings(gomock.Any()).Return(&database.Settings{
					OrganizationInway: "inway-name",
				}, nil).AnyTimes()

				return fields{
					logger:          logger,
					mainProcess:     process.NewProcess(logger),
					configDatabase:  configDatabase,
					directoryClient: generateMockDirectoryClient(t),
				}
			}(),
			expectedResponse: &api.Settings{
				OrganizationInway: "inway-name",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := server.NewManagementService(tt.fields.logger, tt.fields.mainProcess, tt.fields.directoryClient, tt.fields.configDatabase)
			got, err := h.GetSettings(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestManagementService_UpdateSettings(t *testing.T) {
	type fields struct {
		logger          *zap.Logger
		configDatabase  database.ConfigDatabase
		mainProcess     *process.Process
		directoryClient directory.Client
	}

	type args struct {
		ctx context.Context
		req *api.UpdateSettingsRequest
	}

	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse *types.Empty
		expectedError    error
	}{
		{
			name: "when the database call fails",
			fields: func() fields {
				logger := zap.NewNop()

				configDatabase := generateMockConfigDatabase(t)
				configDatabase.EXPECT().UpdateSettings(
					gomock.Any(), gomock.Any(),
				).Return(errors.New("arbitrary error")).AnyTimes()

				return fields{
					logger:          logger,
					mainProcess:     process.NewProcess(logger),
					configDatabase:  configDatabase,
					directoryClient: generateMockDirectoryClient(t),
				}
			}(),
			args: args{
				ctx: context.Background(),
				req: &api.UpdateSettingsRequest{
					OrganizationInway: "inway-name",
				},
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		{
			name: "happy flow",
			fields: func() fields {
				logger := zap.NewNop()

				configDatabase := generateMockConfigDatabase(t)
				configDatabase.EXPECT().UpdateSettings(gomock.Any(), &database.Settings{
					OrganizationInway: "inway-name",
				}).Return(nil).AnyTimes()

				return fields{
					logger:          logger,
					mainProcess:     process.NewProcess(logger),
					configDatabase:  configDatabase,
					directoryClient: generateMockDirectoryClient(t),
				}
			}(),
			args: args{
				ctx: context.Background(),
				req: &api.UpdateSettingsRequest{
					OrganizationInway: "inway-name",
				},
			},
			expectedResponse: &types.Empty{},
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := server.NewManagementService(tt.fields.logger, tt.fields.mainProcess, tt.fields.directoryClient, tt.fields.configDatabase)
			got, err := h.UpdateSettings(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
