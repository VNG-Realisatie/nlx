// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server_test

import (
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/outway/pkg/server"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

func newCertificateBundle() (*common_tls.CertificateBundle, error) {
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	return common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
}

func newService(t *testing.T, mockSigner server.SignFunction) *server.OutwayService {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	logger := zaptest.NewLogger(t)

	bundle, err := newCertificateBundle()
	assert.NoError(t, err)

	s := server.NewOutwayService(
		logger,
		bundle,
		mockSigner,
	)

	return s
}
