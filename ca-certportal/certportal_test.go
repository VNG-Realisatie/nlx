// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package certportal_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/cloudflare/cfssl/signer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	certportal "go.nlx.io/nlx/ca-certportal"
	mock "go.nlx.io/nlx/ca-certportal/mock"
)

func TestRequestCertificate(t *testing.T) {
	tests := map[string]struct {
		csr        string
		setup      func(mocks certportalMocks)
		wantResult certportal.Certificate
		wantErr    error
	}{
		"when_providing_an_invalid_csr": {
			csr:     "invalid",
			wantErr: certportal.ErrFailedToParseCSR,
		},
		"when_providing_a_key_instead_of_csr": {
			csr:     string(getKey()),
			wantErr: certportal.ErrFailedToParseCSR,
		},
		"when_providing_a_csr_without_san": {
			csr: string(getCsrWithoutSAN()),
			setup: func(m certportalMocks) {
				m.signer.EXPECT().Sign(signer.SignRequest{
					Request: string(getCsrWithoutSAN()),
					Hosts:   []string{"hostname.test.local"},
				}).Return([]byte("test_cert"), nil)
			},
			wantResult: []byte("test_cert"),
		},
		"when_the_signer_returns_an_error": {
			csr: string(getCsr()),
			setup: func(m certportalMocks) {
				m.signer.EXPECT().Sign(signer.SignRequest{
					Request: string(getCsr()),
				}).Return(nil, fmt.Errorf("arbitrary error"))
			},
			wantErr: certportal.ErrFailedToSignCSR,
		},
		"happy_flow": {
			csr: string(getCsr()),
			setup: func(m certportalMocks) {
				m.signer.EXPECT().Sign(signer.SignRequest{
					Request: string(getCsr()),
				}).Return([]byte("test_cert"), nil)
			},
			wantResult: []byte("test_cert"),
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mocks := createMocks(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			actual, err := certportal.RequestCertificate(tt.csr, func() (signer.Signer, error) {
				return mocks.signer, nil
			})

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantResult, actual)
		})
	}
}

type certportalMocks struct {
	signer *mock.MockSigner
}

func createMocks(t *testing.T) (mocks certportalMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks = certportalMocks{
		signer: mock.NewMockSigner(ctrl),
	}

	return mocks
}

func getCsr() []byte {
	var pkiDir = filepath.Join("..", "testing", "pki")

	csr, err := ioutil.ReadFile(filepath.Join(pkiDir, "org-nlx-test.csr"))
	if err != nil {
		panic(err)
	}

	return csr
}

func getCsrWithoutSAN() []byte {
	return []byte(`
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
`)
}

func getKey() []byte {
	return []byte(`
-----BEGIN PRIVATE KEY-----
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
nc680wj8nW/CkOxeEoWZzsQy3SxTlg==
-----END PRIVATE KEY-----`)
}
