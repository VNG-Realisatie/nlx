// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/plugins"
)

func TestAuthenticationPlugin(t *testing.T) {
	authenticationPlugin := plugins.NewAuthenticationPlugin()

	var pkiDir = filepath.Join("..", "..", "testing", "pki")

	tests := map[string]struct {
		certifcate         *x509.Certificate
		expectedStatusCode int
		expectedMessage    string
	}{
		"invalid_certificate": {
			certifcate: func() *x509.Certificate {
				cert, _ := common_tls.NewBundleFromFiles(
					filepath.Join(pkiDir, "org-without-name-chain.pem"),
					filepath.Join(pkiDir, "org-without-name-key.pem"),
					filepath.Join(pkiDir, "ca-root.pem"),
				)

				return cert.Certificate()
			}(),
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "nlx-inway: invalid certificate provided: missing organizations attribute in subject\n",
		},
		"happy_flow": {
			certifcate: func() *x509.Certificate {
				cert, _ := common_tls.NewBundleFromFiles(
					filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
					filepath.Join(pkiDir, "org-nlx-test-key.pem"),
					filepath.Join(pkiDir, "ca-root.pem"),
				)

				return cert.Certificate()
			}(),
			expectedStatusCode: http.StatusOK,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			context := fakeContext(&plugins.Destination{}, tt.certifcate, &plugins.AuthInfo{})

			err := authenticationPlugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()
			defer response.Body.Close()

			contents, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedMessage, string(contents))
			assert.Equal(t, tt.expectedStatusCode, response.StatusCode)
		})
	}
}
