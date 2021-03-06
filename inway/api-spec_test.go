// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/config"
)

var pkiDir = filepath.Join("..", "testing", "pki")

//nolint:funlen // this is a test
func TestInwayApiSpec(t *testing.T) {
	cert, err := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	assert.NoError(t, err)

	mockAPISpecEndpoint := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	defer mockAPISpecEndpoint.Close()

	serviceConfig := &config.ServiceConfig{}
	serviceConfig.Services = make(map[string]config.ServiceDetails)
	serviceConfig.Services["mock-service-public"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        mockAPISpecEndpoint.URL,
			AuthorizationModel: config.AuthorizationmodelNone,
		},
	}
	serviceConfig.Services["mock-service-public-apispec"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:                 mockAPISpecEndpoint.URL,
			AuthorizationModel:          config.AuthorizationmodelNone,
			APISpecificationDocumentURL: mockAPISpecEndpoint.URL,
		},
	}
	serviceConfig.Services["mock-service-public-invalid-apispec"] = config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:                 mockAPISpecEndpoint.URL,
			AuthorizationModel:          config.AuthorizationmodelNone,
			APISpecificationDocumentURL: "invalid",
		},
	}

	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	iw, err := NewInway(logger, nil, testProcess, "", "localhost:1812", "localhost:1813", cert, "localhost:1815")
	assert.Nil(t, err)

	apiSpecMockServer := httptest.NewUnstartedServer(http.HandlerFunc(iw.handleAPISpecDocRequest))
	apiSpecMockServer.TLS = iw.orgCertBundle.TLSConfig(iw.orgCertBundle.WithTLSClientAuth())

	apiSpecMockServer.StartTLS()
	defer apiSpecMockServer.Close()

	endPoints := []ServiceEndpoint{}

	for serviceName := range serviceConfig.Services {
		serviceDetails := serviceConfig.Services[serviceName]

		endpoint, errEndpoint := iw.NewHTTPServiceEndpoint(serviceName, &serviceDetails, nil)
		assert.Nil(t, errEndpoint)

		endPoints = append(endPoints, endpoint)
	}

	err = iw.SetServiceEndpoints(endPoints)
	if err != nil {
		t.Fatal("error adding endpoint", err)
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

	client := setupClient(cert)

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
