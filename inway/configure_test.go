// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
package inway

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.nlx.io/nlx/inway/plugins"
	"go.nlx.io/nlx/management-api/api"
	mock_api "go.nlx.io/nlx/management-api/api/mock"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen // this is a test
func TestStartConfigurationPolling(t *testing.T) {
	orgCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	tests := map[string]struct {
		managementClient          func(ctrl *gomock.Controller) *mock_api.MockManagementClient
		expectError               bool
		expectedErrorMessage      string
		expectedService           *plugins.Service
		shouldBeOrganizationInway bool
	}{
		"management_api_unavailable": {
			managementClient: func(ctrl *gomock.Controller) *mock_api.MockManagementClient {
				managementClient := mock_api.NewMockManagementClient(ctrl)
				managementClient.EXPECT().GetInwayConfig(gomock.Any(), gomock.Any()).Return(nil, errManagementAPIUnavailable)

				return managementClient
			},
			expectError:               true,
			expectedErrorMessage:      "managementAPI unavailable",
			shouldBeOrganizationInway: false,
		},
		"get_inway_config_arbitrary error": {
			managementClient: func(ctrl *gomock.Controller) *mock_api.MockManagementClient {
				managementClient := mock_api.NewMockManagementClient(ctrl)
				managementClient.EXPECT().GetInwayConfig(gomock.Any(), gomock.Any()).Return(nil, errors.New("arbitrary error"))

				return managementClient
			},
			expectError:               true,
			expectedErrorMessage:      "arbitrary error",
			shouldBeOrganizationInway: false,
		},
		"happy_flow_organization_inway": {
			managementClient: func(ctrl *gomock.Controller) *mock_api.MockManagementClient {
				managementClient := mock_api.NewMockManagementClient(ctrl)

				managementClient.EXPECT().GetInwayConfig(gomock.Any(), &api.GetInwayConfigRequest{
					Name: "mock-inway",
				}).Return(&api.GetInwayConfigResponse{
					IsOrganizationInway: true,
					Services: []*api.GetInwayConfigResponse_Service{
						{
							Name:                 "mock-service",
							EndpointURL:          "http://endpoint.mock",
							DocumentationURL:     "http://docs.mock",
							ApiSpecificationURL:  "http://api-specs.mock",
							Internal:             false,
							TechSupportContact:   "tech@support.mock",
							PublicSupportContact: "public@support.mock",
							AuthorizationSettings: &api.GetInwayConfigResponse_Service_AuthorizationSettings{
								Authorizations: []*api.GetInwayConfigResponse_Service_AuthorizationSettings_Authorization{
									{
										Organization: &api.Organization{
											SerialNumber: "00000000000000000001",
											Name:         "mock-org",
										},
										PublicKeyHash: "mock-public-key-hash",
										PublicKeyPEM:  "mock-public-key-pem",
									},
								},
							},
						},
					},
				}, nil)
				return managementClient
			},
			expectError: false,
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
						OrganizationSerialNumber: "00000000000000000001",
						PublicKeyFingerprint:     "mock-public-key-hash",
						PublicKeyPEM:             "mock-public-key-pem",
					},
				},
			},
			shouldBeOrganizationInway: true,
		},
		"happy_flow_not_organization_inway": {
			managementClient: func(ctrl *gomock.Controller) *mock_api.MockManagementClient {
				managementClient := mock_api.NewMockManagementClient(ctrl)

				managementClient.EXPECT().GetInwayConfig(gomock.Any(), &api.GetInwayConfigRequest{
					Name: "mock-inway",
				}).Return(&api.GetInwayConfigResponse{
					IsOrganizationInway: false,
					Services:            []*api.GetInwayConfigResponse_Service{},
				}, nil)
				return managementClient
			},
			expectError:               false,
			shouldBeOrganizationInway: false,
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

			params := &Params{
				Context:                         ctx,
				Logger:                          zap.NewNop(),
				Txlogger:                        nil,
				ManagementClient:                tc.managementClient(ctrl),
				ManagementProxy:                 nil,
				Name:                            "mock-inway",
				Address:                         "localhost:1812",
				MonitoringAddress:               "localhost:1813",
				ListenAddressManagementAPIProxy: "",
				OrgCertBundle:                   orgCert,
			}

			iw, err := NewInway(params)
			assert.Nil(t, err)

			err = iw.retrieveAndUpdateConfig()
			if tc.expectError {
				assert.EqualError(t, err, tc.expectedErrorMessage)
			}

			if tc.expectedService != nil {
				service := iw.services[tc.expectedService.Name]
				assert.NotNil(t, service)

				assert.Equal(t, tc.expectedService, service)
			}

			assert.Equal(t, tc.shouldBeOrganizationInway, iw.isOrganizationInway)
		})
	}
}
