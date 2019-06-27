
package inway

import (
    "crypto/tls"
    "testing"
    "net/http"
    "go.nlx.io/nlx/common/orgtls"
    "github.com/stretchr/testify/assert"
)

type testDefinition struct {
	url          string
	logRecordID  string
	statusCode   int
	errorMessage string
}

// SetupClient create a test client with certificates
func SetupClient(tlsOptions orgtls.TLSOptions, t *testing.T) http.Client {
	cert, err := tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	assert.Nil(t, err)
	pool, err := orgtls.LoadRootCert(tlsOptions.NLXRootCert)
	assert.Nil(t, err)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            pool,
			Certificates:       []tls.Certificate{cert}}}

	client := http.Client{
		Transport: tr,
	}
	return client
}

