// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/net/http2"

	"go.nlx.io/nlx/common/httperrors"
	common_tls "go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	outway_http "go.nlx.io/nlx/outway/http"
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
	organizationSerialNumber string
	serviceName              string

	inways          []directoryapi.Inway
	loadBalanceLock sync.Mutex
	count           int

	logger *zap.Logger

	proxies []*httputil.ReverseProxy
}

func newRoundTripHTTPTransport(logger *zap.Logger, tlsConfig *tls.Config) *http.Transport {
	const (
		Timeout               = 6 * time.Second
		KeepAlive             = 6 * time.Second
		MaxIdleConns          = 100
		IdleConnTimeout       = 10 * time.Second
		TLSHandshakeTimeout   = 10 * time.Second
		ExpectContinueTimeout = 1 * time.Second
	)

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   Timeout,
			KeepAlive: KeepAlive,
		}).DialContext,
		MaxIdleConns:          MaxIdleConns,
		IdleConnTimeout:       IdleConnTimeout,
		TLSHandshakeTimeout:   TLSHandshakeTimeout,
		ExpectContinueTimeout: ExpectContinueTimeout,
		TLSClientConfig:       tlsConfig,
	}
	if err := http2.ConfigureTransport(transport); err != nil {
		logger.Error("failed to add http2 to transport")
	}

	return transport
}

func NewRoundRobinLoadBalancedHTTPService(
	logger *zap.Logger,
	cert *common_tls.CertificateBundle,
	organizationSerialNumber,
	serviceName string,
	inways []directoryapi.Inway,
) (*RoundRobinLoadBalancedHTTPService, error) {
	if len(inways) == 0 {
		return nil, errNoInwaysAvailable
	}

	s := &RoundRobinLoadBalancedHTTPService{
		organizationSerialNumber: organizationSerialNumber,
		serviceName:              serviceName,
		count:                    0,
		inways:                   inways,
		proxies:                  make([]*httputil.ReverseProxy, len(inways)),
	}
	s.logger = logger.With(zap.String("outway-service-full-name", s.FullName()))

	tlsConfig := cert.TLSConfig()
	roundTripTransport := newRoundTripHTTPTransport(logger, tlsConfig)

	// index is used instead of `_, inway` to avoid the following error:
	// govet: copylocks: range var inway copies lock: go.nlx.io/nlx/directory-api/directoryapi.Inway contains google.golang.org/protobuf/internal/impl.MessageState contains sync.Mutex
	for i := range inways {
		endpointURL, err := url.Parse("https://" + inways[i].Address)
		if err != nil {
			return nil, errors.Wrap(err, "inway address:"+inways[i].Address+" is not a valid url")
		}

		proxy := httputil.NewSingleHostReverseProxy(endpointURL)
		proxy.Transport = roundTripTransport
		proxy.ErrorHandler = s.LogServiceErrors

		s.proxies[i] = proxy
	}

	return s, nil
}

// FullName returns the name of the service
func (s *RoundRobinLoadBalancedHTTPService) FullName() string {
	return s.organizationSerialNumber + "." + s.serviceName
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
	msg := fmt.Sprintf("failed request to '%s', try again later and check your firewall, check O1 and M1 at https://docs.nlx.io/support/common-errors/", r.URL.String())
	s.logger.Error(msg, zap.Error(e))
	outway_http.WriteError(w, httperrors.O1, httperrors.ServiceUnreachable, msg)
}

// GetInwayAddresses returns the possible inwayaddresses of the httpservice
func (s *RoundRobinLoadBalancedHTTPService) GetInways() []directoryapi.Inway {
	return s.inways
}

func (s *RoundRobinLoadBalancedHTTPService) GetInwayAddresses() []string {
	addresses := []string{}

	// index is used instead of `_, inway` to avoid the following error:
	// govet: copylocks: range var inway copies lock: go.nlx.io/nlx/directory-api/directoryapi.Inway contains google.golang.org/protobuf/internal/impl.MessageState contains sync.Mutex
	for i := range s.inways {
		addresses = append(addresses, s.inways[i].Address)
	}

	return addresses
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
