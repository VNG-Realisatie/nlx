// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inspectionservice

import (
	"context"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

// compile-time interface implementation verification
var _ inspectionapi.DirectoryInspectionServer = &InspectionService{}

// InspectionService handles all requests for a directory inspection api
type InspectionService struct {
	inspectionapi.UnimplementedDirectoryInspectionServer
	logger                                *zap.Logger
	db                                    database.DirectoryDatabase
	getOrganisationInformationFromRequest func(ctx context.Context) (*tls.OrganizationInformation, error)
}

// New sets up a new DirectoryService
func New(logger *zap.Logger, db database.DirectoryDatabase, getOrganisationInformationFromRequest func(ctx context.Context) (*tls.OrganizationInformation, error)) *InspectionService {
	s := &InspectionService{
		logger:                                logger,
		db:                                    db,
		getOrganisationInformationFromRequest: getOrganisationInformationFromRequest,
	}

	return s
}
