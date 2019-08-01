// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi/mock"
)

func TestUpdateServiceList(t *testing.T) {
	// Create mock directory client
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock.NewMockDirectoryInspectionClient(ctrl)

	logger := zap.NewNop()
	mainProcess := process.NewProcess(logger)

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(workDir)

	o := &Outway{
		directoryInspectionClient: client,
		logger:                    logger,
		tlsOptions: orgtls.TLSOptions{
			NLXRootCert: filepath.Join(parent, "testing", "root.crt"),
			OrgCertFile: filepath.Join(parent, "testing", "org-nlx-test.crt"),
			OrgKeyFile:  filepath.Join(parent, "testing", "org-nlx-test.key"),
		},
		process: mainProcess,
	}

	// Make the mock directory client return an error when calling ListServices
	client.EXPECT().ListServices(
		context.Background(),
		&inspectionapi.ListServicesRequest{}).Return(nil, fmt.Errorf("mock error"))

	// Test of updateServiceList generates the correct error
	err = o.updateServiceList()
	assert.EqualError(t, err, "failed to fetch services from directory: mock error")

	mockServiceAInwayAddresses := []string{"mock-service-a-1:123", "mock-service-a-2:123"}
	mockServiceBInwayAddresses := []string{"mock-service-b-1:123", "mock-service-b-2:123"}

	mockServiceAFullName := "mock-org-a.mock-service-a"
	mockServiceBFullName := "mock-org-b.mock-service-b"
	// Make the mock directory client provide a list of services when calling ListServices
	client.EXPECT().ListServices(context.Background(), &inspectionapi.ListServicesRequest{}).Return(
		&inspectionapi.ListServicesResponse{
			Services: []*inspectionapi.ListServicesResponse_Service{
				{ServiceName: "mock-service-a", OrganizationName: "mock-org-a", InwayAddresses: mockServiceAInwayAddresses},
				{ServiceName: "mock-service-b", OrganizationName: "mock-org-b", InwayAddresses: mockServiceBInwayAddresses},
				{ServiceName: "mock-service-c", OrganizationName: "mock-org-c"},
			},
		}, nil)

	// Test of updateServiceList creates a correct o.services map
	err = o.updateServiceList()
	assert.Nil(t, err)

	// mock-service-c should not be included because this service does not have any inwayaddresses
	assert.Len(t, o.services, 2, fmt.Sprintf("%v", o.services))

	tests := []struct {
		serviceName    string
		inwayAddresses []string
	}{
		{
			mockServiceAFullName,
			[]string{"mock-service-a-1:123", "mock-service-a-2:123"},
		},
		{
			mockServiceBFullName,
			[]string{"mock-service-b-1:123", "mock-service-b-2:123"},
		},
	}

	for _, test := range tests {
		assert.Contains(t, o.services, test.serviceName)
		assert.ElementsMatch(
			t,
			o.services[test.serviceName].GetInwayAddresses(),
			test.inwayAddresses)
	}
}
