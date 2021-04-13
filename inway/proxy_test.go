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

func TestInwayNoOrganizationNameInClientCertificate(t *testing.T) {
	mockEndPoint := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

	i := &Inway{
		logger: zap.NewNop(),
		services: map[string]*plugins.Service{
			"mock-service": {EndpointURL: mockEndPoint.URL},
		},
		servicesLock: sync.RWMutex{},
		plugins:      []plugins.Plugin{},
	}

	tests := map[string]struct {
		path                 string
		expectedStatusCode   int
		expectedErrorMessage string
	}{
		"invalid_path": {
			path:                 "/invalid",
			expectedStatusCode:   http.StatusBadRequest,
			expectedErrorMessage: "nlx-inway: invalid path in url\n",
		},
		"service_does_not_exist": {
			path:                 "/non-existing-service/",
			expectedStatusCode:   http.StatusBadRequest,
			expectedErrorMessage: "nlx-inway: no endpoint for service\n",
		},
		"happy_flow": {
			path:                 "/mock-service/",
			expectedStatusCode:   http.StatusOK,
			expectedErrorMessage: "",
		},
	}

	for name, test := range tests {
		tc := test

		t.Run(name, func(t *testing.T) {
			url := fmt.Sprintf("%s%s", "http://localhost", tc.path)

			req, err := http.NewRequest("GET", url, nil)
			assert.Nil(t, err)

			responseRecorder := httptest.NewRecorder()
			i.handleProxyRequest(responseRecorder, req)

			result := responseRecorder.Result()
			defer result.Body.Close()

			bytes, err := ioutil.ReadAll(responseRecorder.Body)
			assert.Nil(t, err)

			assert.Equal(t, tc.expectedStatusCode, result.StatusCode)
			assert.Equal(t, tc.expectedErrorMessage, string(bytes))
		})
	}
}
