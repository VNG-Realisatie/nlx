package inway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway/config"
	"go.uber.org/zap"
)

func TestSetAuthorization(t *testing.T) {
	endpoint := &HTTPServiceEndpoint{}
	endpoint.SetAuthorizationPublic()
	if endpoint.public != true {
		t.Fatalf("result: %v, expected HttpServiceEndpoint to be true", endpoint.public)
	}

	whiteList := []string{"demo-org"}
	endpoint.SetAuthorizationWhitelist(whiteList)
	if endpoint.public != false {
		t.Fatalf("result: %v, expected HttpServiceEndpoint to be false", endpoint.public)
	}

	if len(endpoint.whitelistedOrganizations) != 1 {
		t.Fatalf("result: %d, expected HttpServiceEndpoint.whitelistedOrganizations to have a length of 1", len(endpoint.whitelistedOrganizations))
	}

	if strings.Compare("demo-org", endpoint.whitelistedOrganizations[0]) != 0 {
		t.Fatalf("demo-org not present in endpoint.whitelistedOrganizations")
	}

	var err error
	endpoint.logger = zap.NewNop()
	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "localhost/whitelist", nil)
	if err != nil {
		t.Fatal("error creating request", err)
	}
	reqMD := &RequestMetadata{
		requesterOrganization: "demo-org-fault",
	}
	endpoint.handleRequest(reqMD, httpRecorder, req)

	result := httpRecorder.Result()
	if result.StatusCode != http.StatusForbidden {
		t.Fatalf("result: %d, expected http status code: %d", result.StatusCode, http.StatusForbidden)
	}

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("error parsing result.body", err)
	}

	resultBody := strings.Trim(string(bytes[:len(bytes)]), "\n")
	expected := fmt.Sprintf(`nlx outway: could not handle your request, organization "%s" is not allowed access.`, reqMD.requesterOrganization)
	if strings.Compare(resultBody, expected) != 0 {
		t.Fatalf("result: %s, expected: %s", resultBody, expected)
	}

}

func TestInwayAddServiceEndpoint(t *testing.T) {
	logger := zap.NewNop()

	// Certificate organisation = nlx-test
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org-nlx-test.key"}

	iw, err := NewInway(logger, nil, "localhost:1812", tlsOptions,
		"localhost:1815", nil)
	if err != nil {
		t.Fatal(err)
	}

	p := process.NewProcess(logger)

	endpoint, err := iw.NewHTTPServiceEndpoint(logger, "mock-service", "12://invalid-endpoint", nil)
	if err == nil {
		t.Fatal("no error when adding a service with invalid endpoint")
	}

	if !strings.HasPrefix(err.Error(), "invalid endpoint provided") {
		t.Fatalf("result: %s, expected error message to start with: invalid endpoint provided", err.Error())
	}

	endpoint, err = iw.NewHTTPServiceEndpoint(logger, "mock-service", "127.0.0.1", nil)
	if err != nil {
		t.Fatal("failed to create service", err)
	}

	if strings.Compare(endpoint.ServiceName(), "mock-service") != 0 {
		t.Fatalf("result %s, expected servicename : mock-service", endpoint.ServiceName())
	}

	err = iw.AddServiceEndpoint(p, endpoint, config.ServiceDetails{
		EndpointURL:            "http://127.0.0.1:1813",
		AuthorizationWhitelist: []string{"nlx-forbidden"}})

	if err != nil {
		t.Fatal("failed to add service endpoint", err)
	}

	err = iw.AddServiceEndpoint(p, endpoint, config.ServiceDetails{
		EndpointURL:            "http://127.0.0.1:1813",
		AuthorizationWhitelist: []string{"nlx-forbidden"}})
	if err == nil {
		t.Fatal("result: error is nil, expected error when calling AddServiceEndpoint with a dublicate service")
	}

	if strings.Compare(err.Error(), "service endpoint for a service with the same name has already been registered") != 0 {
		t.Fatalf("result: %s, expected error message: service endpoint for a service with the same name has already been registered", err.Error())
	}

}
