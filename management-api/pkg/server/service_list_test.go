// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // alot of scenario's to test
func TestListServices(t *testing.T) {
	t.Helper()

	databaseServices := []*database.Service{
		{
			Name:   "my-service",
			Inways: []*database.Inway{{Name: "inway.mock"}},
		},
		{
			Name:   "another-service",
			Inways: []*database.Inway{{Name: "another-inway.mock"}},
		},
		{
			Name: "third-service",
		},
	}

	tests := map[string]struct {
		ctx              context.Context
		setup            func(*common_tls.CertificateBundle, serviceMocks)
		expectedResponse *api.ListServicesResponse
		expectedError    error
	}{
		"missing_required_permission": {
			ctx:              testCreateUserWithoutPermissionsContext(),
			setup:            func(_ *common_tls.CertificateBundle, mocks serviceMocks) {},
			expectedResponse: nil,
			expectedError:    status.New(codes.PermissionDenied, "user needs the permission \"permissions.services.read\" to execute this request").Err(),
		},
		"happy_flow": {
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.EXPECT().ListServices(gomock.Any()).Return(databaseServices, nil)

				mocks.db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{}, nil)

				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "my-service").Return([]*database.AccessGrant{}, nil)
				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "another-service").Return([]*database.AccessGrant{}, nil)
				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "third-service").Return([]*database.AccessGrant{}, nil)
			},
			ctx: testCreateAdminUserContext(),
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.ListServicesResponse_Service{
					{
						Name:   "my-service",
						Inways: []string{"inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
					},
					{
						Name:   "another-service",
						Inways: []string{"another-inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
					},
					{
						Name:   "third-service",
						Inways: []string{},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
					},
				},
			},
			expectedError: nil,
		},
		"happy flow with incoming access requests": {
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.EXPECT().ListServices(gomock.Any()).Return(databaseServices, nil)

				mocks.db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{
					"my-service":      2,
					"another-service": 0,
					"third-service":   0,
				}, nil)

				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "my-service").Return([]*database.AccessGrant{}, nil)
				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "another-service").Return([]*database.AccessGrant{}, nil)
				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "third-service").Return([]*database.AccessGrant{}, nil)
			},
			ctx: testCreateAdminUserContext(),
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.ListServicesResponse_Service{
					{
						Name:   "my-service",
						Inways: []string{"inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
						IncomingAccessRequestCount: 2,
					},
					{
						Name:   "another-service",
						Inways: []string{"another-inway.mock"},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
						IncomingAccessRequestCount: 0,
					},
					{
						Name:   "third-service",
						Inways: []string{},
						AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
							Mode:           "whitelist",
							Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{},
						},
						IncomingAccessRequestCount: 0,
					},
				},
			},
			expectedError: nil,
		},
		"when database call for service fails": {
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.EXPECT().ListServices(gomock.Any()).Return(nil, errors.New("arbitrary error"))
			},
			ctx:              testCreateAdminUserContext(),
			expectedResponse: nil,
			expectedError:    status.Error(codes.Internal, "database error"),
		},
		"when database for access grants fails": {
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.EXPECT().ListServices(gomock.Any()).Return(databaseServices, nil)

				mocks.db.EXPECT().GetIncomingAccessRequestCountByService(gomock.Any()).Return(map[string]int{}, nil)

				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "my-service").Return(nil, errors.New("arbitrary error"))
				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "another-service").Return(nil, errors.New("arbitrary error"))
				mocks.db.EXPECT().ListAccessGrantsForService(gomock.Any(), "third-service").Return(nil, errors.New("arbitrary error"))
			},
			ctx: testCreateAdminUserContext(),
			expectedResponse: &api.ListServicesResponse{
				Services: []*api.ListServicesResponse_Service{},
			},
			expectedError: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, bundle, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(bundle, mocks)
			}

			response, err := service.ListServices(tt.ctx, nil)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedResponse, response)
		})
	}
}
