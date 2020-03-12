// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway"
)

func TestNewInwayException(t *testing.T) {
	// Test exceptions NewInway
	logger := zap.NewNop()
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
		OrgCertFile: filepath.Join("..", "testing", "pki", "org-without-name-chain.pem"),
		OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-without-name-key.pem"),
	}

	testProcess := process.NewProcess(logger)

	_, err := inway.NewInway(logger, nil, testProcess, "", "", "localhost:8080", tlsOptions, "")
	assert.NotNil(t, err)

	tests := []struct {
		description          string
		tlsConfig            orgtls.TLSOptions
		monitoringAddress    string
		expectedErrorMessage string
	}{
		{
			"certificates without an organization name",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-without-name-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-without-name-key.pem"),
			},
			"localhost:8080",
			"cannot obtain organization name from self cert",
		},
		{
			"missing organization certificate",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-non-existing-key.pem"),
			},
			"localhost:8080",
			"failed to load organization certificate '../testing/pki/org-nlx-test-chain.pem: open ../testing/pki/org-non-existing-key.pem: no such file or directory",
		},
		{
			"missing monitoring address",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"localhost:8080",
			"unable to create monitoring service: : address required",
		},
	}

	for _, test := range tests {
		_, err = inway.NewInway(logger, nil, testProcess, "", "", "", test.tlsConfig, "")
		assert.EqualError(t, err, test.expectedErrorMessage)
	}

	tlsOptions = orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
		OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
		OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
	}

	testInway, err := inway.NewInway(logger, nil, testProcess, "", "", "localhost:8080", tlsOptions, "")
	assert.Nil(t, err)

	err = testInway.RunServer("invalidlistenaddress")
	assert.EqualError(t, err, "error listening on TLS server: listen tcp: address invalidlistenaddress: missing port in address")
}
