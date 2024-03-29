// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/httperrors"
	"go.nlx.io/nlx/common/transactionlog"
	inway_http "go.nlx.io/nlx/inway/http"
)

type LogRecordPlugin struct {
	organizationSerialNumber string
	txLogger                 transactionlog.TransactionLogger
}

func NewLogRecordPlugin(organizationSerialNumber string, txLogger transactionlog.TransactionLogger) *LogRecordPlugin {
	return &LogRecordPlugin{
		organizationSerialNumber: organizationSerialNumber,
		txLogger:                 txLogger,
	}
}

func (plugin *LogRecordPlugin) Serve(next ServeFunc) ServeFunc {
	return func(context *Context) error {
		logRecordID := context.Request.Header.Get("X-NLX-Logrecord-Id")
		if logRecordID == "" {
			context.Logger.Warn("Received request with missing logrecord id", zap.String("organization_serial_number", context.AuthInfo.OrganizationSerialNumber))

			inway_http.WriteError(context.Response, httperrors.O1, httperrors.MissingLogRecordID())

			return nil
		}

		err := plugin.createLogRecord(context, logRecordID)
		if err != nil {
			context.Logger.Error("failed to store transactionlog record", zap.Error(err))

			inway_http.WriteError(context.Response, httperrors.O1, httperrors.ServerError(nil))

			return nil
		}

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

	return recordData
}

func (plugin *LogRecordPlugin) createLogRecord(context *Context, logRecordID string) error {
	recordData := createRecordData(context.Request.Header, context.Destination.Path)
	organizationSerialNumber, ok := context.LogData["organizationSerialNumber"]

	if !ok {
		return fmt.Errorf("missing organization name in log data")
	}

	record := &transactionlog.Record{
		SrcOrganization:  organizationSerialNumber,
		DestOrganization: context.Destination.Organization,
		ServiceName:      context.Destination.Service.Name,
		TransactionID:    logRecordID,
		Data:             recordData,
	}

	if delegator, ok := context.LogData["delegator"]; ok {
		record.Delegator = delegator
	}

	if orderReference, ok := context.LogData["orderReference"]; ok {
		record.OrderReference = orderReference
	}

	if err := plugin.txLogger.AddRecord(record); err != nil {
		return errors.Wrap(err, "unable to add record to database")
	}

	return nil
}
