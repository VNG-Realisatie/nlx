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

	err = s.configDatabase.CreateInway(ctx, inway)
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
		inway.Services = FilterServices(services, inway)
	}

	return inway, nil
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

	err = s.configDatabase.UpdateInway(ctx, req.Name, req.Inway)
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

	if len(services) > 0 {
		for _, i := range inways {
			inwayServices := FilterServices(services, i)
			i.Services = append(i.Services, inwayServices...)
		}
	}

	return &ListInwaysResponse{
		Inways: inways,
	}, nil
}

// FilterServices returns an array with only services for the given inway
func FilterServices(services []*Service, i *Inway) []*Inway_Service {
	result := []*Inway_Service{}

	for _, service := range services {
		if contains(service.Inways, i.Name) {
			result = append(result, &Inway_Service{Name: service.Name})
		}
	}

	return result
}
