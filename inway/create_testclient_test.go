package inway

import (
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/orgtls"
)

// setupClient create a test client with certificates
func setupClient(t *testing.T, tlsOptions orgtls.TLSOptions) http.Client {
	cert, err := tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	assert.Nil(t, err)
	pool, err := orgtls.LoadRootCert(tlsOptions.NLXRootCert)
	assert.Nil(t, err)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, //nolint
			RootCAs:            pool,
			Certificates:       []tls.Certificate{cert}}}

	client := http.Client{
		Transport: tr,
	}
	return client
}
