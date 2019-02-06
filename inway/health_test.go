package inway

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-monitor/health"
)

func TestHealth(t *testing.T) {
	inway := &Inway{}
	inway.serviceEndpoints = make(map[string]ServiceEndpoint, 0)
	inway.serviceEndpoints["mockservice"] = &HTTPServiceEndpoint{}

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
		request, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Errorf("error creating http request: %s", err)
		}
		inway.handleHealthRequest(recorder, request)
		status := &health.Status{}
		response := recorder.Result()
		if response.StatusCode != http.StatusOK {
			t.Errorf(`result: "%d", expected for: http status code should be "%d" for url "%s"`, response.StatusCode, http.StatusOK, test.url)
		}
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("error reading bytes from response: %s", err)
		}
		err = json.Unmarshal(bytes, status)
		if err != nil {
			t.Errorf("error decoding bytes: %s", err)
		}

		if status.Healthy != test.expected {
			t.Errorf(`result: "%t" expected: status.Healthy to be "%t" for url "%s"`, status.Healthy, test.expected, test.url)
		}
	}
}
