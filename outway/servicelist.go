package outway

import (
	"context"
	"reflect"
	"time"

	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
)

func (o *Outway) keepServiceListUpToDate(process *process.Process) {
	o.wg.Add(1)
	defer o.wg.Done()

	expBackOff := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Factor: 2,
		Max:    20 * time.Second,
	}

	const baseInterval = 30 * time.Second
	interval := baseInterval
	for {
		select {
		case <-process.ShutdownComplete:
			return
		case <-time.After(interval):
			err := o.updateServiceList(process)
			if err != nil {
				o.logger.Warn("failed to update the service list", zap.Error(err))
				interval = expBackOff.Duration() // Changing interval on retry
			} else {
				interval = baseInterval // Resetting interval to base on success update
				expBackOff.Reset()
			}
		}
	}
}

func (o *Outway) updateServiceList(process *process.Process) error {
	services := make(map[string]HTTPService)
	resp, err := o.directoryInspectionClient.ListServices(context.Background(), &inspectionapi.ListServicesRequest{})
	if err != nil {
		return errors.Wrap(err, "failed to fetch services from directory")
	}
	o.servicesLock.Lock()
	defer o.servicesLock.Unlock()
	shutDown := make(chan struct{})
	process.CloseGracefully(func() error {
		close(shutDown)
		return nil
	})
	for _, serviceToImplement := range resp.Services {
		select {
		case <-shutDown:
			// On app shutdown we have no need to update services.
			// So we need to wait until started updated is finished and exit
			return nil
		default:
			// Need default to not to block
		}
		o.logger.Debug("directory listed service", zap.String("service-name", serviceToImplement.ServiceName), zap.String("service-organization-name", serviceToImplement.OrganizationName))

		service, exists := o.services[serviceToImplement.OrganizationName+"."+serviceToImplement.ServiceName]
		if !exists || !reflect.DeepEqual(service.GetInwayAddresses(), serviceToImplement.InwayAddresses) {
			// create the service
			rrlbService, err := NewRoundRobinLoadBalancedHTTPService(o.logger, o.tlsRoots, o.tlsOptions.OrgCertFile, o.tlsOptions.OrgKeyFile, serviceToImplement.OrganizationName, serviceToImplement.ServiceName, serviceToImplement.InwayAddresses)
			if err != nil {
				if err == errNoInwaysAvailable {
					o.logger.Debug("service exists but there are no inwayaddresses available", zap.String("service-organization-name", serviceToImplement.OrganizationName), zap.String("service-name", serviceToImplement.ServiceName))
					continue
				}
				o.logger.Error("failed to create new service", zap.String("service-organization-name", serviceToImplement.OrganizationName), zap.String("service-name", serviceToImplement.ServiceName), zap.Error(err))
				continue
			}
			service = rrlbService
			o.logger.Debug("implemented service", zap.String("service-name", serviceToImplement.ServiceName), zap.String("service-organization-name", serviceToImplement.OrganizationName))
		}
		services[service.FullName()] = service
	}

	o.services = services
	return nil
}

func (o *Outway) getService(organization, service string) HTTPService {
	o.servicesLock.RLock()
	httpService := o.services[organization+"."+service]
	o.servicesLock.RUnlock()
	return httpService
}
