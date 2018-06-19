// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"crypto/tls"
	"net/http"

	"github.com/pkg/errors"
)

// ListenAndServeTLS is a blocking function that listens on provided tcp address to handle requests.
func (i *Inway) ListenAndServeTLS(address string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/.nlx/api-spec-doc/", i.handleAPISpecDocRequest)
	serveMux.HandleFunc("/.nlx/health/", i.handleHealthRequest)
	serveMux.Handle("/.nlx/", http.NotFoundHandler())
	serveMux.HandleFunc("/", i.handleProxyRequest)
	server := &http.Server{
		Addr: address,
		TLSConfig: &tls.Config{
			// only allow clients that present a cert signed by our root CA
			ClientCAs:  i.roots,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
		Handler: serveMux,
	}
	err := server.ListenAndServeTLS(i.orgCertFile, i.orgKeyFile)
	if err != nil {
		return errors.Wrap(err, "failed to run http server")
	}
	return nil
}
