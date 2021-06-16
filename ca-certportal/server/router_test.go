// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

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

	mock "go.nlx.io/nlx/ca-certportal/mock"
	"go.nlx.io/nlx/ca-certportal/server"
)

var pkiDir = filepath.Join("..", "..", "testing", "pki")

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

const key = `
-----BEGIN PRIVATE KEY-----
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
nc680wj8nW/CkOxeEoWZzsQy3SxTlg==
-----END PRIVATE KEY-----`

func TestRouteRequestCertificate(t *testing.T) {
	certPortal, mocks := newService(t)
	assert.NotNil(t, certPortal)

	mockSigner := mocks.s

	srv := httptest.NewServer(certPortal.GetRouter())
	defer srv.Close()

	csrData, err := ioutil.ReadFile(filepath.Join(pkiDir, "org-nlx-test.csr"))
	assert.NoError(t, err)

	csr := string(csrData)

	certificateRequest, err := json.Marshal(&server.CertificateRequest{
		Csr: csr,
	})
	assert.NoError(t, err)

	certificateRequestWithoutSAN, err := json.Marshal(&server.CertificateRequest{
		Csr: csrWithoutSAN,
	})
	assert.NoError(t, err)

	certificateRequestWithKey, err := json.Marshal(&server.CertificateRequest{
		Csr: key,
	})
	assert.NoError(t, err)

	tests := map[string]struct {
		requestBody        []byte
		setupMock          func()
		expectedStatusCode int
		expectedBody       string
	}{
		"without_san": {
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
		"invalid_csr": {
			[]byte("invalid"),
			func() {},
			http.StatusBadRequest,
			"could not decode request body\n",
		},
		"with_key_instead_of_csr": {
			certificateRequestWithKey,
			func() {},
			http.StatusBadRequest,
			"failed to parse csr\n",
		},
		"failed_to_sign": {
			certificateRequest,
			func() {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: csr,
				}).Return(nil, fmt.Errorf("error signing request"))
			},
			http.StatusInternalServerError,
			"Internal Server Error\n",
		},
		"happy_path": {
			certificateRequest,
			func() {
				mockSigner.EXPECT().Sign(signer.SignRequest{
					Request: csr,
				}).Return([]byte("test_cert"), nil)
			},
			http.StatusCreated,
			`{"certificate":"test_cert"}` + "\n",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			tt.setupMock()

			resp, err := http.Post(fmt.Sprintf("%s/api/request_certificate", srv.URL), "application/json", bytes.NewReader(tt.requestBody))
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)

			responseBody, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, string(responseBody))

			resp.Body.Close()
		})
	}
}

func TestRouteRoot(t *testing.T) {
	certPortal, mocks := newService(t)
	assert.NotNil(t, certPortal)

	mockSigner := mocks.s

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

func newService(t *testing.T) (*server.CertPortal, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		s: mock.NewMockSigner(ctrl),
	}

	service := server.NewCertPortal(zap.NewNop(), func() (signer.Signer, error) {
		return mocks.s, nil
	})

	return service, mocks
}

func TestRoutesInvalidSigner(t *testing.T) {
	certPortal := server.NewCertPortal(zap.NewNop(), func() (signer.Signer, error) {
		return nil, fmt.Errorf("unable to create certificate signer")
	})
	assert.NotNil(t, certPortal)

	srv := httptest.NewServer(certPortal.GetRouter())
	defer srv.Close()

	jsonBytesCertificateRequest, err := json.Marshal(&server.CertificateRequest{
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
