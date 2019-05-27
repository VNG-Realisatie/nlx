package certportal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock "go.nlx.io/nlx/ca-certportal/mock"
)

func TestRequestCertificateHandler(t *testing.T) {
	jsonBytesCertificateRequest, err := json.Marshal(&certificateRequest{
		Csr: "csr",
	})
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSigner := mock.NewMockSigner(ctrl)

	tests := []struct {
		requestBody        []byte
		handlerFunc        http.HandlerFunc
		expectedStatusCode int
		expectedBody       string
	}{
		{
			jsonBytesCertificateRequest,
			http.HandlerFunc(requestCertificateHandler(zap.NewNop(), func() (signer.Signer, error) {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: "csr",
				}).Return([]byte("test_cert"), nil)
				return mockSigner, nil
			})),
			http.StatusCreated,
			`{"certificate":"test_cert"}` + "\n",
		},
		{
			[]byte("invalid"),
			http.HandlerFunc(requestCertificateHandler(zap.NewNop(), func() (signer.Signer, error) {
				return mockSigner, nil
			})),
			http.StatusBadRequest,
			"",
		},
		{
			jsonBytesCertificateRequest,
			http.HandlerFunc(requestCertificateHandler(zap.NewNop(), func() (signer.Signer, error) {
				return nil, fmt.Errorf("error creating signer")
			})),
			http.StatusInternalServerError,
			"",
		},
		{
			jsonBytesCertificateRequest,
			http.HandlerFunc(requestCertificateHandler(zap.NewNop(), func() (signer.Signer, error) {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: "csr",
				}).Return(nil, fmt.Errorf("error signing request"))
				return mockSigner, nil
			})),
			http.StatusInternalServerError,
			"",
		},
	}

	for _, test := range tests {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/request_certificate", bytes.NewReader(test.requestBody))
		req.Header.Add("Content-Type", "application/json")
		test.handlerFunc.ServeHTTP(rr, req)
		assert.Equal(t, test.expectedStatusCode, rr.Code)
		assert.Equal(t, test.expectedBody, rr.Body.String())
	}
}

func TestRootCertHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSigner := mock.NewMockSigner(ctrl)

	tests := []struct {
		handlerFunc        http.HandlerFunc
		expectedStatusCode int
		expectedBody       string
	}{
		{
			http.HandlerFunc(rootCertHandler(zap.NewNop(), func() (signer.Signer, error) {
				mockSigner.EXPECT().Info(info.Req{}).Return(
					&info.Resp{
						Certificate: "testCert",
					}, nil)
				return mockSigner, nil
			})),
			http.StatusOK,
			"testCert",
		}, {
			http.HandlerFunc(rootCertHandler(zap.NewNop(), func() (signer.Signer, error) {
				mockSigner.EXPECT().Info(info.Req{}).Return(nil, fmt.Errorf("error getting info"))
				return mockSigner, nil
			})),
			http.StatusInternalServerError,
			"failed to obtain root.crt from cfssl root CA\n",
		},
		{
			http.HandlerFunc(rootCertHandler(zap.NewNop(), func() (signer.Signer, error) {
				return nil, fmt.Errorf("error creating signer")
			})),
			http.StatusInternalServerError,
			"failed to create new cfssl signer\n",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/root.crt", nil)
		rr := httptest.NewRecorder()
		test.handlerFunc.ServeHTTP(rr, req)
		assert.Equal(t, test.expectedStatusCode, rr.Code)
		assert.Equal(t, test.expectedBody, rr.Body.String())
	}
}
