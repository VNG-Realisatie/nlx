// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"crypto/x509"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/common/httperrors"
	"go.nlx.io/nlx/inway/plugins"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

func TestAuthenticationPlugin(t *testing.T) {
	authenticationPlugin := plugins.NewAuthenticationPlugin()

	var pkiDir = filepath.Join("..", "..", "testing", "pki")

	tests := map[string]struct {
		certifcate         *x509.Certificate
		wantHTTPStatusCode int
		wantErr            *httperrors.NLXNetworkError
	}{
		"invalid_certificate": {
			certifcate: func() *x509.Certificate {
				cert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutName)
				require.NoError(t, err)

				return cert.Certificate()
			}(),
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.InvalidCertificateErr,
				Message:  "invalid certificate provided: missing organizations attribute in subject",
			},
		},
		"invalid_certificate_without_serial_number": {
			certifcate: func() *x509.Certificate {
				cert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutSerialNumber)
				require.NoError(t, err)

				return cert.Certificate()
			}(),
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.InvalidCertificateErr,
				Message:  "invalid certificate provided: missing or invalid value for serial number in subject",
			},
		},
		"happy_flow": {
			certifcate: func() *x509.Certificate {
				cert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
				require.NoError(t, err)

				return cert.Certificate()
			}(),
			wantHTTPStatusCode: http.StatusOK,
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

			contents, err := io.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantHTTPStatusCode, response.StatusCode)

			if tt.wantErr != nil {
				gotError := &httperrors.NLXNetworkError{}
				err := json.Unmarshal(contents, gotError)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantErr, gotError)
			}
		})
	}
}
