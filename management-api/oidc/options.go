// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package oidc

// Options contains go-flags fields which can be used to configure oidc
type Options struct {
	SecretKey           string `long:"secret-key" env:"SECRET_KEY" description:"Secret key that is used for signing sessions" required:"true"`
	ClientID            string `long:"oidc-client-id" env:"OIDC_CLIENT_ID" description:"The OIDC client ID" required:"true"`
	ClientSecret        string `long:"oidc-client-secret" env:"OIDC_CLIENT_SECRET" description:"The OIDC client secret" required:"true"`
	DiscoveryURL        string `long:"oidc-discovery-url" env:"OIDC_DISCOVERY_URL" description:"The OIDC discovery URL" required:"true"`
	RedirectURL         string `long:"oidc-redirect-url" env:"OIDC_REDIRECT_URL" description:"The OIDC redirect URL" required:"true"`
	SessionCookieSecure bool   `long:"session-cookie-secure" env:"SESSION_COOKIE_SECURE" description:"Use 'secure' cookies"`
}
