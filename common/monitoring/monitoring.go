// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package monitoring

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Service struct
type Service struct {
	isReady bool
	logger  *zap.Logger
	server  *http.Server
}

// NewMonitoringService creates a new monitoring service
func NewMonitoringService(address string, logger *zap.Logger) (*Service, error) {
	if address == "" {
		return nil, fmt.Errorf("address required")
	}

	if logger == nil {
		return nil, fmt.Errorf("logger required")
	}

	m := &Service{
		isReady: false,
		logger:  logger,
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/health/live", m.handleLivenessRequest)
	serveMux.HandleFunc("/health/ready", m.handleReadinessRequest)
	m.server = &http.Server{
		Addr:    address,
		Handler: serveMux,
	}

	return m, nil
}

// Start is a blocking call which starts the internal http server
func (m *Service) Start() error {
	m.logger.Info("starting monitoring service")
	err := m.server.ListenAndServe()

	if err == http.ErrServerClosed {
		return nil
	}

	return err
}

// SetReady sets the service to ready
// the readiness endpoint will return StatusOK
func (m *Service) SetReady() {
	m.logger.Debug(("monitoring service set to ready"))
	m.isReady = true
}

// SetNotReady sets the service to not ready
// the readiness endpoint will return StatusServiceUnavailable
func (m *Service) SetNotReady() {
	m.logger.Debug(("monitoring service set to not ready"))
	m.isReady = false
}

// Stop will gracefully shutdown the internal http server
func (m *Service) Stop() error {
	m.logger.Info(("stopping monitoring service"))

	localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel() // do not remove. Otherwise it could cause implicit goroutine leak

	return m.server.Shutdown(localCtx)
}

func (m *Service) handleReadinessRequest(w http.ResponseWriter, r *http.Request) {
	m.logger.Debug("received readiness check", zap.String("from-host", r.Host))

	if !m.isReady {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *Service) handleLivenessRequest(w http.ResponseWriter, r *http.Request) {
	m.logger.Debug("received liveness check", zap.String("from-host", r.Host))
	w.WriteHeader(http.StatusOK)
}
