// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/form3tech-oss/jwt-go"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/server"
)

type delegationError struct {
	source  error
	message string
}

func (err *delegationError) Error() string {
	return err.message
}

func newDelegationError(message string, source error) *delegationError {
	return &delegationError{
		source:  source,
		message: message,
	}
}

type claimData struct {
	server.JWTClaims
	Raw string
}

type DelegationPlugin struct {
	claims           sync.Map
	managementClient api.ManagementClient
}

func NewDelegationPlugin(managementClient api.ManagementClient) *DelegationPlugin {
	return &DelegationPlugin{
		claims:           sync.Map{},
		managementClient: managementClient,
	}
}

func isDelegatedRequest(r *http.Request) bool {
	return r.Header.Get("X-NLX-Request-Delegator") != "" ||
		r.Header.Get("X-NLX-Request-OrderReference") != ""
}

func parseRequestMetadata(r *http.Request) (name, orderReference string, err error) {
	name = r.Header.Get("X-NLX-Request-Delegator")
	orderReference = r.Header.Get("X-NLX-Request-OrderReference")

	if name == "" {
		return "", "", errors.New("missing organization-name in delegation headers")
	}

	if orderReference == "" {
		return "", "", errors.New("missing order-reference in delegation headers")
	}

	return
}

func (plugin *DelegationPlugin) requestClaim(name, orderReference string) (*server.JWTClaims, string, error) {
	response, err := plugin.managementClient.RetrieveClaimForOrder(context.Background(), &api.RetrieveClaimForOrderRequest{
		OrderOrganizationName: name,
		OrderReference:        orderReference,
	})
	if err != nil {
		return nil, "", newDelegationError("failed to retrieve claim", err)
	}

	parser := &jwt.Parser{}

	token, _, err := parser.ParseUnverified(response.Claim, &server.JWTClaims{})
	if err != nil {
		return nil, "", newDelegationError("failed to parse JWT", err)
	}

	if err := token.Claims.Valid(); err != nil {
		return nil, "", newDelegationError("failed to validate JWT claims", err)
	}

	return token.Claims.(*server.JWTClaims), token.Raw, nil
}

func (plugin *DelegationPlugin) getOrRequestClaim(name, orderReference string) (*claimData, error) {
	claimKey := fmt.Sprintf("%s/%s", name, orderReference)

	value, ok := plugin.claims.Load(claimKey)
	if !ok || value.(*claimData).Valid() != nil {
		claim, raw, err := plugin.requestClaim(name, orderReference)
		if err != nil {
			return nil, err
		}

		data := &claimData{
			Raw:       raw,
			JWTClaims: *claim,
		}

		plugin.claims.Store(claimKey, data)

		return data, nil
	}

	return value.(*claimData), nil
}

func (plugin *DelegationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		if !isDelegatedRequest(context.Request) {
			return next(context)
		}

		name, orderReference, err := parseRequestMetadata(context.Request)
		if err != nil {
			msg := "failed to parse delegation metadata"

			context.Logger.Error(msg, zap.Error(err))
			http.Error(context.Response, msg, http.StatusInternalServerError)

			return nil
		}

		context.LogData["delegator"] = name
		context.LogData["orderReference"] = orderReference

		claim, err := plugin.getOrRequestClaim(name, orderReference)
		if err != nil {
			msg := fmt.Sprintf("failed to request claim from %s", name)

			if delegationErr, ok := err.(*delegationError); ok {
				msg = delegationErr.message
				err = delegationErr.source
			}

			context.Logger.Error(msg, zap.Error(err))
			http.Error(context.Response, msg, http.StatusInternalServerError)

			return nil
		}

		context.Request.Header.Add("X-NLX-Request-Claim", claim.Raw)

		return next(context)
	}
}
