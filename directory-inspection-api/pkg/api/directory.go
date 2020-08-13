// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
)

// compile-time interface implementation verification
var _ inspectionapi.DirectoryInspectionServer = &InspectionService{}

// InspectionService handles all requests for a directory inspection api
type InspectionService struct {
	*listServicesHandler
	*listOrganizationsHandler
}

// New sets up a new DirectoryService and returns an error when something failed during set.
func New(
	logger *zap.Logger,
	db *sqlx.DB,
	demoEnv,
	demoDomain string,
) (*InspectionService, error) {
	s := &InspectionService{}

	var err error

	s.listServicesHandler, err = newListServicesHandler(db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ListServices handler")
	}
	s.listOrganizationsHandler, err = newListOrganizationsHandler(db, logger, demoEnv, demoDomain)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ListOrganizations handler")
	}

	return s, nil
}

func getOrganisationNameFromRequest(ctx context.Context) (string, error) {
	peerContext, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("failed to obtain peer from context")
	}
	tlsInfo := peerContext.AuthInfo.(credentials.TLSInfo)
	return tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0], nil
}
