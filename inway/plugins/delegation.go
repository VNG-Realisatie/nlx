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

var ErrDelegatorDoesNotHaveAccess = errors.New("delegator does have access")
var ErrCannotParsePublicKeyFromPEM = errors.New("failed to parse PEM block containing the public key")
var ErrRequestingOrganizationIsNotDelegatee = errors.New("requesting organization is not the delegatee")

type DelegationPlugin struct {
}

func NewDelegationPlugin() *DelegationPlugin {
	return &DelegationPlugin{}
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
			validationError, ok := err.(*jwt.ValidationError)
			if !ok {
				context.Logger.Error("casting error to jwt validation error failed", zap.Error(err))
				http.Error(context.Response, "nlx-inway: unable to verify claim", http.StatusInternalServerError)

				return nil
			}

			if errors.Is(validationError.Inner, ErrRequestingOrganizationIsNotDelegatee) {
				context.Logger.Info("requesting organization is not the delegatee", zap.String("delegator", claims.Issuer), zap.String("serviceName", context.Destination.Service.Name))
				http.Error(context.Response, "nlx-inway: no access", http.StatusUnauthorized)

				return nil
			}

			if errors.Is(validationError.Inner, ErrDelegatorDoesNotHaveAccess) {
				context.Logger.Info("delegator does not have access to service", zap.String("delegator", claims.Issuer), zap.String("serviceName", context.Destination.Service.Name))
				http.Error(context.Response, "nlx-inway: no access", http.StatusUnauthorized)

				return nil
			}

			context.Logger.Error("failed to parse jwt", zap.Error(err))
			http.Error(context.Response, "nlx-inway: unable to verify claim", http.StatusInternalServerError)

			return nil
		}

		context.AuthInfo.OrganizationName = claims.Issuer
		context.AuthInfo.PublicKeyFingerprint = publicKeyFingerprint

		context.LogData["delegator"] = claims.Issuer
		context.LogData["orderReference"] = claims.OrderReference

		return next(context)
	}
}

func parsePublicKeyFromPEM(publicKeyPEM string) (crypto.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, ErrCannotParsePublicKeyFromPEM
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}
