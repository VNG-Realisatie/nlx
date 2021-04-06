// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

func (i *Inway) handleAPISpecDocRequest(w http.ResponseWriter, r *http.Request) {
	serviceName := strings.TrimPrefix(r.URL.Path, "/.nlx/api-spec-doc/")
	i.servicesLock.RLock()
	defer i.servicesLock.RUnlock()
	
	service, exists := i.services[serviceName]
	if !exists {
		http.Error(w, "service not found", http.StatusNotFound)
		return
	}

	if service.APISpecificationDocumentURL == "" {
		http.Error(w, "api specification not found for service", http.StatusNotFound)
		return
	}

	i.logger.Info("fetching api spec doc", zap.String("api-spec-doc-url", service.APISpecificationDocumentURL))
	
	httpClient := &http.Client{Transport: newRoundTripHTTPTransport(tlsConfig)},
	resp, err := httpClient.Get(service.APISpecificationDocumentURL)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		i.logger.Error("failed to fetch api specification document", zap.Error(err))

		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		i.logger.Error("copy response body failed", zap.Error(err))

		return
	}
}
