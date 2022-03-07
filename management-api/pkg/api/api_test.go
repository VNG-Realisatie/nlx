package api

import (
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	"go.nlx.io/nlx/management-api/pkg/oidc"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

func TestNewAPI(t *testing.T) {
	logger := zap.NewNop()

	var pkiDir = filepath.Join("..", "..", "..", "testing", "pki")

	tests := map[string]struct {
		cert                 common_testing.CertificateBundleOrganizationName
		orgCert              common_testing.CertificateBundleOrganizationName
		db                   database.ConfigDatabase
		directoryAddress     string
		txlogAddress         string
		expectedErrorMessage string
	}{
		"certificate_is_missing_organization": {
			cert:                 common_testing.OrgWithoutName,
			orgCert:              common_testing.OrgWithoutName,
			db:                   &mock_database.MockConfigDatabase{},
			directoryAddress:     "directory.test:8443",
			txlogAddress:         "txlog.test:8443",
			expectedErrorMessage: "cannot obtain organization name from self cert",
		},
		"postgres_connection_is_missing": {
			cert:                 common_testing.OrgNLXTest,
			orgCert:              common_testing.OrgNLXTest,
			db:                   nil,
			directoryAddress:     "directory.test:8443",
			txlogAddress:         "txlog.test:8443",
			expectedErrorMessage: "database is not configured",
		},
		"directory_inspection_address_is_missing": {
			cert:                 common_testing.OrgNLXTest,
			orgCert:              common_testing.OrgNLXTest,
			db:                   &mock_database.MockConfigDatabase{},
			directoryAddress:     "",
			txlogAddress:         "txlog.test:8443",
			expectedErrorMessage: "directory address is not configured",
		},
		"happy_flow": {
			cert:                 common_testing.OrgNLXTest,
			orgCert:              common_testing.OrgNLXTest,
			db:                   &mock_database.MockConfigDatabase{},
			directoryAddress:     "directory.test:8443",
			txlogAddress:         "txlog.test:8443",
			expectedErrorMessage: "",
		},
	}

	// Test exceptions during management-api creation
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			internalCert, err := common_testing.GetCertificateBundle(pkiDir, tt.cert)
			assert.NoError(t, err)

			orgCert, err := common_testing.GetCertificateBundle(pkiDir, tt.orgCert)
			assert.NoError(t, err)

			args := &NewAPIArgs{
				DB:               tt.db,
				TXlogDB:          nil,
				Logger:           logger,
				InternalCert:     internalCert,
				OrgCert:          orgCert,
				DirectoryAddress: tt.directoryAddress,
				TXLogAddress:     tt.txlogAddress,
				Authenticator:    &oidc.Authenticator{},
				AuditLogger:      mock_auditlog.NewMockLogger(mockCtrl),
			}

			_, err = NewAPI(args)

			if tt.expectedErrorMessage != "" {
				assert.EqualError(t, err, tt.expectedErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
