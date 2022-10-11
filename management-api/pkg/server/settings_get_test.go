// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestManagementService_GetSettings(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, serviceMocks)
		ctx     context.Context
		want    *api.GetSettingsResponse
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
		"when_there_is_no_settings_row_present": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetSettings(gomock.Any()).
					Return(nil, database.ErrNotFound)
			},
			ctx: testCreateAdminUserContext(),
			want: &api.GetSettingsResponse{
				Settings: &api.Settings{
					OrganizationInway:        "",
					OrganizationEmailAddress: "",
				},
			},
			wantErr: nil,
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				settings, err := domain.NewSettings("inway-name", "mock@email.com")
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					GetSettings(gomock.Any()).
					Return(settings, nil)
			},
			ctx: testCreateAdminUserContext(),
			want: &api.GetSettingsResponse{
				Settings: &api.Settings{
					OrganizationInway:        "inway-name",
					OrganizationEmailAddress: "mock@email.com",
				},
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setup(tt.ctx, mocks)

			got, err := service.GetSettings(tt.ctx, &api.GetSettingsRequest{})

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
