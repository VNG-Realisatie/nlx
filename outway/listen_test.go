// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/monitoring"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	mock "go.nlx.io/nlx/outway/mock"
	"go.nlx.io/nlx/outway/plugins"
)

var pkiDir = filepath.Join("..", "testing", "pki")

// testRequests to check for expected reponses.
func testRequests(t *testing.T, tests []struct {
	url          string
	statusCode   int
	errorMessage string
}) {
	client := http.Client{}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal("error creating http request", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("error doing http request", err)
		}
		defer resp.Body.Close()

		assert.Equal(t, test.statusCode, resp.StatusCode)
		bytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			t.Fatal("error parsing result.body", err)
		}

		assert.Equal(t, test.errorMessage, string(bytes))
	}
}

func TestOutwayListen(t *testing.T) {
	logger := zap.NewNop()

	// Create a outway with a mock service
	outway := &Outway{
		servicesHTTP:      make(map[string]HTTPService),
		servicesDirectory: make(map[string]*inspectionapi.ListServicesResponse_Service),
		logger:            logger,
		txlogger:          transactionlog.NewDiscardTransactionLogger(),
	}

	outway.requestHTTPHandler = outway.handleHTTPRequest

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockHTTPService(ctrl)
	mockFailService := mock.NewMockHTTPService(ctrl)

	mockService.EXPECT().ProxyHTTPRequest(gomock.Any(), gomock.Any()).Do(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)

	mockFailService.EXPECT().ProxyHTTPRequest(gomock.Any(), gomock.Any()).Do(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	)

	for i := 0; i < 11; i++ {
		outway.servicesHTTP["mockorg.mockservice"+strconv.Itoa(i)] = mockService
		inwayMessage := inspectionapi.ListServicesResponse_Service{
			ServiceName:      "mockservice" + strconv.Itoa(i),
			OrganizationName: "mockorg",
			InwayAddresses:   []string{"mock-service-a-1:123"},
			HealthyStates:    []bool{true},
		}
		outway.servicesDirectory["mockorg.mockservice"+strconv.Itoa(i)] = &inwayMessage
	}

	// Setup a Failing mock service.
	outway.servicesHTTP["mockorg.mockservicefail"] = mockFailService
	inwayMessage := inspectionapi.ListServicesResponse_Service{
		ServiceName:      "mockservicefail",
		OrganizationName: "mockorg",
		InwayAddresses:   []string{"mock-service-fail-1:123"},
		HealthyStates:    []bool{true},
	}
	outway.servicesDirectory["mockorg.mockservicefail"] = &inwayMessage

	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	// Test http responses
	tests := []struct {
		url          string
		statusCode   int
		errorMessage string
	}{
		{
			fmt.Sprintf("%s/invalidpath", mockServer.URL),
			http.StatusBadRequest,
			"nlx outway: invalid /organization/service/ url: valid organizations : [mockorg]\n",
		}, {
			fmt.Sprintf("%s/mockorg/nonexistingservice/add/", mockServer.URL),
			http.StatusBadRequest,
			"nlx outway: invalid organization/service path: valid services : [mockservice0, mockservice1, mockservice10, mockservice2, mockservice3, mockservice4, mockservice5, mockservice6, mockservice7, mockservice8, mockservice9, mockservicefail]\n",
		}, {
			fmt.Sprintf("%s/mockorg/mockservice0/", mockServer.URL),
			http.StatusOK,
			"",
		}, {
			fmt.Sprintf("%s/mockorg/mockservicefail/", mockServer.URL),
			http.StatusInternalServerError,
			"",
		},
	}

	testRequests(t, tests)
}

func createMockOutway() *Outway {
	return &Outway{
		servicesHTTP:      make(map[string]HTTPService),
		servicesDirectory: make(map[string]*inspectionapi.ListServicesResponse_Service),
		logger:            zap.NewNop(),
		txlogger:          transactionlog.NewDiscardTransactionLogger(),
	}
}

func TestOutwayAsProxy(t *testing.T) {
	// Create a outway with a mock service
	outway := createMockOutway()
	outway.forwardingHTTPProxy = newForwardingProxy()

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockHTTPService(ctrl)
	mockService.EXPECT().ProxyHTTPRequest(gomock.Any(), gomock.Any()).Do(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)

	outway.servicesHTTP["mockorg.mockservice"] = mockService

	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	mockPublicServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer mockPublicServer.Close()

	// Test http responses
	tests := []struct {
		description       string
		url               string
		statusCode        int
		errorMessage      string
		dataSubjectHeader string
		httpHandler       loggerHTTPHandler
	}{
		{
			"request using a services.nlx.local URL",
			"http://mockservice.mockorg.services.nlx.local",
			http.StatusOK,
			"",
			"",
			outway.handleHTTPRequestAsProxy,
		}, {
			"request invalid url",
			"http://invalid.mockservice.mockorg.services.nlx.local",
			http.StatusBadRequest,
			"nlx outway: no valid url expecting: service.organization.service.nlx.local/apipath\n",
			"",
			outway.handleHTTPRequestAsProxy,
		}, {
			"request to public internet",
			mockPublicServer.URL,
			http.StatusOK,
			"",
			"",
			outway.handleHTTPRequestAsProxy,
		}, {
			"outway is running without the use-as-http-proxy flag",
			"http://mockservice.mockorg.services.nlx.local",
			http.StatusInternalServerError,
			"please enable proxy mode by setting the 'use-as-http-proxy' flag to resolve: http://mockservice.mockorg.services.nlx.local/\n",
			"",
			outway.handleHTTPRequest,
		},
	}

	outwayURL, err := url.Parse(mockServer.URL)
	assert.Nil(t, err)

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(outwayURL)}}

	for _, test := range tests {
		outway.requestHTTPHandler = test.httpHandler

		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal("error creating http request", err)
		}

		req.Header.Add("X-NLX-Request-Data-Subject", test.dataSubjectHeader)
		resp, err := client.Do(req)

		if err != nil {
			t.Fatal("error doing http request", err)
		}
		defer resp.Body.Close()

		assert.Equal(t, test.statusCode, resp.StatusCode)

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error parsing result.body", err)
		}

		assert.Equal(t, test.errorMessage, string(bytes))
	}
}

func TestHandleConnectMethodException(t *testing.T) {
	logger := zap.NewNop()
	outway := &Outway{}
	outway.forwardingHTTPProxy = newForwardingProxy()

	recorder := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodConnect, "http://mockservice.mockorg.services.nlx.local", nil)
	outway.handleHTTPRequestAsProxy(logger, recorder, req)

	assert.Equal(t, http.StatusNotImplemented, recorder.Code)

	bytes, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal("error parsing result.body", err)
	}

	assert.Equal(t, "CONNECT method is not supported\n", string(bytes))
}

type failingTransactionLogger struct {
}

func (f *failingTransactionLogger) AddRecord(record *transactionlog.Record) error {
	return errors.New("cannot add transaction record")
}

func TestHandleOnNLXExceptions(t *testing.T) {
	outway := createMockOutway()

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockHTTPService(ctrl)
	outway.servicesHTTP["mockorg.mockservice"] = mockService

	tests := map[string]struct {
		authEnabled          bool
		txLogger             transactionlog.TransactionLogger
		dataSubjectHeader    string
		expectedStatusCode   int
		exectpedErrorMessage string
	}{
		"with failing auth settings": {
			true,
			&transactionlog.DiscardTransactionLogger{},
			"",
			http.StatusInternalServerError,
			"nlx outway: error authorizing request\n",
		},
		"with failing transactionlogger": {
			false,
			&failingTransactionLogger{},
			"",
			http.StatusInternalServerError,
			"nlx outway: server error\n",
		},
		"with invalid datasubject header": {
			false,
			&transactionlog.DiscardTransactionLogger{},
			"invalid",
			http.StatusBadRequest,
			"nlx outway: invalid data subject header\n",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			outway.plugins = []plugins.Plugin{
				plugins.NewDelegationPlugin(nil),
				plugins.NewLogRecordPlugin("TestOrg", tt.txLogger),
				plugins.NewStripHeadersPlugin("TestOrg"),
			}

			if tt.authEnabled {
				outway.plugins = append([]plugins.Plugin{
					plugins.NewAuthorizationPlugin(nil, "", http.Client{}),
				}, outway.plugins...)
			}

			outway.txlogger = tt.txLogger

			req := httptest.NewRequest("GET", "http://mockservice.mockorg.services.nlx.local", nil)
			req.Header.Add("X-NLX-Request-Data-Subject", tt.dataSubjectHeader)

			outway.handleOnNLX(outway.logger, &plugins.Destination{
				Organization: "mockorg",
				Service:      "mockservice",
				Path:         "/",
			}, recorder, req)

			assert.Equal(t, tt.expectedStatusCode, recorder.Code)

			bytes, err := ioutil.ReadAll(recorder.Body)
			if err != nil {
				t.Fatal("error parsing result.body", err)
			}

			assert.Equal(t, tt.exectpedErrorMessage, string(bytes))
		})
	}
}

type failingRoundTripper struct{}

func (failingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("some error")
}

func (o *Outway) setFailingTransport() {
	// Change connection Transport to Failing Transports.
	// for all proxies
	for _, s := range o.servicesHTTP {
		if rrlbs, ok := s.(*RoundRobinLoadBalancedHTTPService); ok {
			for _, p := range rrlbs.proxies {
				p.Transport = new(failingRoundTripper)
			}
		}
	}
}

// TestFailingTransport tests the error handling when there are
// network problems to reach the advertised service from the outway
//
// client -> outway -> [FAIL] inway -> service
//
// The test creates a service with failing transport.
// and expecting a 503 service temporarily unavailable status code
// when service gets called
func TestFailingTransport(t *testing.T) {
	logger := zap.NewNop()
	// during tests: logger, _ := zap.NewDevelopment()
	// defer logger.Sync()

	// Create a outway with a mock service
	outway := &Outway{
		servicesHTTP:      make(map[string]HTTPService),
		servicesDirectory: make(map[string]*inspectionapi.ListServicesResponse_Service),
		logger:            logger,
		txlogger:          transactionlog.NewDiscardTransactionLogger(),
	}

	outway.requestHTTPHandler = outway.handleHTTPRequest

	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	tests := []struct {
		url          string
		statusCode   int
		errorMessage string
	}{
		{
			fmt.Sprintf("%s/mockorg/mockservice/", mockServer.URL),
			http.StatusServiceUnavailable,
			"failed request to https://inway.mockorg/mockservice/ try again later / check firewall? check O1 and M1 at https://docs.nlx.io/support/common-errors/\n",
		},
	}

	inwayMessage := inspectionapi.ListServicesResponse_Service{
		ServiceName:      "mockservice",
		OrganizationName: "mockorg",
		InwayAddresses:   []string{"mock-service-:123"},
		HealthyStates:    []bool{true},
	}

	// Setup mock httpservice
	outway.servicesDirectory["mockorg.mockservice"] = &inwayMessage

	inwayAddresses := []string{"inway.mockorg"}
	healthyStates := []bool{true}

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	l, err := NewRoundRobinLoadBalancedHTTPService(
		zap.NewNop(), cert,
		"mockorg", "mockservice",
		inwayAddresses, healthyStates)

	assert.Nil(t, err)

	outway.servicesHTTP["mockorg.mockservice"] = l
	// set transports to fail.
	outway.setFailingTransport()
	testRequests(t, tests)
}

func TestRunServer(t *testing.T) {
	logger := zap.NewNop()

	certificate, _ := tls.LoadX509KeyPair(
		filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
		filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
	)

	tests := []struct {
		description   string
		listenAddress string
		certificate   *tls.Certificate
		errorMessage  string
	}{
		{
			"invalid listen address",
			"invalid",
			nil,
			"error listening on server: listen tcp: address invalid: missing port in address",
		},
		{
			"invalid listen address with TLS",
			"invalid",
			&certificate,
			"error listening on server: listen tcp: address invalid: missing port in address",
		},
	}

	monitorService, err := monitoring.NewMonitoringService("localhost:8081", logger)
	assert.Nil(t, err)

	testProcess := process.NewProcess(logger)

	for _, test := range tests {
		o := &Outway{
			logger:         logger,
			monitorService: monitorService,
			process:        testProcess,
		}

		err = o.RunServer(test.listenAddress, test.certificate)
		assert.EqualError(t, err, test.errorMessage)
	}
}
