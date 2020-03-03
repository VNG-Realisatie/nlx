// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
package inway

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jpillora/backoff"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/config-api/configapi"
	configmock "go.nlx.io/nlx/config-api/configapi/mock"
	"go.nlx.io/nlx/inway/config"
)

func createInway() (*Inway, error) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: "../testing/pki/ca-root.pem",
		OrgCertFile: "../testing/pki/org-nlx-test-chain.pem",
		OrgKeyFile:  "../testing/pki/org-nlx-test-key.pem",
	}

	return NewInway(logger, nil, testProcess, "mock.inway", "localhost:1812", tlsOptions, "localhost:1815")
}

func createMockService() *configapi.Service {
	return &configapi.Service{
		Name:                 "my-service",
		ApiSpecificationURL:  "https://api.spec.com",
		DocumentationURL:     "https://documentation.com",
		EndpointURL:          "https://endpointurl.com",
		PublicSupportContact: "publicsupport@email.com",
		TechSupportContact:   "techsupport@email.com",
		Internal:             true,
		AuthorizationSettings: &configapi.Service_AuthorizationSettings{
			Mode:          "whitelist",
			Organizations: []string{"demo-org1", "demo-org2"},
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
	assert.Equal(t, service.AuthorizationSettings.Organizations, serviceDetails.AuthorizationWhitelist)
}

func TestConfigApiResponseToEndpoints(t *testing.T) {
	serviceConfig := createMockService()
	response := &configapi.ListServicesResponse{
		Services: []*configapi.Service{
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
	assert.Equal(t, serviceConfig.AuthorizationSettings.Organizations, serviceDetails.AuthorizationWhitelist)

	serviceConfig.AuthorizationSettings = nil
	endpoints = iw.createServiceEndpoints(response)
	assert.Len(t, endpoints, 1)
	endpoint = endpoints[0]
	serviceDetails = endpoint.ServiceDetails()

	assert.NotNil(t, serviceDetails)

	// check if the default value of authorization mode is "whitelist"
	assert.Equal(t, config.AuthorizationmodelWhitelist, serviceDetails.AuthorizationModel)
}

func TestSetConfigAPIAddress(t *testing.T) {
	iw, err := createInway()
	assert.Nil(t, err)
	err = iw.SetConfigAPIAddress("https://configapi.mock")
	assert.Nil(t, err)
	assert.NotNil(t, iw.configAPIClient)
}

func TestGetServicesFromConfigAPI(t *testing.T) {
	i, err := createInway()
	assert.Nil(t, err)
	service := createMockService()
	mockResponse := &configapi.ListServicesResponse{
		Services: []*configapi.Service{
			service,
		},
	}

	expectedRequest := &configapi.ListServicesRequest{
		InwayName: "mock.inway",
	}
	controller := gomock.NewController(t)
	configAPIMockClient := configmock.NewMockConfigApiClient(controller)
	configAPIMockClient.EXPECT().ListServices(context.Background(), expectedRequest).Return(mockResponse, nil)

	i.configAPIClient = configAPIMockClient
	services, err := i.getServicesFromConfigAPI()
	assert.Nil(t, err)
	assert.NotNil(t, services)
	assert.Len(t, services, 1)
	assert.Equal(t, service.Name, services[0].ServiceName())

	configAPIMockClient.EXPECT().ListServices(context.Background(), expectedRequest).Return(nil, fmt.Errorf("error"))

	services, err = i.getServicesFromConfigAPI()
	assert.NotNil(t, err)
	assert.Nil(t, services)
	assert.Equal(t, err.Error(), "error")

	configAPIMockClient.EXPECT().ListServices(context.Background(), expectedRequest).
		Return(nil, status.New(codes.Unavailable, "unavailable").Err())

	services, err = i.getServicesFromConfigAPI()
	assert.NotNil(t, err)
	assert.Nil(t, services)
	assert.Equal(t, err, errConfigAPIUnavailable)
}

func TestStartConfigurationPolling(t *testing.T) {
	iw, err := createInway()
	assert.Nil(t, err)

	controller := gomock.NewController(t)
	configAPIMockClient := configmock.NewMockConfigApiClient(controller)
	configAPIMockClient.EXPECT().CreateInway(context.Background(), &configapi.Inway{
		Name: iw.name,
	}).Return(nil, fmt.Errorf("error create inway"))
	configAPIMockClient.EXPECT().ListServices(context.Background(), &configapi.ListServicesRequest{
		InwayName: iw.name,
	}).Return(nil, fmt.Errorf("error list services"))
	iw.configAPIClient = configAPIMockClient

	err = iw.StartConfigurationPolling()
	assert.NotNil(t, err)
	assert.Equal(t, "error create inway", err.Error())

	configAPIMockClient.EXPECT().CreateInway(context.Background(), &configapi.Inway{
		Name: iw.name,
	}).Return(&configapi.Inway{
		Name: iw.name,
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
	mockResponse := &configapi.ListServicesResponse{
		Services: []*configapi.Service{
			service,
		},
	}

	expectedRequest := &configapi.ListServicesRequest{
		InwayName: "mock.inway",
	}

	controller := gomock.NewController(t)
	configAPIMockClient := configmock.NewMockConfigApiClient(controller)
	configAPIMockClient.EXPECT().ListServices(context.Background(), expectedRequest).Return(mockResponse, nil)

	iw, err := createInway()
	assert.Nil(t, err)

	iw.configAPIClient = configAPIMockClient

	newSleepDuration := iw.updateConfig(expBackOff, sleepDuration)
	assert.Equal(t, sleepDuration, newSleepDuration)

	configAPIMockClient.EXPECT().ListServices(context.Background(), expectedRequest).
		Return(nil, status.New(codes.Unavailable, "unavailable").Err())

	newSleepDuration = iw.updateConfig(expBackOff, sleepDuration)
	assert.Equal(t, expBackOff.Min, newSleepDuration)

	configAPIMockClient.EXPECT().ListServices(context.Background(), expectedRequest).
		Return(nil, fmt.Errorf("error"))

	newSleepDuration = iw.updateConfig(expBackOff, sleepDuration)
	assert.Equal(t, sleepDuration, newSleepDuration)
}

func TestCreateServiceEndpoints(t *testing.T) {
	iw, err := createInway()
	assert.Nil(t, err)
	mockResponse := &configapi.ListServicesResponse{
		Services: []*configapi.Service{
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

func TestNewInwayName(t *testing.T) {
	logger := zap.NewNop()
	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "pki", "ca-root.pem"),
		OrgCertFile: filepath.Join("..", "testing", "pki", "org-nlx-test-chain.pem"),
		OrgKeyFile:  filepath.Join("..", "testing", "pki", "org-nlx-test-key.pem"),
	}

	testProcess := process.NewProcess(logger)
	iw, err := NewInway(logger, nil, testProcess, "", "", tlsOptions, "")
	assert.Nil(t, err)

	assert.Equal(t, "XQpL-03EUOCXDNnc8FCsZXrOp41LkYIJ5U_Udz-1Chk=", iw.name)

	iw, err = NewInway(logger, nil, testProcess, "inway.test", "", tlsOptions, "")
	assert.Nil(t, err)
	assert.Equal(t, "inway.test", iw.name)
}
