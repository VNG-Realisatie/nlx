// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package plugins_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/common/transactionlog"
	mock_transactionlog "go.nlx.io/nlx/common/transactionlog/mock"
	"go.nlx.io/nlx/inway/plugins"
)

//nolint:funlen // this is a test
func TestLoginPlugin(t *testing.T) {
	tests := map[string]struct {
		tranactionLogger   func(ctrl *gomock.Controller) *mock_transactionlog.MockTransactionLogger
		headers            map[string]string
		expectedStatusCode int
		expectedMessage    string
	}{
		"missing_log_record_id": {
			headers: map[string]string{
				"X-NLX-Logrecord-Id": "",
			},
			tranactionLogger: func(ctrl *gomock.Controller) *mock_transactionlog.MockTransactionLogger {
				db := mock_transactionlog.NewMockTransactionLogger(ctrl)

				return db
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "nlx-inway: missing logrecord id\n",
		},
		"transaction_log_db_fails": {
			headers: map[string]string{
				"X-NLX-Logrecord-Id": "mock-log-record-id",
			},
			tranactionLogger: func(ctrl *gomock.Controller) *mock_transactionlog.MockTransactionLogger {
				db := mock_transactionlog.NewMockTransactionLogger(ctrl)
				db.EXPECT().AddRecord(&transactionlog.Record{
					LogrecordID:      "mock-log-record-id",
					SrcOrganization:  "mock-source-organization",
					DestOrganization: "mock-destination-organization",
					ServiceName:      "mock-service",
					Data: map[string]interface{}{
						"request-path": "/path",
					},
				}).Return(fmt.Errorf("arbitrary error"))
				return db
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedMessage:    "nlx inway: server error\n",
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
					LogrecordID:      "mock-log-record-id",
					SrcOrganization:  "mock-source-organization",
					DestOrganization: "mock-destination-organization",
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
			expectedStatusCode: http.StatusOK,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			logRecordPlugin := plugins.NewLogRecordPlugin("mock-org", tt.tranactionLogger(ctrl))

			context := fakeContext(&plugins.Destination{
				Organization: "mock-destination-organization",
				Path:         "/path",
				Service: &plugins.Service{
					Name: "mock-service"},
			}, nil, &plugins.AuthInfo{
				OrganizationName: "mock-org",
			})

			context.LogData["organizationName"] = "mock-source-organization"

			for header, value := range tt.headers {
				context.Request.Header.Add(header, value)
			}

			err := logRecordPlugin.Serve(nopServeFunc)(context)
			assert.NoError(t, err)

			response := context.Response.(*httptest.ResponseRecorder).Result()
			defer response.Body.Close()

			contents, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedMessage, string(contents))
			assert.Equal(t, tt.expectedStatusCode, response.StatusCode)
		})
	}
}
