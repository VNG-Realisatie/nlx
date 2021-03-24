// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"go.uber.org/zap/zaptest/observer"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
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
			name:              "missing requesterOrganization",
			requesterMetadata: &RequestMetadata{},
			want:              want{statusCode: http.StatusBadRequest, body: "nlx-inway: could not handle your request, missing requesterOrganization header.\n"},
		},
		{
			name:              "only certificate",
			requesterMetadata: &RequestMetadata{requesterOrganization: "irrelevant", requesterPublicKeyFingerprint: "only-cert"},
			want:              want{statusCode: http.StatusBadRequest, body: "nlx-inway: missing logrecord id\n"},
		},
		{
			name:              "only organization",
			requesterMetadata: &RequestMetadata{requesterOrganization: "only-org"},
			want:              want{statusCode: http.StatusBadRequest, body: "nlx-inway: missing logrecord id\n"},
		},
		{
			name:              "with name and cert",
			requesterMetadata: &RequestMetadata{requesterOrganization: "with-name-and-cert", requesterPublicKeyFingerprint: "with-cert-and-name"},
			want:              want{statusCode: http.StatusBadRequest, body: "nlx-inway: missing logrecord id\n"},
		},
		{
			name:              "unknown",
			requesterMetadata: &RequestMetadata{requesterOrganization: "unknown"},
			want:              want{statusCode: http.StatusForbidden, body: "nlx-inway: permission denied, organization \"unknown\" or public key \"\" is not allowed access.\n"},
		},
		{
			name:              "name with wrong cert",
			requesterMetadata: &RequestMetadata{requesterOrganization: "with-name-and-cert", requesterPublicKeyFingerprint: "wrong"},
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

//nolint:funlen // this is a test
func TestInwaySetServiceEndpoints(t *testing.T) {
	// Certificate organization = nlx-test
	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	type args struct {
		serviceName    string
		serviceDetails *config.ServiceDetails
	}

	type validatorState struct {
		t        *testing.T
		args     args
		endpoint *HTTPServiceEndpoint
		err      error
		recorded *observer.ObservedLogs
	}

	type want struct {
		validator func(state validatorState)
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "authorization model none",
			args: args{
				serviceName: "public-service",
				serviceDetails: &config.ServiceDetails{
					ServiceDetailsBase: config.ServiceDetailsBase{
						EndpointURL:        "127.0.0.1",
						AuthorizationModel: "none",
					},
				},
			},
			want: want{
				validator: func(state validatorState) {
					assert.NoError(t, state.err)
					assert.NotNil(t, state.endpoint)
					assert.Equal(t, state.args.serviceName, state.endpoint.ServiceName())
				},
			},
		},
		{
			name: "authorization model invalid",
			args: args{
				serviceName: "public-service",
				serviceDetails: &config.ServiceDetails{
					ServiceDetailsBase: config.ServiceDetailsBase{
						EndpointURL:        "127.0.0.1",
						AuthorizationModel: "invalid",
					},
				},
			},
			want: want{
				validator: func(state validatorState) {
					assert.Len(t, state.recorded.FilterMessageSnippet("invalid authorization model").All(), 1)
				},
			},
		},
		{
			name: "invalid EndpointURL",
			args: args{
				serviceName: "invalid-service",
				serviceDetails: &config.ServiceDetails{
					ServiceDetailsBase: config.ServiceDetailsBase{
						EndpointURL:        "12://invalid-endpoint",
						AuthorizationModel: "none",
					},
				},
			},
			want: want{
				validator: func(state validatorState) {
					assert.EqualError(t, state.err, "invalid endpoint provided: parse \"12://invalid-endpoint\": first path segment in URL cannot contain colon")
					assert.Nil(t, state.endpoint)
				},
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			core, recorded := observer.New(zapcore.InfoLevel)
			logger := zap.New(core)
			testProcess := process.NewProcess(logger)
			iw, err := NewInway(logger, nil, testProcess, "", "localhost:1812", "localhost:1813", cert, "localhost:1815")
			assert.NoError(t, err)

			endpoint, err := iw.NewHTTPServiceEndpoint(test.args.serviceName, test.args.serviceDetails, nil)

			test.want.validator(validatorState{
				t:        t,
				args:     test.args,
				endpoint: endpoint,
				err:      err,
				recorded: recorded,
			})
		})
	}
}

func TestInwaySetServiceEnpointDuplicateEndpoint(t *testing.T) {
	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	logger := zaptest.NewLogger(t)
	testProcess := process.NewProcess(logger)
	iw, err := NewInway(logger, nil, testProcess, "", "localhost:1812", "localhost:1813", cert, "localhost:1815")
	assert.NoError(t, err)

	serviceDetails := &config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL:        "127.0.0.1",
			AuthorizationModel: "none",
		},
	}

	endpoint, err := iw.NewHTTPServiceEndpoint("no-duplicates-service", serviceDetails, nil)
	assert.NoError(t, err)

	endpoints := []ServiceEndpoint{
		endpoint,
		endpoint,
	}
	err = iw.SetServiceEndpoints(endpoints)
	assert.EqualError(t, err, "service endpoint for a service with the same name has already been registered", "expected error when calling SetServiceEndpoints with a duplicate service")
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

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	iw, err := NewInway(logger, nil, testProcess, "", "localhost:1812", "localhost:1813", cert, "localhost:1815")
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
