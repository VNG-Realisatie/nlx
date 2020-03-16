// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package monitoring

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewMonitoring(t *testing.T) {
	m, err := NewMonitoringService("", nil)
	assert.Nil(t, m)
	assert.EqualError(t, err, "address required")

	m, err = NewMonitoringService("localhost:8080", nil)
	assert.Nil(t, m)
	assert.EqualError(t, err, "logger required")

	m, err = NewMonitoringService("localhost:8080", zap.NewNop())
	assert.NotNil(t, m)
	assert.Nil(t, err)
}

func TestLiveness(t *testing.T) {
	service := &Service{
		logger: zap.NewNop(),
	}
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/health/live", nil)
	service.handleLivenessRequest(recorder, request)

	response := recorder.Result()
	response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestReadiness(t *testing.T) {
	tests := []struct {
		description        string
		isReady            bool
		expectedStatusCode int
	}{
		{
			description:        "the monitoring service is ready",
			isReady:            true,
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "the monitoring service is not ready",
			isReady:            false,
			expectedStatusCode: http.StatusServiceUnavailable,
		},
	}

	logger := zap.NewNop()
	serviceAdddress := "localhost:8080"

	for _, test := range tests {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/health/ready", nil)
		request.Host = "host"
		service, err := NewMonitoringService(serviceAdddress, logger)
		assert.Nil(t, err)
		assert.NotNil(t, service)

		if test.isReady {
			service.SetReady()
		} else {
			service.SetNotReady()
		}

		service.handleReadinessRequest(recorder, request)
		response := recorder.Result()
		assert.Equal(t, test.expectedStatusCode, response.StatusCode)
		response.Body.Close()
	}
}
