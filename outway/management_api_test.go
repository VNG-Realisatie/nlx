// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package outway

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	mock_directory_api "go.nlx.io/nlx/directory-api/api/mock"
	management_api "go.nlx.io/nlx/management-api/api"
	mock_management_api "go.nlx.io/nlx/management-api/api/mock"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

func TestRegisterToManagementAPI(t *testing.T) {
	orgCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	internalCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.NLXTestInternal)
	require.NoError(t, err)

	publicKeyPEM, err := orgCert.PublicKeyPEM()
	require.NoError(t, err)

	tests := map[string]struct {
		outwayName          string
		outwaySelfAddress   string
		managementAPIClient func(ctx context.Context, ctrl *gomock.Controller) *mock_management_api.MockManagementClient
		expectedError       error
	}{
		"registration_failed": {
			outwayName:        "mock-outway",
			outwaySelfAddress: "outway.address.com",
			managementAPIClient: func(ctx context.Context, ctrl *gomock.Controller) *mock_management_api.MockManagementClient {
				managementClient := mock_management_api.NewMockManagementClient(ctrl)
				managementClient.EXPECT().RegisterOutway(gomock.Any(), &management_api.RegisterOutwayRequest{
					Name:           "mock-outway",
					SelfAddressApi: "outway.address.com",
					Version:        "unknown",
					PublicKeyPem:   publicKeyPEM,
				}).Return(nil, fmt.Errorf("arbitrary error"))

				return managementClient
			},
			expectedError: fmt.Errorf("arbitrary error"),
		},
		"happy_flow": {
			outwayName:        "mock-outway",
			outwaySelfAddress: "outway.address.com",
			managementAPIClient: func(ctx context.Context, ctrl *gomock.Controller) *mock_management_api.MockManagementClient {
				managementClient := mock_management_api.NewMockManagementClient(ctrl)
				managementClient.EXPECT().RegisterOutway(gomock.Any(), &management_api.RegisterOutwayRequest{
					Name:           "mock-outway",
					SelfAddressApi: "outway.address.com",
					Version:        "unknown",
					PublicKeyPem:   publicKeyPEM,
				}).Return(nil, nil)

				return managementClient
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

			args := &NewOutwayArgs{
				Ctx:                 ctx,
				Logger:              zap.NewNop(),
				Txlogger:            nil,
				Name:                tc.outwayName,
				AddressAPI:          tc.outwaySelfAddress,
				MonitoringAddress:   "localhost:1813",
				InternalCert:        internalCert,
				OrgCert:             orgCert,
				DirectoryClient:     mock_directory_api.NewMockDirectoryClient(ctrl),
				ManagementAPIClient: tc.managementAPIClient(ctx, ctrl),
			}

			ow, err := New(args)
			assert.Nil(t, err)

			err = ow.registerToManagementAPI(ctx)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
