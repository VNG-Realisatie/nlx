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
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/inway/plugins"
	common_testing "go.nlx.io/nlx/testing/testingutils"
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
				cert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutName)
				require.NoError(t, err)

				return cert.Certificate()
			}(),
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "nlx-inway: invalid certificate provided: missing organizations attribute in subject\n",
		},
		"invalid_certificate_without_serial_number": {
			certifcate: func() *x509.Certificate {
				cert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutSerialNumber)
				require.NoError(t, err)

				return cert.Certificate()
			}(),
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "nlx-inway: invalid certificate provided: missing value for serial number in subject\n",
		},
		"happy_flow": {
			certifcate: func() *x509.Certificate {
				cert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
				require.NoError(t, err)

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
