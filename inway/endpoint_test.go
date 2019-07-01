// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
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

func TestSetAuthorization(t *testing.T) {
	endpoint := &HTTPServiceEndpoint{}
	// Test if public authorization is set
	endpoint.SetAuthorizationPublic()
	assert.True(t, endpoint.public)

	// Test if whitelist is created
	whiteList := []string{"demo-org"}
	endpoint.SetAuthorizationWhitelist(whiteList)
	assert.False(t, endpoint.public)
	assert.Len(t, endpoint.whitelistedOrganizations, 1)
	assert.Equal(t, whiteList, endpoint.whitelistedOrganizations)

	// Test if a not whitelisted organization will receive a 403 response
	var err error
	endpoint.logger = zap.NewNop()
	httpRecorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/whitelist", nil)
	reqMD := &RequestMetadata{
		requesterOrganization: "demo-org-fault",
	}
	endpoint.handleRequest(reqMD, httpRecorder, req)
	result := httpRecorder.Result()
	assert.Equal(t, http.StatusForbidden, result.StatusCode)

	// Test if missing organization will receive a 400 response
	reqMD2 := &RequestMetadata{}

	endpoint.handleRequest(reqMD2, httpRecorder, req)
	result2 := httpRecorder.Result()
	assert.Equal(t, http.StatusForbidden, result2.StatusCode)

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("error parsing result.body", err)
	}

	assert.Equal(t, fmt.Sprintf("nlx outway: could not handle your request, organization \"%s\" is not allowed access.\n", reqMD.requesterOrganization), string(bytes))
}

func TestInwayAddServiceEndpoint(t *testing.T) {
	logger := zap.NewNop()

	// Certificate organization = nlx-test

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org-nlx-test.key",
	}

	iw, err := NewInway(logger, nil, "localhost:1812", tlsOptions,
		"localhost:1815", nil)
	assert.Nil(t, err)

	p := process.NewProcess(logger)
	// Test NewHTTPServiceEnpoint with invalid url
	endpoint, err := iw.NewHTTPServiceEndpoint(logger, "mock-service", "12://invalid-endpoint", nil)
	assert.EqualError(t, err, "invalid endpoint provided: parse 12://invalid-endpoint: first path segment in URL cannot contain colon")

	// Test NewHTTPServicedEnpoint
	endpoint, err = iw.NewHTTPServiceEndpoint(logger, "mock-service", "127.0.0.1", nil)
	assert.Nil(t, err)
	assert.Equal(t, "mock-service", endpoint.ServiceName())

	// Test if duplicate endpoints are disallowed
	err = iw.AddServiceEndpoint(p, endpoint, config.ServiceDetails{
		EndpointURL:            "http://127.0.0.1:1813",
		AuthorizationWhitelist: []string{"nlx-forbidden"},
	})
	assert.Nil(t, err)

	err = iw.AddServiceEndpoint(p, endpoint, config.ServiceDetails{
		EndpointURL:            "http://127.0.0.1:1813",
		AuthorizationWhitelist: []string{"nlx-forbidden"},
	})
	if err == nil {
		t.Fatal("result: error is nil, expected error when calling AddServiceEndpoint with a duplicate service")
	}
	assert.EqualError(t, err, "service endpoint for a service with the same name has already been registered")

}

func TestHTTPServiceEndpointCreateRecordData(t *testing.T) {
	requestPath := "/demo/mock"
	header := http.Header{}
	processID := "123456"
	dataElement := "mock-element"
	header.Add("X-NLX-Request-Process-Id", processID)
	header.Add("X-NLX-Request-Data-Elements", dataElement)
	endpoint := HTTPServiceEndpoint{}

	recordData := endpoint.createRecordData(requestPath, header)

	tests := []struct {
		doelBindingName  string
		doelBindingValue string
	}{
		{doelBindingName: "doelbinding-process-id",
			doelBindingValue: processID},
		{doelBindingName: "doelbinding-data-elements",
			doelBindingValue: dataElement},
	}

	for _, test := range tests {
		assert.Contains(t, recordData, test.doelBindingName)
		assert.Equal(t, recordData[test.doelBindingName], test.doelBindingValue)
	}
}
