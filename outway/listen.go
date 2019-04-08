// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"encoding/binary"
	"hash/crc64"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.nlx.io/nlx/common/tlsconfig"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/transactionlog"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ListenAndServe is a blocking function that listens on provided tcp address to handle requests.
func (o *Outway) ListenAndServe(process *process.Process, address string) error {
	server := &http.Server{
		Addr:    address,
		Handler: o,
	}

	process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		return server.Shutdown(localCtx)
	})

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return errors.Wrap(err, "failed to run http server")
	}

	o.wg.Wait() // Wait until all async jobs will finish
	return nil
}

// ListenAndServeTLS is a blocking function that listens on provided tcp address to handle requests.
func (o *Outway) ListenAndServeTLS(process *process.Process, address string, certFile, keyFile string) error {
	server := &http.Server{
		Addr:      address,
		Handler:   o,
		TLSConfig: tlsconfig.Defaults(),
	}

	process.CloseGracefully(func() error {
		localCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		return server.Shutdown(localCtx)
	})

	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil && err != http.ErrServerClosed {
		return errors.Wrap(err, "failed to run http server")
	}

	o.wg.Wait() // Wait until all async jobs will finish
	return nil
}

// ServeHTTP handles requests from the organization to the outway, it selects the correct service backend and lets it handle the request further.
func (o *Outway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := o.logger.With(
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)

	urlparts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 3)
	if len(urlparts) != 3 {
		http.Error(w, "nlx outway: invalid path in url", http.StatusBadRequest)
		logger.Warn("received request with invalid path")
		return
	}
	destOrganizationName := urlparts[0]
	destServiceName := urlparts[1]
	requestPath := "/" + urlparts[2] // retain original path
	r.URL.Path = requestPath

	o.servicesLock.RLock()
	service := o.services[destOrganizationName+"."+destServiceName]
	o.servicesLock.RUnlock()
	if service == nil {
		http.Error(w, "nlx outway: unknown service", http.StatusBadRequest)
		logger.Warn("received request for unknown service")
		return
	}

	var recordData = make(map[string]interface{})
	recordData["request-path"] = requestPath

	logrecordIDFlake, err := o.requestFlake.NextID()
	if err != nil {
		logger.Error("could not get new request ID", zap.Error(err))
		http.Error(w, "nlx outway: internal server error", http.StatusInternalServerError)
		return
	}
	logrecordIDFlakeBytes := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(logrecordIDFlakeBytes, logrecordIDFlake)
	logrecordIDNum := crc64.Checksum(logrecordIDFlakeBytes, o.ecmaTable)
	logrecordID := strconv.FormatUint(logrecordIDNum, 32)
	r.Header.Set("X-NLX-Logrecord-Id", logrecordID)

	if processID := r.Header.Get("X-NLX-Request-Process-Id"); processID != "" {
		recordData["doelbinding-process-id"] = processID
	}
	if dataElements := r.Header.Get("X-NLX-Request-Data-Elements"); dataElements != "" {
		recordData["doelbinding-data-elements"] = dataElements
	}

	if userID := r.Header.Get("X-NLX-Request-User-Id"); userID != "" {
		recordData["doelbinding-user-id"] = userID
		r.Header.Del("X-NLX-Request-User-Id")
	}
	if applicationID := r.Header.Get("X-NLX-Request-Application-Id"); applicationID != "" {
		recordData["doelbinding-application-id"] = applicationID
		r.Header.Del("X-NLX-Request-Application-Id")
	}
	if subjectIdentifier := r.Header.Get("X-NLX-Request-Subject-Identifier"); subjectIdentifier != "" {
		recordData["doelbinding-subject-identifier"] = subjectIdentifier
		r.Header.Del("X-NLX-Request-Subject-Identifier")
	}
	var dataSubjects map[string]string
	if dataSubjectsHeader := r.Header.Get("X-NLX-Request-Data-Subject"); dataSubjectsHeader != "" {
		dataSubjects, err = transactionlog.ParseDataSubjectHeader(dataSubjectsHeader)
		if err != nil {
			http.Error(w, "nlx outway: invalid data subject header", http.StatusBadRequest)
			o.logger.Warn("invalid data subject header", zap.Error(err))
			return
		}
		r.Header.Del("X-NLX-Request-Data-Subject")
	}

	err = o.txlogger.AddRecord(&transactionlog.Record{
		SrcOrganization:  o.organizationName,
		DestOrganization: destOrganizationName,
		ServiceName:      destServiceName,
		LogrecordID:      logrecordID,
		Data:             recordData,
		DataSubjects:     dataSubjects,
	})
	if err != nil {
		http.Error(w, "nlx outway: server error", http.StatusInternalServerError)
		o.logger.Error("failed to store transactionlog record", zap.Error(err))
		return
	}

	o.logger.Info("forwarding API request", zap.String("destination-organization", destOrganizationName), zap.String("service", destServiceName), zap.String("logrecord-id", logrecordID))

	service.ProxyHTTPRequest(w, r)
}
