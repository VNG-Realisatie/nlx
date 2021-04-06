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
)

func TestHealth(t *testing.T) {
	inway := &Inway{}
	inway.services = make(map[string]ServiceEndpoint)
	inway.services["mockservice"] = &HTTPServiceEndpoint{}

	// Test health check
	tests := []struct {
		url      string
		expected bool
	}{
		{url: "http://localhost:8080/.nlx/health/mockservice", expected: true},
		{url: "http://localhost:8080/.nlx/health/mockservice1", expected: false},
	}
	inway.logger = zap.NewNop()

	for _, test := range tests {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", test.url, nil)
		inway.handleHealthRequest(recorder, request)
		status := &health.Status{}
		response := recorder.Result()
		assert.Equal(t, http.StatusOK, response.StatusCode)

		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("error reading bytes from response: %s", err)
		}
		response.Body.Close()

		err = json.Unmarshal(bytes, status)
		if err != nil {
			t.Errorf("error decoding bytes: %s", err)
		}

		assert.Equal(t, test.expected, status.Healthy)
	}
}
