// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/inway/config"
	"go.nlx.io/nlx/inway/grpcproxy"
	"go.nlx.io/nlx/management-api/api"
	external_api "go.nlx.io/nlx/management-api/api/external"
)

var errManagementAPIUnavailable = fmt.Errorf("managementAPI unavailable")

// SetupManagementAPI configures the inway to use the NLX Management API instead of the config toml
func (i *Inway) SetupManagementAPI(managementAPIAddress string, cert *common_tls.CertificateBundle) error {
	creds := credentials.NewTLS(cert.TLSConfig())

	connCtx, connCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer connCtxCancel()

	i.logger.Info("creating management api connection", zap.String("management api address", managementAPIAddress))
	conn, err := grpc.DialContext(connCtx, managementAPIAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}

	i.managementClient = api.NewManagementClient(conn)

	p, err := grpcproxy.New(context.TODO(), i.logger, managementAPIAddress, i.orgCertBundle, cert)
	if err != nil {
		return err
	}

	p.RegisterService(external_api.GetAccessRequestServiceDesc())

	i.managementProxy = p

	return nil
}

// StartConfigurationPolling will make the inway retrieve its configuration periodically
func (i *Inway) StartConfigurationPolling() error {
	_, err := i.managementClient.CreateInway(context.Background(), &api.Inway{
		Name:        i.name,
		Version:     version.BuildVersion,
		Hostname:    i.hostname(),
		SelfAddress: i.selfAddress,
	})
	if err != nil {
		return err
	}

	services, err := i.getServicesFromManagementAPI()
	if err != nil && err != errManagementAPIUnavailable {
		return err
	}

	err = i.SetServiceEndpoints(services)
	if err != nil {
		return err
	}

	configRetrievalInterval := 1 * time.Minute

	go func() {
		expBackOff := &backoff.Backoff{
			Min:    100 * time.Millisecond,
			Factor: 2,
			Max:    20 * time.Second,
		}
		sleepDuration := configRetrievalInterval

		for {
			select {
			case <-i.stopInwayChannel:
				i.logger.Info("stopping config polling")
				return
			case <-time.After(sleepDuration):
				sleepDuration = i.updateConfig(expBackOff, configRetrievalInterval)
			}
		}
	}()

	return nil
}

// SetServiceEndpoints configures the inway with the provided endpoints
func (i *Inway) SetServiceEndpoints(endpoints []ServiceEndpoint) error {
	i.logger.Info("setting service endpoints")

	// stop the inway with current serviceEndpoints
	i.stop()

	i.serviceEndpointsLock.Lock()
	defer i.serviceEndpointsLock.Unlock()

	i.stopInwayChannel = make(chan struct{})
	i.serviceEndpoints = make(map[string]ServiceEndpoint)
	for _, endPoint := range endpoints {
		if _, exists := i.serviceEndpoints[endPoint.ServiceName()]; exists {
			return errors.New("service endpoint for a service with the same name has already been registered")
		}
		i.logger.Info("adding service to inway", zap.String("servicename", endPoint.ServiceName()))
		i.serviceEndpoints[endPoint.ServiceName()] = endPoint
		i.announceToDirectory(endPoint)
	}

	i.monitoringService.SetReady()

	return nil
}

func (i *Inway) updateConfig(expBackOff *backoff.Backoff, defaultInterval time.Duration) time.Duration {
	i.logger.Info("retrieving config from the management-api", zap.String("inwayname", i.name))
	services, err := i.getServicesFromManagementAPI()
	if err != nil {
		if err == errManagementAPIUnavailable {
			i.logger.Info("waiting for management-api...", zap.Error(err))
			return expBackOff.Duration()
		}

		i.logger.Error("failed to contact the management-api", zap.Error(err))

		return defaultInterval
	}

	if i.isServiceConfigDifferent(services) {
		i.logger.Info("detected changes in inway service config. updating services")
		err = i.SetServiceEndpoints(services)
		if err != nil {
			i.logger.Error("unable to configure the inway with the management-api response", zap.Error(err))
			return defaultInterval
		}
	}

	i.logger.Info("retrieved config successfully")
	expBackOff.Reset()

	return defaultInterval
}

func (i *Inway) isServiceConfigDifferent(services []ServiceEndpoint) bool {
	i.serviceEndpointsLock.Lock()
	defer i.serviceEndpointsLock.Unlock()

	serviceCount := len(services)

	if serviceCount != len(i.serviceEndpoints) {
		return true
	}

	matches := 0

	for _, inwayService := range i.serviceEndpoints {
		for _, service := range services {
			if reflect.DeepEqual(inwayService.ServiceDetails(), service.ServiceDetails()) && strings.Compare(service.ServiceName(), inwayService.ServiceName()) == 0 {
				matches++
			}
		}
	}

	return matches != serviceCount
}

func serviceConfigToServiceDetails(service *api.Service) *config.ServiceDetails {
	serviceDetails := &config.ServiceDetails{
		ServiceDetailsBase: config.ServiceDetailsBase{
			APISpecificationDocumentURL: service.ApiSpecificationURL,
			DocumentationURL:            service.DocumentationURL,
			EndpointURL:                 service.EndpointURL,
			PublicSupportContact:        service.PublicSupportContact,
			TechSupportContact:          service.TechSupportContact,
			Internal:                    service.Internal,
		},
	}

	if service.AuthorizationSettings != nil {
		serviceDetails.AuthorizationModel = config.AuthorizationModel(service.AuthorizationSettings.Mode)
		for _, authorization := range service.AuthorizationSettings.Authorizations {
			serviceDetails.AuthorizationWhitelist = append(serviceDetails.AuthorizationWhitelist, config.AuthorizationWhitelistItem{
				OrganizationName: authorization.OrganizationName,
				PublicKeyHash:    authorization.PublicKeyHash,
			})
		}
	} else {
		serviceDetails.AuthorizationModel = config.AuthorizationmodelWhitelist
	}

	return serviceDetails
}

func (i *Inway) createServiceEndpoints(response *api.ListServicesResponse) []ServiceEndpoint {
	endPoints := make([]ServiceEndpoint, len(response.Services))
	c := 0

	for _, service := range response.Services {
		endpoint, err := i.NewHTTPServiceEndpoint(service.Name, serviceConfigToServiceDetails(service), common_tls.NewConfig())
		if err != nil {
			i.logger.Error("cannot create HTTPServiceEndpoint from service configuration", zap.Error(err))
			continue
		}

		endPoints[c] = endpoint
		c++
	}

	return endPoints
}

func (i *Inway) getServicesFromManagementAPI() ([]ServiceEndpoint, error) {
	response, err := i.managementClient.ListServices(context.Background(), &api.ListServicesRequest{
		InwayName: i.name,
	})

	if err != nil {
		if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
			return nil, errManagementAPIUnavailable
		}

		return nil, err
	}

	return i.createServiceEndpoints(response), nil
}
