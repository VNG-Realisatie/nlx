// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"errors"
	"fmt"
	"hash/crc64"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	mock "go.nlx.io/nlx/outway/mock"
)

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
		requestFlake:      sonyflake.NewSonyflake(sonyflake.Settings{}),
		ecmaTable:         crc64.MakeTable(crc64.ECMA),
		txlogger:          transactionlog.NewDiscardTransactionLogger(),
	}

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	mockService := mock.NewMockHTTPService(ctrl)
	mockFailService := mock.NewMockHTTPService(ctrl)
	defer ctrl.Finish()

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
		requestFlake:      sonyflake.NewSonyflake(sonyflake.Settings{}),
		ecmaTable:         crc64.MakeTable(crc64.ECMA),
		txlogger:          transactionlog.NewDiscardTransactionLogger(),
	}

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
			"failed request to https://inway.mockorg/mockservice/ try again later / check firewall?\n",
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
	healtyStates := []bool{true}
	certFile := filepath.Join("..", "testing", "org-nlx-test.crt")
	keyFile := filepath.Join("..", "testing", "org-nlx-test.key")

	l, err := NewRoundRobinLoadBalancedHTTPService(
		zap.NewNop(), nil, certFile, keyFile,
		"mockorg", "mockservice",
		inwayAddresses, healtyStates)

	assert.Nil(t, err)
	outway.servicesHTTP["mockorg.mockservice"] = l
	// set transports to fail.
	outway.setFailingTransport()
	testRequests(t, tests)
}

func TestParseURLPath(t *testing.T) {
	destination, err := parseURLPath("/organization/service/path")
	assert.Nil(t, err)
	assert.Equal(t, "organization", destination.Organization)
	assert.Equal(t, "service", destination.Service)
	assert.Equal(t, "path", destination.Path)

	_, err = parseURLPath("/organization/service")
	assert.EqualError(t, err, "invalid path in url expecting: /organization/serivice/path")
}

func TestCreateRecordData(t *testing.T) {
	headers := http.Header{}
	headers.Add("X-NLX-Request-Process-Id", "process-id")
	headers.Add("X-NLX-Request-Data-Elements", "data-elements")
	headers.Add("X-NLX-Requester-User", "user")
	headers.Add("X-NLX-Requester-Claims", "claims")
	headers.Add("X-NLX-Request-User-Id", "user-id")
	headers.Add("X-NLX-Request-Application-Id", "application-id")
	headers.Add("X-NLX-Request-Subject-Identifier", "subject-identifier")
	recordData := createRecordData(headers, "/path")

	assert.Equal(t, "process-id", recordData["doelbinding-process-id"])
	assert.Equal(t, "data-elements", recordData["doelbinding-data-elements"])
	assert.Equal(t, "user", recordData["doelbinding-user"])
	assert.Equal(t, "claims", recordData["doelbinding-claims"])
	assert.Equal(t, "user-id", recordData["doelbinding-user-id"])
	assert.Equal(t, "application-id", recordData["doelbinding-application-id"])
	assert.Equal(t, "subject-identifier", recordData["doelbinding-subject-identifier"])
	assert.Equal(t, "/path", recordData["request-path"])
}
