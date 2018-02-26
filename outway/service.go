// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var errNoInwaysAvailable = errors.New("no inways available")

// Service handles the proxying of a request to the inway
// TODO: if we like this model in the PoC, Service should become an interface with multiple implementations (http/json, grpc, ??) like inway
type Service struct {
	organizationName string
	serviceName      string

	logger *zap.Logger
	roots  *x509.CertPool

	proxy *httputil.ReverseProxy
}

// NewService creates a new Service instance with a single inway to forward requests to.
// This is a PoC shortcut, the real outway will fetch services and their inways* from the directory
// 		(* note the plural; a service can have multiple inways published by directory so that an outway can balance load to multiple inways)
func NewService(logger *zap.Logger, roots *x509.CertPool, certFile string, keyFile string, organizationName, serviceName string, inwayAddresses []string) (*Service, error) {
	if len(inwayAddresses) == 0 {
		return nil, errNoInwaysAvailable
	}
	inwayAddress := inwayAddresses[0] // no loadbalancing etc. yet in poc, just using the first inway available

	s := &Service{
		organizationName: organizationName,
		serviceName:      serviceName,
		roots:            roots,
	}
	s.logger = logger.With(zap.String("outway-service-full-name", s.fullName()))
	endpointURL, err := url.Parse("https://" + inwayAddress)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint provided")
	}
	endpointURL.Path = "/" + serviceName
	s.proxy = httputil.NewSingleHostReverseProxy(endpointURL)
	transport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		// This can happen when the internals of net/http change.
		// Afaik an interface implementation isn't under the Go1 compatibility promise.
		// TODO: consider setting up a custom http.Transport to use as the proxies RoundTripper.
		panic("http.DefaultTransport must be of type *http.Transport")
	}
	// load client certificate
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, errors.Wrap(err, "invalid certificate provided")
	}
	transport.TLSClientConfig = &tls.Config{
		RootCAs:      roots,
		Certificates: []tls.Certificate{cert},
	}
	s.proxy.Transport = transport
	return s, nil
}

func (s *Service) fullName() string {
	return s.organizationName + "." + s.serviceName
}

func (s *Service) proxyRequest(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}
