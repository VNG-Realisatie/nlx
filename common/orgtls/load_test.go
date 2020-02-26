// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package orgtls_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/orgtls"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		description          string
		options              orgtls.TLSOptions
		expectError          bool
		expectedErrorMessage string
	}{
		{
			description: "happy flow",
			options: orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			expectError:          false,
			expectedErrorMessage: "",
		},
		{
			description: "certificate is not signed by provided root",
			options: orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca-second.pem"),
				OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			expectError:          true,
			expectedErrorMessage: "failed to verify certificate: certificate is signed by 'CN=NLX CA,O=NLX CA' and not by provided root CA of 'CN=NLX Second CA,O=NLX Second CA'",
		},
		{
			description: "the organization certification does not exist",
			options: orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-not-existing.pem"),
				OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			expectError:          true,
			expectedErrorMessage: "failed to load organization certificate '../../testing/pki/org-not-existing.pem: open ../../testing/pki/org-not-existing.pem: no such file or directory",
		},
		{
			description: "the organization key does not exist",
			options: orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-not-exist-key.pem"),
			},
			expectError:          true,
			expectedErrorMessage: "failed to load organization certificate '../../testing/pki/org-nlx-test.pem: open ../../testing/pki/org-nlx-test-not-exist-key.pem: no such file or directory",
		},
		{
			description: "the organization certificate is not a valid certifcate",
			options: orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "..", "testing", "pki", "invalid-cert.pem"),
				OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			expectError:          true,
			expectedErrorMessage: "failed to load organization certificate '../../testing/pki/invalid-cert.pem: tls: failed to find any PEM data in certificate input",
		},
		{
			description: "the root certificate does not exist",
			options: orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca-not-exist.pem"),
				OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			expectError:          true,
			expectedErrorMessage: "failed to load root CA certificate '../../testing/pki/ca-not-exist.pem: failed to open and read certificate file `../../testing/pki/ca-not-exist.pem`: open ../../testing/pki/ca-not-exist.pem: no such file or directory",
		},
		{
			description: "the root certificate is not valid",
			options: orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "..", "testing", "pki", "invalid-cert.pem"),
				OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			expectError:          true,
			expectedErrorMessage: "failed to load root CA certificate '../../testing/pki/invalid-cert.pem: unable to decode pem for certificate `../../testing/pki/invalid-cert.pem`",
		},
	}

	for _, test := range tests {
		_, _, err := orgtls.Load(test.options)

		if test.expectError {
			assert.NotNil(t, err)
			assert.Equal(t, test.expectedErrorMessage, err.Error())
		} else {
			assert.Nil(t, err)
		}
	}
}
