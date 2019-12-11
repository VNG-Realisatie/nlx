// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/outway"
)

func TestNewOutwayExeception(t *testing.T) {
	logger := zap.NewNop()
	tests := []struct {
		config               orgtls.TLSOptions
		authServiceURL       string
		authCAPath           string
		expectedErrorMessage string
	}{
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-without-name.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-without-name-key.pem"),
			},
			"",
			"",
			"cannot obtain organization name from self cert",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-non-existing-key.pem"),
			},
			"",
			"",
			"failed to read tls keypair: open ../testing/pki/org-non-existing-key.pem: no such file or directory",
		}, {
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"http://auth.nlx.io",
			"",
			"authorization service URL set but no CA for authorization provided",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"http://auth.nlx.io",
			"/path/to",
			"scheme of authorization service URL is not 'https'",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"https://auth.nlx.io",
			"/path/to/non-existing.crt",
			"failed to read root CA certificate file `/path/to/non-existing.crt`: open /path/to/non-existing.crt: no such file or directory",
		},
	}

	testProcess := process.NewProcess(logger)
	// Test exceptions during outway creation
	for _, test := range tests {
		_, err := outway.NewOutway(logger, nil, testProcess, test.config, "", test.authServiceURL, test.authCAPath)
		assert.EqualError(t, err, test.expectedErrorMessage)
	}
}
