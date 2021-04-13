// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
package inway

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/plugins"
	"go.nlx.io/nlx/management-api/api"
	mock_api "go.nlx.io/nlx/management-api/api/mock"
)

//nolint:funlen // this is a test
func TestStartConfigurationPolling(t *testing.T) {
	hostname, err := os.Hostname()
	assert.Nil(t, err)

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	tests := map[string]struct {
		managementClient     func(ctrl *gomock.Controller) *mock_api.MockManagementClient
		expectError          bool
		expectedErrorMessage string
		expectedService      *plugins.Service
	}{
		"cannot_register_to_management_api": {
			managementClient: func(ctrl *gomock.Controller) *mock_api.MockManagementClient {
				managementClient := mock_api.NewMockManagementClient(ctrl)
				managementClient.EXPECT().CreateInway(gomock.Any(), &api.Inway{
					Name:        "mock.inway",
					Version:     "unknown",
					Hostname:    hostname,
					SelfAddress: "localhost:1812",
				}).Return(nil, fmt.Errorf("arbitrary error"))

				return managementClient
			},
			expectError:          true,
			expectedErrorMessage: "arbitrary error",
		},
		"happy_flow": {
			managementClient: func(ctrl *gomock.Controller) *mock_api.MockManagementClient {
				managementClient := mock_api.NewMockManagementClient(ctrl)
				managementClient.EXPECT().CreateInway(gomock.Any(), &api.Inway{
					Name:        "mock.inway",
					Version:     "unknown",
					Hostname:    hostname,
					SelfAddress: "localhost:1812",
				}).Return(nil, nil)

				managementClient.EXPECT().ListServices(gomock.Any(), &api.ListServicesRequest{
					InwayName: "mock.inway",
				}).Return(&api.ListServicesResponse{
					Services: []*api.ListServicesResponse_Service{
						{
							Name:                 "mock-service",
							EndpointURL:          "http://endpoint.mock",
							DocumentationURL:     "http://docs.mock",
							ApiSpecificationURL:  "http://api-specs.mock",
							Internal:             false,
							TechSupportContact:   "tech@support.mock",
							PublicSupportContact: "public@support.mock",
							AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
								Authorizations: []*api.ListServicesResponse_Service_AuthorizationSettings_Authorization{
									{
										OrganizationName: "mock-org",
										PublicKeyHash:    "mock-public-key-hash",
										PublicKeyPEM:     "mock-public-key-pem",
									},
								},
							},
						},
					},
				}, nil)
				return managementClient
			},
			expectError:          false,
			expectedErrorMessage: "arbitrary error",
			expectedService: &plugins.Service{
				Name:                        "mock-service",
				EndpointURL:                 "http://endpoint.mock",
				DocumentationURL:            "http://docs.mock",
				APISpecificationDocumentURL: "http://api-specs.mock",
				Internal:                    false,
				TechSupportContact:          "tech@support.mock",
				PublicSupportContact:        "public@support.mock",
				Grants: []*plugins.Grant{
					{
						OrganizationName:     "mock-org",
						PublicKeyFingerprint: "mock-public-key-hash",
						PublicKeyPEM:         "mock-public-key-pem",
					},
				},
			},
		},
	}

	for name, test := range tests {
		tc := test

		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			ctrl := gomock.NewController(t)

			t.Cleanup(func() {
				cancel()
				ctrl.Finish()
			})

			iw, err := NewInway(ctx, zap.NewNop(), nil, tc.managementClient(ctrl), nil, "mock.inway", "localhost:1812", "localhost:1813", "", cert, "localhost:1815")
			assert.Nil(t, err)

			err = iw.startConfigurationPolling(ctx)
			if tc.expectError {
				assert.EqualError(t, err, tc.expectedErrorMessage)
			}

			if tc.expectedService != nil {
				service := iw.services[tc.expectedService.Name]
				assert.NotNil(t, service)

				assert.Equal(t, tc.expectedService, service)
			}
		})
	}
}
