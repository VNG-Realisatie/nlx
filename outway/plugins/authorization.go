// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

type authRequest struct {
	Headers      http.Header `json:"headers"`
	Organization string      `json:"organization"`
	Service      string      `json:"service"`
}

type authResponse struct {
	Authorized bool   `json:"authorized"`
	Reason     string `json:"reason,omitempty"`
}

type AuthorizationPlugin struct {
	ca                  *x509.CertPool
	serviceURL          string
	authorizationClient http.Client
}

func NewAuthorizationPlugin(ca *x509.CertPool, serviceURL string, authorizationClient http.Client) *AuthorizationPlugin {
	return &AuthorizationPlugin{
		ca:                  ca,
		serviceURL:          serviceURL,
		authorizationClient: authorizationClient,
	}
}

func (plugin *AuthorizationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context Context) error {
		authResponse, authErr := plugin.authorizeRequest(context.Request.Header, context.Destination)
		if authErr != nil {
			context.Logger.Error("error authorizing request", zap.Error(authErr))
			http.Error(context.Response, "nlx outway: error authorizing request", http.StatusInternalServerError)

			return nil
		}

		context.Logger.Info(
			"authorization result",
			zap.Bool("authorized", authResponse.Authorized),
			zap.String("reason", authResponse.Reason),
		)

		if !authResponse.Authorized {
			http.Error(
				context.Response,
				fmt.Sprintf("nlx outway: authorization failed. reason: %s", authResponse.Reason),
				http.StatusUnauthorized,
			)

			return nil
		}

		return next(context)
	}
}

func (plugin *AuthorizationPlugin) authorizeRequest(h http.Header, d *Destination) (*authResponse, error) {
	req, err := http.NewRequest(http.MethodPost, plugin.serviceURL, nil)
	if err != nil {
		return nil, err
	}

	authRequest := &authRequest{
		Headers:      h,
		Organization: d.Organization,
		Service:      d.Service,
	}

	body, err := json.Marshal(authRequest)
	if err != nil {
		return nil, err
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

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
