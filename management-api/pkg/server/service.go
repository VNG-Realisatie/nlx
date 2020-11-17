//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// CreateService creates a new service
func (s *ManagementService) CreateService(ctx context.Context, service *api.CreateServiceRequest) (*api.CreateServiceResponse, error) {
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

	err = s.configDatabase.CreateService(ctx, model)
	if err != nil {
		logger.Error("error creating service in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return convertToCreateServiceResponseFromCreateServiceRequest(service), nil
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

	service := &database.Service{
		Name:                 req.Name,
		EndpointURL:          req.EndpointURL,
		DocumentationURL:     req.DocumentationURL,
		APISpecificationURL:  req.ApiSpecificationURL,
		Internal:             req.Internal,
		TechSupportContact:   req.TechSupportContact,
		PublicSupportContact: req.PublicSupportContact,
		Inways:               req.Inways,
	}

	err = s.configDatabase.UpdateService(ctx, req.Name, service)
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

	err := s.configDatabase.DeleteService(ctx, req.Name)

	if err != nil {
		logger.Error("error deleting service in DB", zap.Error(err))
		return &types.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &types.Empty{}, nil
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
		response.Services = []*api.Service{}

		for _, service := range filteredServices {
			convertedService := convertFromDatabaseService(service)

			accessGrants, err := s.configDatabase.ListAccessGrantsForService(ctx, service.Name)
			if err != nil {
				s.logger.Error("error getting access grants for service from database", zap.String("servicename", service.Name), zap.Error(err))
				continue
			}

			authorizations := make([]*api.Service_AuthorizationSettings_Authorization, 0)

			for _, accessGrant := range accessGrants {
				if !accessGrant.Revoked() {
					authorizations = append(authorizations, convertAccessGrantToAuthorizationSetting(accessGrant))
				}
			}

			convertedService.AuthorizationSettings.Authorizations = authorizations

			response.Services = append(response.Services, convertedService)
		}
	}

	return response, nil
}

func convertAccessGrantToAuthorizationSetting(accessGrant *database.AccessGrant) *api.Service_AuthorizationSettings_Authorization {
	return &api.Service_AuthorizationSettings_Authorization{
		OrganizationName: accessGrant.OrganizationName,
		PublicKeyHash:    accessGrant.PublicKeyFingerprint,
	}
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
		AuthorizationSettings: &api.Service_AuthorizationSettings{
			Mode: "whitelist",
		},
	}

	return service
}

func convertToGetServiceResponseFromDatabaseService(model *database.Service) *api.GetServiceResponse {
	service := &api.GetServiceResponse{
		Name:                 model.Name,
		EndpointURL:          model.EndpointURL,
		DocumentationURL:     model.DocumentationURL,
		ApiSpecificationURL:  model.APISpecificationURL,
		Internal:             model.Internal,
		TechSupportContact:   model.TechSupportContact,
		PublicSupportContact: model.PublicSupportContact,
		Inways:               model.Inways,
		AuthorizationSettings: &api.GetServiceResponse_AuthorizationSettings{
			Mode: "whitelist",
		},
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
		AuthorizationSettings: &api.CreateServiceResponse_AuthorizationSettings{
			Mode: model.AuthorizationSettings.Mode,
		},
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
