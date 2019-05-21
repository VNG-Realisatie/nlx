// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"fmt"
	"hash/crc64"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/transactionlog"
	mock "go.nlx.io/nlx/outway/mock"
)

func TestOutwayListen(t *testing.T) {
	logger := zap.NewNop()
	// Createa a outway with a mock service
	outway := &Outway{
		services:     make(map[string]HTTPService),
		logger:       logger,
		requestFlake: sonyflake.NewSonyflake(sonyflake.Settings{}),
		ecmaTable:    crc64.MakeTable(crc64.ECMA),
		txlogger:     transactionlog.NewDiscardTransactionLogger(),
	}

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	mockService := mock.NewMockHTTPService(ctrl)
	defer ctrl.Finish()
	mockService.EXPECT().ProxyHTTPRequest(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	outway.services["mockorg.mockservice"] = mockService

	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	// Test http responses
	tests := []struct {
		url          string
		statusCode   int
		errorMessage string
	}{
		{fmt.Sprintf("%s/invalidpath", mockServer.URL), http.StatusBadRequest, "nlx outway: invalid path in url\n"},
		{fmt.Sprintf("%s/mockorg/nonexistingservice/add/", mockServer.URL), http.StatusBadRequest, "nlx outway: unknown service\n"},
		{fmt.Sprintf("%s/mockorg/mockservice/", mockServer.URL), http.StatusOK, ""},
	}
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

		assert.Equal(t, test.statusCode, resp.StatusCode)
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error parsing result.body", err)
		}

		assert.Equal(t, test.errorMessage, string(bytes))
	}
}

func TestParseURLPath(t *testing.T) {
	destination, err := parseURLPath("/organization/service/path")
	assert.Nil(t, err)
	assert.Equal(t, "organization", destination.Organization)
	assert.Equal(t, "service", destination.Service)
	assert.Equal(t, "path", destination.Path)

	destination, err = parseURLPath("/organization/service")
	assert.EqualError(t, err, "invalid path in url")
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
