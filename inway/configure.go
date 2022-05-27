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

	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/inway/plugins"
	"go.nlx.io/nlx/management-api/api"
)

const retryFactorConfigRetrieval = 2
const minRetryDurationConfigRetrieval = 100 * time.Millisecond
const maxRetryDurationConfigRetrieval = 20 * time.Second
const configRetrievalInterval = 10 * time.Second

var errManagementAPIUnavailable = fmt.Errorf("managementAPI unavailable")

func (i *Inway) startConfigurationPolling(ctx context.Context) {
	expBackOff := &backoff.Backoff{
		Min:    minRetryDurationConfigRetrieval,
		Factor: retryFactorConfigRetrieval,
		Max:    maxRetryDurationConfigRetrieval,
	}
	sleepDuration := 0 * time.Second

	for {
		select {
		case <-ctx.Done():
			i.logger.Info("stopping configuration polling")
			return
		case <-time.After(sleepDuration):
			err := i.registerToManagement(ctx)
			if err != nil {
				i.logger.Error("unable to register to local management api", zap.Error(err))

				if errors.Is(err, errManagementAPIUnavailable) {
					sleepDuration = expBackOff.Duration()
					continue
				}

				expBackOff.Reset()

				sleepDuration = configRetrievalInterval

				continue
			}

			err = i.retrieveAndUpdateConfig()
			if err != nil {
				i.logger.Error("unable to retrieve config", zap.Error(err))

				if errors.Is(err, errManagementAPIUnavailable) {
					sleepDuration = expBackOff.Duration()
					continue
				}
			}

			expBackOff.Reset()

			sleepDuration = configRetrievalInterval
		}
	}
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

	return nil
}

func (i *Inway) retrieveAndUpdateConfig() error {
	i.logger.Info("retrieving config from the management-api", zap.String("inwayname", i.name))

	response, err := i.managementClient.GetInwayConfig(context.Background(), &api.GetInwayConfigRequest{
		Name: i.name,
	})
	if err != nil {
		if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
			return errManagementAPIUnavailable
		}

		return err
	}

	services := servicesToPluginService(response.Services)

	if i.isServiceConfigDifferent(services) {
		i.logger.Info("detected changes in inway service config. updating services")

		err = i.SetServiceEndpoints(services)
		if err != nil {
			i.logger.Error("unable to configure the inway with the management-api response", zap.Error(err))
			return err
		}
	}

	i.isOrganizationInway = response.IsOrganizationInway
	i.logger.Debug("fetched settings from management client", zap.Bool("i.isOrganizationInway", i.isOrganizationInway))

	i.logger.Info("retrieved config successfully")

	i.monitoringService.SetReady()

	return err
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

func servicesToPluginService(services []*api.GetInwayConfigResponse_Service) []*plugins.Service {
	pluginServices := make([]*plugins.Service, len(services))
	for i, service := range services {
		pluginServices[i] = serviceToPluginService(service)
	}

	return pluginServices
}

func serviceToPluginService(service *api.GetInwayConfigResponse_Service) *plugins.Service {
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

func (i *Inway) registerToManagement(ctx context.Context) error {
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

	return err
}
