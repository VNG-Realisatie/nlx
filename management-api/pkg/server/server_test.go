// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/metadata"

	common_tls "go.nlx.io/nlx/common/tls"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/outway"
	mock_outway "go.nlx.io/nlx/management-api/pkg/outway/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
	mock_txlog "go.nlx.io/nlx/management-api/pkg/txlog/mock"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
	mock_txlogdb "go.nlx.io/nlx/management-api/pkg/txlogdb/mock"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

var fixtureTime = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

type testClock struct{}

func (tc *testClock) Now() time.Time {
	return fixtureTime
}

func newCertificateBundle() (*common_tls.CertificateBundle, error) {
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	return common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
}

func setProxyMetadata(t *testing.T, ctx context.Context) context.Context {
	orgCert, err := newCertificateBundle()
	require.NoError(t, err)

	return setProxyMetadataWithCertBundle(t, ctx, orgCert)
}

func setProxyMetadataWithCertBundle(t *testing.T, ctx context.Context, certBundle *common_tls.CertificateBundle) context.Context {
	publicKeyDER, err := x509.MarshalPKIXPublicKey(certBundle.PublicKey())
	require.NoError(t, err)

	organizationName := certBundle.Certificate().Subject.Organization[0]
	organizationSerialNumber := certBundle.Certificate().Subject.SerialNumber
	publicKeyDEREncoded := base64.StdEncoding.EncodeToString(publicKeyDER)
	publicKeyFingerprint := common_tls.X509PublicKeyFingerprint(certBundle.Certificate())

	md := metadata.Pairs(
		"nlx-organization-name", organizationName,
		"nlx-organization-serial-number", organizationSerialNumber,
		"nlx-public-key-der", publicKeyDEREncoded,
		"nlx-public-key-fingerprint", publicKeyFingerprint,
	)

	return metadata.NewIncomingContext(ctx, md)
}

type serviceMocks struct {
	db      *mock_database.MockConfigDatabase
	dbTxLog *mock_txlogdb.MockTxlogDatabase
	al      *mock_auditlog.MockLogger
	dc      *mock_directory.MockClient
	tx      *mock_txlog.MockClient
	mc      *mock_management.MockClient
	oc      *mock_outway.MockClient
}

func newService(t *testing.T) (*server.ManagementService, *common_tls.CertificateBundle, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		dc:      mock_directory.NewMockClient(ctrl),
		tx:      mock_txlog.NewMockClient(ctrl),
		al:      mock_auditlog.NewMockLogger(ctrl),
		db:      mock_database.NewMockConfigDatabase(ctrl),
		dbTxLog: mock_txlogdb.NewMockTxlogDatabase(ctrl),
		mc:      mock_management.NewMockClient(ctrl),
		oc:      mock_outway.NewMockClient(ctrl),
	}

	s, orgCert := newServer(t, mocks)

	return s, orgCert, mocks
}

func newServiceWithoutTXLog(t *testing.T) (*server.ManagementService, *common_tls.CertificateBundle, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		dc:      mock_directory.NewMockClient(ctrl),
		tx:      mock_txlog.NewMockClient(ctrl),
		al:      mock_auditlog.NewMockLogger(ctrl),
		dbTxLog: nil,
		db:      mock_database.NewMockConfigDatabase(ctrl),
		mc:      mock_management.NewMockClient(ctrl),
		oc:      mock_outway.NewMockClient(ctrl),
	}

	s, orgCert := newServer(t, mocks)

	return s, orgCert, mocks
}

func newServer(t *testing.T, mocks serviceMocks) (*server.ManagementService, *common_tls.CertificateBundle) {
	logger := zaptest.NewLogger(t)

	orgCert, err := newCertificateBundle()
	assert.NoError(t, err)

	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	internalCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.NLXTestInternal)
	require.NoError(t, err)

	var txLog txlogdb.TxlogDatabase = nil
	if mocks.dbTxLog != nil {
		txLog = mocks.dbTxLog
	}

	return server.NewManagementService(
		&server.NewManagementServiceArgs{
			Logger:          logger,
			DirectoryClient: mocks.dc,
			TxlogClient:     mocks.tx,
			OrgCert:         orgCert,
			InternalCert:    internalCert,
			ConfigDatabase:  mocks.db,
			TxlogDatabase:   txLog,
			AuditLogger:     mocks.al,
			CreateManagementClientFunc: func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error) {
				return mocks.mc, nil
			},
			CreateOutwayClientFunc: func(context.Context, string, *common_tls.CertificateBundle) (outway.Client, error) {
				return mocks.oc, nil
			},
			Clock: &testClock{},
		}), orgCert
}
