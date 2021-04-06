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
	SetAuthorizationWhitelist(whitelistedOrganizations []config.AuthorizationWhitelistItem)
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
	whitelistedOrganizations []config.AuthorizationWhitelistItem
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
func (i *Inway) NewHTTPServiceEndpoint(serviceName string, serviceDetails *config.ServiceDetails, tlsConfig *tls.Config) (*HTTPServiceEndpoint, error) {
	h := &HTTPServiceEndpoint{
		inway:          i,
		serviceName:    serviceName,
		serviceDetails: serviceDetails,
		logger:         i.logger.With(zap.String("inway-service-name", serviceName)),
		httpClient:     &http.Client{Transport: newRoundTripHTTPTransport(tlsConfig)},
	}

	endpointURL, err := url.Parse(serviceDetails.EndpointURL)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint provided")
	}

	h.host = endpointURL.Host
	h.proxy = httputil.NewSingleHostReverseProxy(endpointURL)
	h.proxy.Transport = newRoundTripHTTPTransport(tlsConfig)
	h.proxy.ErrorHandler = i.LogAPIErrors

	return h, nil
}

func (i *Inway) LogAPIErrors(w http.ResponseWriter, r *http.Request, e error) {
	msg := ("nlx-inway: failed internal API request to " +
		r.URL.String() +
		" try again later / service api down/unreachable." +
		" check A1 error at https://docs.nlx.io/support/common-errors/")
	i.logger.Error(msg)
	http.Error(w, msg, http.StatusServiceUnavailable)
}

// SetAuthorizationPublic makes the service publicly available.
func (h *HTTPServiceEndpoint) SetAuthorizationPublic() {
	h.public = true
}

// SetAuthorizationWhitelist makes the service private and sets the whitelisted organizations.
func (h *HTTPServiceEndpoint) SetAuthorizationWhitelist(whitelistedOrganizations []config.AuthorizationWhitelistItem) {
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
			http.Error(w, "nlx-inway: could not handle your request, missing requesterOrganization header.", http.StatusBadRequest)
			h.logger.Info("request blocked, missing requesterOrganization header")

			return
		}

		var authorized = false

		h.logger.Debug("check against whitelist",
			zap.String("requesterOrganization", reqMD.requesterOrganization),
			zap.String("requesterPublicKeyHash", reqMD.requesterPublicKeyFingerprint))

		for _, whitelistedOrg := range h.whitelistedOrganizations {
			h.logger.Debug("whitelistitem",
				zap.String("OrganizationName", whitelistedOrg.OrganizationName),
				zap.String("PublicKeyHash", whitelistedOrg.PublicKeyHash))

			if whitelistedOrg.OrganizationName == "" && whitelistedOrg.PublicKeyHash == "" {
				h.logger.Warn("Whitelist item missing both organization-name and public-key-hash")
				continue
			}

			organizationNameMatches := false
			if whitelistedOrg.OrganizationName == "" {
				organizationNameMatches = true
			} else {
				organizationNameMatches = reqMD.requesterOrganization == whitelistedOrg.OrganizationName
			}

			certificateFingerprintMatches := false
			if whitelistedOrg.PublicKeyHash == "" {
				certificateFingerprintMatches = true
			} else {
				certificateFingerprintMatches = reqMD.requesterPublicKeyFingerprint == whitelistedOrg.PublicKeyHash
			}

			if organizationNameMatches && certificateFingerprintMatches {
				authorized = true
				break
			}
		}

		if !authorized {
			http.Error(w, fmt.Sprintf(`nlx-inway: permission denied, organization "%s" or public key "%s" is not allowed access.`, reqMD.requesterOrganization, reqMD.requesterPublicKeyFingerprint), http.StatusForbidden)
			h.logger.Info("unauthorized request blocked, requester was not whitelisted", zap.String("organization-name", reqMD.requesterOrganization), zap.String("certificate-fingerprint", reqMD.requesterPublicKeyFingerprint))

			return
		}
	}

	// we are public or authorized now.
	r.Host = h.host
	r.URL.Path = reqMD.requestPath
	r.Header.Set("X-NLX-Request-Organization", reqMD.requesterOrganization)

	logrecordID := r.Header.Get("X-NLX-Logrecord-Id")
	if logrecordID == "" {
		http.Error(w, "nlx-inway: missing logrecord id", http.StatusBadRequest)
		h.logger.Warn("Received request with missing logrecord id from " + reqMD.requesterOrganization)

		return
	}

	recordData := h.createRecordData(reqMD.requestPath, r.Header)
	record := &transactionlog.Record{
		SrcOrganization:  reqMD.requesterOrganization,
		DestOrganization: h.inway.organizationName,
		ServiceName:      h.serviceName,
		LogrecordID:      logrecordID,
		Data:             recordData,
	}

	if reqMD.delegateeOrganization != "" {
		record.Delegator = reqMD.requesterOrganization
		record.SrcOrganization = reqMD.delegateeOrganization
		record.OrderReference = reqMD.orderReference
	}

	err := h.inway.txlogger.AddRecord(record)
	if err != nil {
		http.Error(w, "nlx-inway: server error", http.StatusInternalServerError)
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
