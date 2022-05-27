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

func TestUpdateService(t *testing.T) {
	tests := map[string]struct {
		ctx     context.Context
		setup   func(*testing.T, serviceMocks)
		request *api.UpdateServiceRequest
		want    *api.UpdateServiceResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:   testCreateUserWithoutPermissionsContext(),
			setup: func(t *testing.T, mocks serviceMocks) {},
			request: &api.UpdateServiceRequest{
				Name:        "my-service",
				EndpointURL: "my-service.test",
				Inways:      []string{},
			},
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.service.update\" to execute this request").Err(),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(t *testing.T, mocks serviceMocks) {
				mocks.al.EXPECT().ServiceUpdate(gomock.Any(), "admin@example.com", "nlxctl", "my-service")

				dbService := &database.Service{
					Name:        "my-service",
					EndpointURL: "my-service.test",
				}

				mocks.db.EXPECT().GetService(gomock.Any(), "my-service").Return(dbService, nil)

				mocks.db.EXPECT().UpdateServiceWithInways(gomock.Any(), dbService, []string{})
			},
			request: &api.UpdateServiceRequest{
				Name:        "my-service",
				EndpointURL: "my-service.test",
				Inways:      []string{},
			},
			want: &api.UpdateServiceResponse{
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

			want, err := service.UpdateService(tt.ctx, tt.request)
			assert.Equal(t, tt.want, want)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
