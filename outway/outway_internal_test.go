// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/common/monitoring"
	"go.nlx.io/nlx/common/nlxversion"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/api/mock"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

//nolint:funlen // it is a test
func TestUpdateServiceList(t *testing.T) {
	// Create mock directory client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock.NewMockDirectoryClient(ctrl)

	logger := zaptest.NewLogger(t)

	orgCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	require.NoError(t, err)

	o := &Outway{
		ctx:               context.Background(),
		directoryClient:   client,
		logger:            logger,
		orgCert:           orgCert,
		servicesHTTP:      make(map[string]HTTPService),
		servicesDirectory: make(map[string]*directoryapi.ListServicesResponse_Service),
	}

	o.monitorService, err = monitoring.NewMonitoringService("localhost:8080", logger)
	assert.Nil(t, err)

	ctx := context.Background()

	// Make the mock directory client return an error when calling ListServices
	client.EXPECT().ListServices(
		nlxversion.NewGRPCContext(ctx, "outway"),
		&emptypb.Empty{}).Return(nil, fmt.Errorf("mock error"))

	// Test of updateServiceList generates the correct error
	err = o.updateServiceList()
	assert.EqualError(t, err, "failed to fetch services from directory: mock error")

	mockServiceAInways := []*directoryapi.Inway{
		{
			Address: "mock-service-a-1:123",
			State:   directoryapi.Inway_UP,
		},
		{
			Address: "mock-service-a-2:123",
			State:   directoryapi.Inway_DOWN,
		},
	}

	mockServiceBInways := []*directoryapi.Inway{
		{
			Address: "mock-service-b-1:123",
			State:   directoryapi.Inway_UP,
		},
		{
			Address: "mock-service-b-2:123",
			State:   directoryapi.Inway_UP,
		},
	}

	// Make the mock directory client provide a list of services when calling ListServices
	client.EXPECT().ListServices(nlxversion.NewGRPCContext(ctx, "outway"), &emptypb.Empty{}).Return(
		&directoryapi.ListServicesResponse{
			Services: []*directoryapi.ListServicesResponse_Service{
				{
					Name: "mock-service-a",
					Organization: &directoryapi.Organization{
						SerialNumber: "00000000000000000001",
						Name:         "mock-org-a",
					},
					Inways: mockServiceAInways,
				},
				{
					Name: "mock-service-b",
					Organization: &directoryapi.Organization{
						SerialNumber: "00000000000000000002",
						Name:         "mock-org-b",
					},
					Inways: mockServiceBInways,
				},
				{
					Name: "mock-service-c",
					Organization: &directoryapi.Organization{
						SerialNumber: "00000000000000000003",
						Name:         "mock-org-c",
					},
				},
			},
		}, nil)

	// Test of updateServiceList creates a correct o.services map
	err = o.updateServiceList()
	assert.Nil(t, err)

	mockServiceAFullName := "00000000000000000001.mock-service-a"
	mockServiceBFullName := "00000000000000000002.mock-service-b"

	// mock-service-c should not be included because this service does not have any inwayaddresses
	assert.Len(t, o.servicesDirectory, 2, fmt.Sprintf("%v", o.servicesDirectory))

	// create the HttpServices
	o.getService("00000000000000000001", "mock-service-a")
	o.getService("00000000000000000002", "mock-service-b")
	o.getService("00000000000000000003", "mock-service-c")

	// mock-service-c should not be included because this service does not have any inwayaddresses
	t.Log(o.servicesHTTP)
	assert.Len(t, o.servicesHTTP, 2, fmt.Sprintf("%v", o.servicesHTTP))

	tests := []struct {
		serviceName string
		inways      []*directoryapi.Inway
	}{
		{
			serviceName: mockServiceAFullName,
			// only one valid true healthy address
			inways: []*directoryapi.Inway{
				{
					Address: "mock-service-a-1:123",
				},
			},
		},
		{
			mockServiceBFullName,
			mockServiceBInways,
		},
	}

	for _, test := range tests {
		assert.Contains(t, o.servicesHTTP, test.serviceName)
		assert.ElementsMatch(
			t,
			o.servicesHTTP[test.serviceName].GetInwayAddresses(),
			func() []string {
				var addresses []string

				for _, i := range test.inways {
					addresses = append(addresses, i.Address)
				}

				return addresses
			}())
	}
}
