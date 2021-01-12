// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package http

import (
	"net"
	"net/http"
	"time"

	common_tls "go.nlx.io/nlx/common/tls"
)

func NewHTTPClient(certificate *common_tls.CertificateBundle) *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       certificate.TLSConfig(),
	}
	return &http.Client{
		Transport: transport,
	}
}
