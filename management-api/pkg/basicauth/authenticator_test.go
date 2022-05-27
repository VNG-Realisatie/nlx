// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package basicauth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/pkg/api"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"go.nlx.io/nlx/management-api/pkg/basicauth"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
)

type authenticatorMocks struct {
	configDatabase *mock_database.MockConfigDatabase
	auditLogger    *mock_auditlog.MockLogger
}

func newAuthenticator(t *testing.T) (api.Authenticator, authenticatorMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := authenticatorMocks{
		configDatabase: mock_database.NewMockConfigDatabase(ctrl),
		auditLogger:    mock_auditlog.NewMockLogger(ctrl),
	}

	authenticator := basicauth.NewAuthenticator(mocks.configDatabase, mocks.auditLogger, zap.NewNop())

	return authenticator, mocks
}
