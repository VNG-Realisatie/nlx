// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/inway/plugins"
	"go.nlx.io/nlx/management-api/api"
)

const configRetrievalInterval = 10 * time.Second

var errManagementAPIUnavailable = fmt.Errorf("managementAPI unavailable")

func (i *Inway) registerInwayToManagementAPI(ctx context.Context) error {
	hostname, err := os.Hostname()
	if err != nil {
		i.logger.Warn("failed to get inway hostname", zap.Error(err))
	}

	_, err = i.managementClient.RegisterInway(ctx, &api.Inway{
		Name:        i.name,
		Version:     version.BuildVersion,
		Hostname:    hostname,
		SelfAddress: i.address,
	})

	if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
		return errManagementAPIUnavailable
	}

	return err
}

func (i *Inway) startConfigurationPolling(ctx context.Context) error {
	if err := i.registerInwayToManagementAPI(ctx); err != nil {
		return err
	}

	settings, err := i.managementClient.GetSettings(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("error fetching settings: %s", err)
	}

	i.isOrganizationInway = settings.OrganizationInway == i.name
	i.logger.Debug("fetched settings from management client", zap.Any("organization inway", settings.OrganizationInway), zap.Bool("i.isOrganizationInway", i.isOrganizationInway))

	services, err := i.getServicesFromManagementAPI()
	if err != nil && err != errManagementAPIUnavailable {
		return err
	}

	err = i.SetServiceEndpoints(services)
	if err != nil {
		return err
	}

	go func() {
		expBackOff := &backoff.Backoff{
			Min:    100 * time.Millisecond,
			Factor: 2,
			Max:    20 * time.Second,
		}
		sleepDuration := configRetrievalInterval

		for {
			select {
			case <-ctx.Done():
				i.logger.Info("stopping config polling")
				return
			case <-time.After(sleepDuration):
				sleepDuration = i.registerInwayAndUpdateConfig(expBackOff, configRetrievalInterval)
			}
		}
	}()

	return nil
}

// SetServiceEndpoints configures the inway with the provided endpoints
func (i *Inway) SetServiceEndpoints(endpoints []*plugins.Service) error {
	i.logger.Info("setting service endpoints")

	i.servicesLock.Lock()
	defer i.servicesLock.Unlock()

	i.monitoringService.SetNotReady()
	i.services = make(map[string]*plugins.Service)

	for _, endPoint := range endpoints {
		if _, exists := i.services[endPoint.Name]; exists {
			return errors.New("service endpoint for a service with the same name has already been registered")
		}

		i.logger.Info("adding service to inway", zap.String("servicename", endPoint.Name))
		i.services[endPoint.Name] = endPoint
	}

	i.monitoringService.SetReady()

	return nil
}

func (i *Inway) registerInwayAndUpdateConfig(expBackOff *backoff.Backoff, defaultInterval time.Duration) time.Duration {
	i.logger.Info("registering inway to management api")
	err := i.registerInwayToManagementAPI(context.TODO())

	if err != nil {
		if errors.Is(err, errManagementAPIUnavailable) {
			i.logger.Warn("management api not available, waiting for management-api...", zap.Error(err))
			return expBackOff.Duration()
		}

		i.logger.Error("error from registerInwayToManagementAPI", zap.Error(err))

		return defaultInterval
	}

	i.logger.Info("retrieving config from the management-api", zap.String("inwayname", i.name))

	services, err := i.getServicesFromManagementAPI()
	if err != nil {
		if errors.Is(err, errManagementAPIUnavailable) {
			i.logger.Warn("management api not available, waiting for management-api...", zap.Error(err))
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

	settings, err := i.managementClient.GetSettings(context.Background(), &emptypb.Empty{})
	if err != nil {
		i.logger.Error("error fetching settings from managementClient", zap.Error(err))
		return defaultInterval
	}

	i.isOrganizationInway = settings.OrganizationInway == i.name
	i.logger.Debug("fetched settings from management client", zap.Any("organization inway", settings.OrganizationInway), zap.Bool("i.isOrganizationInway", i.isOrganizationInway))

	i.logger.Info("retrieved config successfully")
	expBackOff.Reset()

	return defaultInterval
}

func (i *Inway) isServiceConfigDifferent(services []*plugins.Service) bool {
	i.servicesLock.Lock()
	defer i.servicesLock.Unlock()

	serviceCount := len(services)

	if serviceCount != len(i.services) {
		return true
	}

	matches := 0

	for _, inwayService := range i.services {
		for _, service := range services {
			if reflect.DeepEqual(service, inwayService) {
				matches++
			}
		}
	}

	return matches != serviceCount
}

func serviceToPluginService(service *api.ListServicesResponse_Service) *plugins.Service {
	pluginService := &plugins.Service{
		Name:                        service.Name,
		APISpecificationDocumentURL: service.ApiSpecificationURL,
		DocumentationURL:            service.DocumentationURL,
		EndpointURL:                 service.EndpointURL,
		PublicSupportContact:        service.PublicSupportContact,
		TechSupportContact:          service.TechSupportContact,
		Internal:                    service.Internal,
		OneTimeCosts:                service.OneTimeCosts,
		MonthlyCosts:                service.MonthlyCosts,
		RequestCosts:                service.RequestCosts,
		Grants:                      []*plugins.Grant{},
	}

	for _, auth := range service.AuthorizationSettings.Authorizations {
		pluginService.Grants = append(pluginService.Grants, &plugins.Grant{
			OrganizationSerialNumber: auth.Organization.SerialNumber,
			PublicKeyPEM:             auth.PublicKeyPEM,
			PublicKeyFingerprint:     auth.PublicKeyHash,
		})
	}

	return pluginService
}

func (i *Inway) getServicesFromManagementAPI() ([]*plugins.Service, error) {
	response, err := i.managementClient.ListServices(context.Background(), &api.ListServicesRequest{
		InwayName: i.name,
	})

	if err != nil {
		if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
			return nil, errManagementAPIUnavailable
		}

		return nil, err
	}

	services := make([]*plugins.Service, len(response.Services))
	for i, service := range response.Services {
		services[i] = serviceToPluginService(service)
	}

	return services, nil
}
