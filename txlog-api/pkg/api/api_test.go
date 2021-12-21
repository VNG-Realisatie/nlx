package api

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/txlog-api/domain/txlog/storage"
	mock_storage "go.nlx.io/nlx/txlog-api/domain/txlog/storage/mock"
)

func TestNewAPI(t *testing.T) {
	type certFiles struct {
		certFile, keyFile, rootCertFile string
	}

	var pkiDir = filepath.Join("..", "..", "..", "testing", "pki")

	var tests = map[string]struct {
		logger  *zap.Logger
		cert    *certFiles
		storage storage.Repository
		wantErr error
	}{
		"when_no_certs": {
			logger:  zap.NewNop(),
			cert:    nil,
			storage: &mock_storage.MockRepository{},
			wantErr: errors.New("cert is required"),
		},
		"when_no_storage": {
			logger: zap.NewNop(),
			cert: &certFiles{
				filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-root.pem"),
			},
			storage: nil,
			wantErr: errors.New("storage is required"),
		},
		"when_no_logger": {
			logger: nil,
			cert: &certFiles{
				filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-root.pem"),
			},
			storage: &mock_storage.MockRepository{},
			wantErr: errors.New("logger is required"),
		},
		"happy_flow": {
			logger: zap.NewNop(),
			cert: &certFiles{
				filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
				filepath.Join(pkiDir, "org-nlx-test-key.pem"),
				filepath.Join(pkiDir, "ca-root.pem"),
			},
			storage: &mock_storage.MockRepository{},
			wantErr: nil,
		},
	}

	// Test exceptions during management-api creation
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			var cert *common_tls.CertificateBundle
			if tt.cert != nil {
				var err error
				cert, err = common_tls.NewBundleFromFiles(tt.cert.certFile, tt.cert.keyFile, tt.cert.rootCertFile)
				assert.NoError(t, err)
			}

			_, err := NewAPI(
				tt.logger,
				cert,
				tt.storage,
			)

			if tt.wantErr != nil {
				assert.Equal(t, err, tt.wantErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
