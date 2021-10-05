// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inspectionservice

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

func registerOutwayVersion(ctx context.Context, db database.DirectoryDatabase, version nlxversion.Version) {
	_ = db.RegisterOutwayVersion(ctx, version)
}

func (h *InspectionService) ListServices(ctx context.Context, _ *emptypb.Empty) (*inspectionapi.ListServicesResponse, error) {
	h.logger.Info("rpc request ListServices()")

	// do not log requests coming from grpc-gateway
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if _, ok := md["grpcgateway-internal"]; !ok {
			version := nlxversion.NewFromGRPCContext(ctx)
			go registerOutwayVersion(ctx, h.db, version)
		}
	}

	organization, err := h.getOrganisationInformationFromRequest(ctx)
	if err != nil {
		h.logger.Error("determining organization info from request", zap.Error(err))
		return nil, status.Error(codes.Unknown, "determine organization info")
	}

	h.logger.Debug("querying services", zap.String("organizationSerialNumber", organization.SerialNumber), zap.String("organizationName", organization.Name))

	services, err := h.db.ListServices(ctx, organization.SerialNumber)
	if err != nil {
		h.logger.Error("failed to fetch services from db", zap.Error(err))
		return nil, status.Error(codes.Internal, "db error")
	}

	resp := &inspectionapi.ListServicesResponse{}

	for _, service := range services {
		resp.Services = append(resp.Services, convertFromDatabaseService(service))
	}

	return resp, nil
}

func convertFromDatabaseService(model *database.Service) *inspectionapi.ListServicesResponse_Service {
	service := &inspectionapi.ListServicesResponse_Service{
		Name:                 model.Name,
		Internal:             model.Internal,
		ApiSpecificationType: model.APISpecificationType,
		DocumentationUrl:     model.DocumentationURL,
		PublicSupportContact: model.PublicSupportContact,
	}

	if model.Organization != nil {
		service.Organization = &inspectionapi.Organization{
			SerialNumber: model.Organization.SerialNumber,
			Name:         model.Organization.Name,
		}
	}

	if model.Costs != nil {
		service.Costs = &inspectionapi.Costs{
			OneTime: int32(model.Costs.OneTime),
			Monthly: int32(model.Costs.Monthly),
			Request: int32(model.Costs.Request),
		}
	}

	for _, inway := range model.Inways {
		state := inspectionapi.Inway_State(inspectionapi.Inway_State_value[string(inway.State)])

		service.Inways = append(service.Inways, &inspectionapi.Inway{
			Address: inway.Address,
			State:   state,
		})
	}

	return service
}
