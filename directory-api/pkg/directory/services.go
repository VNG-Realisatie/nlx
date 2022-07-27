// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

func (h *DirectoryService) ListServices(ctx context.Context, _ *emptypb.Empty) (*directoryapi.ListServicesResponse, error) {
	h.logger.Info("rpc request ListServices()")

	organization, err := h.getOrganizationInformationFromRequest(ctx)
	if err != nil {
		h.logger.Error("determining organization info from request", zap.Error(err))
		return nil, status.Error(codes.Unknown, "determine organization info")
	}

	h.logger.Debug("querying services", zap.String("organizationSerialNumber", organization.SerialNumber), zap.String("organizationName", organization.Name))

	services, err := h.repository.ListServices(ctx, organization.SerialNumber)
	if err != nil {
		h.logger.Error("failed to fetch services from db", zap.Error(err))
		return nil, status.Error(codes.Internal, "Database error.")
	}

	return convertFromDatabaseService(services), nil
}

func convertFromDatabaseService(models []*domain.Service) *directoryapi.ListServicesResponse {
	response := &directoryapi.ListServicesResponse{
		Services: make([]*directoryapi.ListServicesResponse_Service, len(models)),
	}

	for i, model := range models {
		service := &directoryapi.ListServicesResponse_Service{
			Name:                 model.Name(),
			Internal:             model.Internal(),
			ApiSpecificationType: string(model.APISpecificationType()),
			DocumentationUrl:     model.DocumentationURL(),
			PublicSupportContact: model.PublicSupportContact(),
			Organization: &directoryapi.Organization{
				Name:         model.Organization().Name(),
				SerialNumber: model.Organization().SerialNumber(),
			},
			Costs: &directoryapi.ListServicesResponse_Costs{
				OneTime: int32(model.Costs().OneTime()),
				Monthly: int32(model.Costs().Monthly()),
				Request: int32(model.Costs().Request()),
			},
		}

		serviceInways := model.Availabilities()

		service.Inways = make([]*directoryapi.Inway, len(serviceInways))
		for i, inway := range serviceInways {
			service.Inways[i] = &directoryapi.Inway{
				Address: inway.InwayAddress(),
				State:   directoryapi.Inway_State(directoryapi.Inway_State_value[string(inway.State())]),
			}
		}

		response.Services[i] = service
	}

	return response
}
