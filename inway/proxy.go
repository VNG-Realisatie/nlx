// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"net/http"
	"strings"

	"go.uber.org/zap"
)

// handleProxyRequest handles requests from an NLX Outway to the Inway.
// It verifies authentication and selects the correct EnvpointService to
// handle the request based on the request's URI.
func (i *Inway) handleProxyRequest(w http.ResponseWriter, r *http.Request) {
	logger := i.logger.With(
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)

	peerCertificates := r.TLS.PeerCertificates

	if len(peerCertificates) == 0 {
		http.Error(w, "nlx inway error: requester is missing certificates", http.StatusBadRequest)
		logger.Warn("received request no certificates")
	}

	peerCertificate := peerCertificates[0]
	organizations := peerCertificate.Subject.Organization

	if len(organizations) == 0 {
		http.Error(w, "nlx inway error: requester certificate is missing organization", http.StatusBadRequest)
		logger.Warn("received request no organization on certificate")
		return
	}
	requesterOrganization := organizations[0]

	if requesterOrganization == "" {
		http.Error(w, "nlx inway error: requester organization is empty", http.StatusBadRequest)
		logger.Warn("received request with empty organization")
		return
	}

	issuer := r.TLS.PeerCertificates[0].Issuer.Organization[0]
	logger = logger.With(
		zap.String("cert-issuer", issuer),
		zap.String("requester", requesterOrganization),
	)

	urlparts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	if len(urlparts) != 2 {
		http.Error(w, "nlx inway error: invalid path in url", http.StatusBadRequest)
		logger.Warn("received request with invalid path")
		return
	}
	serviceName := urlparts[0]
	logrecordID := r.Header.Get("X-NLX-Logrecord-Id")
	logger.Info(
		"received API request",
		zap.String("requester-organization", requesterOrganization),
		zap.String("service", serviceName),
		zap.String("logrecord-id", logrecordID),
	)

	reqMD := &RequestMetadata{
		requesterOrganization: requesterOrganization,
		requestPath:           "/" + urlparts[1],
	}

	logger.Info("servicename: " + serviceName)
	i.serviceEndpointsLock.RLock()
	serviceEndpoint := i.serviceEndpoints[serviceName]
	i.serviceEndpointsLock.RUnlock()
	if serviceEndpoint == nil {
		http.Error(w, "nlx inway error: no endpoint for service", http.StatusBadRequest)
		logger.Warn("received request for service with no known endpoint")
		return
	}

	serviceEndpoint.handleRequest(reqMD, w, r)
}
