// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-registration-api/domain/inway"
	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

// compile-time interface implementation verification
var _ registrationapi.DirectoryRegistrationServer = &DirectoryRegistrationService{}

type DirectoryRegistrationService struct {
	registrationapi.UnimplementedDirectoryRegistrationServer
	logger                         *zap.Logger
	db                             database.DirectoryDatabase
	inwayRepository                inway.Repository
	httpClient                     *http.Client
	getOrganisationNameFromRequest func(ctx context.Context) (string, error)
}

// New sets up a new DirectoryRegistrationService
func New(logger *zap.Logger, db database.DirectoryDatabase, inwayRepository inway.Repository, httpClient *http.Client, getOrganisationNameFromRequest func(ctx context.Context) (string, error)) *DirectoryRegistrationService {
	s := &DirectoryRegistrationService{
		logger:                         logger,
		db:                             db,
		inwayRepository:                inwayRepository,
		httpClient:                     httpClient,
		getOrganisationNameFromRequest: getOrganisationNameFromRequest,
	}

	return s
}
