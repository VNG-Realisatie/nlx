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
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

var pkiDir = filepath.Join("..", "testing", "pki")

func Test_NewInway(t *testing.T) {
	orgCertWithoutName, orgCert := getCertBundles()

	tests := map[string]struct {
		params               *inway.Params
		expectedErrorMessage string
	}{
		"missing_context": {
			params: &inway.Params{
				Context:           nil,
				Address:           "inway.test",
				OrgCertBundle:     orgCert,
				MonitoringAddress: "localhost:8080",
				Name:              "my-inway",
			},
			expectedErrorMessage: "context is nil. needed to close gracefully",
		},
		"certificates_without_an_organization_name": {
			params: &inway.Params{
				Context:           context.Background(),
				Address:           "inway.test",
				OrgCertBundle:     orgCertWithoutName,
				MonitoringAddress: "localhost:8080",
				Name:              "my-inway",
			},
			expectedErrorMessage: "cannot obtain organization name from self cert",
		},
		"missing_monitoring_address": {
			params: &inway.Params{
				Context:           context.Background(),
				Address:           "inway.test",
				OrgCertBundle:     orgCert,
				MonitoringAddress: "",
				Name:              "my-inway",
			},
			expectedErrorMessage: "unable to create monitoring service: address required",
		},
		"self_address_not_in_certicate": {
			params: &inway.Params{
				Context:           context.Background(),
				Address:           "test.com",
				OrgCertBundle:     orgCert,
				MonitoringAddress: "localhost:8080",
				Name:              "my-inway",
			},
			expectedErrorMessage: "'test.com' is not in the list of DNS names of the certificate, [localhost inway.test]",
		},
		"self_address_must_be_valid_host_or_host:port": {
			params: &inway.Params{
				Context:           context.Background(),
				Address:           "localhost:1:2",
				OrgCertBundle:     orgCert,
				MonitoringAddress: "localhost:8080",
				Name:              "my-inway",
			},
			expectedErrorMessage: "failed to parse address hostname from 'localhost:1:2': address localhost:1:2: too many colons in address",
		},
		"missing_name": {
			params: &inway.Params{
				Context:           context.Background(),
				Address:           "inway.test",
				OrgCertBundle:     orgCert,
				MonitoringAddress: "localhost:8080",
				Name:              "",
			},
			expectedErrorMessage: "a valid name is required (alphanumeric & dashes, max. 100 characters)",
		},
		"invalid_name": {
			params: &inway.Params{
				Context:           context.Background(),
				Address:           "inway.test",
				OrgCertBundle:     orgCert,
				MonitoringAddress: "localhost:8080",
				Name:              "#",
			},
			expectedErrorMessage: "a valid name is required (alphanumeric & dashes, max. 100 characters)",
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

func getCertBundles() (orgCertWithoutName, orgCert *common_tls.CertificateBundle) {
	orgCertWithoutName, _ = common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutName)

	orgCert, _ = common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)

	return orgCertWithoutName, orgCert
}
