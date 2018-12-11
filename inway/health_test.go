package inway

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"go.nlx.io.copy/nlx/directory-monitor/health"
)

func TestHealth(t *testing.T) {
	inway := &Inway{}
	inway.serviceEndpoints = make(map[string]ServiceEndpoint, 0)
	inway.serviceEndpoints["testingservice"] = &HTTPServiceEndpoint{}
	tests := []struct {
		url      string
		expected bool
	}{
		{url: "http://localhost:8080/.nlx/health/testingservice", expected: true},
		{url: "http://localhost:8080/.nlx/health/testingservice1", expected: false},
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
			t.Errorf("result: %d expected: http status code should be %d", response.StatusCode, http.StatusOK)
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
			t.Errorf("result: %t expected: status.Healthy to be %t for url %s", status.Healthy, test.expected, test.url)
		}
	}
}
