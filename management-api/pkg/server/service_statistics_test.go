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
)

func TestGetStatisticsOfServices(t *testing.T) {
	tests := map[string]struct {
		ctx     context.Context
		setup   func(*testing.T, serviceMocks)
		want    *api.GetStatisticsOfServicesResponse
		wantErr error
	}{
		"missing_required_permission": {
			ctx:     testCreateUserWithoutPermissionsContext(),
			setup:   func(t *testing.T, mocks serviceMocks) {},
			wantErr: status.New(codes.PermissionDenied, "user needs the permission \"permissions.services_statistics.read\" to execute this request").Err(),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(t *testing.T, mocks serviceMocks) {
				mocks.db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{
					"service-a": 3,
				}, nil)
			},

			want: &api.GetStatisticsOfServicesResponse{
				Services: []*api.ServiceStatistics{
					{
						Name:                       "service-a",
						IncomingAccessRequestCount: 3,
					},
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, _, mocks := newService(t, nil)
			tt.setup(t, mocks)

			want, err := service.GetStatisticsOfServices(tt.ctx, nil)
			assert.Equal(t, tt.want, want)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
