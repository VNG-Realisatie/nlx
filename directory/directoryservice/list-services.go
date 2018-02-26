package directoryservice

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

type listServicesHandler struct {
	store *Store
}

func newListServicesHandler(store *Store, logger *zap.Logger) (*listServicesHandler, error) {
	return &listServicesHandler{
		store: store,
	}, nil
}

func (p *listServicesHandler) ListServices(ctx context.Context, req *directoryapi.ListServicesRequest) (*directoryapi.ListServicesResponse, error) {
	fmt.Println("rpc request ListServices()")
	repl := &directoryapi.ListServicesResponse{}

	p.store.ServicesLock.RLock()
	defer p.store.ServicesLock.RUnlock()

	for serviceName, service := range p.store.Services {
		s := &directoryapi.Service{
			Name:             serviceName,
			OrganizationName: service.OrganizationName,
		}
		for inwayAddress, healthy := range service.InwayAddresses {
			if !healthy {
				continue // skip unhealthy instances in the service list
			}
			s.InwayAddresses = append(s.InwayAddresses, inwayAddress)
		}
		repl.Services = append(repl.Services, s)
	}

	return repl, nil
}
