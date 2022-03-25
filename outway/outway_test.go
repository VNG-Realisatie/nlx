// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	mockdirectory "go.nlx.io/nlx/directory-api/api/mock"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

type authRequest struct {
	Headers      http.Header `json:"headers"`
	Organization string      `json:"organization"`
	Service      string      `json:"service"`
	Path         string      `json:"path"`
}

type authResponse struct {
	Result bool `json:"result"`
}

// nolint:funlen // this is a test
func TestNewOutwayExeception(t *testing.T) {
	orgCertWithoutName, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutName)
	require.NoError(t, err)

	orgCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	orgCertWithoutSerialNumber, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutSerialNumber)
	require.NoError(t, err)

	internalCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.NLXTestInternal)
	require.NoError(t, err)

	tests := map[string]struct {
		args            *NewOutwayArgs
		wantError       error
		wantPluginCount int
	}{
		"certificate_without_organization": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           orgCertWithoutName,
				InternalCert:      internalCert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "",
				AuthCAPath:        "",
			},
			wantError: fmt.Errorf("cannot obtain organization name from self cert"),
		},
		"certificate_without_organization_serial_number": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           orgCertWithoutSerialNumber,
				InternalCert:      internalCert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "",
				AuthCAPath:        "",
			},
			wantError: fmt.Errorf("validation error for subject serial number from cert: cannot be empty"),
		},
		"authorization_service_URL_set_but_no_CA_for_authorization_provided": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           orgCert,
				InternalCert:      internalCert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "http://auth.nlx.io",
				AuthCAPath:        "",
			},
			wantError: fmt.Errorf("authorization service URL set but no CA for authorization provided"),
		},
		"authorization_service_URL_is_not_'https'": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           orgCert,
				InternalCert:      internalCert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "http://auth.nlx.io",
				AuthCAPath:        "/path/to",
			},
			wantError: fmt.Errorf("scheme of authorization service URL is not 'https'"),
		},
		"invalid_monitioring_service_address": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           orgCert,
				InternalCert:      internalCert,
				MonitoringAddress: "",
				AuthServiceURL:    "",
				AuthCAPath:        "",
			},
			wantError: fmt.Errorf("unable to create monitoring service: address required"),
		},
		"directory_client_must_be_not_nil": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           orgCert,
				InternalCert:      internalCert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "https://auth.nlx.io",
				AuthCAPath:        "../testing/pki/ca-root.pem",
			},
			wantError: fmt.Errorf("directory client must be not nil"),
		},
		"happy_flow_with_authorization_plugin": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           orgCert,
				InternalCert:      internalCert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "https://auth.nlx.io",
				AuthCAPath:        "../testing/pki/ca-root.pem",
				DirectoryClient:   mockdirectory.NewMockDirectoryClient(gomock.NewController(t)),
			},
			wantError:       nil,
			wantPluginCount: 4,
		},
		"happy_flow": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           orgCert,
				InternalCert:      internalCert,
				MonitoringAddress: "localhost:8080",
				DirectoryClient:   mockdirectory.NewMockDirectoryClient(gomock.NewController(t)),
			},
			wantError:       nil,
			wantPluginCount: 3,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			outway, err := NewOutway(tt.args)
			if tt.wantError != nil {
				assert.Equal(t, tt.wantError.Error(), err.Error())
			} else {
				assert.NotNil(t, outway)
				assert.Equal(t, tt.wantPluginCount, len(outway.plugins))
			}
		})
	}
}
