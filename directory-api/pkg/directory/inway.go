// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

const maxServiceCount = 250

func (h *DirectoryService) RegisterInway(ctx context.Context, req *directoryapi.RegisterInwayRequest) (*directoryapi.RegisterInwayResponse, error) {
	resp := &directoryapi.RegisterInwayResponse{}

	organizationInformation, err := h.getOrganizationInformationFromRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization information from request: %v", err)
	}

	nlxVersion := nlxversion.NewFromGRPCContext(ctx).Version

	// Created at and Updated at time are the same times when registering new inway
	now := time.Now()

	organizationModel, err := domain.NewOrganization(organizationInformation.Name, organizationInformation.SerialNumber)
	if err != nil {
		msg := fmt.Sprintf("validation failed: %s", err.Error())
		return nil, status.New(codes.InvalidArgument, msg).Err()
	}

	inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
		Name:                req.InwayName,
		Organization:        organizationModel,
		IsOrganizationInway: req.IsOrganizationInway,
		Address:             req.InwayAddress,
		NlxVersion:          nlxVersion,
		CreatedAt:           now,
		UpdatedAt:           now,
	},
	)
	if err != nil {
		msg := fmt.Sprintf("validation failed: %s", err.Error())
		return nil, status.New(codes.InvalidArgument, msg).Err()
	}

	err = h.repository.RegisterInway(inwayModel)
	if err != nil {
		h.logger.Error("register inway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
		return nil, status.New(codes.Internal, "failed to register inway").Err()
	}

	err = h.setOrganizationInway(ctx, organizationInformation, inwayModel)
	if err != nil {
		h.logger.Error("failed to configure organization inway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
		return nil, status.New(codes.Internal, "failed to set organization inway ").Err()
	}

	if len(req.Services) > maxServiceCount {
		return nil, status.New(codes.InvalidArgument, fmt.Sprintf("inway registers more services than allowed (max. %d)", maxServiceCount)).Err()
	}

	for _, s := range req.Services {
		err = h.registerService(organizationInformation, req.InwayAddress, s)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (h *DirectoryService) setOrganizationInway(ctx context.Context, organizationInformation *tls.OrganizationInformation, inwayModel *domain.Inway) error {
	if inwayModel.IsOrganizationInway() {
		h.logger.Debug("is organization inway", zap.String("inway", inwayModel.ToString()))
		err := h.repository.SetOrganizationInway(ctx, organizationInformation.SerialNumber, inwayModel.Address())

		if err != nil {
			h.logger.Error("set organization inway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
			return status.New(codes.Internal, "failed to set organization inway").Err()
		}
	} else {
		h.logger.Debug("is not organization inway", zap.String("inway", inwayModel.ToString()))
		err := h.repository.ClearIfSetAsOrganizationInway(ctx, organizationInformation.SerialNumber, inwayModel.Address())

		if err != nil {
			h.logger.Error("failed to execute ClearIfSetAsOrganizationInway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
			return status.New(codes.Internal, "failed to update organization inway settings").Err()
		}
	}

	return nil
}

func (h *DirectoryService) registerService(organizationInformation *tls.OrganizationInformation, inwayAddress string, service *directoryapi.RegisterInwayRequest_RegisterService) error {
	serviceSpecificationType := getAPISpecificationTypeForService(
		h.httpClient,
		h.logger,
		service.ApiSpecificationDocumentUrl,
		inwayAddress,
		service.Name,
	)

	organization, err := domain.NewOrganization(organizationInformation.Name, organizationInformation.SerialNumber)
	if err != nil {
		msg := fmt.Sprintf("validation for organization with serial number '%s' failed: %s", organizationInformation.SerialNumber, err.Error())
		return status.New(codes.InvalidArgument, msg).Err()
	}

	serviceModel, err := domain.NewService(
		&domain.NewServiceArgs{
			Name:                 service.Name,
			Organization:         organization,
			Internal:             service.Internal,
			DocumentationURL:     service.DocumentationUrl,
			APISpecificationType: serviceSpecificationType,
			PublicSupportContact: service.PublicSupportContact,
			TechSupportContact:   service.TechSupportContact,
			Costs: &domain.NewServiceCostsArgs{
				OneTime: uint(service.OneTimeCosts),
				Monthly: uint(service.MonthlyCosts),
				Request: uint(service.RequestCosts),
			},
		},
	)
	if err != nil {
		msg := fmt.Sprintf("validation for service named '%s' failed: %s", service.Name, err.Error())
		return status.New(codes.InvalidArgument, msg).Err()
	}

	err = h.repository.RegisterService(serviceModel)
	if err != nil {
		h.logger.Error("failed to register service", zap.Error(err))
		return status.New(codes.Internal, "database error").Err()
	}

	return nil
}

func getAPISpecificationTypeForService(httpClient *http.Client, logger *zap.Logger, apiSpecificationDocumentURL, inwayAddress, serviceName string) domain.SpecificationType {
	if len(apiSpecificationDocumentURL) < 1 {
		return ""
	}

	specificationType, err := getAPISpecsTypeViaInway(httpClient, inwayAddress, serviceName)
	if err != nil {
		logger.Error(
			"invalid api specification document provided by inway",
			zap.String("api specification document url", apiSpecificationDocumentURL),
			zap.Error(err),
		)

		return ""
	}

	return domain.SpecificationType(specificationType)
}
