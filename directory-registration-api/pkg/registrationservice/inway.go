// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-registration-api/domain"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

const maxServiceCount = 250

func (h *DirectoryRegistrationService) RegisterInway(ctx context.Context, req *registrationapi.RegisterInwayRequest) (*registrationapi.RegisterInwayResponse, error) {
	logger := h.logger.With(zap.String("handler", "register-inway"))

	resp := &registrationapi.RegisterInwayResponse{}

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

	if inwayModel.IsOrganizationInway() {
		h.logger.Debug("is organization inway", zap.String("inway", inwayModel.ToString()))
		err = h.repository.SetOrganizationInway(ctx, organizationInformation.SerialNumber, inwayModel.Address())

		if err != nil {
			h.logger.Error("set organization inway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
			return nil, status.New(codes.Internal, "failed to set organization inway").Err()
		}
	} else {
		h.logger.Debug("is not organization inway", zap.String("inway", inwayModel.ToString()))
		err = h.repository.ClearIfSetAsOrganizationInway(ctx, organizationInformation.SerialNumber, inwayModel.Address())

		if err != nil {
			h.logger.Error("failed to execute ClearIfSetAsOrganizationInway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
			return nil, status.New(codes.Internal, "failed to update organization inway settings").Err()
		}
	}

	if len(req.Services) > maxServiceCount {
		return nil, status.New(codes.InvalidArgument, fmt.Sprintf("inway registers more services than allowed (max. %d)", maxServiceCount)).Err()
	}

	for _, s := range req.Services {
		s := s

		serviceSpecificationType := getAPISpecificationTypeForService(
			h.httpClient,
			h.logger,
			s.ApiSpecificationDocumentUrl,
			req.InwayAddress,
			s.Name,
		)

		serviceModel, err := domain.NewService(
			&domain.NewServiceArgs{
				Name:                     s.Name,
				OrganizationSerialNumber: organizationInformation.SerialNumber,
				Internal:                 s.Internal,
				DocumentationURL:         s.DocumentationUrl,
				APISpecificationType:     serviceSpecificationType,
				PublicSupportContact:     s.PublicSupportContact,
				TechSupportContact:       s.TechSupportContact,
				OneTimeCosts:             uint(s.OneTimeCosts),
				MonthlyCosts:             uint(s.MonthlyCosts),
				RequestCosts:             uint(s.RequestCosts),
			},
		)
		if err != nil {
			msg := fmt.Sprintf("validation for service named '%s' failed: %s", s.Name, err.Error())
			return nil, status.New(codes.InvalidArgument, msg).Err()
		}

		err = h.repository.RegisterService(serviceModel)
		if err != nil {
			logger.Error("failed to register service", zap.Error(err))
			return nil, status.New(codes.Internal, "database error").Err()
		}
	}

	return resp, nil
}

func getAPISpecificationTypeForService(httpClient *http.Client, logger *zap.Logger, specificationDocumentURL, inwayAddress, serviceName string) domain.SpecificationType {
	if len(specificationDocumentURL) < 1 {
		return ""
	}

	specificationType, err := getAPISpecsTypeViaInway(httpClient, inwayAddress, serviceName)
	if err != nil {
		logger.Error(
			"invalid documentation specification document provided by inway",
			zap.String("documentation url", specificationType),
			zap.Error(err),
		)

		return ""
	}

	return domain.SpecificationType(specificationType)
}
