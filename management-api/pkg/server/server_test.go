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

var (
	testPublicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArN5xGkM73tJsCpKny59e
5lXNRY+eT0sbWyEGsR1qIPRKmLSiRHl3xMsovn5mo6jN3eeK/Q4wKd6Ae5XGzP63
pTG6U5KVVB74eQxSFfV3UEOrDaJ78X5mBZO+Ku21V2QFr44tvMh5IZDX3RbMB/4K
ad6sapmSF00HWrqTVMkrEsZ98DTb5nwGLh3kISnct4tLyVSpsl9s1rtkSgGUcs1T
IvWxS2D2mOsSL1HRdUNcFQmzchbfG87kXPvicoOISAZDJKDqWp3iuH0gJpQ+XMBf
mcD90I7Z/cRQjWP3P93B3V06cJkd00cEIRcIQqF8N+lE01H88Fi+wePhZRy92NP5
4wIDAQAB
-----END PUBLIC KEY-----
`
	testPublicKeyFingerprint = "60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4="
)

func newCertificateBundle() (*common_tls.CertificateBundle, error) {
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	return common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)
}

func setProxyMetadata(ctx context.Context) context.Context {
	// @TODO: generate values based on the actual bundle?
	md := metadata.Pairs(
		"nlx-organization", "organization-a",
		"nlx-public-key-der", "ZHVtbXktcHVibGljLWtleQo=",
		"nlx-public-key-fingerprint", testPublicKeyFingerprint,
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
