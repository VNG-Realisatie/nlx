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
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAxahK/ruBG74MZ2/Z71ll
WS1OMJy6xs9qpQY7YC7C+u59JoqNdoWToV6EZRfPYzh61BWsyKlqvkl11HhD6HVk
WmdidYJtmoJqmFeWm02a6RkP4XbBiOCm9xxX/xlZRWubTCswaLkfmlI2IYxgpLIP
QuvO2nbor8YG4dS7u7u7yrtl1dOBD1utlMCzX5j8vG+BaHUqE1kBWIE5kg9ogeVf
wa/w30EPcD+5gdknn5uGoTFP/xi6WiZ+6MJli1CPjrHX0N73ZMSdgHK+4jk8Kdrz
Fou5sNtCl+CTdzhDhwYJxJv/McsgqPfXsOdk0T3QUcCqWsawJ8VblJYYwyj1WW7l
bJSygJjvOTG+C2+vbht3mKvimKpx/+8S/Zg+g7nen//SvFQhe2wI7Eaottgk/abU
6i3ntvSty4EyxFPnchKa7EXeFAsp4stO0Q5iTE4rEdDotwaWrmcN54UQr2ZOVPJ/
BGGG6SxeciX9jB9I1FHBngMyiXVDgMlgGa9Ke3y1V+Yaqh3LOp6JXnjXp50Ke0nc
CMa7tBd6GGJqV4hl3daYj7yyBWzB3E2d/u+gJx1e9mxqgA0V7nidh2CRelHtczhC
O5/DpYFGnjKm4YMkzSb7CxRDrL2OJeyvM3tKyRZES5eEiedMcpjvm5ULzZeCp2r3
P72Jy9qTigqNYoIHBYMpFzUCAwEAAQ==
-----END PUBLIC KEY-----
`
	testPublicKeyFingerprint = "g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc="
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
