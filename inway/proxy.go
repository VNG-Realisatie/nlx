// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"go.uber.org/zap"

	"go.nlx.io/nlx/inway/plugins"
)

func (i *Inway) handleProxyRequest(w http.ResponseWriter, r *http.Request) {
	logger := i.logger.With(
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)

	urlparts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	if len(urlparts) != 2 {
		http.Error(w, "nlx-inway: invalid path in url", http.StatusBadRequest)
		logger.Warn("received request with invalid path")

		return
	}

	serviceName := urlparts[0]

	i.servicesLock.RLock()
	service, exists := i.services[serviceName]
	i.servicesLock.RUnlock()

	if !exists {
		http.Error(w, "nlx-inway: no endpoint for service", http.StatusBadRequest)
		logger.Warn("received request for service with no known endpoint")

		return
	}

	destination := plugins.Destination{
		Organization: i.organizationName,
		Path:         "/" + urlparts[1],
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
		r.URL.Path = context.Destination.Path

		proxy := httputil.NewSingleHostReverseProxy(endpointURL)
		proxy.Transport = newRoundTripHTTPTransport()
		proxy.ErrorHandler = i.LogAPIErrors
		proxy.ServeHTTP(w, r)

		return nil
	}, i.plugins...)

	if err := chain(context); err != nil {
		logger.Error("error executing plugin chain", zap.Error(err))
		http.Error(w, "inway: error executing plugin chain", http.StatusInternalServerError)
	}
}

func (i *Inway) LogAPIErrors(w http.ResponseWriter, r *http.Request, e error) {
	msg := ("nlx-inway: failed internal API request to " +
		r.URL.String() +
		" try again later / service api down/unreachable." +
		" check A1 error at https://docs.nlx.io/support/common-errors/")
	i.logger.Error(msg)
	http.Error(w, msg, http.StatusServiceUnavailable)
}
