// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inspectionservice

import (
	"context"

	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

// compile-time interface implementation verification
var _ inspectionapi.DirectoryInspectionServer = &InspectionService{}

// InspectionService handles all requests for a directory inspection api
type InspectionService struct {
	logger   *zap.Logger
	database database.DirectoryDatabase
}

// New sets up a new DirectoryService
func New(logger *zap.Logger, database database.DirectoryDatabase) *InspectionService {
	s := &InspectionService{
		logger:   logger,
		database: database,
	}

	return s
}

// TODO: best way to enable using peerContext in tests?
func getOrganisationNameFromRequest(ctx context.Context) (string, error) {
	return "TODO", nil
	//peerContext, ok := peer.FromContext(ctx)
	//if !ok {
	//	return "", errors.New("failed to obtain peer from context")
	//}
	//tlsInfo := peerContext.AuthInfo.(credentials.TLSInfo)
	//return tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0], nil
}
