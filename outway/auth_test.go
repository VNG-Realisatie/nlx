package outway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
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
		"X-NLX-Logrecord-ID",
		"X-NLX-Request-Data-Subject",
		"Proxy-Authorization",
	}
	unsafeHeaders := []string{
		"X-NLX-Requester-User",
		"X-NLX-Requester-Claims",
		"X-NLX-Request-Subject-Identifier",
		"X-NLX-Request-Application-Id",
		"X-NLX-Request-User-Id",
	}

	safeHeaders := []string{
		"X-NLX-Logrecord-ID",
		"X-NLX-Request-Data-Subject",
	}

	tests := []struct {
		name                 string
		receiverOrganization string
		expectHeaders        []string
		disallowedHeaders    []string
	}{
		{
			name:                 "Different Organization",
			receiverOrganization: "differentOrg",
			expectHeaders:        safeHeaders,
			disallowedHeaders:    unsafeHeaders,
		},
		{
			name:                 "Same Organization",
			receiverOrganization: "org",
			expectHeaders:        append(safeHeaders, unsafeHeaders...),
			disallowedHeaders:    nil,
		},
		{
			name:                 "Do not pass Proxy-Authorization",
			receiverOrganization: "differentOrg",
			expectHeaders:        nil,
			disallowedHeaders:    []string{"Proxy-Authorization"},
		},
		{
			name:                 "Never pass Proxy-Authorization",
			receiverOrganization: "org",
			expectHeaders:        nil,
			disallowedHeaders:    []string{"Proxy-Authorization"},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := &http.Request{
				Header: http.Header{},
			}
			for _, header := range headers {
				r.Header.Add(header, header)
			}

			o.stripHeaders(r, tc.receiverOrganization)

			if tc.expectHeaders != nil {
				for _, header := range tc.expectHeaders {
					assert.Equal(t, header, r.Header.Get(header))
				}
			}
			if tc.disallowedHeaders != nil {
				for _, header := range tc.disallowedHeaders {
					assert.Equal(t, "", r.Header.Get(header))
				}
			}
		})
	}
}

func TestAuthListen(t *testing.T) {
	logger := zap.NewNop()
	// Createa a outway with a mock service
	outway := &Outway{
		organizationName: "org",
		servicesHTTP:     make(map[string]HTTPService),
		logger:           logger,
		txlogger:         transactionlog.NewDiscardTransactionLogger(),
	}

	outway.requestHTTPHandler = outway.handleHTTPRequest

	// Setup mock httpservice
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock.NewMockHTTPService(ctrl)
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
			if encodeErr := json.NewEncoder(w).Encode(authResponse); encodeErr != nil {
				panic(encodeErr)
			}
			return
		}

		authResponse.Authorized = false
		authResponse.Reason = "invalid user"
		if encodeErr := json.NewEncoder(w).Encode(authResponse); encodeErr != nil {
			panic(encodeErr)
		}
	}))
	defer mockAuthServer.Close()

	outway.servicesHTTP["mockorg.mockservice"] = mockService
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

		defer resp.Body.Close()

		assert.Equal(t, test.statusCode, resp.StatusCode)

		bytes, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, test.errorMessage, string(bytes))
	}
}
