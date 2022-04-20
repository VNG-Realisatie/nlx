// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/httperrors"
	"go.nlx.io/nlx/inway/plugins"
)

func TestInwayProxy(t *testing.T) {
	tests := map[string]struct {
		path             string
		wantStatusCode   int
		wantErrorMessage string
		wantEndpointPath string
	}{
		"empty_path": {
			path:             "",
			wantStatusCode:   httperrors.StatusNLXNetworkError,
			wantErrorMessage: "nlx-inway: path cannot be empty, must at least contain the service name.\n",
		},
		"service_does_not_exist": {
			path:             "/non-existing-service/",
			wantStatusCode:   httperrors.StatusNLXNetworkError,
			wantErrorMessage: "nlx-inway: no endpoint for service 'non-existing-service'\n",
		},
		"happy_flow": {
			path:             "/mock-service/",
			wantStatusCode:   http.StatusOK,
			wantErrorMessage: "",
			wantEndpointPath: "/",
		},
		"happy_flow_with_path": {
			path:             "/mock-service/custom/path",
			wantStatusCode:   http.StatusOK,
			wantErrorMessage: "",
			wantEndpointPath: "/custom/path",
		},
		"happy_flow_with_path_and_trailing_slash": {
			path:             "/mock-service/custom/path/",
			wantStatusCode:   http.StatusOK,
			wantErrorMessage: "",
			wantEndpointPath: "/custom/path/",
		},
	}

	for name, test := range tests {
		tc := test

		t.Run(name, func(t *testing.T) {
			mockEndPoint := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					assert.Equal(t, tc.wantEndpointPath, r.URL.Path)
				}))
			defer mockEndPoint.Close()

			i := &Inway{
				logger: zap.NewNop(),
				services: map[string]*plugins.Service{
					"mock-service": {EndpointURL: mockEndPoint.URL},
				},
				servicesLock: sync.RWMutex{},
				plugins:      []plugins.Plugin{},
			}

			url := fmt.Sprintf("%s%s", "http://localhost", tc.path)

			req, err := http.NewRequest("GET", url, http.NoBody)
			assert.Nil(t, err)

			responseRecorder := httptest.NewRecorder()
			i.handleProxyRequest(responseRecorder, req)

			result := responseRecorder.Result()
			defer result.Body.Close()

			bytes, err := io.ReadAll(responseRecorder.Body)
			assert.Nil(t, err)

			assert.Equal(t, tc.wantStatusCode, result.StatusCode)
			assert.Equal(t, tc.wantErrorMessage, string(bytes))
		})
	}
}

func TestInwayProxyEndpointNotReachable(t *testing.T) {
	i := &Inway{
		logger: zap.NewNop(),
		services: map[string]*plugins.Service{
			"mock-service": {EndpointURL: "http://non-existing-url"},
		},
		servicesLock: sync.RWMutex{},
		plugins:      []plugins.Plugin{},
	}

	req, err := http.NewRequest("GET", "http://localhost/mock-service", http.NoBody)
	assert.Nil(t, err)

	responseRecorder := httptest.NewRecorder()
	i.handleProxyRequest(responseRecorder, req)

	result := responseRecorder.Result()
	defer result.Body.Close()

	bytes, err := io.ReadAll(responseRecorder.Body)
	assert.Nil(t, err)

	assert.Equal(t, httperrors.StatusNLXNetworkError, result.StatusCode)
	assert.Equal(t, "nlx-inway: failed internal API request to http://non-existing-url try again later. service api down/unreachable. check A1 error at https://docs.nlx.io/support/common-errors/\n", string(bytes))
}
