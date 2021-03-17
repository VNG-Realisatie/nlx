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

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	mock "go.nlx.io/nlx/outway/mock"
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

	tests := []struct {
		description              string
		cert                     *common_tls.CertificateBundle
		monitoringServiceAddress string
		authServiceURL           string
		authCAPath               string
		expectedErrorMessage     string
	}{
		{
			"certificate without organization",
			certOrg,
			"localhost:8080",
			"",
			"",
			"cannot obtain organization name from self cert",
		},
		{
			"authorization service URL set but no CA for authorization provided",
			cert,
			"localhost:8080",
			"http://auth.nlx.io",
			"",
			"authorization service URL set but no CA for authorization provided",
		},
		{
			"authorization service URL is not 'https'",
			cert,
			"localhost:8080",
			"http://auth.nlx.io",
			"/path/to",
			"scheme of authorization service URL is not 'https'",
		},
		{
			"invalid monitioring service address",
			cert,
			"",
			"",
			"",
			"unable to create monitoring service: address required",
		},
	}

	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	// Test exceptions during outway creation
	for _, test := range tests {
		_, err := NewOutway(logger, nil, testProcess, test.monitoringServiceAddress, test.cert, "", test.authServiceURL, test.authCAPath, false)
		assert.EqualError(t, err, test.expectedErrorMessage)
	}
}

func TestAuthListen(t *testing.T) {
	logger := zap.NewNop()
	// Createa a outway with a mock service
	outway := &Outway{
		organizationName: "org",
		servicesHTTP:     make(map[string]HTTPService),
		logger:           logger,
		txlogger:         transactionlog.NewDiscardTransactionLogger(),
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
				panic(encodeErr)
			}
			return
		}

		authResponse.Authorized = false
		authResponse.Reason = "invalid user"
		if encodeErr := json.NewEncoder(w).Encode(authResponse); encodeErr != nil {
			panic(encodeErr)
		}
	}))
	defer mockAuthServer.Close()

	outway.servicesHTTP["mockorg.mockservice"] = mockService
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
		{fmt.Sprintf("%s/mockorg/mockservice/", mockServer.URL), false, http.StatusUnauthorized, "nlx outway: authorization failed. reason: invalid user\n"},
		{fmt.Sprintf("%s/mockorg/mockservice/", mockServer.URL), true, http.StatusOK, ""},
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
