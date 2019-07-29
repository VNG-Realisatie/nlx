package http

import (
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"time"
)

func NewHTTPClient(rootCA *x509.CertPool, certKeyPair *tls.Certificate) *http.Client {
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
		TLSClientConfig: &tls.Config{
			RootCAs:      rootCA,
			Certificates: []tls.Certificate{*certKeyPair},
		},
	}
	return &http.Client{
		Transport: transport,
	}
}
