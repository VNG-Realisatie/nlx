// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/httperrors"
	"go.nlx.io/nlx/common/transactionlog"
	mock_transactionlog "go.nlx.io/nlx/common/transactionlog/mock"
	"go.nlx.io/nlx/inway/plugins"
)

//nolint:funlen // this is a test
func TestLoginPlugin(t *testing.T) {
	tests := map[string]struct {
		tranactionLogger   func(ctrl *gomock.Controller) *mock_transactionlog.MockTransactionLogger
		headers            map[string]string
		wantHTTPStatusCode int
		wantErr            *httperrors.NLXNetworkError
	}{
		"missing_log_record_id": {
			headers: map[string]string{
				"X-NLX-Logrecord-Id": "",
			},
			tranactionLogger: func(ctrl *gomock.Controller) *mock_transactionlog.MockTransactionLogger {
				db := mock_transactionlog.NewMockTransactionLogger(ctrl)

				return db
			},
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.MissingLogRecordID,
				Message:  "missing logrecord id",
			},
		},
		"transaction_log_db_fails": {
			headers: map[string]string{
				"X-NLX-Logrecord-Id": "mock-log-record-id",
			},
			tranactionLogger: func(ctrl *gomock.Controller) *mock_transactionlog.MockTransactionLogger {
				db := mock_transactionlog.NewMockTransactionLogger(ctrl)
				db.EXPECT().AddRecord(&transactionlog.Record{
					TransactionID:    "mock-log-record-id",
					SrcOrganization:  "99999999999999999999",
					DestOrganization: "11111111111111111111",
					ServiceName:      "mock-service",
					Data: map[string]interface{}{
						"request-path": "/path",
					},
				}).Return(fmt.Errorf("arbitrary error"))
				return db
			},
			wantHTTPStatusCode: httperrors.StatusNLXNetworkError,
			wantErr: &httperrors.NLXNetworkError{
				Source:   httperrors.Inway,
				Location: httperrors.O1,
				Code:     httperrors.ServerError,
				Message:  "server error",
			},
		},
		"happy_flow": {
			headers: map[string]string{
				"X-NLX-Logrecord-Id":          "mock-log-record-id",
				"X-NLX-Request-Process-Id":    "process-id",
				"X-NLX-Request-Data-Elements": "data-elements",
				"X-NLX-Requester-User":        "requester-user",
				"X-NLX-Requester-Claims":      "claims",
			},
			tranactionLogger: func(ctrl *gomock.Controller) *mock_transactionlog.MockTransactionLogger {
				db := mock_transactionlog.NewMockTransactionLogger(ctrl)
				db.EXPECT().AddRecord(&transactionlog.Record{
					TransactionID:    "mock-log-record-id",
					SrcOrganization:  "99999999999999999999",
					DestOrganization: "11111111111111111111",
					ServiceName:      "mock-service",
					Data: map[string]interface{}{
						"request-path":              "/path",
						"doelbinding-process-id":    "process-id",
						"doelbinding-data-elements": "data-elements",
						"doelbinding-user":          "requester-user",
						"doelbinding-claims":        "claims",
					},
				}).Return(nil)
				return db
			},
			wantHTTPStatusCode: http.StatusOK,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			logRecordPlugin := plugins.NewLogRecordPlugin("99999999999999999999", tt.tranactionLogger(ctrl))

			context := fakeContext(&plugins.Destination{
				Organization: "11111111111111111111",
				Path:         "/path",
				Service: &plugins.Service{
					Name: "mock-service"},
			}, nil, &plugins.AuthInfo{
				OrganizationSerialNumber: "99999999999999999999",
			})

			context.LogData["organizationName"] = "mock-source-organization"
			context.LogData["organizationSerialNumber"] = "99999999999999999999"

			for header, value := range tt.headers {
				context.Request.Header.Add(header, value)
			}

			err := logRecordPlugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()
			defer response.Body.Close()

			contents, err := io.ReadAll(response.Body)
			assert.NoError(t, err)

			if tt.wantErr != nil {
				gotError := &httperrors.NLXNetworkError{}
				err := json.Unmarshal(contents, gotError)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantErr, gotError)
			}

			assert.Equal(t, tt.wantHTTPStatusCode, response.StatusCode)
		})
	}
}
