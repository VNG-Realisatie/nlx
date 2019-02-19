package outway_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/outway"
)

func TestNewOutwayExeception(t *testing.T) {
	logger := zap.NewNop()
	p := process.NewProcess(logger)
	tests := []struct {
		config               orgtls.TLSOptions
		expectedErrorMessage string
	}{
		{
			orgtls.TLSOptions{
				NLXRootCert: "../testing/root.crt",
				OrgCertFile: "../testing/org_without_name.crt",
				OrgKeyFile:  "../testing/org_without_name.key",
			},
			"cannot obtain organization name from self cert",
		},
		{
			orgtls.TLSOptions{
				NLXRootCert: "../testing/root.crt",
				OrgCertFile: "../testing/org-nlx-test.crt",
				OrgKeyFile:  "../testing/org-non-existing.key",
			},
			"failed to read tls keypair: open ../testing/org-non-existing.key: no such file or directory",
		}, {
			orgtls.TLSOptions{
				NLXRootCert: "../testing/root.crt",
				OrgCertFile: "../testing/org-nlx-test.crt",
				OrgKeyFile:  "../testing/org-nlx-test.key",
			},
			"failed to update internal service directory: failed to fetch services from directory: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = \"transport: Error while dialing dial tcp: missing address\"",
		},
	}
	// Test exceptions during outway creation
	for _, test := range tests {
		_, err := outway.NewOutway(p, logger, nil, test.config, "")
		assert.EqualError(t, err, test.expectedErrorMessage)
	}
}
