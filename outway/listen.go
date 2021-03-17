// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/outway/plugins"
)

const (
	timeOut              = 30 * time.Second
	keepAlive            = 30 * time.Second
	maxIdelCons          = 100
	IdleConnTimeout      = 20 * time.Second
	TLSHandshakeTimeout  = 10 * time.Second
	ExpectConinueTimeout = 1 * time.Second
)

// RunServer is a blocking function that listens on provided tcp address to handle requests.
func (o *Outway) RunServer(listenAddress string, serverCertificate *tls.Certificate) error {
	o.httpServer = &http.Server{
		Addr:    listenAddress,
		Handler: o,
	}

	errorChannel := make(chan error)

	if serverCertificate == nil {
		go func() {
			o.logger.Info(fmt.Sprintf("starting HTTP server on %s", listenAddress))
			errorChannel <- o.httpServer.ListenAndServe()
		}()
	} else {
		tlsConfig := common_tls.NewConfig(common_tls.WithTLS12())
		tlsConfig.Certificates = []tls.Certificate{*serverCertificate}

		o.httpServer.TLSConfig = tlsConfig

		go func() {
			o.logger.Info(fmt.Sprintf("starting HTTPS server on %s", listenAddress))
			errorChannel <- o.httpServer.ListenAndServeTLS("", "")
		}()
	}

	go func() {
		err := o.monitorService.Start()
		if err != nil {
			errorChannel <- errors.Wrap(err, "error listening on monitoring service")
		}
	}()

	o.process.CloseGracefully(func() error {
		o.shutDown()
		return nil
	})

	err := <-errorChannel

	if err == http.ErrServerClosed {
		return nil
	}

	o.shutDown()

	return errors.Wrap(err, "error listening on server")
}

func (o *Outway) shutDown() {
	o.logger.Debug("shutting down")

	o.monitorService.SetNotReady()

	localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := o.httpServer.Shutdown(localCtx)
	if err != nil {
		o.logger.Error("error shutting down server", zap.Error(err))
	}

	err = o.monitorService.Stop()
	if err != nil {
		o.logger.Error("error shutting down monitoring service", zap.Error(err))
	}
}

func createHTTPTransport(tlsConfig *tls.Config) *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeOut,
			KeepAlive: keepAlive,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          maxIdelCons,
		IdleConnTimeout:       IdleConnTimeout,
		TLSHandshakeTimeout:   TLSHandshakeTimeout,
		ExpectContinueTimeout: ExpectConinueTimeout,
		TLSClientConfig:       tlsConfig,
	}
}

// ServeHTTP handles requests from the organization to the outway,
// it selects the correct service backend and lets it handle the request further.
func (o *Outway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := o.logger.With(
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)
	logger.Info("received request")

	o.requestHTTPHandler(logger, w, r)
}

func (o *Outway) handleHTTPRequest(logger *zap.Logger, w http.ResponseWriter, r *http.Request) {
	destination, err := parseURLPath(r.URL.Path)
	if err != nil {
		if isNLXUrl(r.URL) {
			http.Error(w, fmt.Sprintf("please enable proxy mode by setting the 'use-as-http-proxy' flag to resolve: %s", r.URL.String()), http.StatusInternalServerError)
			return
		}

		msg := "no valid url path expecting: organization/service/apipath"
		logger.Error(msg, zap.Error(err))

		o.helpUser(w, msg, nil, r.URL.Path)

		return
	}

	o.handleOnNLX(logger, destination, w, r)
}

func (o *Outway) handleHTTPRequestAsProxy(logger *zap.Logger, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		logger.Error("CONNECT method not supported")
		http.Error(w, "CONNECT method is not supported", http.StatusNotImplemented)

		return
	}

	if !isNLXUrl(r.URL) {
		o.forwardingHTTPProxy.ServeHTTP(w, r)
		return
	}

	destination, err := parseLocalNLXURL(r.URL)
	if err != nil {
		logger.Error("error parsing desination", zap.Error(err))
		http.Error(w, "nlx outway: no valid url expecting: service.organization.service.nlx.local/apipath", http.StatusBadRequest)

		return
	}

	o.handleOnNLX(logger, destination, w, r)
}

func buildChain(serve plugins.ServeFunc, pluginList ...plugins.Plugin) plugins.ServeFunc {
	if len(pluginList) == 0 {
		return serve
	}

	return pluginList[0].Serve(buildChain(serve, pluginList[1:]...))
}

func (o *Outway) handleOnNLX(logger *zap.Logger, destination *plugins.Destination, w http.ResponseWriter, r *http.Request) {
	service := o.getService(destination.Organization, destination.Service)
	if service == nil {
		logger.Warn("received request for unknown service")

		o.helpUser(w, "nlx outway: unknown service", destination, r.URL.Path)

		return
	}

	chain := buildChain(func(context plugins.Context) error {
		context.Request.URL.Path = destination.Path

		service.ProxyHTTPRequest(context.Response, context.Request)
		return nil
	}, o.plugins...)

	ctx := plugins.Context{
		Response:    w,
		Request:     r,
		Logger:      o.logger,
		Destination: destination,
		LogData:     map[string]string{},
	}

	logger.Info(
		"forwarding API request",
		zap.String("service", destination.Service),
		zap.String("destination-organization", destination.Organization),
	)

	if err := chain(ctx); err != nil {
		logger.Error("error while handling API request", zap.Error(err))
	}
}
