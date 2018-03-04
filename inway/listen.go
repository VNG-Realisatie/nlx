// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ListenAndServeTLS is a blocking function that listens on provided tcp address to handle requests.
func (i *Inway) ListenAndServeTLS(address string, roots *x509.CertPool, certFile, keyFile string) error {
	server := &http.Server{
		Addr: address,
		TLSConfig: &tls.Config{
			// only allow clients that present a cert signed by our root CA
			ClientCAs:  roots,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
		Handler: i,
	}
	err := server.ListenAndServeTLS(certFile, keyFile)
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
	subjectOrganization := r.TLS.PeerCertificates[0].Subject.Organization[0]
	issuer := r.TLS.PeerCertificates[0].Issuer.Organization[0]
	logger = logger.With(zap.String("issuer", issuer), zap.String("subject", subjectOrganization))
	logger.Debug("received request")

	// simple health check
	if r.URL.Path == "/health" {
		io.WriteString(w, "ok")
		return
	}

	urlparts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	if len(urlparts) != 2 {
		http.Error(w, "nlx inway error: invalid path in url", http.StatusBadRequest)
		logger.Warn("received request with invalid path")
		return
	}
	serviceName := urlparts[0]
	r.URL.Path = urlparts[1]

	r.Header.Set("X-NLX-Requester-Organization", subjectOrganization)

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
