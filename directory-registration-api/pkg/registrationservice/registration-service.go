// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"crypto/tls"
	"crypto/x509"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

// compile-time interface implementation verification
var _ registrationapi.DirectoryRegistrationServer = &DirectoryRegistrationService{}

// InspectionService handles all requests for a directory inspection api
type DirectoryRegistrationService struct {
	*RegisterInwayHandler
	*SetInsightConfigurationHandler
}

var regExpOrgName = regexp.MustCompile(`^[a-zA-Z0-9-\.\s]{1,100}$`)

// New sets up a new DirectoryRegistrationService and returns an error when something failed during set.
func New(
	logger *zap.Logger,
	db *sqlx.DB, rootCA *x509.CertPool,
	certKeyPair *tls.Certificate,
) (*DirectoryRegistrationService, error) {
	s := &DirectoryRegistrationService{}
	var err error

	s.RegisterInwayHandler, err = newRegisterInwayHandler(db, logger, rootCA, certKeyPair)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup NewRegisterInwayHandler handler")
	}
	s.SetInsightConfigurationHandler, err = newSetInsightConfigurationHandler(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup NewSetInsightConfigurationHandler handler")
	}

	return s, nil
}

func validateName(name string) bool {
	return regExpOrgName.MatchString(name)
}
