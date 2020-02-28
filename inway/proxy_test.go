// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway/config"
)

func newTestEnv(t *testing.T, tlsOptions orgtls.TLSOptions) (proxy, mock *httptest.Server) {

	// Mock endpoint (service)
	mockEndPoint := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	// defer is missing do this in you test!!
	// defer mockEndPoint.Close()

	serviceConfig := &config.ServiceConfig{}
	serviceConfig.Services = make(map[string]config.ServiceDetails)
	serviceConfig.Services["mock-service-whitelist"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "whitelist",
		},
		AuthorizationWhitelist: []config.AuthorizationWhitelistItem{{OrganizationName: "nlx-test"}},
	}
	serviceConfig.Services["mock-service-whitelist-unauthorized"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "whitelist",
		},
		AuthorizationWhitelist: []config.AuthorizationWhitelistItem{{OrganizationName: "nlx-forbidden"}},
	}
	serviceConfig.Services["mock-service-unspecified-unauthorized"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "",
		},
	}
	serviceConfig.Services["mock-service-public"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockEndPoint.URL,
			AuthorizationModel: "none",
		},
	}

	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	iw, err := NewInway(logger, nil, testProcess, "", "localhost:1812", tlsOptions, "localhost:1815")
	assert.Nil(t, err)

	endPoints := []ServiceEndpoint{}
	// Add service endpoints
	for serviceName := range serviceConfig.Services {
		serviceDetails := serviceConfig.Services[serviceName]
		endpoint, errr := iw.NewHTTPServiceEndpoint(serviceName, &serviceDetails, nil)
		if errr != nil {
			t.Fatal("failed to create service endpoint", err)
		}

		endPoints = append(endPoints, endpoint)
	}

	err = iw.SetServiceEndpoints(endPoints)
	assert.Nil(t, err)

	proxyRequestMockServer := httptest.NewUnstartedServer(http.HandlerFunc(iw.handleProxyRequest))
	proxyRequestMockServer.TLS = &tls.Config{
		ClientCAs:  iw.roots,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	return proxyRequestMockServer, mockEndPoint

}

func TestInwayProxyRequest(t *testing.T) {

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
		OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
		OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
	}

	proxyRequestMockServer, mockEndPoint := newTestEnv(t, tlsOptions)
	proxyRequestMockServer.StartTLS()
	defer proxyRequestMockServer.Close()
	defer mockEndPoint.Close()

	client := setupClient(t, tlsOptions)

	//nolint:dupl
	tests := []struct {
		url          string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusOK, ""},
		{fmt.Sprintf("%s/mock-service-whitelist/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusOK, ""},
		{fmt.Sprintf("%s/mock-service-whitelist-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-inway: permission denied, organization \"nlx-test\" or public key \"sha256:eb48a0a7a922688175e1b40274d88f3e188fdd727de6a2c5840148d5e989726e\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service-unspecified-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-inway: permission denied, organization \"nlx-test\" or public key \"sha256:eb48a0a7a922688175e1b40274d88f3e188fdd727de6a2c5840148d5e989726e\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx-inway: invalid path in url\n"},
		{fmt.Sprintf("%s/mock-service/fictive", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx-inway: no endpoint for service\n"},
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "", http.StatusBadRequest, "nlx-inway: missing logrecord id\n"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.url, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.url, nil)
			assert.Nil(t, err)

			req.Header.Add("X-NLX-Logrecord-Id", test.logRecordID)
			resp, err := client.Do(req)
			assert.Nil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, test.statusCode, resp.StatusCode)

			defer resp.Body.Close()
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("error parsing result.body", err)
			}
			assert.Equal(t, test.errorMessage, string(bytes))
		})
	}
}

func TestInwayNoOrgProxyRequest(t *testing.T) {

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
		OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
		OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
	}

	tlsNoOrgOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
		OrgCertFile: filepath.Join("..", "testing", "pki", "org-without-name-chain.pem"),
		OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-without-name-key.pem"),
	}

	// Clients with no organization specified in the certificate
	// should not be allowed on the nlx network.
	proxyRequestMockServer, mockEndPoint := newTestEnv(t, tlsOptions)
	proxyRequestMockServer.StartTLS()
	defer proxyRequestMockServer.Close()
	defer mockEndPoint.Close()

	//nolint:dupl
	tests := []struct {
		url          string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, ""},
		{fmt.Sprintf("%s/mock-service-whitelist/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, ""},
		{fmt.Sprintf("%s/mock-service-whitelist-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-outway: could not handle your request, organization \"nlx-test\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service-unspecified-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx-outway: could not handle your request, organization \"nlx-test\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx inway error: invalid path in url\n"},
		{fmt.Sprintf("%s/mock-service/fictive", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx inway error: no endpoint for service\n"},
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "", http.StatusBadRequest, "nlx-outway: missing logrecord id\n"},
	}

	noOrgClient := setupClient(t, tlsNoOrgOptions)

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		assert.Nil(t, err)

		req.Header.Add("X-NLX-Logrecord-Id", test.logRecordID)
		resp, err := noOrgClient.Do(req)
		assert.Nil(t, err)
		defer resp.Body.Close()

		if resp.StatusCode != 400 {
			t.Fatalf(
				`result: "%d" for url "%s", expected http status code : "%d"`,
				resp.StatusCode, test.url, 400)
		}
	}
}
