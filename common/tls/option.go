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

// WithTLS12 enables TLS v1.2 on a tls.Config instance
//
// Cipher suites are taken from https://ssl-config.mozilla.org/#server=go&config=intermediate&guideline=5.6
// Ordered by the ciphers with the most bits.
func WithTLS12() ConfigOption {
	return func(t *tls.Config) {
		t.MinVersion = tls.VersionTLS12
		t.CipherSuites = []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		}
	}
}
