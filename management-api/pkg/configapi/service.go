//nolint:dupl // service and inway structs look the same
package configapi

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateService creates a new service
func (s *ConfigService) CreateService(ctx context.Context, service *Service) (*Service, error) {
	logger := s.logger.With(zap.String("name", service.Name))
	logger.Info("rpc request CreateService")

	err := service.Validate()
	if err != nil {
		logger.Error("invalid service", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid service: %s", err))
	}

	err = s.configDatabase.CreateService(ctx, service)
	if err != nil {
		logger.Error("error creating service in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return service, nil
}

// GetService returns a specific service
func (s *ConfigService) GetService(ctx context.Context, req *GetServiceRequest) (*Service, error) {
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

	return service, nil
}

// UpdateService updates an existing service
func (s *ConfigService) UpdateService(ctx context.Context, req *UpdateServiceRequest) (*Service, error) {
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

	err = s.configDatabase.UpdateService(ctx, req.Name, req.Service)
	if err != nil {
		logger.Error("error updating service in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req.Service, nil
}

// DeleteService deletes a specific service
func (s *ConfigService) DeleteService(ctx context.Context, req *DeleteServiceRequest) (*Empty, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteService")

	err := s.configDatabase.DeleteService(ctx, req.Name)

	if err != nil {
		logger.Error("error deleting service in DB", zap.Error(err))
		return &Empty{}, status.Error(codes.Internal, "database error")
	}

	return &Empty{}, nil
}

// ListServices returns a list of services
func (s *ConfigService) ListServices(ctx context.Context, req *ListServicesRequest) (*ListServicesResponse, error) {
	s.logger.Info("rpc request ListServices")

	services, err := s.configDatabase.ListServices(ctx)
	if err != nil {
		s.logger.Error("error getting services list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response := &ListServicesResponse{}

	if len(req.InwayName) > 0 {
		for _, service := range services {
			for _, inway := range service.Inways {
				if strings.Compare(req.InwayName, inway) == 0 {
					response.Services = append(response.Services, service)
					break
				}
			}
		}
	} else {
		response.Services = services
	}

	return response, nil
}
