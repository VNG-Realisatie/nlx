package inway_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/inway"
	"go.nlx.io/nlx/inway/config"
)

func TestNewInwayException(t *testing.T) {
	// Test exceptions NewInway
	logger := zap.NewNop()
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org_without_name.crt",
		OrgKeyFile:  "../testing/org_without_name.key",
	}

	serviceConfig := &config.ServiceConfig{}

	_, err := inway.NewInway(logger, nil, "", tlsOptions, "", serviceConfig)
	if err == nil {
		t.Fatal(`result: err is nil, expected err to be set when calling NewInway with invalid certificates`)
	}

	tests := []struct {
		tlsConfig            orgtls.TLSOptions
		expectedErrorMessage string
	}{
		{
			orgtls.TLSOptions{
				NLXRootCert: "../testing/root.crt",
				OrgCertFile: "../testing/org_without_name.crt",
				OrgKeyFile:  "../testing/org_without_name.key",
			}, "cannot obtain organization name from self cert",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: "../testing/root.crt",
				OrgCertFile: "../testing/org-nlx-test.crt",
				OrgKeyFile:  "../testing/org_non_existing.key",
			},
			"failed to read tls keypair: open ../testing/org_non_existing.key: no such file or directory",
		},
	}

	for _, test := range tests {
		_, err = inway.NewInway(logger, nil, "", test.tlsConfig, "", serviceConfig)
		assert.EqualError(t, err, test.expectedErrorMessage)
	}

	tlsOptions = orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org-nlx-test.key",
	}

	inway, err := inway.NewInway(logger, nil, "", tlsOptions, "", serviceConfig)
	if err != nil {
		t.Fatal(err)
	}
	err = inway.ListenAndServeTLS(process.NewProcess(logger), "invalidlistenaddress")
	assert.EqualError(t, err, "failed to run http server: listen tcp: address invalidlistenaddress: missing port in address")
	if err == nil {
		t.Fatal(`result: error is nil, expected error to be set when calling ListenAndServeTLS with an invalid listen adress`)
	}
}
