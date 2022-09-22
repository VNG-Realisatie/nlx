// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package syncer_test

import (
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	common_tls "go.nlx.io/nlx/common/tls"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_management "go.nlx.io/nlx/management-api/pkg/management/mock"
	"go.nlx.io/nlx/management-api/pkg/syncer"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

type syncOutgoingAccessRequestTestCases map[string]struct {
	setup      func(mocks syncMocks)
	createArgs func(mocks syncMocks) *syncer.SyncArgs
	want       error
}

func Test_SyncOutgoingAccessRequests(t *testing.T) {
	testGroups := map[string]syncOutgoingAccessRequestTestCases{
		"received": getReceivedTestCases(t),
		"approved": getApprovedTestCases(t),
	}

	for groupName, testGroup := range testGroups {
		testGroup := testGroup

		for name, test := range testGroup {
			test := test

			t.Run(groupName+" "+name, func(t *testing.T) {
				mocks := newMocks(t)

				test.setup(mocks)
				args := test.createArgs(mocks)

				got := syncer.SyncOutgoingAccessRequests(args)

				assert.Equal(t, test.want, got)
			})
		}
	}
}

type syncMocks struct {
	db *mock_database.MockConfigDatabase
	mc *mock_management.MockClient
}

func newMocks(t *testing.T) syncMocks {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		t.Helper()
		ctrl.Finish()
	})

	mocks := syncMocks{
		db: mock_database.NewMockConfigDatabase(ctrl),
		mc: mock_management.NewMockClient(ctrl),
	}

	return mocks
}

func newCertificateBundle() (*common_tls.CertificateBundle, error) {
	pkiDir := filepath.Join("..", "..", "..", "testing", "pki")

	return common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
}
