// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	"go.nlx.io/nlx/common/httperrors"
	"go.nlx.io/nlx/common/monitoring"
	"go.nlx.io/nlx/common/transactionlog"
	mock_transactionlog "go.nlx.io/nlx/common/transactionlog/mock"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	mock "go.nlx.io/nlx/outway/mock"
	"go.nlx.io/nlx/outway/plugins"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

var pkiDir = filepath.Join("..", "testing", "pki")

func testRequests(t *testing.T, tests map[string]struct {
	url        string
	statusCode int
	wantErr    *httperrors.NLXNetworkError
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
			contents, err := io.ReadAll(resp.Body)

			if err != nil {
				t.Fatal("error reading response body", err)
			}

			if tt.wantErr != nil {
				gotError := &httperrors.NLXNetworkError{}
				err := json.Unmarshal(contents, gotError)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantErr, gotError)
			}
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
			w.WriteHeader(httperrors.StatusNLXNetworkError)
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
					State:   directoryapi.Inway_STATE_UP,
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
				State:   directoryapi.Inway_STATE_UP,
			},
		},
	}
	outway.servicesDirectory["00000000000000000001.mockservicefail"] = &inwayMessage

	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	// Test http responses
	tests := map[string]struct {
		url        string
		statusCode int
		wantErr    *httperrors.NLXNetworkError
	}{
		"when_invalid_path": {
			url:        fmt.Sprintf("%s/invalidpath", mockServer.URL),
			statusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.C1,
				Code:     httperrors.InvalidURLErr,
				Message:  "invalid /serialNumber/service/ url: valid organization serial numbers : [00000000000000000001]",
			},
		},
		"when_service_not_found": {
			url:        fmt.Sprintf("%s/00000000000000000001/nonexistingservice/add/", mockServer.URL),
			statusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.C1,
				Code:     httperrors.InvalidURLErr,
				Message:  "invalid serialNumber/service path: valid services : [mockservice0, mockservice1, mockservice10, mockservice2, mockservice3, mockservice4, mockservice5, mockservice6, mockservice7, mockservice8, mockservice9, mockservicefail]",
			},
		},
		"when_service_fails": {
			url:        fmt.Sprintf("%s/00000000000000000001/mockservicefail/", mockServer.URL),
			statusCode: httperrors.StatusNLXNetworkError,
		},
		"happy_flow": {
			url:        fmt.Sprintf("%s/00000000000000000001/mockservice0/", mockServer.URL),
			statusCode: http.StatusOK,
		},
		"happy_flow_without_trailing_slash": {
			url:        fmt.Sprintf("%s/00000000000000000001/mockservice0", mockServer.URL),
			statusCode: http.StatusOK,
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
	tests := map[string]struct {
		url               string
		statusCode        int
		wantErr           *httperrors.NLXNetworkError
		dataSubjectHeader string
		httpHandler       loggerHTTPHandler
	}{
		"request using a services.nlx.local URL": {
			url:               "http://mockservice.00000000000000000001.services.nlx.local",
			statusCode:        http.StatusOK,
			dataSubjectHeader: "",
			httpHandler:       outway.handleHTTPRequestAsProxy,
		},
		"request invalid url": {
			url:        "http://invalid.mockservice.00000000000000000001.services.nlx.local",
			statusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.C1,
				Code:     httperrors.InvalidURLErr,
				Message:  "no valid url expecting: service.serialNumber.service.nlx.local/apipath",
			},
			dataSubjectHeader: "",
			httpHandler:       outway.handleHTTPRequestAsProxy,
		},
		"request to public internet": {
			url:               mockPublicServer.URL,
			statusCode:        http.StatusOK,
			dataSubjectHeader: "",
			httpHandler:       outway.handleHTTPRequestAsProxy,
		},
		"outway is running without the use-as-http-proxy flag": {
			url:        "http://mockservice.00000000000000000001.services.nlx.local",
			statusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.C1,
				Code:     httperrors.ProxyModeDisabledErr,
				Message:  "please enable proxy mode by setting the 'use-as-http-proxy' flag to resolve: http://mockservice.00000000000000000001.services.nlx.local/",
			},
			dataSubjectHeader: "",
			httpHandler:       outway.handleHTTPRequest,
		},
	}

	outwayURL, err := url.Parse(mockServer.URL)
	assert.Nil(t, err)

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(outwayURL)}}

	for _, tt := range tests {
		outway.requestHTTPHandler = tt.httpHandler

		req, err := http.NewRequest("GET", tt.url, http.NoBody)
		if err != nil {
			t.Fatal("error creating http request", err)
		}

		req.Header.Add("X-NLX-Request-Data-Subject", tt.dataSubjectHeader)
		resp, err := client.Do(req)

		if err != nil {
			t.Fatal("error doing http request", err)
		}
		defer resp.Body.Close()

		assert.Equal(t, tt.statusCode, resp.StatusCode)

		contents, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error parsing result.body", err)
		}

		if tt.wantErr != nil {
			gotError := &httperrors.NLXNetworkError{}
			err := json.Unmarshal(contents, gotError)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantErr, gotError)
		}
	}
}

func TestHandleConnectMethodException(t *testing.T) {
	logger := zap.NewNop()
	outway := &Outway{}
	outway.forwardingHTTPProxy = newForwardingProxy()

	recorder := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodConnect, "http://mockservice.00000000000000000001.services.nlx.local", nil)
	outway.handleHTTPRequestAsProxy(logger, recorder, req)

	assert.Equal(t, httperrors.StatusNLXNetworkError, recorder.Code)

	contents, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal("error parsing result.body", err)
	}

	gotError := &httperrors.NLXNetworkError{}
	err = json.Unmarshal(contents, gotError)
	assert.NoError(t, err)

	assert.Equal(t, &httperrors.NLXNetworkError{
		Source:   httperrors.Outway,
		Location: httperrors.C1,
		Code:     httperrors.UnsupportedMethodErr,
		Message:  "CONNECT method is not supported",
	}, gotError)
}

func TestHandleOnNLXExceptions(t *testing.T) {
	outway := createMockOutway()

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockHTTPService(ctrl)
	outway.servicesHTTP["00000000000000000001.mockservice"] = mockService

	tests := map[string]struct {
		authEnabled        bool
		txLogger           func(ctrl *gomock.Controller) transactionlog.TransactionLogger
		dataSubjectHeader  string
		wantHTTPStatusCode int
		wantErr            *httperrors.NLXNetworkError
	}{
		"with_failing_auth_settings": {
			authEnabled: true,
			txLogger: func(ctrl *gomock.Controller) transactionlog.TransactionLogger {
				txLogger := mock_transactionlog.NewMockTransactionLogger(ctrl)
				return txLogger
			},
			dataSubjectHeader:  "",
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.OAS1,
				Code:     httperrors.ErrorWhileAuthorizingRequestErr,
				Message:  "error authorizing request",
			},
		},
		"with_failing_transactionlogger": {
			authEnabled: false,
			txLogger: func(ctrl *gomock.Controller) transactionlog.TransactionLogger {
				txLogger := mock_transactionlog.NewMockTransactionLogger(ctrl)
				txLogger.EXPECT().AddRecord(gomock.Any()).Return(fmt.Errorf("cannot add transaction record"))

				return txLogger
			},
			dataSubjectHeader:  "",
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.C1,
				Code:     httperrors.ServerErrorErr,
				Message:  "server error: unable to add record to database: cannot add transaction record",
			},
		},
		"with_invalid_datasubject header": {
			authEnabled: false,
			txLogger: func(ctrl *gomock.Controller) transactionlog.TransactionLogger {
				txLogger := mock_transactionlog.NewMockTransactionLogger(ctrl)
				return txLogger
			},
			dataSubjectHeader:  "invalid",
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.C1,
				Code:     httperrors.InvalidDataSubjectHeaderErr,
				Message:  "invalid data subject header",
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			tCtrl := gomock.NewController(t)
			defer tCtrl.Finish()

			outway.txlogger = tt.txLogger(tCtrl)

			outway.plugins = []plugins.Plugin{
				plugins.NewLogRecordPlugin("00000000000000000001", outway.txlogger),
				plugins.NewStripHeadersPlugin("00000000000000000001"),
			}

			if tt.authEnabled {
				outway.plugins = append([]plugins.Plugin{
					plugins.NewAuthorizationPlugin(&plugins.NewAuthorizationPluginArgs{
						CA:                  nil,
						ServiceURL:          "",
						AuthorizationClient: http.Client{},
					}),
				}, outway.plugins...)
			}

			req := httptest.NewRequest("GET", "http://mockservice.00000000000000000001.services.nlx.local", nil)
			req.Header.Add("X-NLX-Request-Data-Subject", tt.dataSubjectHeader)

			outway.handleOnNLX(outway.logger, &plugins.Destination{
				OrganizationSerialNumber: "00000000000000000001",
				Service:                  "mockservice",
				Path:                     "/",
			}, recorder, req)

			assert.Equal(t, tt.wantHTTPStatusCode, recorder.Code)

			contents, err := io.ReadAll(recorder.Body)
			if err != nil {
				t.Fatal("error parsing result.body", err)
			}

			if tt.wantErr != nil {
				gotError := &httperrors.NLXNetworkError{}
				err := json.Unmarshal(contents, gotError)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantErr, gotError)
			}
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
		url        string
		statusCode int
		wantErr    *httperrors.NLXNetworkError
	}{
		"when_request_to_inway_fails": {
			url:        fmt.Sprintf("%s/00000000000000000001/mockservice/", mockServer.URL),
			statusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Outway,
				Location: httperrors.O1,
				Code:     httperrors.ServiceUnreachableErr,
				Message:  "failed API request to https://inway.00000000000000000001/mockservice/ try again later. service api down/unreachable. check error at https://docs.nlx.io/support/common-errors/",
			},
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
				State:   directoryapi.Inway_STATE_UP,
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
			State:   directoryapi.Inway_STATE_UP,
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
