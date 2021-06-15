// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package certportal_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/cloudflare/cfssl/info"
	"github.com/cloudflare/cfssl/signer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	certportal "go.nlx.io/nlx/ca-certportal"
	mock "go.nlx.io/nlx/ca-certportal/mock"
)

var pkiDir = filepath.Join("..", "testing", "pki")

const csrWithoutSAN = `
-----BEGIN CERTIFICATE REQUEST-----
MIIBXTCBxwIBADAeMRwwGgYDVQQDDBNob3N0bmFtZS50ZXN0LmxvY2FsMIGfMA0G
CSqGSIb3DQEBAQUAA4GNADCBiQKBgQCh0Fi/xEALsOBvWTpCtMtmS5UP2pqBFPx8
O0DWaIRNyCi3JyerL9qhjxvrIWJyD3/Aam3fbe17Y6/1hnBBpkJ0WzFdWvdYsXCA
I+vT8GUk8iYL09xwnzxL2Bx1rGG9URSWLBtYuD2lT4sntBACwyag6QQVMT7lbvB/
MbW/pGdziwIDAQABoAAwDQYJKoZIhvcNAQELBQADgYEAVMYCP6vJQbLSSce7LX6A
7YO98Hrvzc7/wZuWmG3EYyM7Sw3dEb8pLxKGiTiZl2rBZZs/rDOB5xz8iGNwHIfl
rPmL0grTgE4AW8cEJqzRNeDs52RR6MnYTdCfUMkNNc54OWsCH8ZgT8PpWpc6dyqH
2B9XFNelZbfv3GHt27eIKYI=
-----END CERTIFICATE REQUEST-----
`

func TestRouteRequestCertificate(t *testing.T) {
	certPortal, mocks := newService(t)
	mockSigner := mocks.s
	assert.NotNil(t, certPortal)

	srv := httptest.NewServer(certPortal.GetRouter())
	defer srv.Close()

	csrData, err := ioutil.ReadFile(filepath.Join(pkiDir, "org-nlx-test.csr"))
	assert.NoError(t, err)

	csr := string(csrData)

	certificateRequest, err := json.Marshal(&certportal.CertificateRequest{
		Csr: csr,
	})
	assert.NoError(t, err)

	certificateRequestWithoutSAN, err := json.Marshal(&certportal.CertificateRequest{
		Csr: csrWithoutSAN,
	})
	assert.NoError(t, err)

	tests := []struct {
		requestBody        []byte
		setupMock          func()
		expectedStatusCode int
		expectedBody       string
	}{
		{
			certificateRequest,
			func() {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: csr,
				}).Return([]byte("test_cert"), nil)
			},
			http.StatusCreated,
			`{"certificate":"test_cert"}` + "\n",
		},
		{
			certificateRequestWithoutSAN,
			func() {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: csrWithoutSAN,
					Hosts:   []string{"hostname.test.local"},
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
			certificateRequest,
			func() {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: csr,
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
		resp.Body.Close()
	}
}

func TestRouteRoot(t *testing.T) {
	certPortal, mocks := newService(t)
	mockSigner := mocks.s
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

		resp.Body.Close()
	}
}

type serviceMocks struct {
	s *mock.MockSigner
}

func newService(t *testing.T) (*certportal.CertPortal, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		s: mock.NewMockSigner(ctrl),
	}

	service := certportal.NewCertPortal(zap.NewNop(), func() (signer.Signer, error) {
		return mocks.s, nil
	})

	return service, mocks
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
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()

	resp, err = http.Get(fmt.Sprintf("%s/root.crt", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	resp.Body.Close()
}
