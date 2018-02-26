package directoryservice

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

// compile-time interface implementation verification
var _ directoryapi.DirectoryServer = &DirectoryService{}

// DirectoryService handles all requests for a directory api
type DirectoryService struct {
	*registerInwayHandler
	*listServicesHandler
}

// New sets up a new DirectoryService and returns an error when something failed during set.
func New(store *Store, logger *zap.Logger) (*DirectoryService, error) {
	s := &DirectoryService{}

	var err error

	s.registerInwayHandler, err = newRegisterInwayHandler(store, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup RegisterInway handler")
	}
	s.listServicesHandler, err = newListServicesHandler(store, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ListServices handler")
	}

	return s, nil
}
