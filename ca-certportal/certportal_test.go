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
