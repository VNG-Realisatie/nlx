// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

const maxServiceCount = 250

func (h *DirectoryRegistrationService) RegisterInway(ctx context.Context, req *registrationapi.RegisterInwayRequest) (*registrationapi.RegisterInwayResponse, error) {
	logger := h.logger.With(zap.String("handler", "register-inway"))

	resp := &registrationapi.RegisterInwayResponse{}

	organizationName, err := h.getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization name from request: %v", err)
	}

	inwayParams := &database.RegisterInwayParams{
		OrganizationName:    organizationName,
		RequestInwayAddress: req.InwayAddress,
		NlxVersion:          nlxversion.NewFromGRPCContext(ctx).Version,
	}

	err = inwayParams.Validate()
	if err != nil {
		msg := fmt.Sprintf("validation failed: %s", err.Error())
		return nil, status.New(codes.InvalidArgument, msg).Err()
	}

	err = h.db.RegisterInway(inwayParams)
	if err != nil {
		h.logger.Error("database error while registering inway", zap.Error(err))
		return nil, status.New(codes.Internal, "failed to register inway").Err()
	}

	if len(req.Services) > maxServiceCount {
		return nil, status.New(codes.InvalidArgument, fmt.Sprintf("inway registers more services than allowed (max. %d)", maxServiceCount)).Err()
	}

	for _, service := range req.Services {
		service := service

		serviceSpecificationType := getAPISpecificationTypeForService(
			h.httpClient,
			h.logger,
			service.ApiSpecificationDocumentUrl,
			req.InwayAddress,
			service.Name,
		)

		serviceParams := &database.RegisterServiceParams{
			OrganizationName:     organizationName,
			Name:                 service.Name,
			Internal:             service.Internal,
			DocumentationURL:     service.DocumentationUrl,
			APISpecificationType: serviceSpecificationType,
			PublicSupportContact: service.PublicSupportContact,
			TechSupportContact:   service.TechSupportContact,
			OneTimeCosts:         service.OneTimeCosts,
			MonthlyCosts:         service.MonthlyCosts,
			RequestCosts:         service.RequestCosts,
		}

		err = serviceParams.Validate()
		if err != nil {
			msg := fmt.Sprintf("validation for service named '%s' failed: %s", serviceParams.Name, err.Error())
			return nil, status.New(codes.InvalidArgument, msg).Err()
		}

		err := h.db.RegisterService(serviceParams)
		if err != nil {
			logger.Error("failed to register service", zap.Error(err))
			return nil, status.New(codes.Internal, "database error").Err()
		}
	}

	return resp, nil
}

func getAPISpecificationTypeForService(httpClient *http.Client, logger *zap.Logger, specificationDocumentURL, inwayAddress, serviceName string) string {
	var result string

	if len(specificationDocumentURL) > 0 {
		specificationType, err := getAPISpecsTypeViaInway(httpClient, inwayAddress, serviceName)
		if err != nil {
			logger.Info(
				"invalid documentation specification document provided by inway",
				zap.String("documentation url", specificationType),
				zap.Error(err),
			)
		}

		result = specificationType
	}

	return result
}
