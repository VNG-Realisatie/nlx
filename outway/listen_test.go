// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"go.nlx.io/nlx/common/monitoring"
	"go.nlx.io/nlx/common/transactionlog"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	mock "go.nlx.io/nlx/outway/mock"
	"go.nlx.io/nlx/outway/plugins"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

var pkiDir = filepath.Join("..", "testing", "pki")

func testRequests(t *testing.T, tests map[string]struct {
	url          string
	statusCode   int
	errorMessage string
}) {
	client := http.Client{}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.url, http.NoBody)
			if err != nil {
				t.Fatal("error creating http request", err)
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal("error making http request", err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			bytes, err := io.ReadAll(resp.Body)

			if err != nil {
				t.Fatal("error reading response body", err)
			}

			assert.Equal(t, tt.errorMessage, string(bytes))
		})
	}
}

//nolint:funlen // test function
func TestOutwayListen(t *testing.T) {
	logger := zap.NewNop()

	outway := &Outway{
		servicesHTTP:      make(map[string]HTTPService),
		servicesDirectory: make(map[string]*directoryapi.ListServicesResponse_Service),
		logger:            logger,
		txlogger:          transactionlog.NewDiscardTransactionLogger(),
	}

	outway.requestHTTPHandler = outway.handleHTTPRequest

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockHTTPService(ctrl)
	mockFailService := mock.NewMockHTTPService(ctrl)

	mockService.EXPECT().ProxyHTTPRequest(gomock.Any(), gomock.Any()).Do(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)

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
		outway.servicesHTTP["00000000000000000001.mockservice"+strconv.Itoa(i)] = mockService
		inwayMessage := directoryapi.ListServicesResponse_Service{
			Name: "mockservice" + strconv.Itoa(i),
			Organization: &directoryapi.Organization{
				SerialNumber: "00000000000000000001",
				Name:         "test-org",
			},
			Inways: []*directoryapi.Inway{
				{
					Address: "mock-service-a-1:123",
					State:   directoryapi.Inway_UP,
				},
			},
		}
		outway.servicesDirectory["00000000000000000001.mockservice"+strconv.Itoa(i)] = &inwayMessage
	}

	// Setup a Failing mock service.
	outway.servicesHTTP["00000000000000000001.mockservicefail"] = mockFailService
	inwayMessage := directoryapi.ListServicesResponse_Service{
		Name: "mockservicefail",
		Organization: &directoryapi.Organization{
			SerialNumber: "00000000000000000001",
			Name:         "test-org",
		},
		Inways: []*directoryapi.Inway{
			{
				Address: "mock-service-fail-1:123",
				State:   directoryapi.Inway_UP,
			},
		},
	}
	outway.servicesDirectory["00000000000000000001.mockservicefail"] = &inwayMessage

	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	// Test http responses
	tests := map[string]struct {
		url          string
		statusCode   int
		errorMessage string
	}{
		"when_invalid_path": {
			fmt.Sprintf("%s/invalidpath", mockServer.URL),
			http.StatusBadRequest,
			"nlx outway: invalid /serialNumber/service/ url: valid organization serial numbers : [00000000000000000001]\n",
		},
		"when_service_not_found": {
			fmt.Sprintf("%s/00000000000000000001/nonexistingservice/add/", mockServer.URL),
			http.StatusBadRequest,
			"nlx outway: invalid serialNumber/service path: valid services : [mockservice0, mockservice1, mockservice10, mockservice2, mockservice3, mockservice4, mockservice5, mockservice6, mockservice7, mockservice8, mockservice9, mockservicefail]\n",
		},
		"when_service_fails": {
			fmt.Sprintf("%s/00000000000000000001/mockservicefail/", mockServer.URL),
			http.StatusInternalServerError,
			"",
		},
		"happy_flow": {
			fmt.Sprintf("%s/00000000000000000001/mockservice0/", mockServer.URL),
			http.StatusOK,
			"",
		},
		"happy_flow_without_trailing_slash": {
			fmt.Sprintf("%s/00000000000000000001/mockservice0", mockServer.URL),
			http.StatusOK,
			"",
		},
	}

	testRequests(t, tests)
}

func createMockOutway() *Outway {
	return &Outway{
		servicesHTTP:      make(map[string]HTTPService),
		servicesDirectory: make(map[string]*directoryapi.ListServicesResponse_Service),
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

	outway.servicesHTTP["00000000000000000001.mockservice"] = mockService

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
			"http://mockservice.00000000000000000001.services.nlx.local",
			http.StatusOK,
			"",
			"",
			outway.handleHTTPRequestAsProxy,
		}, {
			"request invalid url",
			"http://invalid.mockservice.00000000000000000001.services.nlx.local",
			http.StatusBadRequest,
			"nlx outway: no valid url expecting: service.serialNumber.service.nlx.local/apipath\n",
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
			"http://mockservice.00000000000000000001.services.nlx.local",
			http.StatusInternalServerError,
			"please enable proxy mode by setting the 'use-as-http-proxy' flag to resolve: http://mockservice.00000000000000000001.services.nlx.local/\n",
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

	req := httptest.NewRequest(http.MethodConnect, "http://mockservice.00000000000000000001.services.nlx.local", nil)
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
	outway.servicesHTTP["00000000000000000001.mockservice"] = mockService

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
				plugins.NewLogRecordPlugin("00000000000000000001", tt.txLogger),
				plugins.NewStripHeadersPlugin("00000000000000000001"),
			}

			if tt.authEnabled {
				outway.plugins = append([]plugins.Plugin{
					plugins.NewAuthorizationPlugin(nil, "", http.Client{}),
				}, outway.plugins...)
			}

			outway.txlogger = tt.txLogger

			req := httptest.NewRequest("GET", "http://mockservice.00000000000000000001.services.nlx.local", nil)
			req.Header.Add("X-NLX-Request-Data-Subject", tt.dataSubjectHeader)

			outway.handleOnNLX(outway.logger, &plugins.Destination{
				OrganizationSerialNumber: "00000000000000000001",
				Service:                  "mockservice",
				Path:                     "/",
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
		servicesDirectory: make(map[string]*directoryapi.ListServicesResponse_Service),
		logger:            logger,
		txlogger:          transactionlog.NewDiscardTransactionLogger(),
	}

	outway.requestHTTPHandler = outway.handleHTTPRequest

	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	tests := map[string]struct {
		url          string
		statusCode   int
		errorMessage string
	}{
		"when_request_to_inway_fails": {
			fmt.Sprintf("%s/00000000000000000001/mockservice/", mockServer.URL),
			http.StatusServiceUnavailable,
			"failed request to 'https://inway.00000000000000000001/mockservice/', try again later and check your firewall, check O1 and M1 at https://docs.nlx.io/support/common-errors/\n",
		},
	}

	inwayMessage := directoryapi.ListServicesResponse_Service{
		Name: "mockservice",
		Organization: &directoryapi.Organization{
			SerialNumber: "00000000000000000001",
			Name:         "test-org",
		},
		Inways: []*directoryapi.Inway{
			{
				Address: "mock-service-:123",
				State:   directoryapi.Inway_UP,
			},
		},
	}

	// Setup mock httpservice
	outway.servicesDirectory["00000000000000000001.mockservice"] = &inwayMessage

	certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	l, err := NewRoundRobinLoadBalancedHTTPService(
		zap.NewNop(), certBundle,
		"00000000000000000001", "mockservice",
		[]directoryapi.Inway{{
			Address: "inway.00000000000000000001",
			State:   directoryapi.Inway_UP,
		}})

	assert.Nil(t, err)

	outway.servicesHTTP["00000000000000000001.mockservice"] = l
	// set transports to fail.
	outway.setFailingTransport()
	testRequests(t, tests)
}

func TestRunServer(t *testing.T) {
	t.Parallel()

	logger := zap.NewNop()

	certificate, _ := tls.LoadX509KeyPair(
		filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
		filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
	)

	tests := map[string]struct {
		listenAddress            string
		listenAddressGRPC        string
		monitoringServiceAddress string
		certificate              *tls.Certificate
		errorMessage             string
	}{
		"invalid_listen_address": {
			listenAddress:            "invalid",
			listenAddressGRPC:        "127.0.0.1:8082",
			monitoringServiceAddress: "localhost:8081",
			certificate:              nil,
			errorMessage:             "error listening on server: listen tcp: address invalid: missing port in address",
		},
		"invalid_listen_address with TLS": {
			listenAddress:            "invalid",
			listenAddressGRPC:        "127.0.0.1:8083",
			monitoringServiceAddress: "localhost:8082",
			certificate:              &certificate,
			errorMessage:             "error listening on server: listen tcp: address invalid: missing port in address",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			monitorService, err := monitoring.NewMonitoringService(tt.monitoringServiceAddress, logger)
			assert.Nil(t, err)

			o := &Outway{
				ctx:            context.Background(),
				logger:         logger,
				monitorService: monitorService,
				grpcServer:     grpc.NewServer(),
			}

			err = o.RunServer(tt.listenAddress, tt.listenAddressGRPC, tt.certificate)
			assert.EqualError(t, err, tt.errorMessage)

			err = monitorService.Stop()
			assert.NoError(t, err)
		})
	}
}
