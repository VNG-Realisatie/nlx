// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	directoryapi_mock "go.nlx.io/nlx/directory-api/api/mock"
)

func TestGetOrganizationInwayProxyAddress(t *testing.T) {
	organizationSerialNumber := "00000000000000000001"
	tests := map[string]struct {
		directoryClient func(ctrl *gomock.Controller) directoryapi.DirectoryClient
		want            string
		wantErr         error
	}{
		"happy_flow": {
			directoryClient: func(ctrl *gomock.Controller) directoryapi.DirectoryClient {
				client := directoryapi_mock.NewMockDirectoryClient(ctrl)

				client.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directoryapi.GetOrganizationManagementAPIProxyAddressRequest{
					OrganizationSerialNumber: organizationSerialNumber,
				}).Return(&directoryapi.GetOrganizationManagementAPIProxyAddressResponse{
					Address: "localhost:8443",
				}, nil)

				return client
			},
			want:    "localhost:8443",
			wantErr: nil,
		},
		"directory_client_errors": {
			directoryClient: func(ctrl *gomock.Controller) directoryapi.DirectoryClient {
				client := directoryapi_mock.NewMockDirectoryClient(ctrl)

				client.EXPECT().GetOrganizationManagementAPIProxyAddress(gomock.Any(), &directoryapi.GetOrganizationManagementAPIProxyAddressRequest{
					OrganizationSerialNumber: organizationSerialNumber,
				}).Return(nil, fmt.Errorf("arbitrary error"))
				return client
			},
			want:    "",
			wantErr: fmt.Errorf("arbitrary error"),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			directoryClient := tt.directoryClient(ctrl)
			client := &client{
				directoryClient,
			}

			got, err := client.GetOrganizationInwayProxyAddress(context.Background(), organizationSerialNumber)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
