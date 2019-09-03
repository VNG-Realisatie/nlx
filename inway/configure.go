// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"context"
	"crypto/tls"
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

	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/inway/config"
)

var errConfigAPIUnavailable = fmt.Errorf("configAPI unavailable")

// SetConfigAPIAddress configures the inway to use the config API instead of the config toml
func (i *Inway) SetConfigAPIAddress(configAPIAddress string) error {
	orgKeypair, err := tls.LoadX509KeyPair(i.orgCertFile, i.orgKeyFile)
	if err != nil {
		return err
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{orgKeypair},
		RootCAs:      i.roots,
	})

	connCtx, connCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer connCtxCancel()

	i.logger.Info("creating config api connection", zap.String("config api address", configAPIAddress))
	configAPIConn, err := grpc.DialContext(connCtx, configAPIAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}

	i.configAPIClient = configapi.NewConfigApiClient(configAPIConn)

	return nil
}

// StartConfigurationPolling will make the inway retrieve its configuration periodically
func (i *Inway) StartConfigurationPolling() error {
	_, err := i.configAPIClient.CreateInway(context.Background(), &configapi.Inway{
		Name: i.name,
	})
	if err != nil {
		return err
	}

	services, err := i.getServicesFromConfigAPI()
	if err != nil && err != errConfigAPIUnavailable {
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

	return nil
}

func (i *Inway) updateConfig(expBackOff *backoff.Backoff, defaultInterval time.Duration) time.Duration {
	i.logger.Info("retrieving config from the config-api", zap.String("inwayname", i.name))
	services, err := i.getServicesFromConfigAPI()
	if err != nil {
		if err == errConfigAPIUnavailable {
			i.logger.Info("waiting for config-api...", zap.Error(err))
			return expBackOff.Duration()
		}
		i.logger.Error("failed to contact the config-api", zap.Error(err))

		return defaultInterval
	}

	if i.isServiceConfigDifferent(services) {
		i.logger.Info("detected changes in inway service config. updating services")
		err = i.SetServiceEndpoints(services)
		if err != nil {
			i.logger.Error("unable to configure the inway with the config-api response", zap.Error(err))
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
	matches := 0
	for _, inwayService := range i.serviceEndpoints {
		for _, service := range services {

			if reflect.DeepEqual(inwayService.ServiceDetails(), service.ServiceDetails()) && strings.Compare(service.ServiceName(), inwayService.ServiceName()) == 0 {
				matches++
			}
		}
	}

	return matches != len(services)
}

func serviceConfigToServiceDetails(service *configapi.Service) *config.ServiceDetails {
	serviceDetails := &config.ServiceDetails{
		APISpecificationDocumentURL: service.ApiSpecificationURL,
		DocumentationURL:            service.DocumentationURL,
		EndpointURL:                 service.EndpointURL,
		PublicSupportContact:        service.PublicSupportContact,
		TechSupportContact:          service.TechSupportContact,
		Internal:                    service.Internal,
	}

	if service.AuthorizationSettings != nil {
		serviceDetails.AuthorizationModel = config.AuthorizationModel(service.AuthorizationSettings.Mode)
		serviceDetails.AuthorizationWhitelist = service.AuthorizationSettings.Organizations
	} else {
		serviceDetails.AuthorizationModel = config.AuthorizationmodelNone
	}

	return serviceDetails
}

func (i *Inway) createServiceEndpoints(response *configapi.ListServicesResponse) []ServiceEndpoint {
	endPoints := make([]ServiceEndpoint, len(response.Services))
	c := 0
	for _, service := range response.Services {
		endpoint, err := i.NewHTTPServiceEndpoint(service.Name, serviceConfigToServiceDetails(service), &tls.Config{})
		if err != nil {
			i.logger.Error("cannot create HTTPServiceEndpoint from service configuration", zap.Error(err))
			continue
		}
		endPoints[c] = endpoint
		c++
	}

	return endPoints
}

func (i *Inway) getServicesFromConfigAPI() ([]ServiceEndpoint, error) {
	response, err := i.configAPIClient.ListServices(context.Background(), &configapi.ListServicesRequest{
		InwayName: i.name,
	})

	if err != nil {
		if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
			return nil, errConfigAPIUnavailable
		}

		return nil, err
	}

	return i.createServiceEndpoints(response), nil
}
