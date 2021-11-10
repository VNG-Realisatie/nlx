// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	storage "go.nlx.io/nlx/directory-api/domain/directory/storage"
)

// compile-time interface implementation verification
var _ directoryapi.DirectoryServer = &DirectoryService{}

type DirectoryService struct {
	directoryapi.UnimplementedDirectoryServer
	logger                                *zap.Logger
	repository                            storage.Repository
	httpClient                            *http.Client
	getOrganizationInformationFromRequest func(ctx context.Context) (*tls.OrganizationInformation, error)
}

func New(logger *zap.Logger, repository storage.Repository, httpClient *http.Client, getOrganisationInformationFromRequest func(ctx context.Context) (*tls.OrganizationInformation, error)) *DirectoryService {
	s := &DirectoryService{
		logger:                                logger,
		repository:                            repository,
		httpClient:                            httpClient,
		getOrganizationInformationFromRequest: getOrganisationInformationFromRequest,
	}

	return s
}
