package directoryservice

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

type getServicesHandler struct{}

func newGetServicesHandler(logger *zap.Logger) (*getServicesHandler, error) {
	return &getServicesHandler{}, nil
}

func (p *getServicesHandler) GetServices(ctx context.Context, req *directoryapi.GetServicesRequest) (*directoryapi.GetServicesResponse, error) {
	fmt.Println("rpc request GetServices()")
	repl := &directoryapi.GetServicesResponse{}

	store.ServicesLock.RLock()
	defer store.ServicesLock.RUnlock()

	for serviceName, service := range store.Services {
		s := &directoryapi.Service{
			Name:             serviceName,
			OrganizationName: service.OrganizationName,
		}
		for inwayAddress := range service.InwayAddresses {
			s.InwayAddresses = append(s.InwayAddresses, inwayAddress)
		}
		repl.Services = append(repl.Services, s)
	}

	return repl, nil
}
