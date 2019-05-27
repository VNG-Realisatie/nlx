package certportal_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	certportal "go.nlx.io/nlx/ca-certportal"
	mock "go.nlx.io/nlx/ca-certportal/mock"
)

func TestRouteRequestCertificate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSigner := mock.NewMockSigner(ctrl)
	certPortal := certportal.NewCertPortal(zap.NewNop(), func() (signer.Signer, error) {
		return mockSigner, nil
	})
	assert.NotNil(t, certPortal)
	srv := httptest.NewServer(certPortal.GetRouter())
	defer srv.Close()

	jsonBytesCertificateRequest, err := json.Marshal(&certportal.CertificateRequest{
		Csr: "csr",
	})
	assert.NoError(t, err)

	tests := []struct {
		requestBody        []byte
		setupMock          func()
		expectedStatusCode int
		expectedBody       string
	}{
		{
			jsonBytesCertificateRequest,
			func() {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: "csr",
				}).Return([]byte("test_cert"), nil)
			},
			http.StatusCreated,
			`{"certificate":"test_cert"}` + "\n",
		},
		{
			[]byte("invalid"),
			func() {
			},
			http.StatusBadRequest,
			"",
		},
		{
			jsonBytesCertificateRequest,
			func() {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: "csr",
				}).Return(nil, fmt.Errorf("error signing request"))
			},
			http.StatusInternalServerError,
			"",
		},
	}

	for _, test := range tests {
		test.setupMock()
		resp, err := http.Post(fmt.Sprintf("%s/api/request_certificate", srv.URL), "application/json", bytes.NewReader(test.requestBody))
		assert.NoError(t, err)
		assert.Equal(t, test.expectedStatusCode, resp.StatusCode)
		responseBody, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedBody, string(responseBody))
	}
}

func TestRouteRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSigner := mock.NewMockSigner(ctrl)
	certPortal := certportal.NewCertPortal(zap.NewNop(), func() (signer.Signer, error) {
		return mockSigner, nil
	})
	assert.NotNil(t, certPortal)
	srv := httptest.NewServer(certPortal.GetRouter())
	defer srv.Close()

	tests := []struct {
		setupMock          func()
		expectedStatusCode int
		expectedBody       string
	}{
		{
			func() {
				mockSigner.EXPECT().Info(info.Req{}).Return(
					&info.Resp{
						Certificate: "testCert",
					}, nil)
			},
			http.StatusOK,
			"testCert",
		},
		{
			func() {
				mockSigner.EXPECT().Info(info.Req{}).Return(nil, fmt.Errorf("error getting info"))
			},
			http.StatusInternalServerError,
			"failed to obtain root.crt from cfssl root CA\n",
		},
	}

	for _, test := range tests {
		test.setupMock()
		resp, err := http.Get(fmt.Sprintf("%s/root.crt", srv.URL))
		assert.NoError(t, err)
		assert.Equal(t, test.expectedStatusCode, resp.StatusCode)
		responseBody, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedBody, string(responseBody))
	}
}

func TestRoutesInvalidSigner(t *testing.T) {
	certPortal := certportal.NewCertPortal(zap.NewNop(), func() (signer.Signer, error) {
		return nil, fmt.Errorf("unable to create certificate signer")
	})
	assert.NotNil(t, certPortal)
	srv := httptest.NewServer(certPortal.GetRouter())
	defer srv.Close()
	jsonBytesCertificateRequest, err := json.Marshal(&certportal.CertificateRequest{
		Csr: "csr",
	})
	assert.NoError(t, err)
	resp, err := http.Post(fmt.Sprintf("%s/api/request_certificate", srv.URL), "application/json", bytes.NewReader(jsonBytesCertificateRequest))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	resp, err = http.Get(fmt.Sprintf("%s/root.crt", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
