//nolint:dupl // test package
package configservice_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
	"go.nlx.io/nlx/management-api/configapi"
	"go.nlx.io/nlx/management-api/configservice"
	mock_configservice "go.nlx.io/nlx/management-api/configservice/mock"
)

func TestCreateInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	requestInway := &configapi.Inway{
		Name: "inway42.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().CreateInway(ctx, requestInway)

	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	responseInway, err := service.CreateInway(ctx, requestInway)
	if err != nil {
		t.Error("could not create inway", err)
	}

	assert.Equal(t, requestInway, responseInway)
}

func TestGetInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	getInwayRequest := &configapi.GetInwayRequest{
		Name: "inway42.test",
	}

	mockDatabase.EXPECT().GetInway(ctx, "inway42.test")

	_, actualError := service.GetInway(ctx, getInwayRequest)
	expectedError := status.Error(codes.NotFound, "inway not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	mockInwayResponse := &configapi.Inway{
		Name: "inway42.test",
	}
	expectedResponse := &configapi.Inway{
		Name:     "inway42.test",
		Services: []*configapi.Inway_Service{{Name: "forty-two"}},
	}
	mockServices := []*configapi.Service{{Name: "forty-two", Inways: []string{"inway42.test"}}}

	mockDatabase.EXPECT().ListServices(ctx).Return(mockServices, nil)
	mockDatabase.EXPECT().GetInway(ctx, "inway42.test").Return(mockInwayResponse, nil)

	getInwayResponse, err := service.GetInway(ctx, getInwayRequest)
	if err != nil {
		t.Fatal("could not get inway", err)
	}

	assert.Equal(t, expectedResponse, getInwayResponse)
}

func TestUpdateInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockInway := &configapi.Inway{
		Name: "inway42.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().UpdateInway(ctx, "inway42.test", mockInway)

	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	updateInwayRequest := &configapi.UpdateInwayRequest{
		Name:  "inway42.test",
		Inway: mockInway,
	}

	updateInwayResponse, err := service.UpdateInway(ctx, updateInwayRequest)
	if err != nil {
		t.Error("could not update inway", err)
	}

	assert.Equal(t, mockInway, updateInwayResponse)
}

func TestDeleteInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().DeleteInway(ctx, "inway42.test")

	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)

	deleteRequest := &configapi.DeleteInwayRequest{
		Name: "inway42.test",
	}

	_, err := service.DeleteInway(ctx, deleteRequest)
	if err != nil {
		t.Error("could not delete inway", err)
	}
}

func TestListInways(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_configservice.NewMockConfigDatabase(mockCtrl)

	mockListServices := []*configapi.Service{
		{
			Name:   "mock-service",
			Inways: []string{"inway43.test"},
		},
	}

	mockDatabase.EXPECT().ListServices(ctx).Return(mockListServices, nil)

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

	service := configservice.New(logger, testProcess, registrationapi.NewDirectoryRegistrationClient(nil), mockDatabase)
	actualResponse, err := service.ListInways(ctx, &configapi.ListInwaysRequest{})

	if err != nil {
		t.Fatal("could not get list of inways", err)
	}

	expectedResponse := &configapi.ListInwaysResponse{
		Inways: []*configapi.Inway{
			{
				Name: "inway42.test",
			},
			{
				Name: "inway43.test",
				Services: []*configapi.Inway_Service{
					{
						Name: "mock-service",
					},
				},
			},
			{
				Name:        "inway.test",
				Version:     "1.0.0",
				Hostname:    "inway.test.local",
				SelfAddress: "inway.nlx",
			},
		},
	}

	assert.Equal(t, expectedResponse, actualResponse)
}

func TestFilterServices(t *testing.T) {
	type args struct {
		services []*configapi.Service
		inway    *configapi.Inway
	}

	var filterServicesTests = []struct {
		name string
		want []*configapi.Inway_Service
		args args
	}{
		{
			name: "one service",
			args: args{
				services: []*configapi.Service{{
					Name:   "service1",
					Inways: []string{"inway1"},
				}, {
					Name:   "service2",
					Inways: []string{"inway2"},
				}},
				inway: &configapi.Inway{
					Name: "inway1",
				}},
			want: []*configapi.Inway_Service{{
				Name: "service1",
			}},
		},
		{
			name: "two services",
			args: args{
				services: []*configapi.Service{{
					Name:   "service11",
					Inways: []string{"inway1"},
				}, {
					Name:   "service12",
					Inways: []string{"inway1"},
				}, {
					Name:   "service2",
					Inways: []string{"inway2"},
				}},
				inway: &configapi.Inway{
					Name: "inway1",
				}},
			want: []*configapi.Inway_Service{{
				Name: "service11",
			}, {
				Name: "service12",
			}},
		},
		{
			name: "no services",
			args: args{
				services: []*configapi.Service{{
					Name:   "service1",
					Inways: []string{"inway1"},
				}},
				inway: &configapi.Inway{
					Name: "inway2",
				}},
			want: []*configapi.Inway_Service{},
		},
	}

	for _, tt := range filterServicesTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := configservice.FilterServices(tt.args.services, tt.args.inway)
			assert.Equal(t, tt.want, actual)
		})
	}
}
