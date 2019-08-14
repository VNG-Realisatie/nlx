// nolint:dupl
package configservice_test

import (
	"context"
	"testing"

	"go.nlx.io/nlx/config-api/configapi"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/config-api/configservice"
	mock "go.nlx.io/nlx/config-api/configservice/mock"

	"go.nlx.io/nlx/common/process"
	"go.uber.org/zap"
)

func TestListInways(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock.NewMockConfigDatabase(mockCtrl)
	service := configservice.New(logger, testProcess, mockDatabase)

	mockListInways := []*configapi.Inway{
		{Name: "inway42.test"},
		{Name: "inway43.test"},
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
