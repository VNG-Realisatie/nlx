package api

import (
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
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org_without_name.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org_without_name.key"),
			},
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
			"failed to load x509 keypair for organization",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
			},
			"",
			"config API address is not configured",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
			},
			"config-api.test:8443",
			"",
		},
	}

	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	// Test exceptions during management-api creation
	for _, test := range tests {
		_, err := NewAPI(logger, testProcess, test.tlsOptions, test.configAPIAddress)

		if test.expectedErrorMessage != "" {
			assert.EqualError(t, err, test.expectedErrorMessage)
		} else {
			assert.Nil(t, err)
		}
	}
}
