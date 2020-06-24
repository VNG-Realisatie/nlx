//nolint:dupl // service and inway structs look the same
package configapi

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/pkg/database"
)

// CreateInway creates a new inway
func (s *ConfigService) CreateInway(ctx context.Context, inway *Inway) (*Inway, error) {
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

	err := inway.Validate()
	if err != nil {
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

	err = s.configDatabase.CreateInway(ctx, model)
	if err != nil {
		logger.Error("error creating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return inway, nil
}

// GetInway returns a specific inway
func (s *ConfigService) GetInway(ctx context.Context, req *GetInwayRequest) (*Inway, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request GetInway")

	inway, err := s.configDatabase.GetInway(ctx, req.Name)
	if err != nil {
		logger.Error("error getting inway from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if inway == nil {
		logger.Warn("inway not found")
		return nil, status.Error(codes.NotFound, "inway not found")
	}

	services, err := s.configDatabase.ListServices(ctx)
	if err != nil {
		s.logger.Error("error getting services  from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if len(services) > 0 {
		services = database.FilterServices(services, inway)
	}

	response := &Inway{
		Name:        inway.Name,
		Version:     inway.Version,
		Hostname:    inway.Hostname,
		SelfAddress: inway.SelfAddress,
		IpAddress:   inway.IPAddress,
	}

	response.Services = make([]*Inway_Service, len(services))

	for i, service := range services {
		servicesResponse := &Inway_Service{
			Name: service.Name,
		}

		response.Services[i] = servicesResponse
	}

	return response, nil
}

// UpdateInway updates an existing inway
func (s *ConfigService) UpdateInway(ctx context.Context, req *UpdateInwayRequest) (*Inway, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request UpdateInway")

	err := req.Inway.Validate()
	if err != nil {
		logger.Error("invalid inway", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid inway: %s", err))
	}

	inway := &database.Inway{
		Name:        req.Inway.Name,
		Version:     req.Inway.Version,
		Hostname:    req.Inway.Hostname,
		SelfAddress: req.Inway.SelfAddress,
	}

	err = s.configDatabase.UpdateInway(ctx, req.Name, inway)
	if err != nil {
		logger.Error("error updating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req.Inway, nil
}

// DeleteInway deletes a specific inway
func (s *ConfigService) DeleteInway(ctx context.Context, req *DeleteInwayRequest) (*Empty, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteInway")

	err := s.configDatabase.DeleteInway(ctx, req.Name)

	if err != nil {
		logger.Error("error deleting inway in DB", zap.Error(err))
		return &Empty{}, status.Error(codes.Internal, "database error")
	}

	return &Empty{}, nil
}

// ListInways returns a list of inways
func (s *ConfigService) ListInways(ctx context.Context, req *ListInwaysRequest) (*ListInwaysResponse, error) {
	s.logger.Info("rpc request ListInways")

	inways, err := s.configDatabase.ListInways(ctx)
	if err != nil {
		s.logger.Error("error getting inway list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	services, err := s.configDatabase.ListServices(ctx)
	if err != nil {
		s.logger.Error("error getting services  from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response := &ListInwaysResponse{}
	response.Inways = make([]*Inway, len(inways))

	for i, inway := range inways {
		inwayServices := database.FilterServices(services, inway)
		response.Inways[i] = convertFromDatabaseInway(inway, inwayServices)
	}

	return response, nil
}

func convertFromDatabaseInway(model *database.Inway, services []*database.Service) *Inway {
	inway := &Inway{
		Name:        model.Name,
		Version:     model.Version,
		Hostname:    model.Hostname,
		SelfAddress: model.SelfAddress,
		IpAddress:   model.IPAddress,
	}

	if length := len(services); length > 0 {
		inway.Services = make([]*Inway_Service, length)

		for i, service := range services {
			inwayService := &Inway_Service{
				Name: service.Name,
			}

			inway.Services[i] = inwayService
		}
	}

	return inway
}
