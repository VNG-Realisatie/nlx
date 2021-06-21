// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway"
)

var pkiDir = filepath.Join("..", "testing", "pki")

func Test_NewInway(t *testing.T) {
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
		params               *inway.Params
		expectedErrorMessage string
	}{
		"missing_context": {
			params: &inway.Params{
				Context:           nil,
				SelfAddress:       "inway.test",
				OrgCertBundle:     cert,
				MonitoringAddress: "localhost:8080",
			},
			expectedErrorMessage: "context is nil. needed to close gracefully",
		},
		"certificates_without_an_organization_name": {
			params: &inway.Params{
				Context:           context.Background(),
				SelfAddress:       "inway.test",
				OrgCertBundle:     certOrg,
				MonitoringAddress: "localhost:8080",
			},
			expectedErrorMessage: "cannot obtain organization name from self cert",
		},
		"missing_monitoring_address": {
			params: &inway.Params{
				Context:           context.Background(),
				SelfAddress:       "inway.test",
				OrgCertBundle:     cert,
				MonitoringAddress: "",
			},
			expectedErrorMessage: "unable to create monitoring service: address required",
		},
		"self_address_not_in_certicate": {
			params: &inway.Params{
				Context:           context.Background(),
				SelfAddress:       "test.com",
				OrgCertBundle:     cert,
				MonitoringAddress: "localhost:8080",
			},
			expectedErrorMessage: "'test.com' is not in the list of DNS names of the certificate, [localhost inway.test]",
		},
		"self_address_must_be_valid_host_or_host:port": {
			params: &inway.Params{
				Context:           context.Background(),
				SelfAddress:       "localhost:1:2",
				OrgCertBundle:     cert,
				MonitoringAddress: "localhost:8080",
			},
			expectedErrorMessage: "failed to parse selfAddress hostname from 'localhost:1:2': address localhost:1:2: too many colons in address",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			_, err := inway.NewInway(tt.params)
			assert.EqualError(t, err, tt.expectedErrorMessage)
		})
	}
}
