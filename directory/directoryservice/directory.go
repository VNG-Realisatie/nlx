package directoryservice

import (
	"github.com/jmoiron/sqlx"
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
func New(logger *zap.Logger, db *sqlx.DB) (*DirectoryService, error) {
	s := &DirectoryService{}

	var err error

	s.registerInwayHandler, err = newRegisterInwayHandler(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup RegisterInway handler")
	}
	s.listServicesHandler, err = newListServicesHandler(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ListServices handler")
	}

	return s, nil
}
