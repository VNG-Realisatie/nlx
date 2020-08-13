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
