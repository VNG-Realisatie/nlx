package outway

import (
	"encoding/json"
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

func TestStripHeaders(t *testing.T) {
	o := &Outway{
		organizationName: "org",
		logger:           zap.NewNop(),
	}

	headers := []string{
		"X-NLX-Requester-User",
		"X-NLX-Requester-Claims",
		"X-NLX-Request-Subject-Identifier",
		"X-NLX-Request-Application-Id",
		"X-NLX-Request-User-Id",
		"X-NLX-Request-Data-Subject",
	}

	r := &http.Request{
		Header: http.Header{},
	}
	for _, header := range headers {
		r.Header.Add(header, header)
	}

	o.stripHeaders(r, "org")
	for _, header := range headers {
		assert.Equal(t, header, r.Header.Get(header))
	}

	o.stripHeaders(r, "differentOrg")
	for _, header := range headers {
		assert.Equal(t, "", r.Header.Get(header))
	}

	r.Header.Add("Proxy-Authorization", "Proxy-Authorization")
	o.stripHeaders(r, "org")
	assert.Equal(t, "", r.Header.Get("Proxy-Authorization"))
}

func TestAuthListen(t *testing.T) {
	logger := zap.NewNop()
	// Createa a outway with a mock service
	outway := &Outway{
		organizationName: "org",
		services:         make(map[string]HTTPService),
		logger:           logger,
		requestFlake:     sonyflake.NewSonyflake(sonyflake.Settings{}),
		ecmaTable:        crc64.MakeTable(crc64.ECMA),
		txlogger:         transactionlog.NewDiscardTransactionLogger(),
	}

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	mockService := mock.NewMockHTTPService(ctrl)
	defer ctrl.Finish()
	mockService.EXPECT().ProxyHTTPRequest(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mockAuthServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authRequest := &authRequest{}
		err := json.NewDecoder(r.Body).Decode(authRequest)
		assert.Nil(t, err)

		authResponse := &authResponse{}
		if user := authRequest.Headers.Get("Authorization-Proxy"); user == "Bearer token" {
			authResponse.Authorized = true
			json.NewEncoder(w).Encode(authResponse)
			return
		}

		authResponse.Authorized = false
		authResponse.Reason = "invalid user"
		json.NewEncoder(w).Encode(authResponse)
	}))
	defer mockAuthServer.Close()

	outway.services["mockorg.mockservice"] = mockService
	outway.authorizationSettings = &authSettings{
		serviceURL: mockAuthServer.URL,
	}
	outway.authorizationClient = http.Client{}
	// Setup mock http server with the outway as http handler
	mockServer := httptest.NewServer(outway)
	defer mockServer.Close()

	// Test http responses
	tests := []struct {
		url                    string
		setAuthorizationHeader bool
		statusCode             int
		errorMessage           string
	}{
		{fmt.Sprintf("%s/mockorg/mockservice/", mockServer.URL), false, http.StatusUnauthorized, "nlx outway: authorization failed. reason: invalid user\n"},
		{fmt.Sprintf("%s/mockorg/mockservice/", mockServer.URL), true, http.StatusOK, ""},
	}
	client := http.Client{}
	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		assert.Nil(t, err)

		if test.setAuthorizationHeader {
			req.Header.Add("Authorization-Proxy", "Bearer token")
		}

		resp, err := client.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, test.statusCode, resp.StatusCode)

		bytes, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, test.errorMessage, string(bytes))
	}
}
