// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package oidc

type Options struct {
	SecretKey           string
	ClientID            string
	ClientSecret        string
	DiscoveryURL        string
	RedirectURL         string
	SessionCookieSecure bool
}
