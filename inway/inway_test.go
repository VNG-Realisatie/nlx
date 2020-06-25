// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway_test

import (
	"path/filepath"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway"
)

func TestNewInwayException(t *testing.T) {
	// Test exceptions NewInway
	logger := zaptest.Logger(t)
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
		OrgCertFile: filepath.Join("..", "testing", "pki", "org-without-name-chain.pem"),
		OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-without-name-key.pem"),
	}

	testProcess := process.NewProcess(logger)

	_, err := inway.NewInway(logger, nil, testProcess, "", "", "localhost:8080", tlsOptions, "")
	assert.NotNil(t, err)

	tlsOptions = orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
		OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
		OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
	}

	testInway, err := inway.NewInway(logger, nil, testProcess, "", "inway.test", "localhost:8080", tlsOptions, "")
	assert.Nil(t, err)

	err = testInway.RunServer("invalidlistenaddress")
	assert.EqualError(t, err, "error listening on TLS server: listen tcp: address invalidlistenaddress: missing port in address")
}

func TestNewInway(t *testing.T) {
	tests := []struct {
		description          string
		selfAddress          string
		tlsConfig            orgtls.TLSOptions
		monitoringAddress    string
		expectedErrorMessage string
	}{
		{
			"certificates without an organization name",
			"inway.test",
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
			"inway.test",
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
			"inway.test",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"localhost:8080",
			"unable to create monitoring service: address required",
		},
		{
			"selfAddres not in certicate",
			"test.com",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"localhost:8080",
			"'test.com' is not in the list of DNS names of the certificate, [localhost inway.test]",
		},
		{
			"selfAddres must be valid host or host:port",
			"localhost:1:2",
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
				OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
				OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
			},
			"localhost:8080",
			"Failed to parse selfAddress hostname from 'localhost:1:2': address localhost:1:2: too many colons in address",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			logger := zaptest.Logger(t)
			testProcess := process.NewProcess(logger)

			_, err := inway.NewInway(zaptest.Logger(t), nil, testProcess, "", tt.selfAddress, "", tt.tlsConfig, "")
			assert.EqualError(t, err, tt.expectedErrorMessage)
		})
	}
}
