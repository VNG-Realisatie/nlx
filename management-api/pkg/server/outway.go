// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"fmt"
	"net"

	"github.com/jackc/pgtype"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

func (s *ManagementService) RegisterOutway(ctx context.Context, req *api.RegisterOutwayRequest) (*emptypb.Empty, error) {
	logger := s.logger

	logger.Info("rpc request RegisterOutway")

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

	if err := req.Validate(); err != nil {
		logger.Error("invalid outway", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid outway: %s", err))
	}

	ipAddress, err := getCIDRFromTCPAddress(addr)
	if err != nil {
		logger.Error("cannot get CIDR from TCP address", zap.Error(err), zap.Any("TCP address", addr))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid outway: %s", err))
	}

	model := &database.Outway{
		Name: req.Name,
		IPAddress: pgtype.Inet{
			IPNet:  ipAddress,
			Status: pgtype.Present,
		},
		PublicKeyPEM: req.PublicKeyPEM,
		Version:      req.Version,
	}

	if err := s.configDatabase.RegisterOutway(ctx, model); err != nil {
		logger.Error("error creating outway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &emptypb.Empty{}, nil
}

// ListInways returns a list of outways
func (s *ManagementService) ListOutways(ctx context.Context, req *api.ListOutwaysRequest) (*api.ListOutwaysResponse, error) {
	s.logger.Info("rpc request ListInways")

	outways, err := s.configDatabase.ListOutways(ctx)
	if err != nil {
		s.logger.Error("error getting outway list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response := &api.ListOutwaysResponse{}
	response.Outways = make([]*api.Outway, len(outways))

	for i, outway := range outways {
		response.Outways[i] = convertFromDatabaseOutway(outway)
	}

	return response, nil
}

func getCIDRFromTCPAddress(tcpAddress *net.TCPAddr) (*net.IPNet, error) {
	_, ipAddress, err := net.ParseCIDR(fmt.Sprintf("%s/32", tcpAddress.IP.String()))
	if err != nil {
		return nil, err
	}

	return ipAddress, nil
}

func convertFromDatabaseOutway(model *database.Outway) *api.Outway {
	outway := &api.Outway{
		Name:         model.Name,
		PublicKeyPEM: model.PublicKeyPEM,
		Version:      model.Version,
	}

	if model.IPAddress.Status == pgtype.Present {
		outway.IpAddress = model.IPAddress.IPNet.IP.String()
	}

	return outway
}