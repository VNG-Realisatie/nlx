// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test function
func TestAcceptTermsOfServiceStatus(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, serviceMocks)
		ctx     context.Context
		want    *api.AcceptTermsOfServiceResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(ctx context.Context, mocks serviceMocks) {},
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.terms_of_service.accept\" to execute this request").Err(),
		},
		"when_accepting_terms_of_service_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					AcceptTermsOfService(gomock.Any(), "admin@example.com", gomock.Any()).
					Return(false, errors.New("arbitrary error"))
			},
			ctx:     testCreateAdminUserContext(),
			want:    nil,
			wantErr: status.Errorf(codes.Internal, "database error"),
		},
		"when_writing_audit_logs_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					AcceptTermsOfService(gomock.Any(), "admin@example.com", gomock.Any()).
					Return(false, nil)

				mocks.al.
					EXPECT().
					AcceptTermsOfService(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
					).
					Return(errors.New("arbitrary error"))
			},
			ctx:     testCreateAdminUserContext(),
			want:    nil,
			wantErr: status.Error(codes.Internal, "could not create audit log"),
		},
		"happy_flow": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.al.
					EXPECT().
					AcceptTermsOfService(
						gomock.Any(),
						"admin@example.com",
						"nlxctl",
					)

				mocks.db.
					EXPECT().
					AcceptTermsOfService(gomock.Any(), "admin@example.com", gomock.Any()).
					Return(false, nil)
			},
			ctx:     testCreateAdminUserContext(),
			want:    &api.AcceptTermsOfServiceResponse{},
			wantErr: nil,
		},
		"happy_flow_already_accepted": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					AcceptTermsOfService(gomock.Any(), "admin@example.com", gomock.Any()).
					Return(true, nil)
			},
			ctx:     testCreateAdminUserContext(),
			want:    &api.AcceptTermsOfServiceResponse{},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setup(tt.ctx, mocks)

			got, err := service.AcceptTermsOfService(tt.ctx, &api.AcceptTermsOfServiceRequest{})

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGetTermsOfServiceStatus(t *testing.T) {
	tests := map[string]struct {
		setup   func(context.Context, serviceMocks)
		ctx     context.Context
		want    *api.GetTermsOfServiceStatusResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(ctx context.Context, mocks serviceMocks) {},
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.terms_of_service_status.read\" to execute this request").Err(),
		},
		"when_getting_terms_of_service_status_fails": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetTermsOfServiceStatus(gomock.Any()).
					Return(nil, errors.New("arbitrary error"))
			},
			ctx:     testCreateAdminUserContext(),
			want:    nil,
			wantErr: status.Errorf(codes.Internal, "database error"),
		},
		"happy_flow_not_accepted": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					GetTermsOfServiceStatus(gomock.Any()).
					Return(nil, database.ErrNotFound)
			},
			ctx: testCreateAdminUserContext(),
			want: &api.GetTermsOfServiceStatusResponse{
				Accepted: false,
			},
			wantErr: nil,
		},
		"happy_flow_accepted": {
			setup: func(ctx context.Context, mocks serviceMocks) {
				model, err := domain.NewTermsOfServiceStatus(&domain.NewTermsOfServiceStatusArgs{
					Username:  "admin@example.com",
					CreatedAt: time.Now(),
				})
				require.NoError(t, err)

				mocks.db.
					EXPECT().
					GetTermsOfServiceStatus(gomock.Any()).
					Return(model, nil)
			},
			ctx: testCreateAdminUserContext(),
			want: &api.GetTermsOfServiceStatusResponse{
				Accepted: true,
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, _, mocks := newService(t)
			tt.setup(context.Background(), mocks)

			got, err := service.GetTermsOfServiceStatus(tt.ctx, &api.GetTermsOfServiceStatusRequest{})

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
