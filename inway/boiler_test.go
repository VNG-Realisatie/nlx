package inway

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/orgtls"
)

type testDefinition struct {
	url          string
	logRecordID  string
	statusCode   int
	errorMessage string
}

type testEnv struct {
	proxy *httptest.Server
	mock  *httptest.Server
}

// setupClient create a test client with certificates
func setupClient(t *testing.T, tlsOptions orgtls.TLSOptions) http.Client {
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
