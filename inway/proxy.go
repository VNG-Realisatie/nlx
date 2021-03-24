// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
)

type JWTClaims struct {
	jwt.StandardClaims
	Organization   string `json:"organization"`
	OrderReference string `json:"order_reference"`
}

var ErrDelegatorDoesNotHaveAccess = errors.New("delegator does have access")
var ErrCannotParsePublicKeyFromPEM = errors.New("failed to parse PEM block containing the public key")

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
		http.Error(w, "nlx-inway: invalid connection: missing peer certificates", http.StatusBadRequest)
		logger.Warn("received request no certificates")

		return
	}

	peerCertificate := peerCertificates[0]
	organizations := peerCertificate.Subject.Organization

	if len(organizations) == 0 {
		msg := "invalid certificate provided: missing organizations attribute in subject"
		http.Error(w, "nlx-inway: "+msg, http.StatusBadRequest)
		logger.Warn(msg)

		return
	}

	requesterOrganization := organizations[0]

	if requesterOrganization == "" {
		msg := "invalid certificate provided: missing value for organization in subject"
		http.Error(w, "nlx-inway: "+msg, http.StatusBadRequest)
		logger.Warn(msg)

		return
	}

	if len(peerCertificate.Issuer.Organization) == 0 {
		msg := "invalid certificate provided: missing value for issuer organization in issuer"
		http.Error(w, "nlx-inway: "+msg, http.StatusBadRequest)
		logger.Warn(msg)

		return
	}

	issuer := r.TLS.PeerCertificates[0].Issuer.Organization[0]

	logger = logger.With(
		zap.String("cert-issuer", issuer),
		zap.String("requester", requesterOrganization),
	)

	urlparts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	if len(urlparts) != 2 {
		http.Error(w, "nlx-inway: invalid path in url", http.StatusBadRequest)
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

	publicKeyFingerprint := common_tls.PublicKeyFingerprint(peerCertificate)

	reqMD := &RequestMetadata{
		requesterOrganization:         requesterOrganization,
		requesterPublicKeyFingerprint: publicKeyFingerprint,
		requestPath:                   "/" + urlparts[1],
	}

	logger.Info("servicename: " + serviceName)
	i.serviceEndpointsLock.RLock()
	serviceEndpoint := i.serviceEndpoints[serviceName]
	i.serviceEndpointsLock.RUnlock()

	if serviceEndpoint == nil {
		http.Error(w, "nlx-inway: no endpoint for service", http.StatusBadRequest)
		logger.Warn("received request for service with no known endpoint")

		return
	}

	claim := r.Header.Get("X-NLX-Request-Claim")
	if claim != "" {
		claims := &JWTClaims{}
		_, err := jwt.ParseWithClaims(claim, claims, func(token *jwt.Token) (interface{}, error) {
			for _, whitelistItem := range serviceEndpoint.ServiceDetails().AuthorizationWhitelist {
				if whitelistItem.OrganizationName == claims.Issuer {
					return parsePublicKeyFromPEM(whitelistItem.PublicKeyPEM)
				}
			}

			return nil, ErrDelegatorDoesNotHaveAccess
		})

		if err != nil {
			validationError, ok := err.(*jwt.ValidationError)
			if !ok {
				i.logger.Error("casting error to jwt validation error failed", zap.Error(err))
				http.Error(w, "unable to verify claim", http.StatusInternalServerError)

				return
			}

			if errors.Is(validationError.Inner, ErrDelegatorDoesNotHaveAccess) {
				i.logger.Info("delegator does not have access to service", zap.String("delegator", claims.Issuer), zap.String("serviceName", serviceName))
				http.Error(w, "nlx-inway: no access", http.StatusUnauthorized)

				return
			}

			i.logger.Error("failed to parse jwt", zap.Error(err))
			http.Error(w, "nlx-inway: unable to verify claim", http.StatusInternalServerError)

			return
		}

		reqMD.requesterOrganization = claims.Issuer
		reqMD.requesterPublicKeyFingerprint = claims.Issuer
		reqMD.delegatorOrganization = claims.Issuer
		reqMD.orderReference = claims.OrderReference
	}

	serviceEndpoint.handleRequest(reqMD, w, r)
}

func parsePublicKeyFromPEM(publicKeyPEM string) (crypto.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, ErrCannotParsePublicKeyFromPEM
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}
