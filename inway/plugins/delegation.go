// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/delegation"
)

var (
	ErrServiceNotInClaims                   = errors.New("service is not in claims")
	ErrDelegatorDoesNotHaveAccess           = errors.New("delegator does have access")
	ErrRequestingOrganizationIsNotDelegatee = errors.New("requesting organization is not the delegatee")
	ErrCannotParsePublicKeyFromPEM          = errors.New("failed to parse PEM block containing the public key")
)

type DelegationPlugin struct {
}

func NewDelegationPlugin() *DelegationPlugin {
	return &DelegationPlugin{}
}

func isServiceInClaims(claims *delegation.JWTClaims, serviceName, organizationName string) bool {
	for _, service := range claims.Services {
		if service.Service == serviceName &&
			service.Organization == organizationName {
			return true
		}
	}

	return false
}

func (d *DelegationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		claim := context.Request.Header.Get("X-NLX-Request-Claim")
		if claim == "" {
			return next(context)
		}

		var publicKeyFingerprint string

		claims := &delegation.JWTClaims{}

		_, err := jwt.ParseWithClaims(claim, claims, func(token *jwt.Token) (interface{}, error) {
			if claims.Delegatee != context.AuthInfo.OrganizationName {
				return nil, ErrRequestingOrganizationIsNotDelegatee
			}

			if !isServiceInClaims(claims, context.Destination.Service.Name, context.Destination.Organization) {
				return nil, ErrServiceNotInClaims
			}

			for _, grant := range context.Destination.Service.Grants {
				if grant.OrganizationName == claims.Issuer {
					publicKey, err := parsePublicKeyFromPEM(grant.PublicKeyPEM)
					if err != nil {
						return nil, err
					}

					publicKeyFingerprint = grant.PublicKeyFingerprint

					return publicKey, nil
				}
			}

			return nil, ErrDelegatorDoesNotHaveAccess
		})

		if err != nil {
			handleJWTValidationError(context, claims, err)

			return nil
		}

		context.AuthInfo.OrganizationName = claims.Issuer
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
		http.Error(context.Response, "nlx-inway: unable to verify claim", http.StatusInternalServerError)

		return
	}

	if errors.Is(validationError.Inner, ErrRequestingOrganizationIsNotDelegatee) {
		context.Logger.Info("requesting organization is not the delegatee", zap.String("delegator", claims.Issuer), zap.String("serviceName", context.Destination.Service.Name))
		http.Error(context.Response, "nlx-inway: no access", http.StatusUnauthorized)

		return
	}

	if errors.Is(validationError.Inner, ErrDelegatorDoesNotHaveAccess) {
		context.Logger.Info("delegator does not have access to service", zap.String("delegator", claims.Issuer), zap.String("serviceName", context.Destination.Service.Name))
		http.Error(context.Response, "nlx-inway: no access", http.StatusUnauthorized)

		return
	}

	if errors.Is(validationError.Inner, ErrServiceNotInClaims) {
		context.Logger.Info("delegator does have access but not to this service", zap.String("delegator", claims.Issuer), zap.String("serviceName", context.Destination.Service.Name))
		http.Error(context.Response, "nlx-inway: no access", http.StatusUnauthorized)

		return
	}

	context.Logger.Error("failed to parse jwt", zap.Error(err))
	http.Error(context.Response, "nlx-inway: unable to verify claim", http.StatusInternalServerError)
}

func parsePublicKeyFromPEM(publicKeyPEM string) (crypto.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, ErrCannotParsePublicKeyFromPEM
	}

	return x509.ParsePKIXPublicKey(block.Bytes)
}
