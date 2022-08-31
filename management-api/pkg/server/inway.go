// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

//nolint:dupl // service and inway structs look the same
package server

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

// RegisterInway creates a new inway
func (s *ManagementService) RegisterInway(ctx context.Context, inway *api.Inway) (*api.Inway, error) {
	logger := s.logger.With(zap.String("name", inway.Name))

	logger.Info("rpc request RegisterInway")

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
		CreatedAt:   s.clock.Now(),
		UpdatedAt:   s.clock.Now(),
	}

	if err := s.configDatabase.RegisterInway(ctx, model); err != nil {
		logger.Error("error creating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return inway, nil
}

// GetInway returns a specific inway
func (s *ManagementService) GetInway(ctx context.Context, req *api.GetInwayRequest) (*api.Inway, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request GetInway")

	err := s.authorize(ctx, permissions.ReadInway)
	if err != nil {
		return nil, err
	}

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

	err := s.authorize(ctx, permissions.UpdateInway)
	if err != nil {
		return nil, err
	}

	err = req.Inway.Validate()
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
func (s *ManagementService) DeleteInway(ctx context.Context, req *api.DeleteInwayRequest) (*emptypb.Empty, error) {
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteInway")

	userInfo, err := retrieveUserFromContext(ctx)
	if err != nil {
		s.logger.Error("could not retrieve user info for audit log from grpc context", zap.Error(err))
		return nil, status.Error(codes.Internal, "could not retrieve user info to create audit log")
	}

	err = s.authorize(ctx, permissions.DeleteInway)
	if err != nil {
		return nil, err
	}

	err = s.auditLogger.InwayDelete(ctx, userInfo.Email, userInfo.UserAgent, req.Name)
	if err != nil {
		s.logger.Error("failed to write auditlog", zap.Error(err))

		return nil, status.Error(codes.Internal, "failed to write to auditlog")
	}

	err = s.configDatabase.DeleteInway(ctx, req.Name)
	if err != nil {
		logger.Error("error deleting inway in DB", zap.Error(err))
		return &emptypb.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &emptypb.Empty{}, nil
}

// ListInways returns a list of inways
func (s *ManagementService) ListInways(ctx context.Context, req *api.ListInwaysRequest) (*api.ListInwaysResponse, error) {
	s.logger.Info("rpc request ListInways")

	err := s.authorize(ctx, permissions.ReadInways)
	if err != nil {
		return nil, err
	}

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

func (s *ManagementService) GetInwayConfig(ctx context.Context, req *api.GetInwayConfigRequest) (*api.GetInwayConfigResponse, error) {
	s.logger.Info("rpc request GetInwayConfig")

	var services []*database.Service

	inway, err := s.configDatabase.GetInway(ctx, req.Name)
	if err != nil {
		if errIsNotFound(err) {
			return nil, InwayNotFoundError
		}

		s.logger.Error("error getting inway from database", zap.String("name", req.Name), zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	settings, err := s.configDatabase.GetSettings(ctx)
	if err != nil {
		s.logger.Error("could not get the settings from the database", zap.Error(err))

		return nil, status.Error(codes.Internal, "database error")
	}

	response := &api.GetInwayConfigResponse{
		IsOrganizationInway: settings.OrganizationInwayName() == inway.Name,
	}

	services = inway.Services

	if len(services) > 0 {
		response.Services = []*api.GetInwayConfigResponse_Service{}

		for _, service := range services {
			protoService := convertFromDatabaseServiceToInwayService(service)

			accessGrants, err := s.configDatabase.ListAccessGrantsForService(ctx, service.Name)
			if err != nil {
				s.logger.Error("error getting access grants for service from database", zap.String("servicename", service.Name), zap.Error(err))
				continue
			}

			authorizations := make([]*api.GetInwayConfigResponse_Service_AuthorizationSettings_Authorization, 0)

			for _, accessGrant := range accessGrants {
				if !accessGrant.RevokedAt.Valid {
					authorizations = append(authorizations, convertAccessGrantToInwayAuthorizationSetting(accessGrant))
				}
			}

			protoService.AuthorizationSettings.Authorizations = authorizations

			response.Services = append(response.Services, protoService)
		}
	}

	return response, nil
}

func convertFromDatabaseServiceToInwayService(model *database.Service) *api.GetInwayConfigResponse_Service {
	service := &api.GetInwayConfigResponse_Service{
		Name:                  model.Name,
		EndpointURL:           model.EndpointURL,
		DocumentationURL:      model.DocumentationURL,
		ApiSpecificationURL:   model.APISpecificationURL,
		Internal:              model.Internal,
		TechSupportContact:    model.TechSupportContact,
		PublicSupportContact:  model.PublicSupportContact,
		OneTimeCosts:          int32(model.OneTimeCosts),
		MonthlyCosts:          int32(model.MonthlyCosts),
		RequestCosts:          int32(model.RequestCosts),
		AuthorizationSettings: &api.GetInwayConfigResponse_Service_AuthorizationSettings{},
	}

	return service
}

func convertAccessGrantToInwayAuthorizationSetting(accessGrant *database.AccessGrant) *api.GetInwayConfigResponse_Service_AuthorizationSettings_Authorization {
	return &api.GetInwayConfigResponse_Service_AuthorizationSettings_Authorization{
		Organization: &api.Organization{
			Name:         accessGrant.IncomingAccessRequest.Organization.Name,
			SerialNumber: accessGrant.IncomingAccessRequest.Organization.SerialNumber,
		},
		PublicKeyHash: accessGrant.IncomingAccessRequest.PublicKeyFingerprint,
		PublicKeyPEM:  accessGrant.IncomingAccessRequest.PublicKeyPEM,
	}
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
