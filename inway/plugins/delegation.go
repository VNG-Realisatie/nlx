// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/delegation"
	inway_http "go.nlx.io/nlx/inway/http"
)

var (
	ErrDelegatorDoesNotHaveAccess                     = errors.New("delegator does not have access")
	ErrRequestingOrganizationIsNotDelegatee           = errors.New("requesting organization is not the delegatee")
	ErrRequestingOrganizationPublicKeyNotFoundInOrder = errors.New("requesting organization public key is not the public key found in order")
	ErrCannotParsePublicKeyFromPEM                    = errors.New("failed to parse PEM block containing the public key")
)

type DelegationPlugin struct {
}

func NewDelegationPlugin() *DelegationPlugin {
	return &DelegationPlugin{}
}

func (d *DelegationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		claim := context.Request.Header.Get(delegation.HTTPHeaderClaim)
		if claim == "" {
			return next(context)
		}

		var publicKeyFingerprint string

		claims := &delegation.JWTClaims{}

		_, err := jwt.ParseWithClaims(claim, claims, func(token *jwt.Token) (interface{}, error) {
			if claims.Delegatee != context.AuthInfo.OrganizationSerialNumber {
				return nil, ErrRequestingOrganizationIsNotDelegatee
			}

			if claims.DelegateePublicKeyFingerprint != context.AuthInfo.PublicKeyFingerprint {
				return nil, ErrRequestingOrganizationPublicKeyNotFoundInOrder
			}

			for _, grant := range context.Destination.Service.Grants {
				if claims.IsValidFor(context.Destination.Service.Name, claims.Issuer, grant.PublicKeyFingerprint) {
					publicKeyFingerprint = grant.PublicKeyFingerprint
					publicKey, err := parsePublicKeyFromPEM(grant.PublicKeyPEM)
					if err != nil {
						return nil, err
					}

					return publicKey, nil
				}
			}

			return nil, ErrDelegatorDoesNotHaveAccess
		})

		if err != nil {
			handleJWTValidationError(context, claims, err)

			return nil
		}

		context.AuthInfo.OrganizationSerialNumber = claims.Issuer
		context.AuthInfo.PublicKeyFingerprint = publicKeyFingerprint

		context.LogData["delegator"] = claims.Issuer
		context.LogData["orderReference"] = claims.OrderReference

		return next(context)
	}
}

func handleJWTValidationError(context *Context, claims *delegation.JWTClaims, err error) {
	validationError, ok := err.(*jwt.ValidationError)
	if !ok {
		context.Logger.Error("casting error to jwt validation error failed", zap.Error(err))
		inway_http.WriteError(context.Response, "unable to verify claim")

		return
	}

	if errors.Is(validationError.Inner, ErrRequestingOrganizationIsNotDelegatee) {
		context.Logger.Info("requesting organization public key is not the public key found in order", zap.String("delegator", claims.Issuer), zap.String("serviceName", context.Destination.Service.Name))
		inway_http.WriteError(context.Response, "no access. organization serialnumber does not match the delegatee organization serialnumber of the order")

		return
	}

	if errors.Is(validationError.Inner, ErrRequestingOrganizationPublicKeyNotFoundInOrder) {
		context.Logger.Info("requesting organization is not the delegatee", zap.String("delegator", claims.Issuer), zap.String("serviceName", context.Destination.Service.Name))
		inway_http.WriteError(context.Response, "no access. public key of the connection does not match the delegatee public key of the order")

		return
	}

	if errors.Is(validationError.Inner, ErrDelegatorDoesNotHaveAccess) {
		context.Logger.Info("delegator does not have access to service", zap.String("delegator", claims.Issuer), zap.String("serviceName", context.Destination.Service.Name))
		inway_http.WriteError(context.Response, "no access. delegator does not have access to the service for the public key in the claim")

		return
	}

	context.Logger.Error("failed to parse jwt", zap.Error(err))
	inway_http.WriteError(context.Response, "unable to verify claim")
}

func parsePublicKeyFromPEM(publicKeyPEM string) (crypto.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, ErrCannotParsePublicKeyFromPEM
	}

	return x509.ParsePKIXPublicKey(block.Bytes)
}
