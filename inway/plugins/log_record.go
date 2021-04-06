package plugins

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"go.nlx.io/nlx/common/transactionlog"
	"go.uber.org/zap"
)

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
		logRecordID := context.Request.Header.Get("X-NLX-Logrecord-Id")
		if logRecordID == "" {
			http.Error(context.Response, "nlx-inway: missing logrecord id", http.StatusBadRequest)
			context.Logger.Warn("Received request with missing logrecord id from " + context.AuthInfo.OrganizationName)

			return nil
		}

		err := plugin.createLogRecord(context, logRecordID)
		if err != nil {
			context.Logger.Error("failed to store transactionlog record", zap.Error(err))

			http.Error(context.Response, "nlx inway: server error", http.StatusInternalServerError)

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
	organizationName, ok := context.logData["organizationName"]
	if !ok {
		return fmt.Errorf("missing organization name in log data")
	}

	record := &transactionlog.Record{
		SrcOrganization:  organizationName,
		DestOrganization: context.Destination.Organization,
		ServiceName:      context.Destination.Service,
		LogrecordID:      logRecordID,
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
