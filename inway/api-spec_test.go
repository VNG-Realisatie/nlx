// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

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

func TestInwayApiSpec(t *testing.T) {
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org-nlx-test.key"}

	mockAPISpecEndpoint := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	defer mockAPISpecEndpoint.Close()

	serviceConfig := &config.ServiceConfig{}
	serviceConfig.Services = make(map[string]config.ServiceDetails)
	serviceConfig.Services["mock-service-public"] = config.ServiceDetails{
		EndpointURL:        mockAPISpecEndpoint.URL,
		AuthorizationModel: "none",
	}
	serviceConfig.Services["mock-service-public-apispec"] = config.ServiceDetails{
		EndpointURL:                 mockAPISpecEndpoint.URL,
		AuthorizationModel:          "none",
		APISpecificationDocumentURL: mockAPISpecEndpoint.URL,
	}
	serviceConfig.Services["mock-service-public-invalid-apispec"] = config.ServiceDetails{
		EndpointURL:                 mockAPISpecEndpoint.URL,
		AuthorizationModel:          "none",
		APISpecificationDocumentURL: "invalid",
	}

	logger := zap.NewNop()
	iw, err := NewInway(logger, nil, "localhost:1812", tlsOptions,
		"localhost:1815", serviceConfig)
	assert.Nil(t, err)

	p := process.NewProcess(logger)
	apiSpecMockServer := httptest.NewUnstartedServer(http.HandlerFunc(iw.handleAPISpecDocRequest))
	apiSpecMockServer.TLS = &tls.Config{
		ClientCAs:  iw.roots,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	apiSpecMockServer.StartTLS()
	defer apiSpecMockServer.Close()

	for serviceName, serviceDetails := range serviceConfig.Services {
		endpoint, err := iw.NewHTTPServiceEndpoint(logger, serviceName, serviceDetails.EndpointURL, nil)
		assert.Nil(t, err)

		switch serviceDetails.AuthorizationModel {
		case "none", "":
			endpoint.SetAuthorizationPublic()
		case "whitelist":
			endpoint.SetAuthorizationWhitelist(serviceDetails.AuthorizationWhitelist)
		default:
			logger.Fatal(fmt.Sprintf(`invalid authorization model "%s" for service "%s"`, serviceDetails.AuthorizationModel, serviceName))
		}

		err = iw.AddServiceEndpoint(p, endpoint, serviceDetails)

		if err != nil {
			t.Fatal("error adding endpoint", err)
		}
	}

	tests := []struct {
		url          string
		logRecordID  string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/.nlx/api-spec-doc/mock-service-public", apiSpecMockServer.URL),
			"dummy-ID", http.StatusNotFound, "api specification not found for service\n"},
		{fmt.Sprintf("%s/.nlx/api-spec-doc/nonexisting-service", apiSpecMockServer.URL),
			"dummy-ID", http.StatusNotFound, "service not found\n"},
		{fmt.Sprintf("%s/.nlx/api-spec-doc/mock-service-public-invalid-apispec", apiSpecMockServer.URL),
			"dummy-ID", http.StatusInternalServerError, "server error\n"},
		{fmt.Sprintf("%s/.nlx/api-spec-doc/mock-service-public-apispec", apiSpecMockServer.URL), "dummy-ID", http.StatusOK, ""},
	}

	client := setupClient(t, tlsOptions)

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Add("X-NLX-Logrecord-Id", test.logRecordID)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(`error doing http request`, err)
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error parsing result.body", err)
		}
		resp.Body.Close()

		assert.Equal(t, test.statusCode, resp.StatusCode)
		assert.Equal(t, test.errorMessage, string(bytes))
	}
}
