// Copyright © VNG Realisatie 2021
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
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

const (
	ErrMessageDatabase                                 = "database error"
	ErrMessageOrganizationInformationMissing           = "tls certificate does not contain organization information"
	ErrMessageOrganizationInwayNotFound                = "organization not found or the organization has no inway configured"
	ErrMessageOrganizationSerialNumberMissing          = "organization serial number is empty"
	ErrMessageUnableToComputeManagementAPIProxyAddress = "unable to compute organization management API proxy address"
)

func (h *DirectoryService) ClearOrganizationInway(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	logger := h.logger.With(zap.String("handler", "ClearOrganizationInway"))

	organization, err := h.getOrganizationInformationFromRequest(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, ErrMessageOrganizationInformationMissing)
	}

	logger.Debug("clearing organization inway", zap.Any("organization", organization))

	err = h.repository.ClearOrganizationInway(ctx, organization.SerialNumber)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			logger.Info("did not clear organization because the organization could not be found", zap.Any("organization", organization))
			return &emptypb.Empty{}, nil
		}

		logger.Error("failed to clear the organization inway", zap.Error(err))

		return nil, status.New(codes.Internal, ErrMessageDatabase).Err()
	}

	return &emptypb.Empty{}, nil
}

func (h *DirectoryService) GetOrganizationManagementAPIProxyAddress(ctx context.Context, req *directoryapi.GetOrganizationManagementAPIProxyAddressRequest) (*directoryapi.GetOrganizationManagementAPIProxyAddressResponse, error) {
	h.logger.Info("rpc request GetOrganizationManagementAPIProxyAddress")

	serialNumber := req.OrganizationSerialNumber
	if serialNumber == "" {
		return nil, status.New(codes.InvalidArgument, "organization serial number is empty").Err()
	}

	address, err := h.repository.GetOrganizationInwayManagementAPIProxyAddress(ctx, serialNumber)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.New(codes.NotFound, ErrMessageOrganizationInwayNotFound).Err()
		}

		h.logger.Error("failed to select organization inway address from storage", zap.Error(err))

		return nil, status.New(codes.Internal, ErrMessageDatabase).Err()
	}

	resp := &directoryapi.GetOrganizationManagementAPIProxyAddressResponse{
		Address: address,
	}

	return resp, nil
}

// GetOrganizationInway is implemented to ensure that NLX instances running version <= v0.127.0 will still work.
func (h *DirectoryService) GetOrganizationInway(ctx context.Context, req *directoryapi.GetOrganizationInwayRequest) (*directoryapi.GetOrganizationInwayResponse, error) {
	h.logger.Info("rpc request GetOrganizationInway")

	serialNumber := req.OrganizationSerialNumber
	if serialNumber == "" {
		return nil, status.New(codes.InvalidArgument, ErrMessageOrganizationSerialNumberMissing).Err()
	}

	address, err := h.repository.GetOrganizationInwayAddress(ctx, serialNumber)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.New(codes.NotFound, ErrMessageOrganizationInwayNotFound).Err()
		}

		h.logger.Error("failed to select organization inway address from storage", zap.Error(err))

		return nil, status.New(codes.Internal, ErrMessageDatabase).Err()
	}

	return &directoryapi.GetOrganizationInwayResponse{
		Address: address,
	}, nil
}

func (h *DirectoryService) ListOrganizations(ctx context.Context, _ *emptypb.Empty) (*directoryapi.ListOrganizationsResponse, error) {
	h.logger.Info("rpc request ListOrganizations")

	organizations, err := h.repository.ListOrganizations(ctx)
	if err != nil {
		h.logger.Error("failed to select organizations from db", zap.Error(err))
		return nil, status.New(codes.Internal, ErrMessageDatabase).Err()
	}

	return convertFromDatabaseOrganization(organizations), nil
}

func (h *DirectoryService) SetOrganizationContactDetails(ctx context.Context, req *directoryapi.SetOrganizationContactDetailsRequest) (*emptypb.Empty, error) {
	h.logger.Info("rpc request SetOrganizationContactDetailsRequest")

	organization, err := h.getOrganizationFromRequest(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, ErrMessageOrganizationInformationMissing)
	}

	err = req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = h.repository.SetOrganizationEmailAddress(ctx, organization, req.EmailAddress)
	if err != nil {
		h.logger.Error("unable to set organization contact details in database", zap.Error(err))
		return nil, status.Error(codes.Internal, ErrMessageDatabase)
	}

	return &emptypb.Empty{}, nil
}

func (h *DirectoryService) getOrganizationFromRequest(ctx context.Context) (*domain.Organization, error) {
	org, err := h.getOrganizationInformationFromRequest(ctx)
	if err != nil {
		return nil, err
	}

	organization, err := domain.NewOrganization(org.Name, org.SerialNumber)
	if err != nil {
		return nil, err
	}

	return organization, nil
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
