// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/common/monitoring"
	"go.nlx.io/nlx/common/nlxversion"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi/mock"
)

//nolint:funlen // it is a test
func TestUpdateServiceList(t *testing.T) {
	// Create mock directory client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock.NewMockDirectoryInspectionClient(ctrl)

	logger := zaptest.NewLogger(t)

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	o := &Outway{
		ctx:                       context.Background(),
		directoryInspectionClient: client,
		logger:                    logger,
		orgCert:                   cert,
		servicesHTTP:              make(map[string]HTTPService),
		servicesDirectory:         make(map[string]*inspectionapi.ListServicesResponse_Service),
	}

	var err error
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

	mockServiceAInwayAddresses := []string{"mock-service-a-1:123", "mock-service-a-2:123"}
	mockServiceBInwayAddresses := []string{"mock-service-b-1:123", "mock-service-b-2:123"}
	healthyStatesA := []bool{true, false}
	healthyStatesB := []bool{true, true}

	// Make the mock directory client provide a list of services when calling ListServices
	client.EXPECT().ListServices(nlxversion.NewGRPCContext(ctx, "outway"), &emptypb.Empty{}).Return(
		&inspectionapi.ListServicesResponse{
			Services: []*inspectionapi.ListServicesResponse_Service{
				{
					Name:             "mock-service-a",
					OrganizationName: "mock-org-a",
					InwayAddresses:   mockServiceAInwayAddresses,
					HealthyStates:    healthyStatesA,
				},
				{
					Name:             "mock-service-b",
					OrganizationName: "mock-org-b",
					InwayAddresses:   mockServiceBInwayAddresses,
					HealthyStates:    healthyStatesB,
				},
				{
					Name:             "mock-service-c",
					OrganizationName: "mock-org-c",
				},
			},
		}, nil)

	// Test of updateServiceList creates a correct o.services map
	err = o.updateServiceList()
	assert.Nil(t, err)

	mockServiceAFullName := "mock-org-a.mock-service-a"
	mockServiceBFullName := "mock-org-b.mock-service-b"

	// mock-service-c should not be included because this service does not have any inwayaddresses
	assert.Len(t, o.servicesDirectory, 2, fmt.Sprintf("%v", o.servicesDirectory))

	// create the HttpServices
	o.getService("mock-org-a", "mock-service-a")
	o.getService("mock-org-b", "mock-service-b")
	o.getService("mock-org-c", "mock-service-c")

	// mock-service-c should not be included because this service does not have any inwayaddresses
	t.Log(o.servicesHTTP)
	assert.Len(t, o.servicesHTTP, 2, fmt.Sprintf("%v", o.servicesHTTP))

	tests := []struct {
		serviceName    string
		inwayAddresses []string
	}{
		{
			mockServiceAFullName,
			// only one valid true healthy address
			[]string{"mock-service-a-1:123"},
		},
		{
			mockServiceBFullName,
			mockServiceBInwayAddresses,
		},
	}

	for _, test := range tests {
		assert.Contains(t, o.servicesHTTP, test.serviceName)
		assert.ElementsMatch(
			t,
			o.servicesHTTP[test.serviceName].GetInwayAddresses(),
			test.inwayAddresses)
	}
}
