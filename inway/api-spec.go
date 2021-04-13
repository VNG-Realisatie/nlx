// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

const timeout = 30 * time.Second
const maxIdleConns = 100
const idleConnTimeout = 90 * time.Second
const tlsHandshakeTimeout = 10 * time.Second

func newRoundTripHTTPTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: timeout,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          maxIdleConns,
		MaxIdleConnsPerHost:   maxIdleConns,
		IdleConnTimeout:       idleConnTimeout,
		TLSHandshakeTimeout:   tlsHandshakeTimeout,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

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

	httpClient := &http.Client{Transport: newRoundTripHTTPTransport()}

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
