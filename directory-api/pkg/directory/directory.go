// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	storage "go.nlx.io/nlx/directory-api/domain/directory/storage"
)

// compile-time interface implementation verification
var _ directoryapi.DirectoryServer = &DirectoryService{}

type OrganizationInformationExtractor func(ctx context.Context) (*tls.OrganizationInformation, error)

type Clock interface {
	Now() time.Time
}

type DirectoryService struct {
	directoryapi.UnimplementedDirectoryServer
	directoryapi.UnimplementedDirectoryRegistrationServer
	directoryapi.UnimplementedDirectoryInspectionServer
	logger                                *zap.Logger
	repository                            storage.Repository
	httpClient                            *http.Client
	termsOfServiceURL                     string
	getOrganizationInformationFromRequest OrganizationInformationExtractor
	clock                                 Clock
	version                               string
}

type NewDirectoryArgs struct {
	Logger                                *zap.Logger
	TermsOfServiceURL                     string
	Repository                            storage.Repository
	HTTPClient                            *http.Client
	GetOrganizationInformationFromRequest OrganizationInformationExtractor
	Clock                                 Clock
	Version                               string
}

func New(args *NewDirectoryArgs) *DirectoryService {
	version := args.Version
	if args.Version == "" {
		version = "unknown"
	}

	s := &DirectoryService{
		logger:                                args.Logger,
		repository:                            args.Repository,
		httpClient:                            args.HTTPClient,
		termsOfServiceURL:                     args.TermsOfServiceURL,
		getOrganizationInformationFromRequest: args.GetOrganizationInformationFromRequest,
		clock:                                 args.Clock,
		version:                               version,
	}

	return s
}
