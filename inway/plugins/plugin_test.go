// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/http/httptest"

	"go.uber.org/zap"

	"go.nlx.io/nlx/inway/plugins"
)

func nopServeFunc(context *plugins.Context) error {
	return nil
}

func fakeContext(dest *plugins.Destination, cert *x509.Certificate, authInfo *plugins.AuthInfo) *plugins.Context {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)
	request.TLS = &tls.ConnectionState{
		PeerCertificates: []*x509.Certificate{cert},
	}

	return &plugins.Context{
		Destination: dest,
		Request:     request,
		Response:    recorder,
		Logger:      zap.NewNop(),
		LogData:     map[string]string{},
		AuthInfo:    authInfo,
	}
}
