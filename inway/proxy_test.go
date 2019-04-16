package inway

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway/config"
)

func TestInWayProxyRequest(t *testing.T) {
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org-nlx-test.key",
	}
	pool, err := orgtls.LoadRootCert(tlsOptions.NLXRootCert)
	assert.Nil(t, err)

	// Mock endpoint (service)
	mockEndPoint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer mockEndPoint.Close()

	serviceConfig := &config.ServiceConfig{}
	serviceConfig.Services = make(map[string]config.ServiceDetails)
	serviceConfig.Services["mock-service-whitelist"] = config.ServiceDetails{
		EndpointURL:            mockEndPoint.URL,
		AuthorizationWhitelist: []string{"nlx-test"},
		AuthorizationModel:     "whitelist",
	}
	serviceConfig.Services["mock-servicewhitelist-unauthorized"] = config.ServiceDetails{
		EndpointURL:            mockEndPoint.URL,
		AuthorizationWhitelist: []string{"nlx-forbidden"},
		AuthorizationModel:     "whitelist",
	}
	serviceConfig.Services["mock-service-public"] = config.ServiceDetails{
		EndpointURL:        mockEndPoint.URL,
		AuthorizationModel: "none",
	}

	logger := zap.NewNop()
	iw, err := NewInway(logger, nil, "localhost:1812", tlsOptions,
		"localhost:1815", serviceConfig)
	assert.Nil(t, err)

	p := process.NewProcess(logger)

	proxyRequestMockServer := httptest.NewUnstartedServer(http.HandlerFunc(iw.handleProxyRequest))
	proxyRequestMockServer.TLS = &tls.Config{
		ClientCAs:  iw.roots,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	proxyRequestMockServer.StartTLS()
	defer proxyRequestMockServer.Close()

	// Add service endpoints
	for serviceName, serviceDetails := range serviceConfig.Services {
		endpoint, err := iw.NewHTTPServiceEndpoint(logger, serviceName, serviceDetails.EndpointURL, nil)
		if err != nil {
			t.Fatal("failed to create service endpoint", err)
		}

		switch serviceDetails.AuthorizationModel {
		case "none", "":
			endpoint.SetAuthorizationPublic()
		case "whitelist":
			endpoint.SetAuthorizationWhitelist(serviceDetails.AuthorizationWhitelist)
		default:
			logger.Fatal(fmt.Sprintf(`invalid authorization model "%s" for service "%s"`, serviceDetails.AuthorizationModel, serviceName))
		}

		err = iw.AddServiceEndpoint(p, endpoint, serviceDetails)
		assert.Nil(t, err)
	}

	cert, err := tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	assert.Nil(t, err)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: pool, Certificates: []tls.Certificate{cert}}}
	client := http.Client{
		Transport: tr,
	}

	// Test http responses
	tests := []struct {
		url          string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusOK, ""},
		{fmt.Sprintf("%s/mock-service-whitelist/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusOK, ""},
		{fmt.Sprintf("%s/mock-servicewhitelist-unauthorized/dummy", proxyRequestMockServer.URL), "dummy-ID", http.StatusForbidden, "nlx outway: could not handle your request, organization \"nlx-test\" is not allowed access.\n"},
		{fmt.Sprintf("%s/mock-service", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx inway error: invalid path in url\n"},
		{fmt.Sprintf("%s/mock-service/fictive", proxyRequestMockServer.URL), "dummy-ID", http.StatusBadRequest, "nlx inway error: no endpoint for service\n"},
		{fmt.Sprintf("%s/mock-service-public/dummy", proxyRequestMockServer.URL), "", http.StatusBadRequest, "nlx outway: missing logrecord id\n"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		assert.Nil(t, err)

		req.Header.Add("X-NLX-Logrecord-Id", test.logRecordID)
		resp, err := client.Do(req)
		assert.Nil(t, err)

		assert.Equal(t, test.statusCode, resp.StatusCode)

		bytes, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		assert.Equal(t, test.errorMessage, string(bytes))
	}
}
