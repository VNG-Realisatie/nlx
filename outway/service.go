// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var errNoInwaysAvailable = errors.New("no inways available")

// HTTPService abstracts HTTP-based services to a single usable format.
type HTTPService interface {
	FullName() string
	ProxyHTTPRequest(w http.ResponseWriter, r *http.Request)
	GetInwayAddresses() []string
}

// RoundRobinLoadBalancedHTTPService handles the proxying of a request to the inway
type RoundRobinLoadBalancedHTTPService struct {
	organizationName string
	serviceName      string

	inwayAddresses  []string
	loadBalanceLock sync.Mutex
	count           int
	endPoints       []*url.URL

	logger *zap.Logger
	roots  *x509.CertPool

	proxies []*httputil.ReverseProxy
}

// NewRoundRobinLoadBalancedHTTPService creates a RoundRobinLoadBalancedHTTPService
func NewRoundRobinLoadBalancedHTTPService(logger *zap.Logger, roots *x509.CertPool, certFile string, keyFile string, organizationName, serviceName string, inwayAddresses []string) (*RoundRobinLoadBalancedHTTPService, error) {
	if len(inwayAddresses) == 0 {
		return nil, errNoInwaysAvailable
	}

	s := &RoundRobinLoadBalancedHTTPService{
		organizationName: organizationName,
		serviceName:      serviceName,
		roots:            roots,
		count:            0,
		inwayAddresses:   inwayAddresses,
		proxies:          make([]*httputil.ReverseProxy, len(inwayAddresses)),
	}
	s.logger = logger.With(zap.String("outway-service-full-name", s.FullName()))
	for i, inwayAddress := range inwayAddresses {
		endpointURL, err := url.Parse("https://" + inwayAddress)
		if err != nil {
			return nil, errors.Wrap(err, "invalid endpoint provided")
		}
		endpointURL.Path = "/" + serviceName
		proxy := httputil.NewSingleHostReverseProxy(endpointURL)
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
		proxy.Transport = transport
		s.proxies[i] = proxy
	}

	return s, nil
}

// FullName returns the name of the service
func (s *RoundRobinLoadBalancedHTTPService) FullName() string {
	return s.organizationName + "." + s.serviceName
}

// ProxyHTTPRequest procies the HTTP request to the proper endpoint.
func (s *RoundRobinLoadBalancedHTTPService) ProxyHTTPRequest(w http.ResponseWriter, r *http.Request) {
	s.getProxy().ServeHTTP(w, r)
}

// GetInwayAddresses returns the possible inwayaddresses of the httpservice
func (s *RoundRobinLoadBalancedHTTPService) GetInwayAddresses() []string {
	return s.inwayAddresses
}

func (s *RoundRobinLoadBalancedHTTPService) getProxy() *httputil.ReverseProxy {
	if len(s.proxies) == 0 {
		return nil
	}
	s.loadBalanceLock.Lock()
	proxy := s.proxies[s.count]
	s.count = (s.count + 1) % len(s.proxies)
	s.loadBalanceLock.Unlock()
	return proxy
}
