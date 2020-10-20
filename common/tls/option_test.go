// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package tls_test

import (
	"crypto/tls"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	common_tls "go.nlx.io/nlx/common/tls"
)

func TestWithTLSClientAuth(t *testing.T) {
	c, err := common_tls.NewBundleFromFiles(
		path.Join(pkiDir, "org-nlx-test-chain.pem"),
		path.Join(pkiDir, "org-nlx-test-key.pem"),
		path.Join(pkiDir, "ca-root.pem"),
	)

	assert.NoError(t, err)

	tc := &tls.Config{} //nolint:gosec // test

	o := c.WithTLSClientAuth()
	o(tc)

	assert.Equal(t, tc.ClientAuth, tls.RequireAndVerifyClientCert)
	assert.Equal(t, tc.ClientCAs, c.RootCAs())
}

func TestWithTLS12(t *testing.T) {
	config := common_tls.NewConfig(common_tls.WithTLS12())

	exptectedCipherSuites := []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
	}

	assert.Equal(t, uint16(tls.VersionTLS12), config.MinVersion)
	assert.Equal(t, exptectedCipherSuites, config.CipherSuites)
}
