// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
package inway

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jpillora/backoff"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/config"
	"go.nlx.io/nlx/management-api/api"
	api_mock "go.nlx.io/nlx/management-api/api/mock"
)

func createInway() (*Inway, error) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	return NewInway(logger, nil, testProcess, "mock.inway", "localhost:1812", "localhost:1813", cert, "localhost:1815")
}

func createMockService() *api.Service {
	return &api.Service{
		Name:                 "my-service",
		ApiSpecificationURL:  "https://api.spec.com",
		DocumentationURL:     "https://documentation.com",
		EndpointURL:          "https://endpointurl.com",
		PublicSupportContact: "publicsupport@email.com",
		TechSupportContact:   "techsupport@email.com",
		Internal:             true,
		AuthorizationSettings: &api.Service_AuthorizationSettings{
			Mode: "whitelist",
			Authorizations: []*api.Service_AuthorizationSettings_Authorization{
				{OrganizationName: "demo-org1"},
				{OrganizationName: "demo-org2"},
				{PublicKeyHash: "sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
			},
		},
	}
}
func TestServiceToServiceDetails(t *testing.T) {
	service := createMockService()
	serviceDetails := serviceConfigToServiceDetails(service)
	assert.Equal(t, service.ApiSpecificationURL, serviceDetails.APISpecificationDocumentURL)
	assert.Equal(t, service.DocumentationURL, serviceDetails.DocumentationURL)
	assert.Equal(t, service.EndpointURL, serviceDetails.EndpointURL)
	assert.Equal(t, service.PublicSupportContact, serviceDetails.PublicSupportContact)
	assert.Equal(t, service.TechSupportContact, serviceDetails.TechSupportContact)
	assert.Equal(t, service.Internal, serviceDetails.Internal)
	assert.Equal(t, service.AuthorizationSettings.Mode, string(serviceDetails.AuthorizationModel))
	assert.Equal(t, service.AuthorizationSettings.Authorizations[0].OrganizationName, serviceDetails.AuthorizationWhitelist[0].OrganizationName)
	assert.Equal(t, service.AuthorizationSettings.Authorizations[1].OrganizationName, serviceDetails.AuthorizationWhitelist[1].OrganizationName)
	assert.Equal(t, service.AuthorizationSettings.Authorizations[2].PublicKeyHash, serviceDetails.AuthorizationWhitelist[2].PublicKeyHash)
}

func TestManagementAPIResponseToEndpoints(t *testing.T) {
	serviceConfig := createMockService()
	response := &api.ListServicesResponse{
		Services: []*api.Service{
			serviceConfig,
		},
	}

	iw, err := createInway()
	assert.Nil(t, err)
	endpoints := iw.createServiceEndpoints(response)
	assert.Len(t, endpoints, 1)
	endpoint := endpoints[0]
	assert.Equal(t, serviceConfig.Name, endpoint.ServiceName())
	serviceDetails := endpoint.ServiceDetails()
	assert.NotNil(t, serviceDetails)
	assert.Equal(t, serviceConfig.ApiSpecificationURL, serviceDetails.APISpecificationDocumentURL)
	assert.Equal(t, serviceConfig.DocumentationURL, serviceDetails.DocumentationURL)
	assert.Equal(t, serviceConfig.EndpointURL, serviceDetails.EndpointURL)
	assert.Equal(t, serviceConfig.Internal, serviceDetails.Internal)
	assert.Equal(t, serviceConfig.PublicSupportContact, serviceDetails.PublicSupportContact)
	assert.Equal(t, serviceConfig.TechSupportContact, serviceDetails.TechSupportContact)
	assert.Equal(t, serviceConfig.AuthorizationSettings.Mode, string(serviceDetails.AuthorizationModel))
	assert.Equal(t, serviceConfig.AuthorizationSettings.Authorizations[0].OrganizationName, serviceDetails.AuthorizationWhitelist[0].OrganizationName)
	assert.Equal(t, serviceConfig.AuthorizationSettings.Authorizations[1].OrganizationName, serviceDetails.AuthorizationWhitelist[1].OrganizationName)
	assert.Equal(t, serviceConfig.AuthorizationSettings.Authorizations[2].PublicKeyHash, serviceDetails.AuthorizationWhitelist[2].PublicKeyHash)

	serviceConfig.AuthorizationSettings = nil
	endpoints = iw.createServiceEndpoints(response)
	assert.Len(t, endpoints, 1)
	endpoint = endpoints[0]
	serviceDetails = endpoint.ServiceDetails()

	assert.NotNil(t, serviceDetails)

	// check if the default value of authorization mode is "whitelist"
	assert.Equal(t, config.AuthorizationmodelWhitelist, serviceDetails.AuthorizationModel)
}

func TestSetupManagementAPI(t *testing.T) {
	iw, err := createInway()
	assert.Nil(t, err)

	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	err = iw.SetupManagementAPI("https://managementapi.mock", cert)
	assert.Nil(t, err)
	assert.NotNil(t, iw.managementClient)
}

func TestGetServicesFromManagementAPI(t *testing.T) {
	iw, err := createInway()
	assert.Nil(t, err)
	service := createMockService()
	mockResponse := &api.ListServicesResponse{
		Services: []*api.Service{
			service,
		},
	}

	expectedRequest := &api.ListServicesRequest{
		InwayName: "mock.inway",
	}
	controller := gomock.NewController(t)

	managementClient := api_mock.NewMockManagementClient(controller)
	managementClient.EXPECT().ListServices(context.Background(), expectedRequest).Return(mockResponse, nil)

	iw.managementClient = managementClient
	services, err := iw.getServicesFromManagementAPI()
	assert.Nil(t, err)
	assert.NotNil(t, services)
	assert.Len(t, services, 1)
	assert.Equal(t, service.Name, services[0].ServiceName())

	managementClient.EXPECT().ListServices(context.Background(), expectedRequest).Return(nil, fmt.Errorf("error"))

	services, err = iw.getServicesFromManagementAPI()
	assert.NotNil(t, err)
	assert.Nil(t, services)
	assert.Equal(t, err.Error(), "error")

	managementClient.EXPECT().ListServices(context.Background(), expectedRequest).
		Return(nil, status.New(codes.Unavailable, "unavailable").Err())

	services, err = iw.getServicesFromManagementAPI()
	assert.NotNil(t, err)
	assert.Nil(t, services)
	assert.Equal(t, err, errManagementAPIUnavailable)
}

func TestStartConfigurationPolling(t *testing.T) {
	iw, err := createInway()
	assert.Nil(t, err)

	hostname, err := os.Hostname()
	assert.Nil(t, err)

	controller := gomock.NewController(t)
	managementClient := api_mock.NewMockManagementClient(controller)
	managementClient.EXPECT().CreateInway(context.Background(), &api.Inway{
		Name:        "mock.inway",
		Version:     "unknown",
		Hostname:    hostname,
		SelfAddress: "localhost:1812",
	}).Return(nil, fmt.Errorf("error create inway"))

	managementClient.EXPECT().ListServices(context.Background(), &api.ListServicesRequest{
		InwayName: iw.name,
	}).Return(nil, fmt.Errorf("error list services"))

	iw.managementClient = managementClient

	err = iw.StartConfigurationPolling()
	assert.NotNil(t, err)
	assert.Equal(t, "error create inway", err.Error())

	managementClient.EXPECT().CreateInway(context.Background(), &api.Inway{
		Name:        "mock.inway",
		Version:     "unknown",
		Hostname:    hostname,
		SelfAddress: "localhost:1812",
	}).Return(&api.Inway{
		Name:        "mock.inway",
		Version:     "unknown",
		Hostname:    hostname,
		SelfAddress: "localhost:1812",
	}, nil)

	err = iw.StartConfigurationPolling()
	assert.NotNil(t, err)
	assert.Equal(t, "error list services", err.Error())
}

func TestUpdateConfig(t *testing.T) {
	expBackOff := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Factor: 2,
		Max:    20 * time.Second,
	}

	sleepDuration := 10 * time.Second
	service := createMockService()
	mockResponse := &api.ListServicesResponse{
		Services: []*api.Service{
			service,
		},
	}

	expectedRequest := &api.ListServicesRequest{
		InwayName: "mock.inway",
	}

	controller := gomock.NewController(t)
	managementClient := api_mock.NewMockManagementClient(controller)
	managementClient.EXPECT().ListServices(context.Background(), expectedRequest).Return(mockResponse, nil)

	iw, err := createInway()
	assert.Nil(t, err)

	iw.managementClient = managementClient

	newSleepDuration := iw.updateConfig(expBackOff, sleepDuration)
	assert.Equal(t, sleepDuration, newSleepDuration)

	managementClient.EXPECT().ListServices(context.Background(), expectedRequest).
		Return(nil, status.New(codes.Unavailable, "unavailable").Err())

	newSleepDuration = iw.updateConfig(expBackOff, sleepDuration)
	assert.Equal(t, expBackOff.Min, newSleepDuration)

	managementClient.EXPECT().ListServices(context.Background(), expectedRequest).
		Return(nil, fmt.Errorf("error"))

	newSleepDuration = iw.updateConfig(expBackOff, sleepDuration)
	assert.Equal(t, sleepDuration, newSleepDuration)
}

func TestCreateServiceEndpoints(t *testing.T) {
	iw, err := createInway()
	assert.Nil(t, err)
	mockResponse := &api.ListServicesResponse{
		Services: []*api.Service{
			createMockService(),
		},
	}
	services := iw.createServiceEndpoints(mockResponse)
	err = iw.SetServiceEndpoints(services)
	assert.Nil(t, err)

	isDifferent := iw.isServiceConfigDifferent(services)
	assert.False(t, isDifferent)

	mockResponse.Services[0].ApiSpecificationURL = "differenturl"

	services = iw.createServiceEndpoints(mockResponse)
	isDifferent = iw.isServiceConfigDifferent(services)
	assert.True(t, isDifferent)
}

func TestDeleteServiceEndpoints(t *testing.T) {
	iw, err := createInway()
	assert.Nil(t, err)

	endpointA, _ := iw.NewHTTPServiceEndpoint("service-a", &config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL: "https://api-a.test",
		},
	}, common_tls.NewConfig())

	endpointB, _ := iw.NewHTTPServiceEndpoint("service-b", &config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			EndpointURL: "https://api-b.test",
		},
	}, common_tls.NewConfig())

	initEndpoints := []ServiceEndpoint{endpointA, endpointB}

	err = iw.SetServiceEndpoints(initEndpoints)
	assert.Nil(t, err)

	mockResponse := &api.ListServicesResponse{
		Services: []*api.Service{
			{
				Name:                "service-a",
				ApiSpecificationURL: "https://api-a.test",
			},
		},
	}

	controller := gomock.NewController(t)
	managementClient := api_mock.NewMockManagementClient(controller)
	managementClient.EXPECT().CreateInway(gomock.Any(), gomock.Any()).Return(nil, nil)
	managementClient.EXPECT().ListServices(gomock.Any(), gomock.Any()).Return(mockResponse, nil)

	iw.managementClient = managementClient

	assert.Len(t, initEndpoints, 2)

	err = iw.StartConfigurationPolling()
	assert.Nil(t, err)

	assert.Len(t, iw.serviceEndpoints, 1)
}

func TestNewInwayName(t *testing.T) {
	logger := zap.NewNop()
	cert, _ := common_tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)

	testProcess := process.NewProcess(logger)
	iw, err := NewInway(logger, nil, testProcess, "", "inway.test", "localhost:1813", cert, "")
	assert.Nil(t, err)

	assert.Equal(t, "XQpL-03EUOCXDNnc8FCsZXrOp41LkYIJ5U_Udz-1Chk=", iw.name)

	iw, err = NewInway(logger, nil, testProcess, "inway.test", "inway.test", "localhost:1813", cert, "")
	assert.Nil(t, err)
	assert.Equal(t, "inway.test", iw.name)
}
