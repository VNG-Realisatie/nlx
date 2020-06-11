package api

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/oidc"
)

var tests = []struct {
	name                         string
	tlsOptions                   orgtls.TLSOptions
	etcdConnectionString         string
	directoryRegistrationAddress string
	directoryEndpointURL         string
	expectedErrorMessage         string
}{
	{
		"1",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca-root.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-without-name-chain.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-without-name-key.pem"),
		},
		"",
		"",
		"",
		"cannot obtain organization name from self cert",
	},
	{
		"2",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca-root.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-non-existing-key.pem"),
		},
		"",
		"",
		"",
		"failed to load tls certs: failed to load organization certificate '../../testing/pki/org-nlx-test.pem: open ../../testing/pki/org-non-existing-key.pem: no such file or directory",
	},
	{
		"3",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca-root.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test-chain.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
		},
		"",
		"",
		"",
		"etcd connection string is not configured",
	},
	{
		"4",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca-root.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test-chain.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
		},
		"etcd.test:2379",
		"",
		"",
		"directory registration address is not configured",
	},
	{
		"5",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca-root.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test-chain.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
		},
		"etcd.test:2379",
		"directory-registration.test:8443",
		"",
		"directory endpoint URL is not configured",
	},
	{
		"6",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca-root.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test-chain.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
		},
		"etcd.test:2379",
		"directory-registration.test:8443",
		"https://directory.test/",
		"",
	},
}

func TestNewAPI(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	// Test exceptions during management-api creation
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			fmt.Printf("%+v", test.tlsOptions)
			_, err := NewAPI(logger, testProcess, test.tlsOptions, test.etcdConnectionString, test.directoryRegistrationAddress, test.directoryEndpointURL, &oidc.Authenticator{})

			if test.expectedErrorMessage != "" {
				assert.EqualError(t, err, test.expectedErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
