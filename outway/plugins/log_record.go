// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package plugins

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.nlx.io/nlx/common/transactionlog"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	logRecordSize        = 16
	logRecordStringSizes = logRecordSize * 2
)

type LogRecordID [logRecordSize]byte

func (l *LogRecordID) String() string {
	b := make([]byte, logRecordStringSizes)
	hex.Encode(b, l[:])

	return string(b)
}

func NewLogRecordID() (*LogRecordID, error) {
	var l LogRecordID

	if _, err := io.ReadFull(rand.Reader, l[:]); err != nil {
		return nil, fmt.Errorf("failed to read random bytes: %v", err)
	}

	return &l, nil
}

type LogRecordPlugin struct {
	organizationName string
	txLogger         transactionlog.TransactionLogger
}

func NewLogRecordPlugin(organizationName string, txLogger transactionlog.TransactionLogger) *LogRecordPlugin {
	return &LogRecordPlugin{
		organizationName: organizationName,
		txLogger:         txLogger,
	}
}

func (plugin *LogRecordPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		logRecordID, err := plugin.createLogRecord(context)
		if err != nil {
			context.Logger.Error("failed to store transactionlog record", zap.Error(err))

			if strings.Contains(err.Error(), "invalid data subject header") {
				http.Error(context.Response, "nlx outway: invalid data subject header", http.StatusBadRequest)
			} else {
				http.Error(context.Response, "nlx outway: server error", http.StatusInternalServerError)
			}

			return nil
		}

		context.Request.Header.Set("X-NLX-Logrecord-Id", logRecordID.String())

		return next(context)
	}
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

func (plugin *LogRecordPlugin) createLogRecord(context *Context) (*LogRecordID, error) {
	logRecordID, err := NewLogRecordID()
	if err != nil {
		return nil, errors.Wrap(err, "could not get new request ID")
	}

	dataSubjects := map[string]string{}

	if dataSubjectsHeader := context.Request.Header.Get("X-NLX-Request-Data-Subject"); dataSubjectsHeader != "" {
		var err error

		dataSubjects, err = transactionlog.ParseDataSubjectHeader(dataSubjectsHeader)
		if err != nil {
			return nil, errors.Wrap(err, "invalid data subject header")
		}
	}

	recordData := createRecordData(context.Request.Header, context.Destination.Path)
	record := &transactionlog.Record{
		SrcOrganization:  plugin.organizationName,
		DestOrganization: context.Destination.Organization,
		ServiceName:      context.Destination.Service,
		LogrecordID:      logRecordID.String(),
		Data:             recordData,
		DataSubjects:     dataSubjects,
	}

	if delegator, ok := context.LogData["delegator"]; ok {
		record.Delegator = delegator
	}

	if orderReference, ok := context.LogData["orderReference"]; ok {
		record.OrderReference = orderReference
	}

	if err := plugin.txLogger.AddRecord(record); err != nil {
		return nil, errors.Wrap(err, "unable to add record to database")
	}

	return logRecordID, nil
}
