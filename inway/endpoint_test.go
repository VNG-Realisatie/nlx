// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

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
	whiteList := []config.AuthorizationWhitelistItem{{OrganizationName: "demo-org"}, {PublicKeyHash: "demo-cert"}}
	endpoint.SetAuthorizationWhitelist(whiteList)
	assert.False(t, endpoint.public)
	assert.Len(t, endpoint.whitelistedOrganizations, 2)
	assert.Equal(t, whiteList, endpoint.whitelistedOrganizations)

	// Test if a not whitelisted organization will receive a 403 response
	var err error

	endpoint.logger = zaptest.NewLogger(t)
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
	result2.Body.Close()

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("error parsing result.body", err)
	}
	result.Body.Close()

	assert.Equal(
		t,
		fmt.Sprintf("nlx-inway: permission denied, organization \"%s\" or public key \"\" is not allowed access.\n",
			reqMD.requesterOrganization),
		string(bytes),
	)
}

func TestWhitelist(t *testing.T) {
	endpoint := &HTTPServiceEndpoint{
		whitelistedOrganizations: []config.AuthorizationWhitelistItem{
			{OrganizationName: "only-org"},
			{PublicKeyHash: "only-cert"},
			{OrganizationName: "with-name-and-cert", PublicKeyHash: "with-cert-and-name"},
			{}, // This would be a anomaly but we don't want it to be an allow all rule
		},
		logger: zaptest.NewLogger(t),
	}
	req := httptest.NewRequest("GET", "/whitelist", nil)

	type want struct {
		statusCode int
		body       string
	}

	tests := []struct {
		name              string
		requesterMetadata *RequestMetadata
		want              want
	}{
		{
			name:              "only certificate",
			requesterMetadata: &RequestMetadata{requesterOrganization: "irrelevant", requesterPublicKeyHash: "only-cert"},
			want:              want{statusCode: http.StatusBadRequest, body: "nlx-inway: missing logrecord id\n"},
		},
		{
			name:              "only organization",
			requesterMetadata: &RequestMetadata{requesterOrganization: "only-org"},
			want:              want{statusCode: http.StatusBadRequest, body: "nlx-inway: missing logrecord id\n"},
		},
		{
			name:              "with name and cert",
			requesterMetadata: &RequestMetadata{requesterOrganization: "with-name-and-cert", requesterPublicKeyHash: "with-cert-and-name"},
			want:              want{statusCode: http.StatusBadRequest, body: "nlx-inway: missing logrecord id\n"},
		},
		{
			name:              "unknown",
			requesterMetadata: &RequestMetadata{requesterOrganization: "unknown"},
			want:              want{statusCode: http.StatusForbidden, body: "nlx-inway: permission denied, organization \"unknown\" or public key \"\" is not allowed access.\n"},
		},
		{
			name:              "name with wrong cert",
			requesterMetadata: &RequestMetadata{requesterOrganization: "with-name-and-cert", requesterPublicKeyHash: "wrong"},
			want:              want{statusCode: http.StatusForbidden, body: "nlx-inway: permission denied, organization \"with-name-and-cert\" or public key \"wrong\" is not allowed access.\n"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			responseRecorder := httptest.NewRecorder()
			endpoint.handleRequest(test.requesterMetadata, responseRecorder, req)
			statusCode, body := getResponseStatusAndBody(t, responseRecorder)
			assert.Equal(t, test.want.statusCode, statusCode)
			assert.Equal(t, test.want.body, body)
		})
	}
}

func getResponseStatusAndBody(t *testing.T, httpRecorder *httptest.ResponseRecorder) (statusCode int, body string) {
	result := httpRecorder.Result()

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("error parsing result.body", err)
	}

	result.Body.Close()

	return result.StatusCode, string(bytes)
}

func TestInwaySetServiceEndpoints(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	// Certificate organization = nlx-test

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/pki/ca-root.pem",
		OrgCertFile: "../testing/pki/org-nlx-test-chain.pem",
		OrgKeyFile:  "../testing/pki/org-nlx-test-key.pem",
	}

	iw, err := NewInway(logger, nil, testProcess, "", "localhost:1812", tlsOptions, "localhost:1815")
	assert.Nil(t, err)

	serviceDetails := &config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        "12://invalid-endpoint",
			AuthorizationModel: "none",
		},
	}

	// Test NewHTTPServiceEnpoint with invalid url
	_, err = iw.NewHTTPServiceEndpoint("mock-service", serviceDetails, nil)
	assert.EqualError(
		t,
		err,
		"invalid endpoint provided: parse \"12://invalid-endpoint\": first path segment in URL cannot contain colon")

	serviceDetails = &config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        "127.0.0.1",
			AuthorizationModel: "none",
		},
	}

	// Test NewHTTPServiceEndpoint
	endpoint, err := iw.NewHTTPServiceEndpoint("mock-service", serviceDetails, nil)
	assert.Nil(t, err)
	assert.Equal(t, "mock-service", endpoint.ServiceName())

	endpoints := []ServiceEndpoint{
		endpoint,
		endpoint,
	}

	err = iw.SetServiceEndpoints(endpoints)
	if err == nil {
		t.Fatal("result: error is nil, expected error when calling AddServiceEndpoint with a duplicate service")
	}
	assert.EqualError(t, err, "service endpoint for a service with the same name has already been registered")

}

type failingRoundTripper struct{}

func (failingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("some error")
}

// Test if a failing api service results in clear logs about the error
func TestInwayLoggingBadService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	// Certificate organization = nlx-test

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/pki/ca-root.pem",
		OrgCertFile: "../testing/pki/org-nlx-test-chain.pem",
		OrgKeyFile:  "../testing/pki/org-nlx-test-key.pem",
	}

	iw, err := NewInway(logger, nil, testProcess, "", "localhost:1812", tlsOptions, "localhost:1815")
	assert.Nil(t, err)

	serviceDetails := &config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL: "127.0.0.1",
		},
	}

	// Test NewHTTPServiceEndpoint
	endpoint, err := iw.NewHTTPServiceEndpoint(
		"mock-service", serviceDetails, nil)
	endpoint.SetAuthorizationPublic()

	assert.Nil(t, err)
	assert.Equal(t, "mock-service", endpoint.ServiceName())
	// replacing the transport with an always failing one.
	endpoint.proxy.Transport = new(failingRoundTripper)

	httpRecorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/mock-service/", nil)
	req.Header.Add("X-NLX-Logrecord-Id", "dummy-id")

	reqMD := &RequestMetadata{
		requesterOrganization: "demo-org-fault",
	}
	endpoint.handleRequest(reqMD, httpRecorder, req)

	result := httpRecorder.Result()
	defer result.Body.Close()
	bytes, err := ioutil.ReadAll(result.Body)
	t.Log(string(bytes))
	assert.Equal(t, http.StatusServiceUnavailable, result.StatusCode)

	if err != nil {
		t.Fatal("error parsing result.body", err)
	}
	result.Body.Close()

	assert.Equal(
		t,
		"nlx-inway: failed internal API request to 127.0.0.1/ try again later / service api down/unreachable. check A1 error at https://docs.nlx.io/support/common-errors/\n",
		string(bytes),
	)
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
