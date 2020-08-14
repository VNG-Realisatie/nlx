// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package inspectionservice_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"go.nlx.io/nlx/directory-inspection-api/pkg/database/mock"
)

func generateMockDirectoryDatabase(t *testing.T) *mock.MockDirectoryDatabase {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return mock.NewMockDirectoryDatabase(mockCtrl)
}
