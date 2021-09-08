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
	"google.golang.org/grpc/metadata"
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
		"failed_to_retrieve_user_info_from_context": {
			ctx: context.Background(),
			request: &api.DeleteServiceRequest{
				Name: "my-service",
			},
			expectedError: status.Error(codes.Internal, "could not retrieve user info to create audit log"),
		},
		"failed_to_create_audit_log": {
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.al.EXPECT().
					ServiceDelete(gomock.Any(), "Jane Doe", "nlxctl", "my-service").
					Return(fmt.Errorf("error"))
			},
			request: &api.DeleteServiceRequest{
				Name: "my-service",
			},
			expectedError: status.Error(codes.Internal, "could not create audit log"),
		},
		"failed_to_delete_service_from_database": {
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.EXPECT().
					DeleteService(gomock.Any(), "my-service", "nlx-test").
					Return(fmt.Errorf("error"))

				mocks.al.EXPECT().
					ServiceDelete(gomock.Any(), "Jane Doe", "nlxctl", "my-service")
			},
			request: &api.DeleteServiceRequest{
				Name: "my-service",
			},
			expectedError: status.Error(codes.Internal, "database error"),
		},
		"happy_flow": {
			ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"username":               "Jane Doe",
				"grpcgateway-user-agent": "nlxctl",
			})),
			setup: func(_ *common_tls.CertificateBundle, mocks serviceMocks) {
				mocks.db.
					EXPECT().
					DeleteService(gomock.Any(), "my-service", "nlx-test")

				mocks.al.
					EXPECT().
					ServiceDelete(gomock.Any(), "Jane Doe", "nlxctl", "my-service")
			},
			request: &api.DeleteServiceRequest{
				Name: "my-service",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			service, bundle, mocks := newService(t)

			if tt.setup != nil {
				tt.setup(bundle, mocks)
			}

			_, err := service.DeleteService(tt.ctx, tt.request)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
