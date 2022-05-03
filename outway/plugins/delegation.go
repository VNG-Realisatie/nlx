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

	"go.nlx.io/nlx/common/delegation"
	"go.nlx.io/nlx/common/grpcerrors"
	"go.nlx.io/nlx/management-api/api"
	outway_http "go.nlx.io/nlx/outway/http"
	"go.nlx.io/nlx/outway/pkg/httperrors"
)

type delegationError struct {
	source  error
	message string
}

func (err *delegationError) Error() string {
	return fmt.Sprintf("%s: %s", err.message, err.source)
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

func (plugin *DelegationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		if !isDelegatedRequest(context.Request) {
			return next(context)
		}

		serialNumber, orderReference, err := parseRequestMetadata(context.Request)
		if err != nil {
			msg := "failed to parse delegation metadata"

			context.Logger.Error(msg, zap.Error(err))

			outway_http.WriteError(context.Response, msg)

			return nil
		}

		context.LogData["delegator"] = serialNumber
		context.LogData["orderReference"] = orderReference

		claim, err := plugin.getOrRequestClaim(serialNumber, orderReference, context.Destination.OrganizationSerialNumber, context.Destination.Service)
		if err != nil {
			msg := fmt.Sprintf("failed to request claim from %s: %s", serialNumber, err)

			if delegationErr, ok := err.(*delegationError); ok {
				msg = delegationErr.Error()
			}

			outway_http.WriteError(context.Response, msg)

			return nil
		}

		context.Request.Header.Add(delegation.HTTPHeaderClaim, claim.Raw)

		return next(context)
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

func (plugin *DelegationPlugin) getOrRequestClaim(orderOrganizationSerialNumber, orderReference, serviceOrganizationSerialNumber, serviceName string) (*claimData, error) {
	claimKey := fmt.Sprintf("%s/%s/%s", orderOrganizationSerialNumber, orderReference, serviceName)

	value, ok := plugin.claims.Load(claimKey)
	if !ok || value.(*claimData).Valid() != nil {
		claim, raw, err := plugin.requestClaim(orderOrganizationSerialNumber, orderReference, serviceOrganizationSerialNumber, serviceName)
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

func (plugin *DelegationPlugin) requestClaim(orderOrganizationSerialNumber, orderReference, serviceOrganizationSerialNumber, serviceName string) (*delegation.JWTClaims, string, error) {
	response, err := plugin.managementClient.RetrieveClaimForOrder(context.Background(), &api.RetrieveClaimForOrderRequest{
		OrderOrganizationSerialNumber:   orderOrganizationSerialNumber,
		OrderReference:                  orderReference,
		ServiceOrganizationSerialNumber: serviceOrganizationSerialNumber,
		ServiceName:                     serviceName,
	})
	if err != nil {
		if grpcerrors.Equal(err, api.ErrorReason_ORDER_NOT_FOUND) {
			return nil, "", fmt.Errorf("order does not exist for organization")
		}

		if grpcerrors.Equal(err, api.ErrorReason_ORDER_REVOKED) {
			return nil, "", fmt.Errorf("order is revoked")
		}

		return nil, "", httperrors.NewFromGRPCError(err)
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
