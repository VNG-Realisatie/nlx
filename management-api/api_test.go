package api

import (
	"fmt"
	"path/filepath"
	"testing"

	"go.nlx.io/nlx/common/process"
	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/common/orgtls"
)

func TestNewAPI(t *testing.T) {
	tests := []struct {
		tlsOptions           orgtls.TLSOptions
		configAPIAddress     string
		expectedErrorMessage string
	}{
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-without-name.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-without-name-key.pem"),
			},
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
			"failed to load x509 keypair for organization: open ../testing/pki/org-non-existing-key.pem: no such file or directory",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"",
			"config API address is not configured",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"config-api.test:8443",
			"",
		},
	}

	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	// Test exceptions during management-api creation
	for _, test := range tests {
		fmt.Printf("%+v", test.tlsOptions)
		_, err := NewAPI(logger, testProcess, test.tlsOptions, test.configAPIAddress)

		if test.expectedErrorMessage != "" {
			assert.EqualError(t, err, test.expectedErrorMessage)
		} else {
			assert.Nil(t, err)
		}
	}
}
