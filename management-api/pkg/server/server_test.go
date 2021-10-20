// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"path/filepath"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"

	common_tls "go.nlx.io/nlx/common/tls"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

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
	db *mock_database.MockConfigDatabase
	al *mock_auditlog.MockLogger
	dc *mock_directory.MockClient
	mc *mock_management.MockClient
}

func newService(t *testing.T) (*server.ManagementService, *common_tls.CertificateBundle, serviceMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := serviceMocks{
		dc: mock_directory.NewMockClient(ctrl),
		al: mock_auditlog.NewMockLogger(ctrl),
		db: mock_database.NewMockConfigDatabase(ctrl),
		mc: mock_management.NewMockClient(ctrl),
	}

	logger := zaptest.Logger(t)

	bundle, err := newCertificateBundle()
	assert.NoError(t, err)

	s := server.NewManagementService(
		logger,
		mocks.dc,
		bundle,
		mocks.db,
		nil,
		mocks.al,
		func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error) {
			return mocks.mc, nil
		},
	)

	return s, bundle, mocks
}
