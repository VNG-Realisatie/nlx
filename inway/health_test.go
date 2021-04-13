// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-monitor/health"
	"go.nlx.io/nlx/inway/plugins"
)

func TestHealth(t *testing.T) {
	inway := &Inway{
		services: map[string]*plugins.Service{
			"mock-service": {},
		},
	}

	tests := map[string]struct {
		url      string
		expected bool
	}{
		"healthy_endpoint":   {url: "http://localhost:8080/.nlx/health/mock-service", expected: true},
		"unhealthy_endpoint": {url: "http://localhost:8080/.nlx/health/mock-service-not-exist", expected: false},
	}
	inway.logger = zap.NewNop()

	for name, test := range tests {
		tc := test

		t.Run(name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", tc.url, nil)
			inway.handleHealthRequest(recorder, request)

			status := &health.Status{}
			response := recorder.Result()
			assert.Equal(t, http.StatusOK, response.StatusCode)

			bytes, err := ioutil.ReadAll(response.Body)
			assert.Nil(t, err)

			response.Body.Close()

			err = json.Unmarshal(bytes, status)
			assert.Nil(t, err)

			assert.Equal(t, tc.expected, status.Healthy)
		})
	}
}
