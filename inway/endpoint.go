// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/inway/config"
)

// ServiceEndpoint handles the proxying of a request to the organization's API endpoint
type ServiceEndpoint interface {
	ServiceName() string
	ServiceDetails() *config.ServiceDetails
	SetAuthorizationWhitelist(whitelistedOrganizations []string)
	GetAPISpec() (*http.Response, error)
	handleRequest(reqMD *RequestMetadata, w http.ResponseWriter, r *http.Request)
}

// HTTPServiceEndpoint implements a ServiceEndpoint for plain HTTP requests.
type HTTPServiceEndpoint struct {
	inway *Inway

	serviceName    string
	serviceDetails *config.ServiceDetails
	logger         *zap.Logger

	host       string
	proxy      *httputil.ReverseProxy
	httpClient *http.Client

	public                   bool
	whitelistedOrganizations []string
}

func newRoundTripHTTPTransport(tlsConfig *tls.Config) *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}
}

var _ ServiceEndpoint = &HTTPServiceEndpoint{} // compile-time interface validation

// NewHTTPServiceEndpoint creates a new ServiceEndpoint using a simple HTTP reverse proxy backend.
func (iw *Inway) NewHTTPServiceEndpoint(serviceName string, serviceDetails *config.ServiceDetails, tlsConfig *tls.Config) (*HTTPServiceEndpoint, error) {
	h := &HTTPServiceEndpoint{
		inway:          iw,
		serviceName:    serviceName,
		serviceDetails: serviceDetails,
		logger:         iw.logger.With(zap.String("inway-service-name", serviceName)),
		httpClient:     &http.Client{Transport: newRoundTripHTTPTransport(tlsConfig)},
	}
	endpointURL, err := url.Parse(serviceDetails.EndpointURL)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint provided")
	}
	h.host = endpointURL.Host
	h.proxy = httputil.NewSingleHostReverseProxy(endpointURL)
	h.proxy.Transport = newRoundTripHTTPTransport(tlsConfig)
	h.proxy.ErrorHandler = iw.LogAPIErrors
	return h, nil
}

func (iw *Inway) LogAPIErrors(w http.ResponseWriter, r *http.Request, e error) {
	msg := "nlx-inway: failed internal API request to " + r.URL.String() + " try again later / service api down/unreachable"
	iw.logger.Error(msg)
	http.Error(w, msg, http.StatusServiceUnavailable)

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

// ServiceDetails returns the config that this endpoint is based upon
func (h *HTTPServiceEndpoint) ServiceDetails() *config.ServiceDetails {
	return h.serviceDetails
}

// GetAPISpec returns the response of the request to the API Spec Documentation URL
func (h *HTTPServiceEndpoint) GetAPISpec() (*http.Response, error) {
	return h.httpClient.Get(h.serviceDetails.APISpecificationDocumentURL)
}

func (h *HTTPServiceEndpoint) handleRequest(reqMD *RequestMetadata, w http.ResponseWriter, r *http.Request) {
	if !h.public {

		if reqMD.requesterOrganization == "" {
			http.Error(w, fmt.Sprint(`nlx-outway: could not handle your request, missing requesterOrganization header.`, reqMD.requesterOrganization), http.StatusBadRequest)
			h.logger.Info("request blocked, missing requesterOrganization header")
			return
		}

		for _, whitelistedOrg := range h.whitelistedOrganizations {
			h.logger.Info("org: " + whitelistedOrg)
			if reqMD.requesterOrganization == whitelistedOrg {
				goto Authorized
			}
		}
		http.Error(w, fmt.Sprintf(`nlx-outway: could not handle your request, organization "%s" is not allowed access.`, reqMD.requesterOrganization), http.StatusForbidden)
		h.logger.Info("unauthorized request blocked, requester was not whitelisted")
		return
	}
	// we are public or authorized now.

Authorized:

	r.Host = h.host
	r.URL.Path = reqMD.requestPath
	r.Header.Set("X-NLX-Request-Organization", reqMD.requesterOrganization)

	logrecordID := r.Header.Get("X-NLX-Logrecord-Id")
	if logrecordID == "" {
		http.Error(w, "nlx-outway: missing logrecord id", http.StatusBadRequest)
		h.logger.Warn("Received request with missing logrecord id from " + reqMD.requesterOrganization)
		return
	}

	recordData := h.createRecordData(reqMD.requestPath, r.Header)
	err := h.inway.txlogger.AddRecord(&transactionlog.Record{
		SrcOrganization:  reqMD.requesterOrganization,
		DestOrganization: h.inway.organizationName,
		ServiceName:      h.serviceName,
		LogrecordID:      logrecordID,
		Data:             recordData,
	})
	if err != nil {
		http.Error(w, "nlx-outway: server error", http.StatusInternalServerError)
		h.logger.Error("failed to store transactionlog record", zap.Error(err))
		return
	}
	h.proxy.ServeHTTP(w, r)
}

func (h *HTTPServiceEndpoint) createRecordData(requestPath string, header http.Header) map[string]interface{} {
	var recordData = make(map[string]interface{})
	if processID := header.Get("X-NLX-Request-Process-Id"); processID != "" {
		recordData["doelbinding-process-id"] = processID
	}
	if dataElements := header.Get("X-NLX-Request-Data-Elements"); dataElements != "" {
		recordData["doelbinding-data-elements"] = dataElements
	}

	if userData := header.Get("X-NLX-Requester-User"); userData != "" {
		recordData["doelbinding-user"] = userData
	}

	if claims := header.Get("X-NLX-Requester-Claims"); claims != "" {
		recordData["doelbinding-claims"] = claims
	}
	recordData["request-path"] = requestPath

	return recordData
}
