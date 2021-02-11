// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

//nolint:dupl // test package
package server_test

import (
	"context"
	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	"net"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_directory "go.nlx.io/nlx/management-api/pkg/directory/mock"
	"go.nlx.io/nlx/management-api/pkg/server"
)

type args struct {
	peer     *peer.Peer
	database *database.Inway
	request  *api.Inway
}

var createInwayTests = []struct {
	name string
	args args
	want *api.Inway
}{
	{
		name: "ip address from context",
		args: args{
			database: &database.Inway{Name: "inway42.basic", IPAddress: "127.1.1.1"},
			request:  &api.Inway{Name: "inway42.basic"},
			peer:     &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(127, 1, 1, 1)}}},
		want: &api.Inway{
			Name:      "inway42.basic",
			IpAddress: "127.1.1.1",
		},
	},
	{
		name: "ip address from request is ignored",
		args: args{
			database: &database.Inway{Name: "inway42.ignore-ip", IPAddress: "127.1.1.1"},
			request:  &api.Inway{Name: "inway42.ignore-ip", IpAddress: "127.2.2.2"},
			peer:     &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv4(127, 1, 1, 1)}}},
		want: &api.Inway{
			Name:      "inway42.ignore-ip",
			IpAddress: "127.1.1.1",
		},
	},
	{
		name: "the connection context must contain an address",
		args: args{
			database: &database.Inway{Name: "inway42.ip-context-required"},
			request:  &api.Inway{Name: "inway42.ip-context-required"},
			peer:     &peer.Peer{Addr: nil}},
	},
	{
		name: "ipv6",
		args: args{
			database: &database.Inway{Name: "inway42.ipv6", IPAddress: "::1"},
			request:  &api.Inway{Name: "inway42.ipv6"},
			peer:     &peer.Peer{Addr: &net.TCPAddr{IP: net.IPv6loopback}}},
		want: &api.Inway{
			Name:      "inway42.ipv6",
			IpAddress: "::1",
		},
	},
}

func TestCreateInway(t *testing.T) {
	for _, tt := range createInwayTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			logger := zaptest.Logger(t)
			testProcess := process.NewProcess(logger)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := peer.NewContext(context.Background(), tt.args.peer)
			mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
			if tt.want != nil {
				mockDatabase.EXPECT().CreateInway(ctx, tt.args.database)
			}
			service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, mock_auditlog.NewMockLogger(mockCtrl))

			response, err := service.CreateInway(ctx, tt.args.request)
			if tt.want != nil {
				assert.NoError(t, err, "could not create inway")
				assert.Equal(t, tt.want, response)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, mock_auditlog.NewMockLogger(mockCtrl))

	getInwayRequest := &api.GetInwayRequest{
		Name: "inway42.test",
	}

	mockDatabase.EXPECT().GetInway(ctx, "inway42.test").Return(nil, database.ErrNotFound)

	_, actualError := service.GetInway(ctx, getInwayRequest)
	expectedError := status.Error(codes.NotFound, "inway not found")
	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)

	mockInwayResponse := &database.Inway{
		Name:      "inway42.test",
		IPAddress: "",
		Services: []*database.Service{{
			Name: "forty-two",
		}},
	}

	mockDatabase.EXPECT().GetInway(ctx, "inway42.test").Return(mockInwayResponse, nil)

	getInwayResponse, err := service.GetInway(ctx, getInwayRequest)
	assert.NoError(t, err)

	expectedResponse := &api.Inway{
		Name:     "inway42.test",
		Services: []*api.Inway_Service{{Name: "forty-two"}},
	}

	assert.Equal(t, expectedResponse, getInwayResponse)
}

func TestUpdateInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockInway := &database.Inway{
		ID:   1,
		Name: "inway42.test",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().UpdateInway(ctx, mockInway)
	mockDatabase.EXPECT().GetInway(ctx, "inway42.test").Return(mockInway, nil)

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, mock_auditlog.NewMockLogger(mockCtrl))

	updateInwayRequest := &api.UpdateInwayRequest{
		Name: "inway42.test",
		Inway: &api.Inway{
			Name: "inway42.test",
		},
	}

	updateInwayResponse, err := service.UpdateInway(ctx, updateInwayRequest)
	assert.NoError(t, err)

	expectedResponse := &api.Inway{
		Name: "inway42.test",
	}

	assert.Equal(t, expectedResponse, updateInwayResponse)
}

func TestDeleteInway(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)
	mockDatabase.EXPECT().DeleteInway(ctx, "inway42.test")

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, mock_auditlog.NewMockLogger(mockCtrl))

	deleteRequest := &api.DeleteInwayRequest{
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

	mockDatabase := mock_database.NewMockConfigDatabase(mockCtrl)

	mockListInways := []*database.Inway{
		{Name: "inway42.test"},
		{Name: "inway43.test"},
		{
			Name:        "inway.test",
			Version:     "1.0.0",
			Hostname:    "inway.test.local",
			SelfAddress: "inway.nlx",
			Services: []*database.Service{
				{
					Name: "mock-service",
					Inways: []*database.Inway{{
						Name: "inway.test"},
					},
				},
			},
		},
	}

	mockDatabase.EXPECT().ListInways(ctx).Return(mockListInways, nil)

	service := server.NewManagementService(logger, testProcess, mock_directory.NewMockClient(mockCtrl), nil, mockDatabase, mock_auditlog.NewMockLogger(mockCtrl))
	actualResponse, err := service.ListInways(ctx, &api.ListInwaysRequest{})
	assert.NoError(t, err)

	expectedResponse := &api.ListInwaysResponse{
		Inways: []*api.Inway{
			{
				Name: "inway42.test",
			},
			{
				Name: "inway43.test",
			},
			{
				Name:        "inway.test",
				Version:     "1.0.0",
				Hostname:    "inway.test.local",
				SelfAddress: "inway.nlx",
				Services: []*api.Inway_Service{
					{
						Name: "mock-service",
					},
				},
			},
		},
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
