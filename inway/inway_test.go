// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway"
)

var pkiDir = filepath.Join("..", "testing", "pki")

func TestNewInwayException(t *testing.T) {
	logger := zaptest.Logger(t)

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-without-name-chain.pem"),
		filepath.Join(pkiDir, "org-without-name-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	ctx := context.Background()

	_, err := inway.NewInway(ctx, logger, nil, "", "", "localhost:8080", cert, "")
	assert.NotNil(t, err)

	cert, _ = common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	testInway, err := inway.NewInway(ctx, logger, nil, "", "inway.test", "localhost:8080", cert, "")
	assert.Nil(t, err)

	err = testInway.RunServer("invalidlistenaddress", "managementAddress")
	assert.EqualError(t, err, "error listening on TLS server: listen tcp: address invalidlistenaddress: missing port in address")
}

func TestNewInway(t *testing.T) {
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
		description          string
		selfAddress          string
		cert                 *common_tls.CertificateBundle
		monitoringAddress    string
		expectedErrorMessage string
	}{
		{
			"certificates without an organization name",
			"inway.test",
			certOrg,
			"localhost:8080",
			"cannot obtain organization name from self cert",
		},
		{
			"missing monitoring address",
			"inway.test",
			cert,
			"localhost:8080",
			"unable to create monitoring service: address required",
		},
		{
			"selfAddres not in certicate",
			"test.com",
			cert,
			"localhost:8080",
			"'test.com' is not in the list of DNS names of the certificate, [localhost inway.test]",
		},
		{
			"selfAddres must be valid host or host:port",
			"localhost:1:2",
			cert,
			"localhost:8080",
			"Failed to parse selfAddress hostname from 'localhost:1:2': address localhost:1:2: too many colons in address",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			logger := zaptest.Logger(t)
			testProcess := process.NewProcess(logger)

			_, err := inway.NewInway(zaptest.Logger(t), nil, testProcess, "", tt.selfAddress, "", tt.cert, "")
			assert.EqualError(t, err, tt.expectedErrorMessage)
		})
	}
}
