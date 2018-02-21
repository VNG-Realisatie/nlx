package directoryservice

import (
	"context"
	"fmt"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
	"go.uber.org/zap"
)

type registerInwayHandler struct{}

func newRegisterInwayHandler(logger *zap.Logger) (*registerInwayHandler, error) {
	return &registerInwayHandler{}, nil
}

func (p *registerInwayHandler) RegisterInway(ctx context.Context, req *directoryapi.RegisterInwayRequest) (*directoryapi.RegisterInwayResponse, error) {
	fmt.Printf("rpc request RegisterInway(%s, %s)\n", req.OrganizationName, req.InwayAddress)
	repl := &directoryapi.RegisterInwayResponse{}

	store.ServicesLock.Lock()
	defer store.SaveAndUnlock()

	for _, serviceName := range req.ServiceNames {
		service, exists := store.Services[serviceName]
		if !exists {
			service = NewStoredService(req.OrganizationName, serviceName)
			store.Services[serviceName] = service
		}
		service.InwayAddresses[req.InwayAddress] = true
	}

	return repl, nil
}
