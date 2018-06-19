// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/transactionlog"
)

// ServiceEndpoint handles the proxying of a request to the organization's API endpoint
type ServiceEndpoint interface {
	ServiceName() string
	SetAuthorizationWhitelist(whitelistedOrganizations []string)
	handleRequest(reqMD *RequestMetadata, w http.ResponseWriter, r *http.Request)
}

// HTTPServiceEndpoint implements a ServiceEndpoint for plain HTTP requests.
type HTTPServiceEndpoint struct {
	inway *Inway

	serviceName string
	logger      *zap.Logger

	host  string
	proxy *httputil.ReverseProxy

	public                   bool
	whitelistedOrganizations []string
}

var _ ServiceEndpoint = &HTTPServiceEndpoint{} // compile-time interface validation

// NewHTTPServiceEndpoint creates a new ServiceEndpoint using a simple HTTP reverse proxy backend.
func (iw *Inway) NewHTTPServiceEndpoint(logger *zap.Logger, serviceName, endpoint string) (*HTTPServiceEndpoint, error) {
	h := &HTTPServiceEndpoint{
		inway:       iw,
		serviceName: serviceName,
		logger:      logger.With(zap.String("inway-service-name", serviceName)),
	}
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint provided")
	}
	h.host = endpointURL.Host
	h.proxy = httputil.NewSingleHostReverseProxy(endpointURL)
	return h, nil
}

// SetAuthorizationPublic makes the service publicly available.
func (h *HTTPServiceEndpoint) SetAuthorizationPublic() {
	h.public = true
}

// SetAuthorizationWhitelist makes the service private and sets the whitelisted organizations.
func (h *HTTPServiceEndpoint) SetAuthorizationWhitelist(whitelistedOrganizations []string) {
	h.public = false
	h.whitelistedOrganizations = whitelistedOrganizations
}

// ServiceName returns the service name that the attached endpoint handles
func (h *HTTPServiceEndpoint) ServiceName() string {
	return h.serviceName
}

func (h *HTTPServiceEndpoint) handleRequest(reqMD *RequestMetadata, w http.ResponseWriter, r *http.Request) {
	if !h.public {
		for _, whitelistedOrg := range h.whitelistedOrganizations {
			if reqMD.requesterOrganization == whitelistedOrg {
				goto Authorized
			}
		}
		h.logger.Info("unauthorized request blocked, requester was not whitelisted")
		http.Error(w, fmt.Sprintf(`We could not handle your request, organization "%s" is not allowed access.`, reqMD.requesterOrganization), http.StatusForbidden)
		return
	}

Authorized:

	r.Host = h.host
	r.URL.Path = reqMD.requestPath
	r.Header.Set("X-NLX-Request-Organization", reqMD.requesterOrganization)

	logrecordID := r.Header.Get("X-NLX-Logrecord-Id")
	if logrecordID == "" {
		http.Error(w, "missing logrecord id", http.StatusBadRequest)
		h.logger.Warn("Received request with missing logrecord id from " + reqMD.requesterOrganization)
		return
	}

	var recordData = make(map[string]interface{})
	if processID := r.Header.Get("X-NLX-Request-Process-Id"); processID != "" {
		recordData["doelbinding-process-id"] = processID
	}
	if dataElements := r.Header.Get("X-NLX-Request-Data-Elements"); dataElements != "" {
		recordData["doelbinding-data-elements"] = dataElements
	}
	recordData["request-path"] = reqMD.requestPath

	err := h.inway.txlogger.AddRecord(&transactionlog.Record{
		SrcOrganization:  reqMD.requesterOrganization,
		DestOrganization: h.inway.organizationName,
		ServiceName:      h.serviceName,
		LogrecordID:      logrecordID,
		Data:             recordData,
	})
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		h.logger.Error("failed to store transactionlog record", zap.Error(err))
		return
	}

	h.proxy.ServeHTTP(w, r)
}
