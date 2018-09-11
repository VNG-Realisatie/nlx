package directoryservice

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory/directoryapi"
)

// compile-time interface implementation verification
var _ directoryapi.DirectoryServer = &DirectoryService{}

// DirectoryService handles all requests for a directory api
type DirectoryService struct {
	*registerInwayHandler
	*listServicesHandler
	*listOrganizationsHandler
	*getServiceAPISpecHandler
}

// New sets up a new DirectoryService and returns an error when something failed during set.
func New(logger *zap.Logger, db *sqlx.DB, rootCA *x509.CertPool, certKeyPair tls.Certificate, demoEnv string, demoDomain string) (*DirectoryService, error) {
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
	s.listOrganizationsHandler, err = newListOrganizationsHandler(db, logger, demoEnv, demoDomain)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ListOrganizations handler")
	}
	s.getServiceAPISpecHandler, err = newGetServiceAPISpecHandler(db, logger, rootCA, certKeyPair)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup GetServiceAPISpecHandler handler")
	}

	return s, nil
}
