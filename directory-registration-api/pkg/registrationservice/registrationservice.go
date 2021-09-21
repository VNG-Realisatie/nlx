// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-registration-api/domain/directory"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

// compile-time interface implementation verification
var _ registrationapi.DirectoryRegistrationServer = &DirectoryRegistrationService{}

type DirectoryRegistrationService struct {
	registrationapi.UnimplementedDirectoryRegistrationServer
	logger                                *zap.Logger
	repository                            directory.Repository
	httpClient                            *http.Client
	getOrganizationInformationFromRequest func(ctx context.Context) (*tls.OrganizationInformation, error)
}

func New(logger *zap.Logger, repository directory.Repository, httpClient *http.Client, getOrganisationInformationFromRequest func(ctx context.Context) (*tls.OrganizationInformation, error)) *DirectoryRegistrationService {
	s := &DirectoryRegistrationService{
		logger:                                logger,
		repository:                            repository,
		httpClient:                            httpClient,
		getOrganizationInformationFromRequest: getOrganisationInformationFromRequest,
	}

	return s
}
