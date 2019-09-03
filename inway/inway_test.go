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
		NLXRootCert: filepath.Join("..", "testing", "root.crt"),
		OrgCertFile: filepath.Join("..", "testing", "org_without_name.crt"),
		OrgKeyFile:  filepath.Join("..", "testing", "org_without_name.key"),
	}

	testProcess := process.NewProcess(logger)

	_, err := inway.NewInway(logger, nil, testProcess, "", "", tlsOptions, "")
	assert.NotNil(t, err)

	tests := []struct {
		tlsConfig            orgtls.TLSOptions
		expectedErrorMessage string
	}{
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org_without_name.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org_without_name.key"),
			}, "cannot obtain organization name from self cert",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: filepath.Join("..", "testing", "root.crt"),
				OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
				OrgKeyFile:  filepath.Join("..", "testing", "org-non-existing.key"),
			},
			"failed to read tls keypair: open ../testing/org-non-existing.key: no such file or directory",
		},
	}

	for _, test := range tests {
		_, err = inway.NewInway(logger, nil, testProcess, "", "", test.tlsConfig, "")
		assert.EqualError(t, err, test.expectedErrorMessage)
	}

	tlsOptions = orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "root.crt"),
		OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
		OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
	}

	testInway, err := inway.NewInway(logger, nil, testProcess, "", "", tlsOptions, "")
	if err != nil {
		t.Fatal(err)
	}
	err = testInway.ListenAndServeTLS("invalidlistenaddress")
	assert.EqualError(t, err, "failed to run http server: listen tcp: address invalidlistenaddress: missing port in address")
	if err == nil {
		t.Fatal(`result: error is nil, expected error to be set when calling ListenAndServeTLS with an invalid listen address`)
	}
}
