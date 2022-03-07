package api

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	common_testing "go.nlx.io/nlx/testing/testingutils"
	"go.nlx.io/nlx/txlog-api/domain/txlog/storage"
	mock_storage "go.nlx.io/nlx/txlog-api/domain/txlog/storage/mock"
)

func TestNewAPI(t *testing.T) {
	type certFiles struct {
		certFile, keyFile, rootCertFile string
	}

	var pkiDir = filepath.Join("..", "..", "..", "testing", "pki")

	certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	var tests = map[string]struct {
		logger  *zap.Logger
		cert    *common_tls.CertificateBundle
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
			logger:  zap.NewNop(),
			cert:    certBundle,
			storage: nil,
			wantErr: errors.New("storage is required"),
		},
		"when_no_logger": {
			logger:  nil,
			cert:    certBundle,
			storage: &mock_storage.MockRepository{},
			wantErr: errors.New("logger is required"),
		},
		"happy_flow": {
			logger:  zap.NewNop(),
			cert:    certBundle,
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

			_, err := NewAPI(
				tt.logger,
				tt.cert,
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
