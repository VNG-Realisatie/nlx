// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway"
)

var pkiDir = filepath.Join("..", "testing", "pki")

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

	tests := map[string]struct {
		context              context.Context
		selfAddress          string
		cert                 *common_tls.CertificateBundle
		monitoringAddress    string
		expectedErrorMessage string
	}{
		"missing_context": {
			nil,
			"inway.test",
			cert,
			"localhost:8080",
			"context is nil. needed to close gracefully",
		},
		"certificates_without_an_organization_name": {
			context.Background(),
			"inway.test",
			certOrg,
			"localhost:8080",
			"cannot obtain organization name from self cert",
		},
		"missing_monitoring_address": {
			context.Background(),
			"inway.test",
			cert,
			"",
			"unable to create monitoring service: address required",
		},
		"self_address_not_in_certicate": {
			context.Background(),
			"test.com",
			cert,
			"localhost:8080",
			"'test.com' is not in the list of DNS names of the certificate, [localhost inway.test]",
		},
		"self_address_must_be_valid_host_or_host:port": {
			context.Background(),
			"localhost:1:2",
			cert,
			"localhost:8080",
			"failed to parse selfAddress hostname from 'localhost:1:2': address localhost:1:2: too many colons in address",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			_, err := inway.NewInway(tt.context, zap.NewNop(), nil, nil, nil, "", tt.selfAddress, tt.monitoringAddress, "", tt.cert, "")
			assert.EqualError(t, err, tt.expectedErrorMessage)
		})
	}
}
