package api

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	mock_authorization "go.nlx.io/nlx/management-api/authorization/mock"
	mock_session "go.nlx.io/nlx/management-api/session/mock"
)

var tests = []struct {
	name                 string
	tlsOptions           orgtls.TLSOptions
	configAPIAddress     string
	expectedErrorMessage string
}{
	{
		"1",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-without-name.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-without-name-key.pem"),
		},
		"",
		"cannot obtain organization name from self cert",
	},
	{
		"2",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-non-existing-key.pem"),
		},
		"",
		"failed to load x509 keypair for organization: open ../../testing/pki/org-non-existing-key.pem: no such file or directory",
	},
	{
		"3",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
		},
		"",
		"config API address is not configured",
	},
	{
		"4",
		orgtls.TLSOptions{
			NLXRootCert: filepath.Join("..", "..", "testing", "pki", "ca.pem"),
			OrgCertFile: filepath.Join("..", "..", "testing", "pki", "org-nlx-test.pem"),
			OrgKeyFile:  filepath.Join("..", "..", "testing", "pki", "org-nlx-test-key.pem"),
		},
		"config-api.test:8443",
		"",
	},
}

func TestNewAPI(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authenticationManager := mock_session.NewMockAuthenticationManager(mockCtrl)
	authorizer := mock_authorization.NewMockAuthorizer(mockCtrl)

	// Test exceptions during management-api creation
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			fmt.Printf("%+v", test.tlsOptions)
			_, err := NewAPI(logger, testProcess, test.tlsOptions, test.configAPIAddress, authenticationManager, authorizer)

			if test.expectedErrorMessage != "" {
				assert.EqualError(t, err, test.expectedErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
