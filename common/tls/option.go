// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package tls

import (
	"crypto/tls"
)

type ConfigOption func(c *tls.Config)

func (c *CertificateBundle) WithTLSClientAuth() ConfigOption {
	return func(t *tls.Config) {
		t.ClientAuth = tls.RequireAndVerifyClientCert
		t.ClientCAs = c.rootCAs
	}
}
