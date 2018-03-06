package directoryservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type registerInwayHandler struct {
	store *Store
}

func newRegisterInwayHandler(store *Store, logger *zap.Logger) (*registerInwayHandler, error) {
	return &registerInwayHandler{
		store: store,
	}, nil
}

func (p *registerInwayHandler) RegisterInway(ctx context.Context, req *directoryapi.RegisterInwayRequest) (*directoryapi.RegisterInwayResponse, error) {
	fmt.Printf("rpc request RegisterInway(%s)\n", req.InwayAddress)
	repl := &directoryapi.RegisterInwayResponse{}

	peer, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("failed to obtain peer from context")
	}
	tlsInfo := peer.AuthInfo.(credentials.TLSInfo)
	organizationName := tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0]
	// TODO: when administrative (client-tls mandatory) and inspection (client-tls optional) endpoints have been seperated,
	// use proper grpc authentication via middleware and context (based on client-tls fields (CN, O) like we do here)

	p.store.ServicesLock.Lock()
	defer p.store.SaveAndUnlock()

	for _, serviceName := range req.ServiceNames {
		service, exists := p.store.Services[serviceName]
		if !exists {
			service = NewStoredService(organizationName, serviceName)
			p.store.Services[serviceName] = service
		}
		service.DocumentationURL = req.DocumentationUrl
		service.InwayAddresses[req.InwayAddress] = true
	}

	return repl, nil
}
