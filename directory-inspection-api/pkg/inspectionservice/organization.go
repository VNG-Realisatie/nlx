// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inspectionservice

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/emptypb"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

func (h *InspectionService) ListOrganizations(ctx context.Context, _ *emptypb.Empty) (*inspectionapi.ListOrganizationsResponse, error) {
	h.logger.Info("rpc request ListOrganizations")

	resp := &inspectionapi.ListOrganizationsResponse{}
	organizations, err := h.db.ListOrganizations(ctx)
	if err != nil {
		h.logger.Error("failed to select organizations from db", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	for _, organization := range organizations {
		resp.Organizations = append(resp.Organizations, convertFromDatabaseOrganization(organization))
	}

	return resp, nil
}

func (h *InspectionService) GetOrganizationInway(ctx context.Context, req *inspectionapi.GetOrganizationInwayRequest) (*inspectionapi.GetOrganizationInwayResponse, error) {
	h.logger.Info("rpc request GetOrganizationInwayAddress")

	name := req.OrganizationName
	if name == "" {
		return nil, status.New(codes.InvalidArgument, "organization name is empty").Err()
	}

	address, err := h.db.GetOrganizationInwayAddress(ctx, name)
	if err != nil {
		if errors.Is(err, database.ErrNoOrganization) {
			return nil, status.New(codes.NotFound, "organization has no inway").Err()
		}

		h.logger.Error("failed to select organization inway address from db", zap.Error(err))

		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	resp := &inspectionapi.GetOrganizationInwayResponse{
		Address: address,
	}

	return resp, nil
}

func convertFromDatabaseOrganization(model *database.Organization) *inspectionapi.ListOrganizationsResponse_Organization {
	organization := &inspectionapi.ListOrganizationsResponse_Organization{
		Name: model.Name,
	}

	return organization
}
