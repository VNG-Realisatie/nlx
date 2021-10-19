// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins

import (
	"net/http"

	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
)

type AuthenticationPlugin struct {
}

func NewAuthenticationPlugin() *AuthenticationPlugin {
	return &AuthenticationPlugin{}
}

func (d *AuthenticationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		logger := context.Logger.With(
			zap.String("request-path", context.Request.URL.Path),
			zap.String("request-remote-address", context.Request.RemoteAddr),
		)

		peerCertificates := context.Request.TLS.PeerCertificates

		if len(peerCertificates) == 0 {
			http.Error(context.Response, "nlx-inway: invalid connection: missing peer certificates", http.StatusBadRequest)
			logger.Warn("received request no certificates")

			return nil
		}

		peerCertificate := peerCertificates[0]
		organizations := peerCertificate.Subject.Organization

		if len(organizations) == 0 {
			msg := "invalid certificate provided: missing organizations attribute in subject"
			http.Error(context.Response, "nlx-inway: "+msg, http.StatusBadRequest)
			logger.Warn(msg)

			return nil
		}

		requesterOrganizationName := organizations[0]

		if requesterOrganizationName == "" {
			msg := "invalid certificate provided: missing value for organization in subject"
			http.Error(context.Response, "nlx-inway: "+msg, http.StatusBadRequest)
			logger.Warn(msg)

			return nil
		}

		requesterOrganizationSerialNumber := peerCertificate.Subject.SerialNumber

		err := common_tls.ValidateSerialNumber(requesterOrganizationSerialNumber)
		if err != nil {
			msg := "invalid certificate provided: missing value for serial number in subject"
			http.Error(context.Response, "nlx-inway: "+msg, http.StatusBadRequest)
			logger.Warn(msg)

			return nil
		}

		if len(peerCertificate.Issuer.Organization) == 0 {
			msg := "invalid certificate provided: missing value for issuer organization in issuer"
			http.Error(context.Response, "nlx-inway: "+msg, http.StatusBadRequest)
			logger.Warn(msg)

			return nil
		}

		context.Request.Header.Set("X-NLX-Request-Organization", requesterOrganizationSerialNumber)

		context.AuthInfo.OrganizationSerialNumber = requesterOrganizationSerialNumber
		context.AuthInfo.PublicKeyFingerprint = common_tls.X509PublicKeyFingerprint(peerCertificate)
		context.LogData["organizationName"] = requesterOrganizationName
		context.LogData["organizationSerialNumber"] = requesterOrganizationSerialNumber

		return next(context)
	}
}
