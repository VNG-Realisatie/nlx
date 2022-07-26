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

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
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

	fingerPrint, err := common_tls.PemPublicKeyFingerprint([]byte(req.PublicKeyPEM))
	if err != nil {
		logger.Error("unable to generate public key fingerprint", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, "invalid public key")
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
		SelfAddressAPI:       req.SelfAddressAPI,
		PublicKeyPEM:         req.PublicKeyPEM,
		PublicKeyFingerprint: fingerPrint,
		Version:              req.Version,
	}

	if err := s.configDatabase.RegisterOutway(ctx, model); err != nil {
		logger.Error("error creating outway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &emptypb.Empty{}, nil
}

// ListOutways returns a list of outways
func (s *ManagementService) ListOutways(ctx context.Context, req *api.ListOutwaysRequest) (*api.ListOutwaysResponse, error) {
	err := s.authorize(ctx, permissions.ReadOutways)
	if err != nil {
		return nil, err
	}

	s.logger.Info("rpc request ListOutways")

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

// DeleteOutway deletes a specific outway
func (s *ManagementService) DeleteOutway(ctx context.Context, req *api.DeleteOutwayRequest) (*emptypb.Empty, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteOutway")

	userInfo, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.authorize(ctx, permissions.DeleteOutway)
	if err != nil {
		return nil, err
	}

	err = s.auditLogger.OutwayDelete(ctx, userInfo.Email, userInfo.UserAgent, req.Name)
	if err != nil {
		s.logger.Error("failed to write auditlog", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to write to auditlog")
	}

	err = s.configDatabase.DeleteOutway(ctx, req.Name)
	if err != nil {
		logger.Error("error deleting outway in DB", zap.Error(err))
		return &emptypb.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &emptypb.Empty{}, nil
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
		Name:                 model.Name,
		PublicKeyPEM:         model.PublicKeyPEM,
		PublicKeyFingerprint: model.PublicKeyFingerprint,
		SelfAddressAPI:       model.SelfAddressAPI,
		Version:              model.Version,
	}

	if model.IPAddress.Status == pgtype.Present {
		outway.IpAddress = model.IPAddress.IPNet.IP.String()
	}

	return outway
}
