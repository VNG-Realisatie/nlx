// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
)

//nolint:funlen // this is a test function
func TestDeleteService(t *testing.T) {
	tests := map[string]struct {
		request       *api.DeleteServiceRequest
		ctx           context.Context
		setup         func(*common_tls.CertificateBundle, serviceMocks)
		expectedError error
	}{
		"missing_required_permission": {
			ctx:           testCreateUserWithoutPermissionsContext(),
			setup:         func(_ *common_tls.CertificateBundle, mocks serviceMocks) {},
			expectedError: status.New(codes.PermissionDenied, "user needs the permission \"permissions.service.delete\" to execute this request").Err(),
		},
		"failed_to_create_audit_log": {
			ctx: testCreateAdminUserContext(),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.al.EXPECT().
					ServiceDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-service").
					Return(fmt.Errorf("error"))
			},
			request: &api.DeleteServiceRequest{
				Name: "my-service",
			},
			expectedError: status.Error(codes.Internal, "could not create audit log"),
		},
		"failed_to_delete_service_from_database": {
			ctx: testCreateAdminUserContext(),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.EXPECT().
					DeleteService(gomock.Any(), "my-service", "00000000000000000001").
					Return(fmt.Errorf("error"))

				mocks.al.EXPECT().
					ServiceDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-service")
			},
			request: &api.DeleteServiceRequest{
				Name: "my-service",
			},
			expectedError: status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			ctx: testCreateAdminUserContext(),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					DeleteService(gomock.Any(), "my-service", "00000000000000000001")

				mocks.al.
					EXPECT().
					ServiceDelete(gomock.Any(), "admin@example.com", "nlxctl", "my-service")
			},
			request: &api.DeleteServiceRequest{
				Name: "my-service",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, bundle, mocks := newService(t, nil)

			if tt.setup != nil {
				tt.setup(bundle, mocks)
			}

			_, err := service.DeleteService(tt.ctx, tt.request)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
