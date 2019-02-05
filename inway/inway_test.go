package inway_test

import (
	"strings"
	"testing"

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

	if strings.Compare(err.Error(), "cannot obtain organization name from self cert") != 0 {
		t.Fatalf(`result: %s, expected error to be "cannot obtain organization name from self cert"`, err.Error())
	}

	tlsOptions = orgtls.TLSOptions{
		NLXRootCert: "../testing/root.crt",
		OrgCertFile: "../testing/org-nlx-test.crt",
		OrgKeyFile:  "../testing/org_non_existing.key",
	}

	_, err = inway.NewInway(logger, nil, "", tlsOptions, "", serviceConfig)
	if err == nil {
		t.Fatal(`result: err is nil, expected err to be set when calling NewInway without a existing .key file`)
	}

	if !strings.HasPrefix(err.Error(), "failed to read tls keypair") {
		t.Fatalf(`result: %s, expected error to start with "failed to read tls keypair"`, err.Error())
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
	if err == nil {
		t.Fatal(`result: error is nil, expected error to be set when calling ListenAndServeTLS with an invalid listen adress`)
	}

	if !strings.HasPrefix(err.Error(), "failed to run http server") {
		t.Fatalf(`result: %s, expected error to start with "failed to run http server"`, err.Error())
	}
}
