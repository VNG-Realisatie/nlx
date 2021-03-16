// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/management"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

var testPublicKeyPEM = `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEArN5xGkM73tJsCpKny59e5lXNRY+eT0sbWyEGsR1qIPRKmLSiRHl3
xMsovn5mo6jN3eeK/Q4wKd6Ae5XGzP63pTG6U5KVVB74eQxSFfV3UEOrDaJ78X5m
BZO+Ku21V2QFr44tvMh5IZDX3RbMB/4Kad6sapmSF00HWrqTVMkrEsZ98DTb5nwG
Lh3kISnct4tLyVSpsl9s1rtkSgGUcs1TIvWxS2D2mOsSL1HRdUNcFQmzchbfG87k
XPvicoOISAZDJKDqWp3iuH0gJpQ+XMBfmcD90I7Z/cRQjWP3P93B3V06cJkd00cE
IRcIQqF8N+lE01H88Fi+wePhZRy92NP54wIDAQAB
-----END RSA PUBLIC KEY-----
`

func newCertificateBundle() (*common_tls.CertificateBundle, error) {
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	return common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)
}

func setProxyMetadata(ctx context.Context) context.Context {
	md := metadata.Pairs(
		"nlx-organization", "organization-a",
		"nlx-public-key-der", "ZHVtbXktcHVibGljLWtleQo=",
		"nlx-public-key-fingerprint", "1655A0AB68576280",
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
	proc := process.NewProcess(logger)

	bundle, err := newCertificateBundle()
	assert.NoError(t, err)

	s := server.NewManagementService(
		logger,
		proc,
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
