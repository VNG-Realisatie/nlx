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
	"go.nlx.io/nlx/common/httperrors"
	"go.nlx.io/nlx/common/tls"
	directory_api "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/management"
	outway_http "go.nlx.io/nlx/outway/http"
)

type claimData struct {
	delegation.JWTClaims
	Raw string
}

type createManagementClientFunc func(context.Context, string, *tls.CertificateBundle) (management.Client, error)

type DelegationPlugin struct {
	logger                     *zap.Logger
	orgCertificate             *tls.CertificateBundle
	claims                     sync.Map
	directoryClient            directory_api.DirectoryClient
	createManagementClientFunc createManagementClientFunc
}

type NewDelegationPluginArgs struct {
	Logger                     *zap.Logger
	OrgCertificate             *tls.CertificateBundle
	DirectoryClient            directory_api.DirectoryClient
	CreateManagementClientFunc createManagementClientFunc
}

func NewDelegationPlugin(args *NewDelegationPluginArgs) *DelegationPlugin {
	return &DelegationPlugin{
		logger:                     args.Logger,
		claims:                     sync.Map{},
		orgCertificate:             args.OrgCertificate,
		directoryClient:            args.DirectoryClient,
		createManagementClientFunc: args.CreateManagementClientFunc,
	}
}

func (plugin *DelegationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(requestContext *Context) error {
		if !isDelegatedRequest(requestContext.Request) {
			return next(requestContext)
		}

		delegatorSerialNumber, orderReference, err := parseRequestMetadata(requestContext.Request)
		if err != nil {
			requestContext.Logger.Error("failed to parse delegation metadata", zap.Error(err))

			outway_http.WriteError(requestContext.Response, httperrors.C1, httperrors.UnableToParseDelegationMetadata())

			return nil
		}

		requestContext.LogData["delegator"] = delegatorSerialNumber
		requestContext.LogData["orderReference"] = orderReference

		claim := plugin.getClaimFromMemory(delegatorSerialNumber, orderReference, requestContext.Destination.Service)
		if claim == nil {
			externalManagementClient, err := plugin.setupExternalManagementClient(requestContext.Request.Context(), delegatorSerialNumber)
			if err != nil {
				plugin.logger.Error("unable to setup external management client", zap.String("delegatorSerialNumber", delegatorSerialNumber), zap.Error(err))
				outway_http.WriteError(requestContext.Response, httperrors.O1, httperrors.UnableToSetupManagementClient())

				return nil
			}

			response, err := externalManagementClient.RequestClaim(requestContext.Request.Context(), &external.RequestClaimRequest{
				OrderReference:                  orderReference,
				ServiceOrganizationSerialNumber: requestContext.Destination.OrganizationSerialNumber,
				ServiceName:                     requestContext.Destination.Service,
			})
			if err != nil {
				plugin.logger.Error("unable to request claim", zap.String("orderReference", orderReference), zap.String("serviceOrganizationSerialNumber", requestContext.Destination.OrganizationSerialNumber), zap.String("serviceName", requestContext.Destination.Service), zap.Error(err))

				if grpcerrors.Equal(err, external.ErrorReason_ERROR_REASON_ORDER_NOT_FOUND) {
					outway_http.WriteError(requestContext.Response, httperrors.O1, httperrors.OrderNotFound())

					return nil
				}

				if grpcerrors.Equal(err, external.ErrorReason_ERROR_REASON_ORDER_NOT_FOUND_FOR_ORG) {
					outway_http.WriteError(requestContext.Response, httperrors.O1, httperrors.OrderDoesNotExistForYourOrganization())

					return nil
				}

				if grpcerrors.Equal(err, external.ErrorReason_ERROR_REASON_ORDER_REVOKED) {
					outway_http.WriteError(requestContext.Response, httperrors.O1, httperrors.OrderRevoked())

					return nil
				}

				if grpcerrors.Equal(err, external.ErrorReason_ERROR_REASON_ORDER_EXPIRED) {
					outway_http.WriteError(requestContext.Response, httperrors.O1, httperrors.OrderExpired())

					return nil
				}

				if grpcerrors.Equal(err, external.ErrorReason_ERROR_REASON_ORDER_DOES_NOT_CONTAIN_SERVICE) {
					outway_http.WriteError(requestContext.Response, httperrors.O1, httperrors.OrderDoesNotContainService(requestContext.Destination.Service))

					return nil
				}

				outway_http.WriteError(requestContext.Response, httperrors.O1, httperrors.UnableToRequestClaim(delegatorSerialNumber))

				return nil
			}

			claim, err = parseClaimFromResponse(response)
			if err != nil {
				plugin.logger.Error("unable to parse received claim", zap.String("claim", response.Claim), zap.Error(err))

				outway_http.WriteError(requestContext.Response, httperrors.O1, httperrors.ReceivedInvalidClaim(delegatorSerialNumber))

				return nil
			}

			plugin.storeClaimInMemory(delegatorSerialNumber, orderReference, requestContext.Destination.Service, claim)
		}

		requestContext.Request.Header.Add(delegation.HTTPHeaderClaim, claim.Raw)

		return next(requestContext)
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

func (plugin *DelegationPlugin) setupExternalManagementClient(ctx context.Context, delegatorSerialNumber string) (management.Client, error) {
	response, err := plugin.directoryClient.GetOrganizationManagementAPIProxyAddress(ctx, &directory_api.GetOrganizationManagementAPIProxyAddressRequest{
		OrganizationSerialNumber: delegatorSerialNumber,
	})
	if err != nil {
		return nil, err
	}

	externalManagementClient, err := plugin.createManagementClientFunc(ctx, response.Address, plugin.orgCertificate)
	if err != nil {
		return nil, err
	}

	return externalManagementClient, nil
}

func parseClaimFromResponse(response *external.RequestClaimResponse) (*claimData, error) {
	parser := &jwt.Parser{}

	token, _, err := parser.ParseUnverified(response.Claim, &delegation.JWTClaims{})
	if err != nil {
		return nil, err
	}

	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	return &claimData{
		Raw:       token.Raw,
		JWTClaims: *token.Claims.(*delegation.JWTClaims),
	}, nil
}

func (plugin *DelegationPlugin) getClaimFromMemory(delegatorSerialNumber, orderReference, serviceName string) *claimData {
	claimKey := fmt.Sprintf("%s/%s/%s", delegatorSerialNumber, orderReference, serviceName)

	value, ok := plugin.claims.Load(claimKey)
	if !ok || value.(*claimData).Valid() != nil {
		return nil
	}

	return value.(*claimData)
}

func (plugin *DelegationPlugin) storeClaimInMemory(delegatorSerialNumber, orderReference, serviceName string, claimData *claimData) {
	claimKey := fmt.Sprintf("%s/%s/%s", delegatorSerialNumber, orderReference, serviceName)
	plugin.claims.Store(claimKey, claimData)
}
