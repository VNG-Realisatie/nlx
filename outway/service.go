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

// HTTPService abstracts HTTP-based services to a single usable format.
type HTTPService interface {
	FullName() string
	ProxyHTTPRequest(w http.ResponseWriter, r *http.Request)
}

// SimpleHTTPService handles the proxying of a request to the inway
type SimpleHTTPService struct {
	organizationName string
	serviceName      string

	logger *zap.Logger
	roots  *x509.CertPool

	proxy *httputil.ReverseProxy
}

// NewSimpleHTTPService creates a SimpleHTTPService instance with a single inway to forward requests to.
func NewSimpleHTTPService(logger *zap.Logger, roots *x509.CertPool, certFile string, keyFile string, organizationName, serviceName string, inwayAddresses []string) (*SimpleHTTPService, error) {
	if len(inwayAddresses) == 0 {
		return nil, errNoInwaysAvailable
	}
	inwayAddress := inwayAddresses[0] // no loadbalancing etc. yet in poc, just using the first inway available

	s := &SimpleHTTPService{
		organizationName: organizationName,
		serviceName:      serviceName,
		roots:            roots,
	}
	s.logger = logger.With(zap.String("outway-service-full-name", s.FullName()))
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
		// TODO: #209 consider setting up a custom http.Transport to use as the proxies RoundTripper.
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

// FullName returns the name of the service
func (s *SimpleHTTPService) FullName() string {
	return s.organizationName + "." + s.serviceName
}

// ProxyHTTPRequest procies the HTTP request to the proper endpoint.
func (s *SimpleHTTPService) ProxyHTTPRequest(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}
