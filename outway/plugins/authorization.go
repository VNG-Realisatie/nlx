// Copyright © VNG Realisatie 2018
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

	"go.nlx.io/nlx/common/httperrors"
	outway_http "go.nlx.io/nlx/outway/http"
)

const (
	HTTPHeaderAuthorization = "X-NLX-Authorization"
)

type authRequest struct {
	Input *authRequestInput `json:"input"`
}

type authRequestInput struct {
	Headers                  http.Header `json:"headers"`
	Path                     string      `json:"path"`
	OrganizationSerialNumber string      `json:"organization_serial_number"`
	Service                  string      `json:"service"`
}

type authResponse struct {
	Result bool `json:"result"`
}

type AuthorizationPlugin struct {
	ca                  *x509.CertPool
	serviceURL          string
	authorizationClient http.Client
}

type NewAuthorizationPluginArgs struct {
	CA                  *x509.CertPool
	ServiceURL          string
	AuthorizationClient http.Client
}

func NewAuthorizationPlugin(args *NewAuthorizationPluginArgs) *AuthorizationPlugin {
	return &AuthorizationPlugin{
		ca:                  args.CA,
		serviceURL:          args.ServiceURL,
		authorizationClient: args.AuthorizationClient,
	}
}

func (plugin *AuthorizationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		authResponse, authErr := plugin.authorizeRequest(context.Request.Header, context.Destination)
		if authErr != nil {
			context.Logger.Error("error authorizing request", zap.Error(authErr))

			outway_http.WriteError(context.Response, httperrors.OAS1, httperrors.ErrorWhileAuthorizingRequest, "error authorizing request")

			return nil
		}

		context.Logger.Info(
			"authorization result",
			zap.Bool("authorized", authResponse.Result),
		)

		if !authResponse.Result {
			outway_http.WriteError(context.Response, httperrors.OAS1, httperrors.Unauthorized, "authorization server denied request")

			return nil
		}

		return next(context)
	}
}

func (plugin *AuthorizationPlugin) authorizeRequest(h http.Header, d *Destination) (*authResponse, error) {
	req, err := http.NewRequest(http.MethodPost, plugin.serviceURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(&authRequest{
		Input: &authRequestInput{
			Headers:                  h,
			Path:                     d.Path,
			OrganizationSerialNumber: d.OrganizationSerialNumber,
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

	authResponse := &authResponse{}

	err = json.NewDecoder(resp.Body).Decode(authResponse)
	if err != nil {
		return nil, err
	}

	return authResponse, nil
}
