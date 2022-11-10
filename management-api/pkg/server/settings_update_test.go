// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestManagementService_UpdateSettings(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, serviceMocks)
		ctx     context.Context
		req     *api.UpdateSettingsRequest
		want    *api.UpdateSettingsResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:   testCreateUserWithoutPermissionsContext(),
			setup: func(context.Context, serviceMocks) {},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			want:    nil,
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.organization_settings.update\" to execute this request").Err(),
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
			want:    nil,
			wantErr: status.Error(codes.Internal, "could not create audit log"),
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
					Return(&directoryapi.SetOrganizationContactDetailsResponse{}, nil)

				mocks.al.
					EXPECT().
					OrganizationSettingsUpdate(ctx, "admin@example.com", "nlxctl").
					Return(nil)

			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "database error"),
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
					Return(&directoryapi.SetOrganizationContactDetailsResponse{}, nil)

				mocks.al.
					EXPECT().
					OrganizationSettingsUpdate(ctx, "admin@example.com", "nlxctl").
					Return(nil)
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "inway not found"),
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
			want:    nil,
			wantErr: status.Error(codes.Internal, "nlx directory unreachable"),
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
					Return(&directoryapi.SetOrganizationContactDetailsResponse{}, nil)

				mocks.al.
					EXPECT().
					OrganizationSettingsUpdate(ctx, "admin@example.com", "nlxctl")
			},
			req: &api.UpdateSettingsRequest{
				OrganizationInway:        "inway-name",
				OrganizationEmailAddress: "mock@email.com",
			},
			want:    &api.UpdateSettingsResponse{},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t, nil)
			tt.setup(tt.ctx, mocks)

			got, err := service.UpdateSettings(tt.ctx, tt.req)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
