// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"net/http"
	"regexp"

	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

// compile-time interface implementation verification
var _ registrationapi.DirectoryRegistrationServer = &DirectoryRegistrationService{}

// InspectionService handles all requests for a directory inspection api
type DirectoryRegistrationService struct {
	logger                         *zap.Logger
	db                             database.DirectoryDatabase
	httpClient                     *http.Client
	getOrganisationNameFromRequest func(ctx context.Context) (string, error)
}

var regExpOrgName = regexp.MustCompile(`^[a-zA-Z0-9-\.\s]{1,100}$`)

// New sets up a new DirectoryRegistrationService
func New(logger *zap.Logger, db database.DirectoryDatabase, httpClient *http.Client, getOrganisationNameFromRequest func(ctx context.Context) (string, error)) *DirectoryRegistrationService {
	s := &DirectoryRegistrationService{
		logger:                         logger,
		db:                             db,
		httpClient:                     httpClient,
		getOrganisationNameFromRequest: getOrganisationNameFromRequest,
	}

	return s
}

func isValidOrganizationName(name string) bool {
	return regExpOrgName.MatchString(name)
}

func isValidServiceName(name string) bool {
	return regExpOrgName.MatchString(name)
}
