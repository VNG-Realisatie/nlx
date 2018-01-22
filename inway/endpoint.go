// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ServiceEndpoint handles the proxying of a request to the organization's API endpoint
type ServiceEndpoint interface {
	ServiceName() string
	sendRequest(w http.ResponseWriter, r *http.Request)
}

// HTTPServiceEndpoint implements an ServiceEndpoint for plain HTTP requests.
type HTTPServiceEndpoint struct {
	serviceName string
	logger      *zap.Logger

	proxy *httputil.ReverseProxy
}

var _ ServiceEndpoint = &HTTPServiceEndpoint{} // copmile-time interface validation

// NewHTTPServiceEndpoint creates a new ServiceEndpoint using a simple http reverse proxy backend.
func NewHTTPServiceEndpoint(logger *zap.Logger, serviceName, endpoint string) (*HTTPServiceEndpoint, error) {
	h := &HTTPServiceEndpoint{
		serviceName: serviceName,
		logger:      logger.With(zap.String("inway-service-name", serviceName)),
	}
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint provided")
	}
	h.proxy = httputil.NewSingleHostReverseProxy(endpointURL)
	return h, nil
}

// ServiceName returns the service name that the attached endpoint handles
func (h *HTTPServiceEndpoint) ServiceName() string {
	return h.serviceName
}

func (h *HTTPServiceEndpoint) sendRequest(w http.ResponseWriter, r *http.Request) {
	h.proxy.ServeHTTP(w, r)
}
