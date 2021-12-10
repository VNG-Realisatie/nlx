// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/directory"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	"go.nlx.io/nlx/management-api/pkg/server"
)

func TestManagementService_GetSettings(t *testing.T) {
	tests := map[string]struct {
		db               func(ctrl *gomock.Controller) database.ConfigDatabase
		expectedResponse *api.Settings
		expectedError    error
	}{
		"when_the_database_call_fails": {
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				db.EXPECT().GetSettings(gomock.Any()).Return(nil, errors.New("arbitrary error"))

				return db
			},

			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		"happy flow": {
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)
				db.EXPECT().GetSettings(gomock.Any()).Return(settings, nil)

				return db
			},

			expectedResponse: &api.Settings{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedError: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l := zap.NewNop()
			d := mock_directory.NewMockClient(ctrl)

			h := server.NewManagementService(l, d, nil, nil, tt.db(ctrl), nil, mock_auditlog.NewMockLogger(ctrl), management.NewClient)
			got, err := h.GetSettings(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

//nolint:funlen // alot of scenario's to test
func TestManagementService_UpdateSettings(t *testing.T) {
	tests := map[string]struct {
		db               func(ctrl *gomock.Controller) database.ConfigDatabase
		directoryClient  func(ctrl *gomock.Controller) directory.Client
		auditLog         func(ctrl *gomock.Controller) auditlog.Logger
		req              *api.UpdateSettingsRequest
		ctx              context.Context
		expectedResponse *emptypb.Empty
		expectedError    error
	}{
		"when_writing_audit_log_fails": {
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				return mock_database.NewMockConfigDatabase(ctrl)
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				return mock_directory.NewMockClient(ctrl)
			},
			auditLog: func(ctrl *gomock.Controller) auditlog.Logger {
				auditLogger := mock_auditlog.NewMockLogger(ctrl)
				auditLogger.EXPECT().OrganizationSettingsUpdate(gomock.Any(), "Jane Doe", "nlxctl").Return(fmt.Errorf("arbitrary error"))
				return auditLogger
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "could not create audit log"),
		},
		"when_update_settings_database_call_fails": {
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)

				db.EXPECT().UpdateSettings(
					gomock.Any(), settings,
				).Return(errors.New("arbitrary error"))

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)
				directoryClient.EXPECT().SetOrganizationContactDetails(gomock.Any(), &directoryapi.SetOrganizationContactDetailsRequest{
					EmailAddress: "mock@email.com",
				}).Return(&emptypb.Empty{}, nil)
				return directoryClient
			},
			auditLog: func(ctrl *gomock.Controller) auditlog.Logger {
				auditLogger := mock_auditlog.NewMockLogger(ctrl)
				auditLogger.EXPECT().OrganizationSettingsUpdate(gomock.Any(), "Jane Doe", "nlxctl")
				return auditLogger
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		"inway_does_not_exist": {
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)

				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)

				db.EXPECT().UpdateSettings(gomock.Any(), settings).Return(database.ErrInwayNotFound)

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)
				directoryClient.EXPECT().SetOrganizationContactDetails(gomock.Any(), &directoryapi.SetOrganizationContactDetailsRequest{
					EmailAddress: "mock@email.com",
				}).Return(&emptypb.Empty{}, nil)

				return directoryClient
			},
			auditLog: func(ctrl *gomock.Controller) auditlog.Logger {
				auditLogger := mock_auditlog.NewMockLogger(ctrl)
				auditLogger.EXPECT().OrganizationSettingsUpdate(gomock.Any(), "Jane Doe", "nlxctl")
				return auditLogger
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.InvalidArgument, "inway not found"),
		},
		"call_to_directory_fails": {
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)
				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)
				directoryClient.EXPECT().SetOrganizationContactDetails(gomock.Any(), &directoryapi.SetOrganizationContactDetailsRequest{
					EmailAddress: "mock@email.com",
				}).Return(nil, fmt.Errorf("arbitrary error"))

				return directoryClient
			},
			auditLog: func(ctrl *gomock.Controller) auditlog.Logger {
				auditLogger := mock_auditlog.NewMockLogger(ctrl)
				auditLogger.EXPECT().OrganizationSettingsUpdate(gomock.Any(), "Jane Doe", "nlxctl")
				return auditLogger
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "nlx directory unreachable"),
		},
		"happy_flow": {
			db: func(ctrl *gomock.Controller) database.ConfigDatabase {
				db := mock_database.NewMockConfigDatabase(ctrl)

				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)

				db.EXPECT().UpdateSettings(gomock.Any(), settings).Return(nil)

				return db
			},
			directoryClient: func(ctrl *gomock.Controller) directory.Client {
				directoryClient := mock_directory.NewMockClient(ctrl)
				directoryClient.EXPECT().SetOrganizationContactDetails(gomock.Any(), &directoryapi.SetOrganizationContactDetailsRequest{
					EmailAddress: "mock@email.com",
				}).Return(&emptypb.Empty{}, nil)

				return directoryClient
			},
			auditLog: func(ctrl *gomock.Controller) auditlog.Logger {
				auditLogger := mock_auditlog.NewMockLogger(ctrl)
				auditLogger.EXPECT().OrganizationSettingsUpdate(gomock.Any(), "Jane Doe", "nlxctl")
				return auditLogger
			},
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: &emptypb.Empty{},
			expectedError:    nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l := zap.NewNop()

			h := server.NewManagementService(l, tt.directoryClient(ctrl), nil, nil, tt.db(ctrl), nil, tt.auditLog(ctrl), management.NewClient)
			got, err := h.UpdateSettings(tt.ctx, tt.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
