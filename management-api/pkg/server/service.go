//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// CreateService creates a new service
func (s *ManagementService) CreateService(ctx context.Context, request *api.CreateServiceRequest) (*api.CreateServiceResponse, error) {
	logger := s.logger.With(zap.String("name", request.Name))
	logger.Info("rpc request Createrequest")

	err := request.Validate()
	if err != nil {
		logger.Error("invalid request", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid request: %s", err))
	}

	service := &database.Service{
		Name:                 request.Name,
		EndpointURL:          request.EndpointURL,
		DocumentationURL:     request.DocumentationURL,
		APISpecificationURL:  request.ApiSpecificationURL,
		Internal:             request.Internal,
		TechSupportContact:   request.TechSupportContact,
		PublicSupportContact: request.PublicSupportContact,
	}

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.ServiceCreate(ctx, userInfo.username, userInfo.userAgent, request.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	err = s.configDatabase.CreateServiceWithInways(ctx, service, request.Inways)
	if err != nil {
		logger.Error("error creating request in DB", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	return convertToCreateServiceResponseFromCreateServiceRequest(request), nil
}

// GetService returns a specific service
func (s *ManagementService) GetService(ctx context.Context, req *api.GetServiceRequest) (*api.GetServiceResponse, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request GetService")

	service, err := s.configDatabase.GetService(ctx, req.Name)
	if err != nil {
		if errIsNotFound(err) {
			logger.Warn("service not found")
			return nil, status.Error(codes.NotFound, "service not found")
		}

		logger.Error("error getting service from DB", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	response := convertToGetServiceResponseFromDatabaseService(service)

	return response, nil
}

// UpdateService updates an existing service
func (s *ManagementService) UpdateService(ctx context.Context, req *api.UpdateServiceRequest) (*api.UpdateServiceResponse, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request UpdateService")

	err := req.Validate()
	if err != nil {
		logger.Error("invalid service", zap.Error(err))

		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid service: %s", err))
	}

	service, err := s.configDatabase.GetService(ctx, req.Name)
	if err != nil {
		if errIsNotFound(err) {
			return nil, status.Error(codes.NotFound, "service not found")
		}

		logger.Error("failed to get service by name", zap.String("name", req.Name), zap.Error(err))

		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to get service: %s", err))
	}

	service.EndpointURL = req.EndpointURL
	service.DocumentationURL = req.DocumentationURL
	service.APISpecificationURL = req.ApiSpecificationURL
	service.Internal = req.Internal
	service.TechSupportContact = req.TechSupportContact
	service.PublicSupportContact = req.PublicSupportContact

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.ServiceUpdate(ctx, userInfo.username, userInfo.userAgent, service.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	err = s.configDatabase.UpdateServiceWithInways(ctx, service, req.Inways)
	if err != nil {
		if errIsNotFound(err) {
			logger.Warn("service not found")
			return nil, status.Error(codes.NotFound, "service not found")
		}

		logger.Error("error updating service in DB", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	return convertToUpdateServiceResponseFromUpdateServiceRequest(req), nil
}

// DeleteService deletes a specific service
func (s *ManagementService) DeleteService(ctx context.Context, req *api.DeleteServiceRequest) (*types.Empty, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteService")

	userInfo, err := retrieveUserInfoFromGRPCContext(ctx)
	if err != nil {
		logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.auditLogger.ServiceDelete(ctx, userInfo.username, userInfo.userAgent, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not create audit log")
	}

	err = s.configDatabase.DeleteService(ctx, req.Name)
	if err != nil {
		logger.Error("error deleting service in DB", zap.Error(err))
		return &types.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &types.Empty{}, nil
}

// ListServices returns a list of services
func (s *ManagementService) ListServices(ctx context.Context, req *api.ListServicesRequest) (*api.ListServicesResponse, error) {
	s.logger.Info("rpc request ListServices")

	var (
		err      error
		services []*database.Service
	)

	if req.InwayName == "" {
		services, err = s.configDatabase.ListServices(ctx)
		if err != nil {
			s.logger.Error("error getting services list from database", zap.Error(err))

			return nil, status.Error(codes.Internal, "database error")
		}
	} else {
		inway, err := s.configDatabase.GetInway(ctx, req.InwayName)
		if err != nil {
			if errIsNotFound(err) {
				return nil, status.Error(codes.NotFound, "inway not found")
			}

			s.logger.Error("error getting inway from database", zap.String("name", req.InwayName), zap.Error(err))

			return nil, status.Error(codes.Internal, "database error")
		}

		services = inway.Services
	}

	response := &api.ListServicesResponse{}

	if len(services) > 0 {
		response.Services = []*api.ListServicesResponse_Service{}

		counts, err := s.configDatabase.GetIncomingAccessRequestCountByService(ctx)
		if err != nil {
			s.logger.Error("error getting incoming access requests count from database", zap.Error(err))
			return nil, status.Error(codes.Internal, "database error")
		}

		for _, service := range services {
			protoService := convertFromDatabaseService(service)
			protoService.IncomingAccessRequestCount = uint32(counts[service.Name])

			accessGrants, err := s.configDatabase.ListAccessGrantsForService(ctx, service.Name)
			if err != nil {
				s.logger.Error("error getting access grants for service from database", zap.String("servicename", service.Name), zap.Error(err))
				continue
			}

			authorizations := make([]*api.ListServicesResponse_Service_AuthorizationSettings_Authorization, 0)

			for _, accessGrant := range accessGrants {
				if !accessGrant.RevokedAt.Valid {
					authorizations = append(authorizations, convertAccessGrantToAuthorizationSetting(accessGrant))
				}
			}

			protoService.AuthorizationSettings.Authorizations = authorizations

			response.Services = append(response.Services, protoService)
		}
	}

	return response, nil
}

// GetStatisticsOfServices return statistics per service
func (s *ManagementService) GetStatisticsOfServices(ctx context.Context, request *api.GetStatisticsOfServicesRequest) (*api.GetStatisticsOfServicesResponse, error) {
	s.logger.Info("rpc request GetStatsOfServices")

	countPerService, err := s.configDatabase.GetIncomingAccessRequestCountByService(ctx)
	if err != nil {
		s.logger.Error("error getting incoming access request count per service from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return convertToGetStatsOfServicesResponse(countPerService), err
}

func convertAccessGrantToAuthorizationSetting(accessGrant *database.AccessGrant) *api.ListServicesResponse_Service_AuthorizationSettings_Authorization {
	return &api.ListServicesResponse_Service_AuthorizationSettings_Authorization{
		OrganizationName: accessGrant.IncomingAccessRequest.OrganizationName,
		PublicKeyHash:    accessGrant.IncomingAccessRequest.PublicKeyFingerprint,
	}
}

func convertFromDatabaseService(model *database.Service) *api.ListServicesResponse_Service {
	inwayNames := []string{}

	for _, inway := range model.Inways {
		inwayNames = append(inwayNames, inway.Name)
	}

	service := &api.ListServicesResponse_Service{
		Name:                 model.Name,
		EndpointURL:          model.EndpointURL,
		DocumentationURL:     model.DocumentationURL,
		ApiSpecificationURL:  model.APISpecificationURL,
		Internal:             model.Internal,
		TechSupportContact:   model.TechSupportContact,
		PublicSupportContact: model.PublicSupportContact,
		Inways:               inwayNames,
		AuthorizationSettings: &api.ListServicesResponse_Service_AuthorizationSettings{
			Mode: "whitelist",
		},
	}

	return service
}

func convertToGetServiceResponseFromDatabaseService(model *database.Service) *api.GetServiceResponse {
	inwayNames := []string{}

	for _, inway := range model.Inways {
		inwayNames = append(inwayNames, inway.Name)
	}

	service := &api.GetServiceResponse{
		Name:                 model.Name,
		EndpointURL:          model.EndpointURL,
		DocumentationURL:     model.DocumentationURL,
		ApiSpecificationURL:  model.APISpecificationURL,
		Internal:             model.Internal,
		TechSupportContact:   model.TechSupportContact,
		PublicSupportContact: model.PublicSupportContact,
		Inways:               inwayNames,
	}

	return service
}

func convertToCreateServiceResponseFromCreateServiceRequest(model *api.CreateServiceRequest) *api.CreateServiceResponse {
	service := &api.CreateServiceResponse{
		Name:                 model.Name,
		EndpointURL:          model.EndpointURL,
		DocumentationURL:     model.DocumentationURL,
		ApiSpecificationURL:  model.ApiSpecificationURL,
		Internal:             model.Internal,
		TechSupportContact:   model.TechSupportContact,
		PublicSupportContact: model.PublicSupportContact,
		Inways:               model.Inways,
	}

	return service
}

func convertToUpdateServiceResponseFromUpdateServiceRequest(model *api.UpdateServiceRequest) *api.UpdateServiceResponse {
	service := &api.UpdateServiceResponse{
		Name:                 model.Name,
		EndpointURL:          model.EndpointURL,
		DocumentationURL:     model.DocumentationURL,
		ApiSpecificationURL:  model.ApiSpecificationURL,
		Internal:             model.Internal,
		TechSupportContact:   model.TechSupportContact,
		PublicSupportContact: model.PublicSupportContact,
		Inways:               model.Inways,
	}

	return service
}

func convertToGetStatsOfServicesResponse(accessRequestCountPerService map[string]int) *api.GetStatisticsOfServicesResponse {
	response := &api.GetStatisticsOfServicesResponse{
		Services: make([]*api.ServiceStatistics, len(accessRequestCountPerService)),
	}

	i := 0

	for serviceName, count := range accessRequestCountPerService {
		response.Services[i] = &api.ServiceStatistics{
			Name:                       serviceName,
			IncomingAccessRequestCount: uint32(count),
		}
		i++
	}

	return response
}
