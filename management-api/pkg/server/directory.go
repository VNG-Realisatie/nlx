// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/environment"
)

type DirectoryService struct {
	api.UnimplementedDirectoryServer

	logger      *zap.Logger
	environment *environment.Environment

	directoryClient directory.Client
	configDatabase  database.ConfigDatabase
}

var inwayStateToDirectoryState = map[directoryapi.Inway_State]api.DirectoryService_State{
	directoryapi.Inway_UNKNOWN: api.DirectoryService_unknown,
	directoryapi.Inway_UP:      api.DirectoryService_up,
	directoryapi.Inway_DOWN:    api.DirectoryService_down,
}

func NewDirectoryService(logger *zap.Logger, e *environment.Environment, directoryClient directory.Client, configDatabase database.ConfigDatabase) *DirectoryService {
	return &DirectoryService{
		logger:          logger,
		environment:     e,
		directoryClient: directoryClient,
		configDatabase:  configDatabase,
	}
}

func (s DirectoryService) ListServices(ctx context.Context, _ *emptypb.Empty) (*api.DirectoryListServicesResponse, error) {
	s.logger.Info("rpc request ListServices")

	resp, err := s.directoryClient.ListServices(ctx, &emptypb.Empty{})
	if err != nil {
		s.logger.Error("error getting services list from directory", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "directory error")
	}

	services := make([]*api.DirectoryService, len(resp.Services))

	for i, service := range resp.Services {
		convertedService := convertDirectoryService(service)

		accessRequestStates, err := getLatestAccessRequestStates(ctx, s.configDatabase, convertedService.Organization.SerialNumber, convertedService.ServiceName)
		if err != nil {
			s.logger.Error("error getting latest access request states", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "database error")
		}

		convertedService.AccessStates = accessRequestStates

		services[i] = convertedService
	}

	return &api.DirectoryListServicesResponse{Services: services}, nil
}

func (s DirectoryService) GetOrganizationService(ctx context.Context, request *api.GetOrganizationServiceRequest) (*api.DirectoryService, error) {
	logger := s.logger.With(zap.String("organizationSerialNumber", request.OrganizationSerialNumber), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request GetOrganizationService")

	service, err := s.getService(ctx, logger, request.OrganizationSerialNumber, request.ServiceName)
	if err != nil {
		return nil, err
	}

	directoryService := convertDirectoryService(service)

	accessRequestStates, err := getLatestAccessRequestStates(ctx, s.configDatabase, directoryService.Organization.SerialNumber, directoryService.ServiceName)
	if err != nil {
		s.logger.Error("error getting latest access request states", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "database error")
	}

	directoryService.AccessStates = accessRequestStates

	return directoryService, nil
}

func (s DirectoryService) getService(ctx context.Context, logger *zap.Logger, organizationSerialNumber, serviceName string) (*directoryapi.ListServicesResponse_Service, error) {
	resp, err := s.directoryClient.ListServices(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, status.Error(codes.Internal, "directory not available")
	}

	for _, s := range resp.Services {
		if s.Organization.SerialNumber == organizationSerialNumber && s.Name == serviceName {
			return s, nil
		}
	}

	logger.Warn("service not found")

	return nil, status.Error(codes.NotFound, "service not found")
}

func (s DirectoryService) RequestAccessToService(ctx context.Context, request *api.RequestAccessToServiceRequest) (*api.OutgoingAccessRequest, error) {
	logger := s.logger.With(zap.String("organizationSerialNumber", request.OrganizationSerialNumber), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request RequestAccessToService")

	ar := &database.OutgoingAccessRequest{
		Organization: database.Organization{
			SerialNumber: request.OrganizationSerialNumber,
		},
		ServiceName: request.ServiceName,
	}

	accessRequest, err := s.configDatabase.CreateOutgoingAccessRequest(ctx, ar)
	if err != nil {
		return nil, err
	}

	response := convertOutgoingAccessRequest(accessRequest)

	service, err := s.getService(ctx, logger, request.OrganizationSerialNumber, request.ServiceName)
	if err != nil {
		return nil, err
	}

	logger = logger.With(zap.Any("service", service))

	logger.Debug("send access request to inway")

	return response, nil
}

func (s DirectoryService) GetTermsOfService(ctx context.Context, _ *emptypb.Empty) (*api.GetTermsOfServiceResponse, error) {
	s.logger.Info("rpc request GetTermsOfService")

	response, err := s.directoryClient.GetTermsOfService(ctx, &emptypb.Empty{})
	if err != nil {
		s.logger.Debug("unable to get terms of service from directory", zap.Error(err))
		return nil, status.Error(codes.Internal, "unable to get terms of service from directory")
	}

	return &api.GetTermsOfServiceResponse{
		Enabled: response.Enabled,
		Url:     response.Url,
	}, nil
}

func DetermineDirectoryServiceState(inways []*directoryapi.Inway) api.DirectoryService_State {
	serviceState := api.DirectoryService_unknown

	if len(inways) == 0 {
		return serviceState
	}

	stateMap := map[directoryapi.Inway_State]int{}

	for _, i := range inways {
		stateMap[i.State]++
	}

	if len(stateMap) > 1 {
		return api.DirectoryService_degraded
	}

	for state := range stateMap {
		serviceState = inwayStateToDirectoryState[state]
	}

	return serviceState
}

func convertDirectoryService(s *directoryapi.ListServicesResponse_Service) *api.DirectoryService {
	serviceState := DetermineDirectoryServiceState(s.Inways)

	service := &api.DirectoryService{
		ServiceName:          s.Name,
		ApiSpecificationType: s.ApiSpecificationType,
		DocumentationURL:     s.DocumentationUrl,
		PublicSupportContact: s.PublicSupportContact,
		State:                serviceState,
	}

	if s.Organization != nil {
		service.Organization = &api.Organization{
			SerialNumber: s.Organization.SerialNumber,
			Name:         s.Organization.Name,
		}
	}

	// @TODO: Use costs object in api.DirectoryService
	if s.Costs != nil {
		service.OneTimeCosts = s.Costs.OneTime
		service.MonthlyCosts = s.Costs.Monthly
		service.RequestCosts = s.Costs.Request
	}

	return service
}

func getLatestAccessRequestStates(ctx context.Context, configDatabase database.ConfigDatabase, organizationSerialNumber, serviceName string) ([]*api.DirectoryService_AccessState, error) {
	outgoingAccessRequests, err := configDatabase.ListLatestOutgoingAccessRequests(ctx, organizationSerialNumber, serviceName)
	if err != nil {
		return nil, err
	}

	accessRequestStates := make([]*api.DirectoryService_AccessState, len(outgoingAccessRequests))

	for i, outgoingAccessRequest := range outgoingAccessRequests {
		accessRequestState := &api.DirectoryService_AccessState{
			AccessRequest: convertOutgoingAccessRequest(outgoingAccessRequest),
		}
		accessRequestStates[i] = accessRequestState

		accessProof, err := configDatabase.GetAccessProofForOutgoingAccessRequest(ctx, outgoingAccessRequest.ID)
		if err != nil && !errors.Is(err, database.ErrNotFound) {
			return nil, err
		}

		if accessProof != nil {
			accessRequestState.AccessProof = convertAccessProof(accessProof)
		}
	}

	return accessRequestStates, nil
}

func convertAccessProof(accessProof *database.AccessProof) *api.AccessProof {
	createdAt := timestamppb.New(accessProof.CreatedAt)

	var revokedAt *timestamp.Timestamp

	if accessProof.RevokedAt.Valid {
		revokedAt = timestamppb.New(accessProof.RevokedAt.Time)
	}

	return &api.AccessProof{
		Id: uint64(accessProof.ID),
		Organization: &api.Organization{
			SerialNumber: accessProof.OutgoingAccessRequest.Organization.SerialNumber,
			Name:         accessProof.OutgoingAccessRequest.Organization.Name,
		},
		ServiceName:     accessProof.OutgoingAccessRequest.ServiceName,
		CreatedAt:       createdAt,
		RevokedAt:       revokedAt,
		AccessRequestId: uint64(accessProof.OutgoingAccessRequest.ID),
	}
}
