// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/outway"
)

var pkiDir = filepath.Join("..", "testing", "pki")

func TestNewOutwayExeception(t *testing.T) {
	certOrg, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-without-name-chain.pem"),
		filepath.Join(pkiDir, "org-without-name-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	tests := []struct {
		description              string
		cert                     *common_tls.CertificateBundle
		monitoringServiceAddress string
		authServiceURL           string
		authCAPath               string
		expectedErrorMessage     string
	}{
		{
			"certificate without organization",
			certOrg,
			"localhost:8080",
			"",
			"",
			"cannot obtain organization name from self cert",
		},
		{
			"authorization service URL set but no CA for authorization provided",
			cert,
			"localhost:8080",
			"http://auth.nlx.io",
			"",
			"authorization service URL set but no CA for authorization provided",
		},
		{
			"authorization service URL is not 'https'",
			cert,
			"localhost:8080",
			"http://auth.nlx.io",
			"/path/to",
			"scheme of authorization service URL is not 'https'",
		},
		{
			"invalid monitioring service address",
			cert,
			"",
			"",
			"",
			"unable to create monitoring service: address required",
		},
	}

	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	// Test exceptions during outway creation
	for _, test := range tests {
		_, err := outway.NewOutway(logger, nil, testProcess, test.monitoringServiceAddress, test.cert, "", test.authServiceURL, test.authCAPath, false)
		assert.EqualError(t, err, test.expectedErrorMessage)
	}
}
