// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	mock "go.nlx.io/nlx/outway/mock"

	mockdirectory "go.nlx.io/nlx/directory-api/api/mock"
	"go.nlx.io/nlx/outway/plugins"
)

type authRequest struct {
	Headers      http.Header `json:"headers"`
	Organization string      `json:"organization"`
	Service      string      `json:"service"`
}

type authResponse struct {
	Authorized bool   `json:"authorized"`
	Reason     string `json:"reason,omitempty"`
}

// nolint:funlen // this is a test
func TestNewOutwayExeception(t *testing.T) {
	certOrg, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-without-name-chain.pem"),
		filepath.Join(pkiDir, "org-without-name-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	certWithoutSerialNumber, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-without-serial-number-chain.pem"),
		filepath.Join(pkiDir, "org-without-serial-number-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	tests := map[string]struct {
		args            *NewOutwayArgs
		wantError       error
		wantPluginCount int
	}{
		"certificate_without_organization": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           certOrg,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "",
				AuthCAPath:        "",
			},
			wantError: fmt.Errorf("cannot obtain organization name from self cert"),
		},
		"certificate_without_organization_serial_number": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           certWithoutSerialNumber,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "",
				AuthCAPath:        "",
			},
			wantError: fmt.Errorf("validation error for subject serial number from cert: cannot be empty"),
		},
		"authorization_service_URL_set_but_no_CA_for_authorization_provided": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           cert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "http://auth.nlx.io",
				AuthCAPath:        "",
			},
			wantError: fmt.Errorf("authorization service URL set but no CA for authorization provided"),
		},
		"authorization_service_URL_is_not_'https'": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           cert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "http://auth.nlx.io",
				AuthCAPath:        "/path/to",
			},
			wantError: fmt.Errorf("scheme of authorization service URL is not 'https'"),
		},
		"invalid_monitioring_service_address": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           cert,
				MonitoringAddress: "",
				AuthServiceURL:    "",
				AuthCAPath:        "",
			},
			wantError: fmt.Errorf("unable to create monitoring service: address required"),
		},
		"directory_client_must_be_not_nil": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           cert,
				MonitoringAddress: "localhost:8080",
				AuthServiceURL:    "https://auth.nlx.io",
				AuthCAPath:        "../testing/pki/ca-root.pem",
			},
			wantError: fmt.Errorf("directory client must be not nil"),
		},
		"happy_flow_with_authorization_plugin": {
			args: &NewOutwayArgs{
				Logger:            zap.NewNop(),
				OrgCert:           cert,
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
				OrgCert:           cert,
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

func TestAuthListen(t *testing.T) {
	logger := zap.NewNop()
	// Create an outway with a mock service
	outway := &Outway{
		organization: &Organization{
			serialNumber: "00000000000000000001",
			name:         "org",
		},
		servicesHTTP: make(map[string]HTTPService),
		logger:       logger,
		txlogger:     transactionlog.NewDiscardTransactionLogger(),
	}

	outway.requestHTTPHandler = outway.handleHTTPRequest

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockHTTPService(ctrl)
	mockService.EXPECT().ProxyHTTPRequest(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mockAuthServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authRequest := &authRequest{}
		err := json.NewDecoder(r.Body).Decode(authRequest)
		assert.Nil(t, err)

		authResponse := &authResponse{}
		if user := authRequest.Headers.Get("Authorization-Proxy"); user == "Bearer token" {
			authResponse.Authorized = true
			if encodeErr := json.NewEncoder(w).Encode(authResponse); encodeErr != nil {
				t.Fatal(encodeErr)
			}
			return
		}

		authResponse.Authorized = false
		authResponse.Reason = "invalid user"
		if encodeErr := json.NewEncoder(w).Encode(authResponse); encodeErr != nil {
			t.Fatal(encodeErr)
		}
	}))
	defer mockAuthServer.Close()

	outway.servicesHTTP["00000000000000000001.mockservice"] = mockService
	outway.plugins = append([]plugins.Plugin{
		plugins.NewAuthorizationPlugin(nil, mockAuthServer.URL, http.Client{}),
	}, outway.plugins...)

	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	// Test http responses
	tests := []struct {
		url                    string
		setAuthorizationHeader bool
		statusCode             int
		errorMessage           string
	}{
		{fmt.Sprintf("%s/00000000000000000001/mockservice/", mockServer.URL), false, http.StatusUnauthorized, "nlx outway: authorization failed. reason: invalid user\n"},
		{fmt.Sprintf("%s/00000000000000000001/mockservice/", mockServer.URL), true, http.StatusOK, ""},
	}
	client := http.Client{}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		assert.Nil(t, err)

		if test.setAuthorizationHeader {
			req.Header.Add("Authorization-Proxy", "Bearer token")
		}

		resp, err := client.Do(req)
		assert.Nil(t, err)

		defer resp.Body.Close()

		assert.Equal(t, test.statusCode, resp.StatusCode)

		bytes, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, test.errorMessage, string(bytes))
	}
}
