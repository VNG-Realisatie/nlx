// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"encoding/json"
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
		wantErr          *httperrors.NLXNetworkError
		wantEndpointPath string
	}{
		"empty_path": {
			path:           "",
			wantStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.EmptyPathErr,
				Message:  "path cannot be empty, must at least contain the service name.",
			},
		},
		"service_does_not_exist": {
			path:           "/non-existing-service/",
			wantStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.ServiceDoesNotExistErr,
				Message:  "no endpoint for service 'non-existing-service'",
			},
		},
		"happy_flow": {
			path:             "/mock-service/",
			wantStatusCode:   http.StatusOK,
			wantEndpointPath: "/",
		},
		"happy_flow_with_path": {
			path:             "/mock-service/custom/path",
			wantStatusCode:   http.StatusOK,
			wantEndpointPath: "/custom/path",
		},
		"happy_flow_with_path_and_trailing_slash": {
			path:             "/mock-service/custom/path/",
			wantStatusCode:   http.StatusOK,
			wantEndpointPath: "/custom/path/",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mockEndPoint := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					assert.Equal(t, tt.wantEndpointPath, r.URL.Path)
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

			url := fmt.Sprintf("%s%s", "http://localhost", tt.path)

			req, err := http.NewRequest("GET", url, http.NoBody)
			assert.Nil(t, err)

			responseRecorder := httptest.NewRecorder()
			i.handleProxyRequest(responseRecorder, req)

			result := responseRecorder.Result()
			defer result.Body.Close()

			contents, err := io.ReadAll(responseRecorder.Body)
			assert.Nil(t, err)

			assert.Equal(t, tt.wantStatusCode, result.StatusCode)

			if tt.wantErr != nil {
				gotError := &httperrors.NLXNetworkError{}
				err := json.Unmarshal(contents, gotError)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantErr, gotError)
			}
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

	contents, err := io.ReadAll(responseRecorder.Body)
	assert.Nil(t, err)

	assert.Equal(t, httperrors.StatusNLXNetworkError, result.StatusCode)

	gotError := &httperrors.NLXNetworkError{}
	err = json.Unmarshal(contents, gotError)
	assert.NoError(t, err)

	assert.Equal(t, &httperrors.NLXNetworkError{
		Source:   httperrors.Inway,
		Location: httperrors.A1,
		Code:     httperrors.ServiceUnreachableErr,
		Message:  "failed API request to http://non-existing-url try again later. service api down/unreachable. check error at https://docs.nlx.io/support/common-errors/",
	}, gotError)
}
