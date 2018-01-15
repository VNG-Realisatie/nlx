// Copyright 2018 VNG Realisatie. All rights reserved.
// Use of this source code is governed by the EUPL
// license that can be found in the LICENSE.md file.

package outway

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (o *Outway) ListenAndServe(address string) error {
	err := http.ListenAndServe(address, o)
	if err != nil {
		return errors.Wrap(err, "failed to run http server")
	}
	return nil
}

// ServeHTTP handles requests from the organization to the outway, it selects the correct service backend and lets it handle the request further.
func (o *Outway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := o.logger.With(
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)
	logger.Debug("received request")
	urlparts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 3)
	if len(urlparts) != 3 {
		http.Error(w, "nlx outway error: invalid path in url", http.StatusBadRequest)
		logger.Warn("received request with invalid path")
		return
	}
	organizationName := urlparts[0]
	serviceName := urlparts[1]
	r.URL.Path = urlparts[2] // retain original path

	o.servicesLock.RLock()
	service := o.services[organizationName+"."+serviceName]
	o.servicesLock.RUnlock()
	if service == nil {
		http.Error(w, "nlx outway error: unknown service", http.StatusBadRequest)
		logger.Warn("received request for unknown service")
		return
	}

	service.proxyRequest(w, r)
}
