// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
)

type JWTClaims struct {
	jwt.StandardClaims
	Organization   string `json:"organization"`
	OrderReference string `json:"order_reference"`
}

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

	claim := r.Header.Get("X-NLX-Request-Claim")
	if claim != "" {
		verifyClaimResponse, err := i.delegationClient.VerifyClaim(r.Context(), &api.VerifyClaimRequest{
			ServiceName: serviceName,
			Claim:       claim,
		})

		if err != nil {
			grpcStatus := status.Convert(err)
			if grpcStatus.Code() == codes.Unauthenticated {
				logger.Error("received request with an invalid claim", zap.Error(err))
				http.Error(w, fmt.Sprintf("nlx-inway: claim is invalid. error: %s", grpcStatus.Message()), http.StatusUnauthorized)

				return
			}

			logger.Error("failed to verify claim", zap.Error(err))
			http.Error(w, "nlx-inway: failed to verify claim", http.StatusInternalServerError)

			return
		}

		reqMD.delegatorOrganization = verifyClaimResponse.OrderOrganizationName
		reqMD.orderReference = verifyClaimResponse.OrderReference
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

	serviceEndpoint.handleRequest(reqMD, w, r)
}
