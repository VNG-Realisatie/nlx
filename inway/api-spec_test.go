// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/plugins"
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

	logger := zap.NewNop()

	ctx := context.Background()
	iw, err := NewInway(ctx, logger, nil, nil, nil, "my.inway", "localhost:1811", "localhost:1812", "localhost:1813", cert, "localhost:1815")
	assert.Nil(t, err)

	apiSpecMockServer := httptest.NewUnstartedServer(http.HandlerFunc(iw.handleAPISpecDocRequest))
	defer apiSpecMockServer.Close()

	apiSpecMockServer.Start()

	services := []*plugins.Service{{
		Name:                        "mock-service",
		APISpecificationDocumentURL: mockAPISpecEndpoint.URL,
	}, {
		Name: "mock-service-without-api-spec",
	}, {
		Name:                        "mock-service-invalid-api-spec",
		APISpecificationDocumentURL: "invalid",
	}}

	err = iw.SetServiceEndpoints(services)
	assert.Nil(t, err)

	tests := map[string]struct {
		url          string
		statusCode   int
		errorMessage string
	}{
		"without_api_specification_url": {
			url:          fmt.Sprintf("%s/.nlx/api-spec-doc/mock-service-without-api-spec", apiSpecMockServer.URL),
			statusCode:   http.StatusNotFound,
			errorMessage: "api specification not found for service\n",
		},
		"service_not_found": {
			url:          fmt.Sprintf("%s/.nlx/api-spec-doc/nonexisting-service", apiSpecMockServer.URL),
			statusCode:   http.StatusNotFound,
			errorMessage: "service not found\n"},
		"invalid_api_specification": {
			url:          fmt.Sprintf("%s/.nlx/api-spec-doc/mock-service-invalid-api-spec", apiSpecMockServer.URL),
			statusCode:   http.StatusInternalServerError,
			errorMessage: "server error\n"},
		"happy_flow": {
			url:          fmt.Sprintf("%s/.nlx/api-spec-doc/mock-service", apiSpecMockServer.URL),
			statusCode:   http.StatusOK,
			errorMessage: ""},
	}

	client := http.Client{}

	for name, test := range tests {
		tc := test

		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.url, nil)
			assert.Nil(t, err)

			resp, err := client.Do(req)
			assert.Nil(t, err)

			bytes, err := ioutil.ReadAll(resp.Body)
			assert.Nil(t, err)

			resp.Body.Close()

			assert.Equal(t, tc.statusCode, resp.StatusCode)
			assert.Equal(t, tc.errorMessage, string(bytes))
		})
	}
}
