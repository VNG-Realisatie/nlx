// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inspectionservice

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

func (h *InspectionService) ListOrganizations(ctx context.Context, req *inspectionapi.ListOrganizationsRequest) (*inspectionapi.ListOrganizationsResponse, error) {
	h.logger.Info("rpc request ListOrganizations")
	resp := &inspectionapi.ListOrganizationsResponse{}

	organizations, err := h.database.ListOrganizations(ctx)
	if err != nil {
		h.logger.Error("failed to select organizations from database", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	for _, organization := range organizations {
		resp.Organizations = append(resp.Organizations, convertFromDatabaseOrganization(organization))
	}

	return resp, nil
}

func convertFromDatabaseOrganization(model *database.Organization) *inspectionapi.ListOrganizationsResponse_Organization {
	organization := &inspectionapi.ListOrganizationsResponse_Organization{
		Name:                model.Name,
		InsightIrmaEndpoint: model.InsightIrmaEndpoint,
		InsightLogEndpoint:  model.InsightLogEndpoint,
	}

	return organization
}
