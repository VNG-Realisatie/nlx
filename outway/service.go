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
	"golang.org/x/net/http2"
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
	healthyStates   []bool
	loadBalanceLock sync.Mutex
	count           int

	logger *zap.Logger
	roots  *x509.CertPool

	proxies []*httputil.ReverseProxy
}

func newRoundTripHTTPTransport(logger *zap.Logger, tlsConfig *tls.Config) *http.Transport {
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
	if err := http2.ConfigureTransport(transport); err != nil {
		logger.Error("failed to add http2 to transport")
	}

	return transport
}

// NewRoundRobinLoadBalancedHTTPService creates a RoundRobinLoadBalancedHTTPService
func NewRoundRobinLoadBalancedHTTPService(
	logger *zap.Logger,
	roots *x509.CertPool,
	certFile,
	keyFile,
	organizationName,
	serviceName string,
	inwayAddresses []string,
	healthyStates []bool,
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
		healthyStates:    healthyStates,
		proxies:          make([]*httputil.ReverseProxy, len(inwayAddresses)),
	}
	s.logger = logger.With(zap.String("outway-service-full-name", s.FullName()))

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, errors.Wrap(err, "invalid certificate provided")
	}
	roundTripTransport := newRoundTripHTTPTransport(
		logger,
		&tls.Config{
			RootCAs:      roots,
			Certificates: []tls.Certificate{cert},
		})

	for i, inwayAddress := range inwayAddresses {
		endpointURL, err := url.Parse("https://" + inwayAddress)
		if err != nil {
			return nil, errors.Wrap(err, "inway address:"+inwayAddress+" is not a valid url")
		}
		endpointURL.Path = "/" + serviceName
		proxy := httputil.NewSingleHostReverseProxy(endpointURL)
		proxy.Transport = roundTripTransport
		proxy.ErrorHandler = s.LogServiceErrors
		s.proxies[i] = proxy
	}

	return s, nil
}

// FullName returns the name of the service
func (s *RoundRobinLoadBalancedHTTPService) FullName() string {
	return s.organizationName + "." + s.serviceName
}

// ProxyHTTPRequest process the HTTP request to the proper endpoint.
func (s *RoundRobinLoadBalancedHTTPService) ProxyHTTPRequest(w http.ResponseWriter, r *http.Request) {
	s.getProxy().ServeHTTP(w, r)
}

// Used for testing purposes to change transport
func (s *RoundRobinLoadBalancedHTTPService) GetProxies() []*httputil.ReverseProxy {
	return s.proxies
}

// LogServiceErrors request failed but service was announced to directory
// log the error and return some helpful text.
// set 503 Status Service Temporarily Unavailable response.
func (s *RoundRobinLoadBalancedHTTPService) LogServiceErrors(w http.ResponseWriter, r *http.Request, e error) {
	msg := "failed request to " + r.URL.String() + " try again later / check firewall?"
	s.logger.Error(msg)
	http.Error(w, msg, http.StatusServiceUnavailable)
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
