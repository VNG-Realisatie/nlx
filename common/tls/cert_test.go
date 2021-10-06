// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package tls_test

import (
	"crypto/tls"
	"errors"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	common_tls "go.nlx.io/nlx/common/tls"
)

var pkiDir = filepath.Join("..", "..", "testing", "pki")

type files struct {
	cert     string
	key      string
	rootCert string
}

func TestNewBundle(t *testing.T) {
	testCases := []struct {
		name        string
		files       files
		expectError error
	}{
		{
			name: "valid_files",
			files: files{
				filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-root.pem"),
			},
			expectError: nil,
		},
		{
			name: "cert_does_not_exist",
			files: files{
				filepath.Join(pkiDir, "org-not-existing.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-root-second.pem"),
			},
			expectError: errors.New("failed to read certificate file: open ../../testing/pki/org-not-existing.pem: no such file or directory"),
		},
		{
			name: "cert_key_does_not_exist",
			files: files{
				filepath.Join(pkiDir, "org-nlx-test.pem"),
				filepath.Join(pkiDir, "org-nlx-test-not-exist-key.pem"),
				filepath.Join(pkiDir, "ca-root.pem"),
			},
			expectError: errors.New("failed to read private key file: open ../../testing/pki/org-nlx-test-not-exist-key.pem: no such file or directory"),
		},
		{
			name: "root_cert_does_not_exist",
			files: files{
				filepath.Join(pkiDir, "org-nlx-test.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-not-exist.pem"),
			},
			expectError: errors.New("failed to read root certificate file: open ../../testing/pki/ca-not-exist.pem: no such file or directory"),
		},
		{
			name: "cert_is_not_valid",
			files: files{
				filepath.Join(pkiDir, "invalid-cert.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-root.pem"),
			},
			expectError: errors.New("failed to parse certificate/key pair: tls: failed to find any PEM data in certificate input"),
		},
		{
			name: "root_cert_is_not_valid",
			files: files{
				filepath.Join(pkiDir, "org-nlx-test.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "invalid-cert.pem"),
			},
			expectError: errors.New("failed to parse root CA certificate: unable to decode pem for certificate"),
		},
		{
			name: "not_signed_by_provided_root",
			files: files{
				filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-root-second.pem"),
			},
			expectError: errors.New("failed to verify certificate: certificate is signed by 'CN=NLX Intermediate CA,O=NLX Intermediate CA' and not by provided root CA of 'CN=NLX Second CA,O=NLX Second CA'"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			c, err := common_tls.NewBundleFromFiles(
				tc.files.cert,
				tc.files.key,
				tc.files.rootCert,
			)

			if tc.expectError != nil {
				assert.Nil(t, c)
				assert.EqualError(t, err, tc.expectError.Error())
			} else {
				assert.NotNil(t, c)
				assert.NoError(t, err)
			}
		})
	}
}

func TestBundle(t *testing.T) {
	c, err := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	assert.NoError(t, err)
	assert.Equal(t, "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc=", c.PublicKeyFingerprint())
	assert.Equal(t, uint16(tls.VersionTLS13), c.TLSConfig().MinVersion)
}

func TestVerifyPrivateKeyPermissions(t *testing.T) {
	tests := []struct {
		name          string
		permissions   os.FileMode
		expectedError error
	}{
		{
			"execute",
			0700,
			errors.New("file permissions too open. the file should not allow execution or be readable and writeable for everybody"),
		},
		{
			"write_for_all",
			0604,
			errors.New("file permissions too open. the file should not allow execution or be readable and writeable for everybody"),
		},
		{
			"read_for_all",
			0602,
			errors.New("file permissions too open. the file should not allow execution or be readable and writeable for everybody"),
		},
		{
			"write_and_read_for_group",
			0660,
			nil,
		},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			filePath := path.Join(tempDir, "file.permission")

			newFile, err := os.Create(filePath)
			assert.NoError(t, err)

			defer newFile.Close()

			err = os.Chmod(filePath, tc.permissions)
			assert.NoError(t, err)

			err = common_tls.VerifyPrivateKeyPermissions(filePath)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestPemPublicKeyFingerprint(t *testing.T) {
	// copied from `/management-api/pkg/server/claim_test.go`
	testPublicKeyPEM := `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAxahK/ruBG74MZ2/Z71ll
WS1OMJy6xs9qpQY7YC7C+u59JoqNdoWToV6EZRfPYzh61BWsyKlqvkl11HhD6HVk
WmdidYJtmoJqmFeWm02a6RkP4XbBiOCm9xxX/xlZRWubTCswaLkfmlI2IYxgpLIP
QuvO2nbor8YG4dS7u7u7yrtl1dOBD1utlMCzX5j8vG+BaHUqE1kBWIE5kg9ogeVf
wa/w30EPcD+5gdknn5uGoTFP/xi6WiZ+6MJli1CPjrHX0N73ZMSdgHK+4jk8Kdrz
Fou5sNtCl+CTdzhDhwYJxJv/McsgqPfXsOdk0T3QUcCqWsawJ8VblJYYwyj1WW7l
bJSygJjvOTG+C2+vbht3mKvimKpx/+8S/Zg+g7nen//SvFQhe2wI7Eaottgk/abU
6i3ntvSty4EyxFPnchKa7EXeFAsp4stO0Q5iTE4rEdDotwaWrmcN54UQr2ZOVPJ/
BGGG6SxeciX9jB9I1FHBngMyiXVDgMlgGa9Ke3y1V+Yaqh3LOp6JXnjXp50Ke0nc
CMa7tBd6GGJqV4hl3daYj7yyBWzB3E2d/u+gJx1e9mxqgA0V7nidh2CRelHtczhC
O5/DpYFGnjKm4YMkzSb7CxRDrL2OJeyvM3tKyRZES5eEiedMcpjvm5ULzZeCp2r3
P72Jy9qTigqNYoIHBYMpFzUCAwEAAQ==
-----END PUBLIC KEY-----
`
	expectedPublicKeyFingerprint := "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc="

	fingerprint, err := common_tls.PemPublicKeyFingerprint([]byte(testPublicKeyPEM))
	assert.NoError(t, err)
	assert.Equal(t, expectedPublicKeyFingerprint, fingerprint)
}
