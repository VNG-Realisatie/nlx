// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.nlx.io/nlx/common/tlsconfig"
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
	// endPoints       []*url.URL

	logger *zap.Logger
	roots  *x509.CertPool

	proxies []*httputil.ReverseProxy
}

func newRoundTripHTTPTransport(tlsConfig *tls.Config) *http.Transport {
	tlsconfig.ApplyDefaults(tlsConfig)
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
		TLSClientConfig:       tlsConfig,
	}
	return transport
}

// NewRoundRobinLoadBalancedHTTPService creates a RoundRobinLoadBalancedHTTPService
func NewRoundRobinLoadBalancedHTTPService(
	logger *zap.Logger,
	roots *x509.CertPool,
	certFile, keyFile,
	organizationName,
	serviceName string,
	inwayAddresses []string,
) (*RoundRobinLoadBalancedHTTPService, error) {

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

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, errors.Wrap(err, "invalid certificate provided")
	}
	roundTripTransport := newRoundTripHTTPTransport(&tls.Config{
		RootCAs:      roots,
		Certificates: []tls.Certificate{cert},
	})

	for i, inwayAddress := range inwayAddresses {
		endpointURL, err := url.Parse("https://" + inwayAddress)
		if err != nil {
			return nil, errors.Wrap(err, "invalid endpoint provided")
		}
		endpointURL.Path = "/" + serviceName
		proxy := httputil.NewSingleHostReverseProxy(endpointURL)
		proxy.Transport = roundTripTransport
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
	// TODO: unnecessary check? seems impossible that len(s.proxies) could be 0
	if len(s.proxies) == 0 {
		return nil
	}
	s.loadBalanceLock.Lock()
	proxy := s.proxies[s.count]
	s.count = (s.count + 1) % len(s.proxies)
	s.loadBalanceLock.Unlock()
	return proxy
}
