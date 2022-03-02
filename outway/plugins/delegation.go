// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/delegation"
	"go.nlx.io/nlx/management-api/api"
)

const errMessageOrderRevoked = "order is revoked"

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
	delegation.JWTClaims
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
	return r.Header.Get(delegation.HTTPHeaderDelegator) != "" ||
		r.Header.Get(delegation.HTTPHeaderOrderReference) != ""
}

func parseRequestMetadata(r *http.Request) (serialNumber, orderReference string, err error) {
	serialNumber = r.Header.Get(delegation.HTTPHeaderDelegator)
	orderReference = r.Header.Get(delegation.HTTPHeaderOrderReference)

	if serialNumber == "" {
		return "", "", errors.New("missing organization serial number in delegation headers")
	}

	if orderReference == "" {
		return "", "", errors.New("missing order-reference in delegation headers")
	}

	return
}

func (plugin *DelegationPlugin) requestClaim(serialNumber, orderReference, serviceName string) (*delegation.JWTClaims, string, error) {
	response, err := plugin.managementClient.RetrieveClaimForOrder(context.Background(), &api.RetrieveClaimForOrderRequest{
		OrderOrganizationSerialNumber: serialNumber,
		OrderReference:                orderReference,
		ServiceName:                   serviceName,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			if st.Message() == errMessageOrderRevoked {
				return nil, "", newDelegationError("order is revoked", err)
			}
		}

		return nil, "", newDelegationError("failed to retrieve claim", err)
	}

	parser := &jwt.Parser{}

	token, _, err := parser.ParseUnverified(response.Claim, &delegation.JWTClaims{})
	if err != nil {
		return nil, "", newDelegationError("failed to parse JWT", err)
	}

	if err := token.Claims.Valid(); err != nil {
		return nil, "", newDelegationError("failed to validate JWT claims", err)
	}

	return token.Claims.(*delegation.JWTClaims), token.Raw, nil
}

func (plugin *DelegationPlugin) getOrRequestClaim(serialNumber, orderReference, serviceName string) (*claimData, error) {
	claimKey := fmt.Sprintf("%s/%s/%s", serialNumber, orderReference, serviceName)

	value, ok := plugin.claims.Load(claimKey)
	if !ok || value.(*claimData).Valid() != nil {
		claim, raw, err := plugin.requestClaim(serialNumber, orderReference, serviceName)
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

		serialNumber, orderReference, err := parseRequestMetadata(context.Request)
		if err != nil {
			msg := "failed to parse delegation metadata"

			context.Logger.Error(msg, zap.Error(err))
			http.Error(context.Response, msg, http.StatusInternalServerError)

			return nil
		}

		context.LogData["delegator"] = serialNumber
		context.LogData["orderReference"] = orderReference

		claim, err := plugin.getOrRequestClaim(serialNumber, orderReference, context.Destination.Service)
		if err != nil {
			msg := fmt.Sprintf("failed to request claim from %s", serialNumber)

			httpStatus := http.StatusInternalServerError

			if delegationErr, ok := err.(*delegationError); ok {
				msg = delegationErr.message

				if msg == errMessageOrderRevoked {
					httpStatus = http.StatusUnauthorized
				}
			}

			http.Error(context.Response, msg, httpStatus)

			return nil
		}

		context.Request.Header.Add(delegation.HTTPHeaderClaim, claim.Raw)

		return next(context)
	}
}
