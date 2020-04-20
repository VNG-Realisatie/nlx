// nolint:dupl
package configservice_test

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/config-api/configservice"
	mock "go.nlx.io/nlx/config-api/configservice/mock"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

func TestListInways(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock.NewMockConfigDatabase(mockCtrl)
	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	mockListInways := []*configapi.Inway{
		{Name: "inway42.test"},
		{Name: "inway43.test"},
		{
			Name:        "inway.test",
			Version:     "1.0.0",
			Hostname:    "inway.test.local",
			SelfAddress: "inway.nlx",
		},
	}

	mockDatabase.EXPECT().ListInways(ctx).Return(mockListInways, nil)

	actualResponse, err := service.ListInways(ctx, &configapi.ListInwaysRequest{})
	if err != nil {
		t.Fatal("could not get list of inways", err)
	}

	expectedResponse := &configapi.ListInwaysResponse{
		Inways: mockListInways,
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
