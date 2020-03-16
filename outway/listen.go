// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/tlsconfig"
	"go.nlx.io/nlx/common/transactionlog"
)

// RunServer is a blocking function that listens on provided tcp address to handle requests.
func (o *Outway) RunServer(listenAddress, listenAddressTLS string, tlsOptions orgtls.TLSOptions) error {
	o.serverPlain = &http.Server{
		Addr:    listenAddress,
		Handler: o,
	}

	o.serverTLS = &http.Server{
		Addr:      listenAddressTLS,
		Handler:   o,
		TLSConfig: tlsconfig.Defaults(),
	}

	errorChannel := make(chan error)

	go func() {
		err := o.serverPlain.ListenAndServe()
		if err != http.ErrServerClosed {
			errorChannel <- errors.Wrap(err, "error listening on server")
		}
	}()

	go func() {
		err := o.serverTLS.ListenAndServeTLS(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
		if err != http.ErrServerClosed {
			errorChannel <- errors.Wrap(err, "error listening on TLS server")
		}
	}()

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

	o.shutDown()

	return err
}

func (o *Outway) shutDown() {
	o.logger.Debug("shutting down")

	o.monitorService.SetNotReady()

	wg := sync.WaitGroup{}

	numberOfServers := 2
	wg.Add(numberOfServers)

	go func() {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		err := o.serverTLS.Shutdown(localCtx)
		if err != nil {
			o.logger.Error("error shutting down server tls", zap.Error(err))
		}

		wg.Done()
	}()

	go func() {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		err := o.serverPlain.Shutdown(localCtx)
		if err != nil {
			o.logger.Error("error shutting down server tls", zap.Error(err))
		}

		wg.Done()
	}()

	wg.Wait()

	err := o.monitorService.Stop()
	if err != nil {
		o.logger.Error("error shutting down monitoring service", zap.Error(err))
	}

}

func createHTTPTransport(tlsConfig *tls.Config) *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       20 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
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

	destination, err := parseURLPath(r.URL.Path)

	if err != nil {
		msg := "no valid url path expecting: organization/service/apipathL"
		logger.Error(msg, zap.Error(err))
		o.helpUser(w, msg, nil, r.URL.Path)
		return
	}

	// Authorize request with plugged authorization service if authorization settings are set.
	if o.authorizationSettings != nil {
		authResponse, authErr := o.authorizeRequest(r.Header, destination)
		if authErr != nil {
			logger.Error("error authorizing request", zap.Error(authErr))
			http.Error(w, "nlx outway: error authorizing request", http.StatusInternalServerError)
			return
		}

		o.logger.Info("authorization result", zap.Bool("authorized", authResponse.Authorized), zap.String("reason", authResponse.Reason))
		if !authResponse.Authorized {
			http.Error(w, fmt.Sprintf("nlx outway: authorization failed. reason: %s", authResponse.Reason), http.StatusUnauthorized)
			return
		}
	}

	r.URL.Path = destination.Path

	recordData := createRecordData(r.Header, destination.Path)
	service := o.getService(destination.Organization, destination.Service)

	if service == nil {
		msg := "nlx outway: unknown service"
		logger.Warn("received request for unknown service")
		o.helpUser(w, msg, destination, r.URL.Path)
		return
	}

	l, err := NewLogRecordID()
	if err != nil {
		logger.Error("could not get new request ID", zap.Error(err))
		http.Error(w, "nlx outway: internal server error", http.StatusInternalServerError)
		return
	}

	logRecordID := l.String()
	r.Header.Set("X-NLX-Logrecord-Id", logRecordID)

	dataSubjects, err := parseDataSubjects(r)
	if err != nil {
		http.Error(w, "nlx outway: invalid data subject header", http.StatusBadRequest)
		o.logger.Warn("invalid data subject header", zap.Error(err))
		return
	}

	o.stripHeaders(r, destination.Organization)

	err = o.txlogger.AddRecord(&transactionlog.Record{
		SrcOrganization:  o.organizationName,
		DestOrganization: destination.Organization,
		ServiceName:      destination.Service,
		LogrecordID:      logRecordID,
		Data:             recordData,
		DataSubjects:     dataSubjects,
	})
	if err != nil {
		http.Error(w, "nlx outway: server error", http.StatusInternalServerError)
		o.logger.Error("failed to store transactionlog record", zap.Error(err))
		return
	}

	o.logger.Info("forwarding API request", zap.String("destination-organization", destination.Organization), zap.String("service", destination.Service), zap.String("logrecord-id", logRecordID))

	service.ProxyHTTPRequest(w, r)
}

func createRecordData(h http.Header, p string) map[string]interface{} {
	recordData := make(map[string]interface{})
	recordData["request-path"] = p
	if processID := h.Get("X-NLX-Request-Process-Id"); processID != "" {
		recordData["doelbinding-process-id"] = processID
	}
	if dataElements := h.Get("X-NLX-Request-Data-Elements"); dataElements != "" {
		recordData["doelbinding-data-elements"] = dataElements
	}

	if userData := h.Get("X-NLX-Requester-User"); userData != "" {
		recordData["doelbinding-user"] = userData
	}

	if claims := h.Get("X-NLX-Requester-Claims"); claims != "" {
		recordData["doelbinding-claims"] = claims
	}

	if userID := h.Get("X-NLX-Request-User-Id"); userID != "" {
		recordData["doelbinding-user-id"] = userID
	}

	if applicationID := h.Get("X-NLX-Request-Application-Id"); applicationID != "" {
		recordData["doelbinding-application-id"] = applicationID
	}

	if subjectIdentifier := h.Get("X-NLX-Request-Subject-Identifier"); subjectIdentifier != "" {
		recordData["doelbinding-subject-identifier"] = subjectIdentifier
	}

	return recordData
}

type destination struct {
	Organization string
	Service      string
	Path         string
}

func parseURLPath(urlPath string) (*destination, error) {
	pathParts := strings.SplitN(strings.TrimPrefix(urlPath, "/"), "/", 3)

	if len(pathParts) != 3 {
		return nil, fmt.Errorf("invalid path in url expecting: /organization/serivice/path")
	}

	return &destination{
		Organization: pathParts[0],
		Service:      pathParts[1],
		Path:         pathParts[2],
	}, nil
}

func parseDataSubjects(r *http.Request) (map[string]string, error) {
	if dataSubjectsHeader := r.Header.Get("X-NLX-Request-Data-Subject"); dataSubjectsHeader != "" {
		return transactionlog.ParseDataSubjectHeader(dataSubjectsHeader)
	}

	return map[string]string{}, nil
}
