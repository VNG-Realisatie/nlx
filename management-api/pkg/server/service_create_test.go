// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

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

func TestCreateService(t *testing.T) {
	tests := map[string]struct {
		ctx     context.Context
		setup   func(*testing.T, serviceMocks)
		request *api.CreateServiceRequest
		want    *api.CreateServiceResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(t *testing.T, mocks serviceMocks) {},
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.service.create\" to execute this request").Err(),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(t *testing.T, mocks serviceMocks) {
				mocks.al.EXPECT().ServiceCreate(gomock.Any(), "admin@example.com", "nlxctl", "my-service")

				mocks.db.EXPECT().CreateServiceWithInways(gomock.Any(), &database.Service{
					Name:        "my-service",
					EndpointURL: "my-service.test",
				}, []string{})
			},
			request: &api.CreateServiceRequest{
				Name:        "my-service",
				EndpointURL: "my-service.test",
				Inways:      []string{},
			},
			want: &api.CreateServiceResponse{
				Name:        "my-service",
				EndpointURL: "my-service.test",
				Inways:      []string{},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, _, mocks := newService(t)
			tt.setup(t, mocks)

			want, err := service.CreateService(tt.ctx, tt.request)
			assert.Equal(t, tt.want, want)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
