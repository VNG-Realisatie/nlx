// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"
	"fmt"

	"github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-registration-api/pkg/database"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

func (h *DirectoryRegistrationService) RegisterInway(ctx context.Context, req *registrationapi.RegisterInwayRequest) (*registrationapi.RegisterInwayResponse, error) {
	logger := h.logger.With(zap.String("handler", "register-inway"))

	logger.Info("rpc request RegisterInway", zap.String("inway address", req.InwayAddress))

	resp := &registrationapi.RegisterInwayResponse{}

	organizationName, err := h.getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization name from request: %v", err)
	}

	if !IsValidOrganizationName(organizationName) {
		logger.Info("invalid organization name", zap.String("organization name", organizationName))
		return nil, status.New(codes.InvalidArgument, "Invalid organization name").Err()
	}

	for _, service := range req.Services {
		service := service

		if !IsValidServiceName(service.Name) {
			logger.Info("invalid service name", zap.String("service name", service.Name))
			return nil, status.New(codes.InvalidArgument, "Invalid service name").Err()
		}

		// NOTE: we get the documentation spec doc via the inway, not directly. This field could probably be dropped form the communication to hte directory.
		logger.Info("service documentation url", zap.String("documentation url", service.ApiSpecificationDocumentUrl))

		var inwayAPISpecificationType string

		if len(service.ApiSpecificationDocumentUrl) > 0 {
			inwayAPISpecificationType, err = getInwayAPISpecsType(h.httpClient, req.InwayAddress, service.Name)
			if err != nil {
				logger.Info("invalid documentation specification document provided by inway", zap.String("documentation url", service.ApiSpecificationDocumentUrl), zap.Error(err))
				// DO NOT STOP WHEN  documentation fails.
				// return nil, status.New(codes.InvalidArgument, "Invalid documentation specification document provided").Err()
				inwayAPISpecificationType = ""
			}

			logger.Info("detected api spec", zap.String("apispectype", inwayAPISpecificationType))
		}

		params := &database.InsertAvailabilityParams{
			OrganizationName:            organizationName,
			ServiceName:                 service.Name,
			ServiceInternal:             service.Internal,
			ServiceDocumentationURL:     service.DocumentationUrl,
			InwayAPISpecificationType:   inwayAPISpecificationType,
			RequestInwayAddress:         req.InwayAddress,
			ServiceInsightAPIURL:        service.InsightApiUrl,
			ServiceIrmaAPIURL:           service.IrmaApiUrl,
			ServicePublicSupportContact: service.PublicSupportContact,
			ServiceTechSupportContact:   service.TechSupportContact,
			NlxVersion:                  nlxversion.NewFromGRPCContext(ctx).Version,
		}

		err := h.db.InsertAvailability(params)
		if err != nil {
			logger.Error("failed to execute stmtInsertAvailability", zap.Error(err))

			pqErr, ok := err.(*pq.Error)
			if ok {
				if pqErr.Constraint == "services_check_typespec" {
					invalidSpecificationTypeErrorMessage := fmt.Sprintf("invalid api-specification-type '%s' configured for service '%s'", service.ApiSpecificationType, service.Name)
					return nil, status.New(codes.InvalidArgument, invalidSpecificationTypeErrorMessage).Err()
				}
			}

			return nil, status.New(codes.Internal, "database error").Err()
		}
	}

	return resp, nil
}
