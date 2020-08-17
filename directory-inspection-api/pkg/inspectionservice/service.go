// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inspectionservice

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

func registerOutwayVersion(ctx context.Context, db database.DirectoryDatabase, version nlxversion.NlxVersion) {
	_ = db.RegisterOutwayVersion(ctx, version)
}

func (h *InspectionService) ListServices(ctx context.Context, req *inspectionapi.ListServicesRequest) (*inspectionapi.ListServicesResponse, error) {
	h.logger.Info("rpc request ListServices()")

	// do not log requests coming from grpc-gateway
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if _, ok := md["grpcgateway-internal"]; !ok {
			version := nlxversion.GetNlxVersionFromContext(ctx)
			go registerOutwayVersion(ctx, h.db, version)
		}
	}

	resp := &inspectionapi.ListServicesResponse{}
	organizationName, err := h.getOrganisationNameFromRequest(ctx)

	if err != nil {
		return nil, err
	}

	h.logger.Debug("querying services", zap.String("organizationName", organizationName))

	services, err := h.db.ListServices(ctx, organizationName)
	if err != nil {
		h.logger.Error("failed to fetch services from db", zap.Error(err))
		return nil, status.Error(codes.Internal, "db error")
	}

	for _, service := range services {
		resp.Services = append(resp.Services, convertFromDatabaseService(service))
	}

	return resp, nil
}

func convertFromDatabaseService(model *database.Service) *inspectionapi.ListServicesResponse_Service {
	service := &inspectionapi.ListServicesResponse_Service{
		OrganizationName:     model.Organization,
		ServiceName:          model.Name,
		Internal:             model.Internal,
		InwayAddresses:       model.InwayAddresses,
		ApiSpecificationType: model.APISpecificationType,
		HealthyStates:        model.HealthyStates,
		PublicSupportContact: model.PublicSupportContact,
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
