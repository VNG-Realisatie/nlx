package plugins

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type AuthorizationPlugin struct {
}

func NewAuthorizationPlugin() *AuthorizationPlugin {
	return &AuthorizationPlugin{}
}

func (d *AuthorizationPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		for _, grant := range context.Destination.Service.Grants {
			if grant.OrganizationName == context.AuthInfo.OrganizationName && grant.PublicKeyFingerprint == context.AuthInfo.PublicKeyFingerprint {
				return next(context)
			}
		}

		http.Error(context.Response, fmt.Sprintf(`nlx-inway: permission denied, organization "%s" or public key fingerprint "%s" is not allowed access.`, context.AuthInfo.OrganizationName, context.AuthInfo.PublicKeyFingerprint), http.StatusForbidden)
		context.Logger.Info("unauthorized request blocked, permission denied", zap.String("organization-name", context.AuthInfo.OrganizationName), zap.String("certificate-fingerprint", context.AuthInfo.PublicKeyFingerprint))

		return nil
	}
}
