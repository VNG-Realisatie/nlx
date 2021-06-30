// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package basicauth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/pkg/api"
	"go.nlx.io/nlx/management-api/pkg/basicauth"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
)

type authenticatorMocks struct {
	configDatabase *mock_database.MockConfigDatabase
}

func newAuthenticator(t *testing.T) (api.Authenticator, authenticatorMocks) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := authenticatorMocks{
		configDatabase: mock_database.NewMockConfigDatabase(ctrl),
	}

	authenticator := basicauth.NewAuthenticator(mocks.configDatabase, zap.NewNop())

	return authenticator, mocks
}
