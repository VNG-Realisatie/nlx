//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// CreateService creates a new service
func (s *ManagementService) CreateService(ctx context.Context, service *api.Service) (*api.Service, error) {
	logger := s.logger.With(zap.String("name", service.Name))
	logger.Info("rpc request CreateService")

	err := service.Validate()
	if err != nil {
		logger.Error("invalid service", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid service: %s", err))
	}

	model := &database.Service{
		Name:                 service.Name,
		EndpointURL:          service.EndpointURL,
		DocumentationURL:     service.DocumentationURL,
		APISpecificationURL:  service.ApiSpecificationURL,
		Internal:             service.Internal,
		TechSupportContact:   service.TechSupportContact,
		PublicSupportContact: service.PublicSupportContact,
		Inways:               service.Inways,
	}

	if service.AuthorizationSettings != nil {
		model.AuthorizationSettings = &database.ServiceAuthorizationSettings{
			Mode: service.AuthorizationSettings.Mode,
		}
	}

	err = s.configDatabase.CreateService(ctx, model)
	if err != nil {
		logger.Error("error creating service in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return service, nil
}

// GetService returns a specific service
func (s *ManagementService) GetService(ctx context.Context, req *api.GetServiceRequest) (*api.Service, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request GetService")

	service, err := s.configDatabase.GetService(ctx, req.Name)
	if err != nil {
		logger.Error("error getting service from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if service == nil {
		logger.Warn("service not found")
		return nil, status.Error(codes.NotFound, "service not found")
	}

	response := convertFromDatabaseService(service)

	return response, nil
}

// UpdateService updates an existing service
func (s *ManagementService) UpdateService(ctx context.Context, req *api.UpdateServiceRequest) (*api.Service, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request UpdateService")

	err := req.Service.Validate()
	if err != nil {
		logger.Error("invalid service", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid service: %s", err))
	}

	if req.Name != req.Service.Name {
		return nil, status.Error(codes.InvalidArgument, "changing the service name is not allowed")
	}

	service := &database.Service{
		Name:                 req.Service.Name,
		EndpointURL:          req.Service.EndpointURL,
		DocumentationURL:     req.Service.DocumentationURL,
		APISpecificationURL:  req.Service.ApiSpecificationURL,
		Internal:             req.Service.Internal,
		TechSupportContact:   req.Service.TechSupportContact,
		PublicSupportContact: req.Service.PublicSupportContact,
		Inways:               req.Service.Inways,
	}

	if req.Service.AuthorizationSettings != nil {
		service.AuthorizationSettings = &database.ServiceAuthorizationSettings{
			Mode: req.Service.AuthorizationSettings.Mode,
		}
	}

	err = s.configDatabase.UpdateService(ctx, req.Name, service)
	if err != nil {
		logger.Error("error updating service in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req.Service, nil
}

// DeleteService deletes a specific service
func (s *ManagementService) DeleteService(ctx context.Context, req *api.DeleteServiceRequest) (*api.Empty, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteService")

	err := s.configDatabase.DeleteService(ctx, req.Name)

	if err != nil {
		logger.Error("error deleting service in DB", zap.Error(err))
		return &api.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &api.Empty{}, nil
}

// ListServices returns a list of services
func (s *ManagementService) ListServices(ctx context.Context, req *api.ListServicesRequest) (*api.ListServicesResponse, error) {
	s.logger.Info("rpc request ListServices")

	services, err := s.configDatabase.ListServices(ctx)
	if err != nil {
		s.logger.Error("error getting services list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response := &api.ListServicesResponse{}

	var filteredServices []*database.Service

	if len(req.InwayName) > 0 {
		for _, service := range services {
			for _, inway := range service.Inways {
				if strings.Compare(req.InwayName, inway) == 0 {
					filteredServices = append(filteredServices, service)
					break
				}
			}
		}
	} else {
		filteredServices = services
	}

	if length := len(filteredServices); length > 0 {
		response.Services = make([]*api.Service, length)

		for i, service := range filteredServices {
			response.Services[i] = convertFromDatabaseService(service)
		}
	}

	return response, nil
}

func convertFromDatabaseService(model *database.Service) *api.Service {
	service := &api.Service{
		Name:                 model.Name,
		EndpointURL:          model.EndpointURL,
		DocumentationURL:     model.DocumentationURL,
		ApiSpecificationURL:  model.APISpecificationURL,
		Internal:             model.Internal,
		TechSupportContact:   model.TechSupportContact,
		PublicSupportContact: model.PublicSupportContact,
		Inways:               model.Inways,
	}

	if model.AuthorizationSettings != nil {
		service.AuthorizationSettings = &api.Service_AuthorizationSettings{
			Mode: model.AuthorizationSettings.Mode,
		}
	}

	return service
}
