// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"

	inway_http "go.nlx.io/nlx/inway/http"
)

type AuthRequest struct {
	Input *AuthRequestInput `json:"input"`
}

type AuthRequestInput struct {
	Headers                  http.Header `json:"headers"`
	Path                     string      `json:"path"`
	OrganizationSerialNumber string      `json:"organization_serial_number"`
	Service                  *Service    `json:"service"`
}

type AuthResponse struct {
	Result bool `json:"result"`
}

type AuthorizationPlugin struct {
	authServerEnabled   bool
	ca                  *x509.CertPool
	serviceURL          string
	authorizationClient *http.Client
}

type NewAuthorizationPluginArgs struct {
	CA                  *x509.CertPool
	AuthorizationClient *http.Client
	ServiceURL          string
	AuthServerEnabled   bool
}

func NewAuthorizationPlugin(args *NewAuthorizationPluginArgs) *AuthorizationPlugin {
	return &AuthorizationPlugin{
		authServerEnabled:   args.AuthServerEnabled,
		ca:                  args.CA,
		serviceURL:          args.ServiceURL,
		authorizationClient: args.AuthorizationClient,
	}
}

func (plugin *AuthorizationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		var foundValidGrant bool

		for _, grant := range context.Destination.Service.Grants {
			if grant.OrganizationSerialNumber == context.AuthInfo.OrganizationSerialNumber && grant.PublicKeyFingerprint == context.AuthInfo.PublicKeyFingerprint {
				foundValidGrant = true
				break
			}
		}

		if !foundValidGrant {
			inway_http.WriteError(context.Response, fmt.Sprintf(`permission denied, organization %q or public key fingerprint %q is not allowed access.`, context.AuthInfo.OrganizationSerialNumber, context.AuthInfo.PublicKeyFingerprint))
			context.Logger.Info("unauthorized request blocked, permission denied", zap.String("organization-serial-number", context.AuthInfo.OrganizationSerialNumber), zap.String("certificate-fingerprint", context.AuthInfo.PublicKeyFingerprint))

			return nil
		}

		if plugin.authServerEnabled {
			authResponse, authErr := plugin.authorizeRequest(context.Request.Header, context.Destination)
			if authErr != nil {
				context.Logger.Error("error authorizing request", zap.Error(authErr))
				inway_http.WriteError(context.Response, "error authorizing request")

				return nil
			}

			context.Logger.Info(
				"authorization result",
				zap.Bool("authorized", authResponse.Result),
			)

			if !authResponse.Result {
				inway_http.WriteError(
					context.Response,
					"authorization server denied request.")

				return nil
			}
		}

		return next(context)
	}
}

func (plugin *AuthorizationPlugin) authorizeRequest(h http.Header, d *Destination) (*AuthResponse, error) {
	req, err := http.NewRequest(http.MethodPost, plugin.serviceURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(&AuthRequest{
		Input: &AuthRequestInput{
			Headers:                  h,
			Path:                     d.Path,
			OrganizationSerialNumber: d.Organization,
			Service:                  d.Service,
		},
	})
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(body))

	resp, err := plugin.authorizationClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authorization service return non 200 status code. status code: %d", resp.StatusCode)
	}

	authResponse := &AuthResponse{}

	err = json.NewDecoder(resp.Body).Decode(authResponse)
	if err != nil {
		return nil, err
	}

	return authResponse, nil
}
