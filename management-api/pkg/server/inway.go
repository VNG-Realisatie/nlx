//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"fmt"
	"net"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

// CreateInway creates a new inway
func (s *ManagementService) CreateInway(ctx context.Context, inway *api.Inway) (*api.Inway, error) {
	logger := s.logger.With(zap.String("name", inway.Name))

	logger.Info("rpc request CreateInway")

	p, ok := peer.FromContext(ctx)
	if !ok {
		logger.Error("peer context cannot be found")
		return nil, status.Error(codes.Internal, "peer context cannot be found")
	}

	addr, ok := p.Addr.(*net.TCPAddr)
	if !ok {
		logger.Error("peer addr is invalid")
		return nil, status.Error(codes.Internal, "peer addr is invalid")
	}

	inway.IpAddress = addr.IP.String()

	if err := inway.Validate(); err != nil {
		logger.Error("invalid inway", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid inway: %s", err))
	}

	model := &database.Inway{
		Name:        inway.Name,
		Version:     inway.Version,
		Hostname:    inway.Hostname,
		SelfAddress: inway.SelfAddress,
		IPAddress:   inway.IpAddress,
	}

	if err := s.configDatabase.CreateInway(ctx, model); err != nil {
		logger.Error("error creating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return inway, nil
}

// GetInway returns a specific inway
func (s *ManagementService) GetInway(ctx context.Context, req *api.GetInwayRequest) (*api.Inway, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request GetInway")

	inway, err := s.configDatabase.GetInway(ctx, req.Name)
	if err != nil {
		if errIsNotFound(err) {
			logger.Warn("inway not found")
			return nil, status.Error(codes.NotFound, "inway not found")
		}

		logger.Error("error getting inway from DB", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	response := &api.Inway{
		Name:        inway.Name,
		Version:     inway.Version,
		Hostname:    inway.Hostname,
		SelfAddress: inway.SelfAddress,
		IpAddress:   inway.IPAddress,
	}

	response.Services = make([]*api.Inway_Service, len(inway.Services))

	for i, service := range inway.Services {
		servicesResponse := &api.Inway_Service{
			Name: service.Name,
		}

		response.Services[i] = servicesResponse
	}

	return response, nil
}

// UpdateInway updates an existing inway
func (s *ManagementService) UpdateInway(ctx context.Context, req *api.UpdateInwayRequest) (*api.Inway, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request UpdateInway")

	err := req.Inway.Validate()
	if err != nil {
		logger.Error("invalid inway", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid inway: %s", err))
	}

	inwayInDB, err := s.configDatabase.GetInway(ctx, req.Inway.Name)
	if err != nil {
		if errIsNotFound(err) {
			logger.Error("inway does not exist")
			return nil, status.Error(codes.NotFound, fmt.Sprintf("inway with the name %s does not exist", req.Name))
		}

		logger.Error("error getting inway from database", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	inway := &database.Inway{
		ID:          inwayInDB.ID,
		Name:        req.Inway.Name,
		Version:     req.Inway.Version,
		Hostname:    req.Inway.Hostname,
		SelfAddress: req.Inway.SelfAddress,
	}

	err = s.configDatabase.UpdateInway(ctx, inway)
	if err != nil {
		logger.Error("error updating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req.Inway, nil
}

// DeleteInway deletes a specific inway
func (s *ManagementService) DeleteInway(ctx context.Context, req *api.DeleteInwayRequest) (*types.Empty, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteInway")

	err := s.configDatabase.DeleteInway(ctx, req.Name)
	if err != nil {
		logger.Error("error deleting inway in DB", zap.Error(err))
		return &types.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &types.Empty{}, nil
}

// ListInways returns a list of inways
func (s *ManagementService) ListInways(ctx context.Context, req *api.ListInwaysRequest) (*api.ListInwaysResponse, error) {
	s.logger.Info("rpc request ListInways")

	inways, err := s.configDatabase.ListInways(ctx)
	if err != nil {
		s.logger.Error("error getting inway list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response := &api.ListInwaysResponse{}
	response.Inways = make([]*api.Inway, len(inways))

	for i, inway := range inways {
		response.Inways[i] = convertFromDatabaseInway(inway)
	}

	return response, nil
}

func convertFromDatabaseInway(model *database.Inway) *api.Inway {
	inway := &api.Inway{
		Name:        model.Name,
		Version:     model.Version,
		Hostname:    model.Hostname,
		SelfAddress: model.SelfAddress,
		IpAddress:   model.IPAddress,
	}

	if len(model.Services) > 0 {
		inway.Services = make([]*api.Inway_Service, len(model.Services))

		for i, service := range model.Services {
			inway.Services[i] = &api.Inway_Service{
				Name: service.Name,
			}
		}
	}

	return inway
}
