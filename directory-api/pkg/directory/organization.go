// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package directory

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
	storage "go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func (h *DirectoryService) ClearOrganizationInway(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	logger := h.logger.With(zap.String("handler", "ClearOrganizationInway"))

	organization, err := h.getOrganizationInformationFromRequest(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "determine organization")
	}

	logger.Debug("clearing organization inway", zap.Any("organization", organization))

	err = h.repository.ClearOrganizationInway(ctx, organization.SerialNumber)
	if err != nil {
		if errors.Is(err, storage.ErrOrganizationNotFound) {
			logger.Info("did not clear organization because the organization could not be found", zap.Any("organization", organization))
			return &emptypb.Empty{}, nil
		}

		logger.Error("failed to clear the organization inway", zap.Error(err))

		return nil, status.New(codes.Internal, "database error").Err()
	}

	return &emptypb.Empty{}, nil
}

func (h *DirectoryService) ListOrganizations(ctx context.Context, _ *emptypb.Empty) (*directoryapi.ListOrganizationsResponse, error) {
	h.logger.Info("rpc request ListOrganizations")

	organizations, err := h.repository.ListOrganizations(ctx)
	if err != nil {
		h.logger.Error("failed to select organizations from db", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	return convertFromDatabaseOrganization(organizations), nil
}

func convertFromDatabaseOrganization(model []*domain.Organization) *directoryapi.ListOrganizationsResponse {
	result := &directoryapi.ListOrganizationsResponse{
		Organizations: make([]*directoryapi.Organization, len(model)),
	}

	for i, organization := range model {
		result.Organizations[i] = &directoryapi.Organization{
			Name:         organization.Name(),
			SerialNumber: organization.SerialNumber(),
		}
	}

	return result
}
