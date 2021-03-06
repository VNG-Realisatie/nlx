package api

import (
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/oidc"
)

type certFiles struct {
	certFile, keyFile, rootCertFile string
}

var pkiDir = filepath.Join("..", "..", "..", "testing", "pki")

var tests = []struct {
	name                         string
	cert                         certFiles
	orgCert                      certFiles
	db                           database.ConfigDatabase
	directoryInspectionAddress   string
	directoryRegistrationAddress string
	expectedErrorMessage         string
}{
	{
		"certificate_is_missing_organization",
		certFiles{
			filepath.Join(pkiDir, "org-without-name-chain.pem"),
			filepath.Join(pkiDir, "org-without-name-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		certFiles{
			filepath.Join(pkiDir, "org-without-name-chain.pem"),
			filepath.Join(pkiDir, "org-without-name-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		&mock_database.MockConfigDatabase{},
		"",
		"",
		"cannot obtain organization name from self cert",
	},
	{
		"postgres_connection_is_missing",
		certFiles{
			filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
			filepath.Join(pkiDir, "org-nlx-test-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		certFiles{
			filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
			filepath.Join(pkiDir, "org-nlx-test-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		nil,
		"",
		"",
		"database is not configured",
	},
	{
		"directory_inspection_address_is_missing",
		certFiles{
			filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
			filepath.Join(pkiDir, "org-nlx-test-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		certFiles{
			filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
			filepath.Join(pkiDir, "org-nlx-test-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		&mock_database.MockConfigDatabase{},
		"",
		"",
		"directory inspection address is not configured",
	},
	{
		"directory_registration_address_is_missing",
		certFiles{
			filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
			filepath.Join(pkiDir, "org-nlx-test-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		certFiles{
			filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
			filepath.Join(pkiDir, "org-nlx-test-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		&mock_database.MockConfigDatabase{},
		"directory-inspection.test:8443",
		"",
		"directory registration address is not configured",
	},
	{
		"happy_flow",
		certFiles{
			filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
			filepath.Join(pkiDir, "org-nlx-test-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		certFiles{
			filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
			filepath.Join(pkiDir, "org-nlx-test-key.pem"),
			filepath.Join(pkiDir, "ca-root.pem"),
		},
		&mock_database.MockConfigDatabase{},
		"directory-inspection.test:8443",
		"directory-registration.test:8443",
		"",
	},
}

func TestNewAPI(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	// Test exceptions during management-api creation
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			cert, err := common_tls.NewBundleFromFiles(test.cert.certFile, test.cert.keyFile, test.cert.rootCertFile)
			assert.NoError(t, err)

			orgCert, err := common_tls.NewBundleFromFiles(test.orgCert.certFile, test.orgCert.keyFile, test.orgCert.rootCertFile)
			assert.NoError(t, err)

			_, err = NewAPI(
				test.db,
				nil,
				logger,
				testProcess,
				cert,
				orgCert,
				test.directoryInspectionAddress,
				test.directoryRegistrationAddress,
				&oidc.Authenticator{},
				mock_auditlog.NewMockLogger(mockCtrl),
			)

			if test.expectedErrorMessage != "" {
				assert.EqualError(t, err, test.expectedErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
