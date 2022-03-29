// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

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
			wantStatusCode:   http.StatusBadRequest,
			wantErrorMessage: "nlx-inway: path cannot be empty, must at least contain the service name.\n",
		},
		"service_does_not_exist": {
			path:             "/non-existing-service/",
			wantStatusCode:   http.StatusBadRequest,
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

			req, err := http.NewRequest("GET", url, nil)
			assert.Nil(t, err)

			responseRecorder := httptest.NewRecorder()
			i.handleProxyRequest(responseRecorder, req)

			result := responseRecorder.Result()
			defer result.Body.Close()

			bytes, err := ioutil.ReadAll(responseRecorder.Body)
			assert.Nil(t, err)

			assert.Equal(t, tc.wantStatusCode, result.StatusCode)
			assert.Equal(t, tc.wantErrorMessage, string(bytes))
		})
	}
}
