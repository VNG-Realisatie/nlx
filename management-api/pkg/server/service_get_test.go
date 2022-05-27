// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func TestGetService(t *testing.T) {
	tests := map[string]struct {
		ctx     context.Context
		setup   func(*testing.T, serviceMocks)
		req     *api.GetServiceRequest
		want    *api.GetServiceResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(t *testing.T, mocks serviceMocks) {},
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.service.read\" to execute this request").Err(),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(t *testing.T, mocks serviceMocks) {
				mocks.db.EXPECT().GetService(gomock.Any(), "my-service").Return(&database.Service{
					Name:   "my-service",
					Inways: []*database.Inway{},
				}, nil)
			},
			req: &api.GetServiceRequest{
				Name: "my-service",
			},
			want: &api.GetServiceResponse{
				Name:   "my-service",
				Inways: []string{},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, _, mocks := newService(t)
			tt.setup(t, mocks)

			want, err := service.GetService(tt.ctx, tt.req)
			assert.Equal(t, tt.want, want)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
