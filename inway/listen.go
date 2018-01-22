// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (i *Inway) ListenAndServe(address string) error {
	err := http.ListenAndServe(address, i)
	if err != nil {
		return errors.Wrap(err, "failed to run http server")
	}
	return nil
}

// ServeHTTP handles requests from an NLX Outway to the Inway. It verifies authentication and selects the correct EnvpointService to handle the request based on the request's URI.
func (i *Inway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := i.logger.With(
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)
	logger.Debug("received request")
	urlparts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	if len(urlparts) != 2 {
		http.Error(w, "nlx inway error: invalid path in url", http.StatusBadRequest)
		logger.Warn("received request with invalid path")
		return
	}
	serviceName := urlparts[0]
	r.URL.Path = urlparts[1]

	i.serviceEndpointsLock.RLock()
	serviceEndpoint := i.serviceEndpoints[serviceName]
	i.serviceEndpointsLock.RUnlock()
	if serviceEndpoint == nil {
		http.Error(w, "nlx inway error: no endpoint for service", http.StatusBadRequest)
		logger.Warn("received request for service with no known endpoint")
		return
	}

	serviceEndpoint.sendRequest(w, r)
}
