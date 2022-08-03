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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestManagementService_GetSettings(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, serviceMocks)
		ctx     context.Context
		want    *api.Settings
		wantErr error
	}{
		"missing_required_permission": {
			setup:   func(ctx context.Context, mocks serviceMocks) {},
			ctx:     testCreateUserWithoutPermissionsContext(),
			want:    nil,
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.organization_settings.read\" to execute this request").Err(),
		},
		"when_the_database_call_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetSettings(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))

			},
			ctx:     testCreateAdminUserContext(),
			want:    nil,
			wantErr: status.Error(codes.Internal, "database error"),
		},
		"happy flow": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					GetSettings(gomock.Any()).
					Return(settings, nil)
			},
			ctx: testCreateAdminUserContext(),
			want: &api.Settings{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setup(tt.ctx, mocks)

			got, err := service.GetSettings(tt.ctx, &emptypb.Empty{})

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

//nolint:funlen // alot of scenario's to test
func TestManagementService_UpdateSettings(t *testing.T) {
	tests := map[string]struct {
		setup            func(context.Context, serviceMocks)
		ctx              context.Context
		req              *api.UpdateSettingsRequest
		expectedResponse *emptypb.Empty
		expectedError    error
	}{
		"missing_required_permission": {
			ctx:   testCreateUserWithoutPermissionsContext(),
			setup: func(context.Context, serviceMocks) {},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.New(codes.PermissionDenied, "user needs the permission \"permissions.organization_settings.update\" to execute this request").Err(),
		},
		"when_writing_audit_log_fails": {
			ctx: testCreateAdminUserContext(),
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					OrganizationSettingsUpdate(ctx, "admin@example.com", "nlxctl").
					Return(fmt.Errorf("arbitrary error"))
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "could not create audit log"),
		},
		"when_update_settings_database_call_fails": {
			ctx: testCreateAdminUserContext(),
			setup: func(ctx context.Context, mocks serviceMocks) {
				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					UpdateSettings(ctx, settings).
					Return(errors.New("arbitrary error"))

				mocks.dc.
					EXPECT().
					SetOrganizationContactDetails(ctx, &directoryapi.SetOrganizationContactDetailsRequest{
						EmailAddress: "mock@email.com",
					}).
					Return(&emptypb.Empty{}, nil)

				mocks.al.
					EXPECT().
					OrganizationSettingsUpdate(ctx, "admin@example.com", "nlxctl").
					Return(nil)

			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		"inway_does_not_exist": {
			ctx: testCreateAdminUserContext(),
			setup: func(ctx context.Context, mocks serviceMocks) {
				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					UpdateSettings(ctx, settings).
					Return(database.ErrInwayNotFound)

				mocks.dc.
					EXPECT().
					SetOrganizationContactDetails(ctx, &directoryapi.SetOrganizationContactDetailsRequest{
						EmailAddress: "mock@email.com",
					}).
					Return(&emptypb.Empty{}, nil)

				mocks.al.
					EXPECT().
					OrganizationSettingsUpdate(ctx, "admin@example.com", "nlxctl").
					Return(nil)
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.InvalidArgument, "inway not found"),
		},
		"call_to_directory_fails": {
			ctx: testCreateAdminUserContext(),
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.dc.
					EXPECT().
					SetOrganizationContactDetails(ctx, &directoryapi.SetOrganizationContactDetailsRequest{
						EmailAddress: "mock@email.com",
					}).
					Return(nil, fmt.Errorf("arbitrary error"))

				mocks.al.
					EXPECT().
					OrganizationSettingsUpdate(ctx, "admin@example.com", "nlxctl").
					Return(nil)
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "nlx directory unreachable"),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(ctx context.Context, mocks serviceMocks) {
				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					UpdateSettings(ctx, settings).
					Return(nil)

				mocks.dc.
					EXPECT().
					SetOrganizationContactDetails(ctx, &directoryapi.SetOrganizationContactDetailsRequest{
						EmailAddress: "mock@email.com",
					}).
					Return(&emptypb.Empty{}, nil)

				mocks.al.
					EXPECT().
					OrganizationSettingsUpdate(ctx, "admin@example.com", "nlxctl")
			},
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
			service, _, mocks := newService(t)
			tt.setup(tt.ctx, mocks)

			got, err := service.UpdateSettings(tt.ctx, tt.req)

			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
