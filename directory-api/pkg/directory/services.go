// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/common/nlxversion"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
	"go.nlx.io/nlx/directory-api/domain/directory/storage"
)

func registerOutwayVersion(ctx context.Context, logger *zap.Logger, db storage.Repository, version nlxversion.Version) {
	err := db.RegisterOutwayVersion(ctx, version)
	if err != nil {
		logger.Error("failed to register outway version", zap.Error(err))
	}
}

func (h *DirectoryService) ListServices(ctx context.Context, _ *emptypb.Empty) (*directoryapi.ListServicesResponse, error) {
	h.logger.Info("rpc request ListServices()")

	// do not log requests coming from grpc-gateway
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if _, ok := md["grpcgateway-internal"]; !ok {
			version := nlxversion.NewFromGRPCContext(ctx)
			go registerOutwayVersion(ctx, h.logger, h.repository, version)
		}
	}

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

		serviceInways := model.Inways()

		service.Inways = make([]*directoryapi.Inway, len(serviceInways))
		for i, inway := range serviceInways {
			service.Inways[i] = &directoryapi.Inway{
				Address: inway.Address(),
				State:   directoryapi.Inway_State(directoryapi.Inway_State_value[string(inway.State())]),
			}
		}

		response.Services[i] = service
	}

	return response
}
