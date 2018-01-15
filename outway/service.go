// Copyright 2018 VNG Realisatie. All rights reserved.
// Use of this source code is governed by the EUPL
// license that can be found in the LICENSE.md file.

package outway

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Service handles the proxying of a request to the inway
// TODO: if we like this model in the PoC, Service should become an interface with multiple implementations (http/json, grpc, ??) like inway
type Service struct {
	organizationName string
	serviceName      string

	logger *zap.Logger

	proxy *httputil.ReverseProxy
}

// NewService creates a new Service instance with a single inway to forward requests to.
// This is a PoC shortcut, the real outway will fetch services and their inways* from the directory
// 		(* note the plural; a service can have multiple inways published by directory so that an outway can balance load to multiple inways)
func NewService(logger *zap.Logger, organizationName, serviceName string, inwayAddress string) (*Service, error) {
	s := &Service{
		organizationName: organizationName,
		serviceName:      serviceName,
	}
	s.logger = logger.With(zap.String("outway-service-full-name", s.fullName()))
	endpointURL, err := url.Parse(inwayAddress)
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint provided")
	}
	endpointURL.Path = "/" + serviceName
	s.proxy = httputil.NewSingleHostReverseProxy(endpointURL)
	return s, nil
}

func (s *Service) fullName() string {
	return s.organizationName + "." + s.serviceName
}

func (s *Service) proxyRequest(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}
