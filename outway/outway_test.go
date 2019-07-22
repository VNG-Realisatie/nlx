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
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org_without_name.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org_without_name.key"),
			},
			"",
			"",
			"cannot obtain organization name from self cert",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org-non-existing.key"),
			},
			"",
			"",
			"failed to read tls keypair: open ../testing/org-non-existing.key: no such file or directory",
		}, {
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
			},
			"http://auth.nlx.io",
			"",
			"authorization service URL set but no CA for authorization provided",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
			},
			"http://auth.nlx.io",
			"/path/to",
			"scheme of authorization service URL is not 'https'",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
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
