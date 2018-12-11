package inway_test

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway"
	"go.nlx.io/nlx/inway/config"
)

func TestSetupInway(t *testing.T) {
	logger := zap.NewNop()

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org-nlx-test.key"}

	serviceConfig := &config.ServiceConfig{}
	serviceConfig.Services = make(map[string]config.ServiceDetails)
	serviceConfig.Services["mock-service-whitelist"] = config.ServiceDetails{
		EndpointURL:            "http://localhost:1813",
		AuthorizationWhitelist: []string{"nlx-test"},
		AuthorizationModel:     "whitelist",
	}
	serviceConfig.Services["mock-servicewhitelist-unauthorized"] = config.ServiceDetails{
		EndpointURL:            "http://localhost:1813",
		AuthorizationWhitelist: []string{"nlx-forbidden"},
		AuthorizationModel:     "whitelist",
	}
	serviceConfig.Services["mock-service-public"] = config.ServiceDetails{
		EndpointURL:        "http://localhost:1813",
		AuthorizationModel: "none",
	}
	serviceConfig.Services["mock-service-public-apispec"] = config.ServiceDetails{
		EndpointURL:                 "http://localhost:1813",
		AuthorizationModel:          "none",
		APISpecificationDocumentURL: "http://localhost:1813/dummy",
	}
	serviceConfig.Services["mock-service-public-invalid-apispec"] = config.ServiceDetails{
		EndpointURL:                 "http://localhost:1813",
		AuthorizationModel:          "none",
		APISpecificationDocumentURL: "invalid",
	}

	iw, err := inway.NewInway(logger, nil, "localhost:1812", tlsOptions,
		"localhost:1815", serviceConfig)
	if err != nil {
		t.Fatal("error creating inway", err)
	}

	p := process.NewProcess(logger)

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
		if err != nil {
			t.Fatal("error adding endpoint", err)
		}
	}

	go func() {
		err = iw.ListenAndServeTLS(p, "localhost:1812")
		if err != nil {
			t.Fatal("error listening", err)
		}
	}()

	server := http.NewServeMux()
	server.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	go func() {
		if err := http.ListenAndServe("localhost:1813", server); err != nil {
			t.Fatal("error serving dummy endpoint", err)
		}
	}()
}

func TestNewInwayException(t *testing.T) {
	logger := zap.NewNop()
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org_without_name.crt",
		OrgKeyFile:  "../testing/org_without_name.key"}

	serviceConfig := &config.ServiceConfig{}
	_, err := inway.NewInway(logger, nil, "", tlsOptions, "", serviceConfig)
	if err == nil {
		t.Fatal("error is nil but should be set")
	}

	if strings.Compare(err.Error(), "cannot obtain organization name from self cert") != 0 {
		t.Fatalf(`result: %s, expected error to be "cannot obtain organization name from self cert"`, err.Error())
	}

	tlsOptions = orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org_non_existing.key"}

	_, err = inway.NewInway(logger, nil, "", tlsOptions, "", serviceConfig)
	if err == nil {
		t.Fatal(`result: error is nil, expected error to be set when call NewInway without a existing .key file`)
	}

	if !strings.HasPrefix(err.Error(), "failed to read tls keypair") {
		t.Fatalf(`result: %s, expected error to start with "failed to read tls keypair"`, err.Error())
	}

	tlsOptions = orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org-nlx-test.key"}

	inway, err := inway.NewInway(logger, nil, "", tlsOptions, "", serviceConfig)
	if err != nil {
		t.Fatal(err)
	}
	err = inway.ListenAndServeTLS(process.NewProcess(logger), "invalidlistenaddress")
	if err == nil {
		t.Fatal(`result: error is nil, expected error to be set when calling ListenAndServeTLS with invalid listen adress`)
	}

	if !strings.HasPrefix(err.Error(), "failed to run http server") {
		t.Fatalf(`result: %s, expected error to start with "failed to run http server"`, err.Error())
	}
}

func TestInWay(t *testing.T) {
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org-nlx-test.key"}
	pool, err := orgtls.LoadRootCert(tlsOptions.NLXRootCert)
	if err != nil {
		t.Fatal("error loading root certificate", err)
	}

	cert, err := tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	if err != nil {
		t.Fatal(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: pool, Certificates: []tls.Certificate{cert}}}
	client := http.Client{
		Transport: tr,
	}

	tests := []struct {
		url          string
		logRecordId  string
		statusCode   int
		errorMessage string
	}{
		{"https://localhost:1812/mock-service-public/dummy", "dummy-ID", http.StatusOK, ""},
		{"https://localhost:1812/mock-service-whitelist/dummy", "dummy-ID", http.StatusOK, ""},
		{"https://localhost:1812/mock-servicewhitelist-unauthorized/dummy", "dummy-ID", http.StatusForbidden, `nlx outway: could not handle your request, organization "nlx-test" is not allowed access.`},
		{"https://localhost:1812/mock-service", "dummy-ID", http.StatusBadRequest, `nlx inway error: invalid path in url`},
		{"https://localhost:1812/mock-service/fictive", "dummy-ID", http.StatusBadRequest, `nlx inway error: no endpoint for service`},
		{"https://localhost:1812/mock-service-public/dummy", "", http.StatusBadRequest, `nlx outway: missing logrecord id`},
		{"https://localhost:1812/.nlx/api-spec-doc/mock-service-public", "dummy-ID", http.StatusNotFound, `api specification not found for service`},
		{"https://localhost:1812/.nlx/api-spec-doc/nonexisting-service", "dummy-ID", http.StatusNotFound, `service not found`},
		{"https://localhost:1812/.nlx/api-spec-doc/mock-service-public-invalid-apispec", "dummy-ID", http.StatusInternalServerError, `server error`},
		{"https://localhost:1812/.nlx/api-spec-doc/mock-service-public-apispec", "dummy-ID", http.StatusOK, ``},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("X-NLX-Logrecord-Id", test.logRecordId)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("error doing http request", err)
		}

		if resp.StatusCode != test.statusCode {
			t.Fatalf("result: %d, expected http status code :%d", resp.StatusCode, test.statusCode)
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error parsing result.body", err)
		}

		resultBody := strings.Trim(string(bytes[:len(bytes)]), "\n")
		if strings.Compare(resultBody, test.errorMessage) != 0 {
			t.Fatalf("result: %s, expected http body: %s", resultBody, test.errorMessage)
		}
	}
}
