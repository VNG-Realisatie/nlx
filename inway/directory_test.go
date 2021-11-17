package inway

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/nlxversion"
	common_tls "go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	mock_directory "go.nlx.io/nlx/directory-api/api/mock"
)

func TestRegisterToDirectory(t *testing.T) {
	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	tests := map[string]struct {
		directoryClient func(ctx context.Context, ctrl *gomock.Controller) *mock_directory.MockDirectoryClient
		expectedError   error
	}{
		"registration_failed": {
			directoryClient: func(ctx context.Context, ctrl *gomock.Controller) *mock_directory.MockDirectoryClient {
				nlxVersion := nlxversion.NewGRPCContext(ctx, "inway")

				directoryClient := mock_directory.NewMockDirectoryClient(ctrl)
				directoryClient.EXPECT().RegisterInway(nlxVersion, &directoryapi.RegisterInwayRequest{
					InwayName:           "mock-inway",
					InwayAddress:        "localhost:1812",
					IsOrganizationInway: false,
					Services:            []*directoryapi.RegisterInwayRequest_RegisterService{},
				}).Return(nil, fmt.Errorf("arbitrary error"))

				return directoryClient
			},
			expectedError: fmt.Errorf("arbitrary error"),
		},
		"error_in_register_inway_response": {
			directoryClient: func(ctx context.Context, ctrl *gomock.Controller) *mock_directory.MockDirectoryClient {
				nlxVersion := nlxversion.NewGRPCContext(ctx, "inway")

				directoryClient := mock_directory.NewMockDirectoryClient(ctrl)
				directoryClient.EXPECT().RegisterInway(nlxVersion, &directoryapi.RegisterInwayRequest{
					InwayName:           "mock-inway",
					InwayAddress:        "localhost:1812",
					IsOrganizationInway: false,
					Services:            []*directoryapi.RegisterInwayRequest_RegisterService{},
				}).Return(&directoryapi.RegisterInwayResponse{
					Error: "call failed",
				}, nil)
				return directoryClient
			},
			expectedError: fmt.Errorf("call failed"),
		},
		"happy_flow": {
			directoryClient: func(ctx context.Context, ctrl *gomock.Controller) *mock_directory.MockDirectoryClient {
				nlxVersion := nlxversion.NewGRPCContext(ctx, "inway")

				directoryClient := mock_directory.NewMockDirectoryClient(ctrl)
				directoryClient.EXPECT().RegisterInway(nlxVersion, &directoryapi.RegisterInwayRequest{
					InwayName:           "mock-inway",
					InwayAddress:        "localhost:1812",
					IsOrganizationInway: false,
					Services:            []*directoryapi.RegisterInwayRequest_RegisterService{},
				}).Return(&directoryapi.RegisterInwayResponse{}, nil)

				return directoryClient
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

			params := &Params{
				Context:                 ctx,
				Logger:                  zap.NewNop(),
				Txlogger:                nil,
				ManagementProxy:         nil,
				Name:                    "mock-inway",
				Address:                 "localhost:1812",
				MonitoringAddress:       "localhost:1813",
				ListenManagementAddress: "",
				OrgCertBundle:           cert,
				DirectoryClient:         tc.directoryClient(ctx, ctrl),
			}

			iw, err := NewInway(params)
			assert.Nil(t, err)

			err = iw.registerToDirectory(ctx)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
