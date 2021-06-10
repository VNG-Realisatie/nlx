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

const maxServiceCount = 250

func (h *DirectoryRegistrationService) RegisterInway(ctx context.Context, req *registrationapi.RegisterInwayRequest) (*registrationapi.RegisterInwayResponse, error) {
	logger := h.logger.With(zap.String("handler", "register-inway"))

	logger.Info("rpc request RegisterInway", zap.String("inway address", req.InwayAddress))

	resp := &registrationapi.RegisterInwayResponse{}

	organizationName, err := h.getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization name from request: %v", err)
	}

	if len(req.Services) > maxServiceCount {
		return nil, status.New(codes.ResourceExhausted, "inway registers more services than allowed").Err()
	}

	for _, service := range req.Services {
		service := service

		// NOTE: we get the documentation spec doc via the inway, not directly. This field could probably be dropped form the communication to the directory.
		logger.Info("service documentation url", zap.String("documentation url", service.ApiSpecificationDocumentUrl))

		var serviceSpecificationType string

		if len(service.ApiSpecificationDocumentUrl) > 0 {
			serviceSpecificationType, err = getAPISpecsTypeViaInway(h.httpClient, req.InwayAddress, service.Name)
			if err != nil {
				logger.Info("invalid documentation specification document provided by inway", zap.String("documentation url", service.ApiSpecificationDocumentUrl), zap.Error(err))
				// DO NOT STOP WHEN  documentation fails.
				// return nil, status.New(codes.InvalidArgument, "Invalid documentation specification document provided").Err()
				serviceSpecificationType = ""
			}

			logger.Info("detected api spec", zap.String("apispectype", serviceSpecificationType))
		}

		params := &database.InsertAvailabilityParams{
			OrganizationName:            organizationName,
			ServiceName:                 service.Name,
			ServiceInternal:             service.Internal,
			ServiceDocumentationURL:     service.DocumentationUrl,
			InwayAPISpecificationType:   serviceSpecificationType,
			RequestInwayAddress:         req.InwayAddress,
			ServicePublicSupportContact: service.PublicSupportContact,
			ServiceTechSupportContact:   service.TechSupportContact,
			NlxVersion:                  nlxversion.NewFromGRPCContext(ctx).Version,
			OneTimeCosts:                service.OneTimeCosts,
			MonthlyCosts:                service.MonthlyCosts,
			RequestCosts:                service.RequestCosts,
		}

		if err := params.Validate(); err != nil {
			return nil, status.New(codes.InvalidArgument, fmt.Sprintf("validation failed: %s", err.Error())).Err()
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
