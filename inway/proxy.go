// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"go.uber.org/zap"

	inway_http "go.nlx.io/nlx/inway/http"
	"go.nlx.io/nlx/inway/plugins"
)

func (i *Inway) handleProxyRequest(w http.ResponseWriter, r *http.Request) {
	logger := i.logger.With(
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)

	if r.URL.Path == "" {
		logger.Warn("received request with invalid path")
		inway_http.WriteError(w, "path cannot be empty, must at least contain the service name.")

		return
	}

	cleanedPath := strings.TrimPrefix(r.URL.Path, "/")
	indexSlash := strings.Index(cleanedPath, "/")

	var serviceName = cleanedPath

	var path = ""

	if indexSlash > -1 {
		serviceName = cleanedPath[0:indexSlash]
		path = cleanedPath[indexSlash:]
	}

	i.servicesLock.RLock()
	service, exists := i.services[serviceName]
	i.servicesLock.RUnlock()

	if !exists {
		logger.Warn("received request for service with no known endpoint")
		inway_http.WriteError(w, fmt.Sprintf("no endpoint for service '%s'", serviceName))

		return
	}

	destination := plugins.Destination{
		Organization: i.organization.SerialNumber,
		Path:         path,
		Service:      service,
	}

	context := &plugins.Context{
		Logger:      logger,
		Destination: &destination,
		Response:    w,
		Request:     r,
		LogData:     map[string]string{},
		AuthInfo:    &plugins.AuthInfo{},
	}

	chain := plugins.BuildChain(func(context *plugins.Context) error {
		endpointURL, err := url.Parse(context.Destination.Service.EndpointURL)
		if err != nil {
			return err
		}

		r.Host = endpointURL.Host
		proxy := &httputil.ReverseProxy{
			// Use custom director because the default Director appends a trailing slash when the request path is empty. The trailing slash made requests to a SOAP API fail.
			// See: https://github.com/golang/go/issues/50337
			Director: func(req *http.Request) {
				req.URL.Scheme = endpointURL.Scheme
				req.URL.Host = endpointURL.Host
				req.URL.Path = endpointURL.Path + destination.Path

				if endpointURL.RawQuery == "" || req.URL.RawQuery == "" {
					req.URL.RawQuery = endpointURL.RawQuery + req.URL.RawQuery
				} else {
					req.URL.RawQuery = endpointURL.RawQuery + "&" + req.URL.RawQuery
				}

				if _, ok := req.Header["User-Agent"]; !ok {
					// explicitly disable User-Agent so it's not set to default value
					req.Header.Set("User-Agent", "")
				}
			},
		}

		proxy.Transport = newRoundTripHTTPTransport()
		proxy.ErrorHandler = i.LogAPIErrors
		proxy.ServeHTTP(w, r)

		return nil
	}, i.plugins...)

	if err := chain(context); err != nil {
		logger.Error("error executing plugin chain", zap.Error(err))

		inway_http.WriteError(w, "error executing plugin chain")
	}
}

func (i *Inway) LogAPIErrors(w http.ResponseWriter, r *http.Request, e error) {
	msg := fmt.Sprintf("failed internal API request to %s try again later. service api down/unreachable. check A1 error at https://docs.nlx.io/support/common-errors/", r.URL.String())
	i.logger.Error(msg)

	inway_http.WriteError(w, msg)
}
