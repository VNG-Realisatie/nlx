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
		description              string
		config                   orgtls.TLSOptions
		monitoringServiceAddress string
		authServiceURL           string
		authCAPath               string
		expectedErrorMessage     string
	}{
		{
			"certificate without organization",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-without-name-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-without-name-key.pem"),
			},
			"localhost:8080",
			"",
			"",
			"cannot obtain organization name from self cert",
		},
		{
			"trying to load a non existing certificate",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-non-existing-key.pem"),
			},
			"localhost:8080",
			"",
			"",
			"failed to load organization certificate '../testing/pki/org-nlx-test-chain.pem: open ../testing/pki/org-non-existing-key.pem: no such file or directory",
		}, {
			"authorization service URL set but no CA for authorization provided",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"localhost:8080",
			"http://auth.nlx.io",
			"",
			"authorization service URL set but no CA for authorization provided",
		},
		{
			"authorization service URL is not 'https'",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"localhost:8080",
			"http://auth.nlx.io",
			"/path/to",
			"scheme of authorization service URL is not 'https'",
		},
		{
			"invalid monitioring service address",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"",
			"",
			"",
			"unable to create monitoring service: address required",
		},
	}

	testProcess := process.NewProcess(logger)
	// Test exceptions during outway creation
	for _, test := range tests {
		_, err := outway.NewOutway(logger, nil, testProcess, test.monitoringServiceAddress, test.config, "", test.authServiceURL, test.authCAPath)
		assert.EqualError(t, err, test.expectedErrorMessage)
	}
}
