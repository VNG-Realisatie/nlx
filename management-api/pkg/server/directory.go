// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
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

var inwayStateToDirectoryState = map[inspectionapi.Inway_State]api.DirectoryService_State{
	inspectionapi.Inway_UNKNOWN: api.DirectoryService_unknown,
	inspectionapi.Inway_UP:      api.DirectoryService_up,
	inspectionapi.Inway_DOWN:    api.DirectoryService_down,
}

func NewDirectoryService(logger *zap.Logger, e *environment.Environment, directoryClient directory.Client, configDatabase database.ConfigDatabase) *DirectoryService {
	return &DirectoryService{
		logger:          logger,
		environment:     e,
		directoryClient: directoryClient,
		configDatabase:  configDatabase,
	}
}

// ListServices returns all services known to the directory except those with the same organization
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

		accessRequest, accessProof, err := getLatestAccessRequestAndAccessGrant(ctx, s.configDatabase, convertedService.OrganizationName, convertedService.ServiceName)
		if err != nil {
			s.logger.Error("error getting latest access request and access proof", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "database error")
		}

		convertedService.LatestAccessRequest = accessRequest
		convertedService.LatestAccessProof = accessProof

		services[i] = convertedService
	}

	return &api.DirectoryListServicesResponse{Services: services}, nil
}

// GetOrganizationService returns a specific service of and organization
func (s DirectoryService) GetOrganizationService(ctx context.Context, request *api.GetOrganizationServiceRequest) (*api.DirectoryService, error) {
	logger := s.logger.With(zap.String("organizationName", request.OrganizationName), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request GetOrganizationService")

	service, err := s.getService(ctx, logger, request.OrganizationName, request.ServiceName)
	if err != nil {
		return nil, err
	}

	directoryService := convertDirectoryService(service)

	accessRequest, accessProof, err := getLatestAccessRequestAndAccessGrant(ctx, s.configDatabase, directoryService.OrganizationName, directoryService.ServiceName)
	if err != nil {
		s.logger.Error("error getting latest access request and access proof", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "database error")
	}

	directoryService.LatestAccessRequest = accessRequest
	directoryService.LatestAccessProof = accessProof

	return directoryService, nil
}

// @TODO: organizationName to serialNumber
func (s DirectoryService) getService(ctx context.Context, logger *zap.Logger, organizationName, serviceName string) (*inspectionapi.ListServicesResponse_Service, error) {
	resp, err := s.directoryClient.ListServices(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, status.Error(codes.Internal, "directory not available")
	}

	for _, s := range resp.Services {
		if s.Organization.Name == organizationName && s.Name == serviceName {
			return s, nil
		}
	}

	logger.Warn("service not found")

	return nil, status.Error(codes.NotFound, "service not found")
}

// RequestAccessToService records an access request and sends it to the organization
func (s DirectoryService) RequestAccessToService(ctx context.Context, request *api.RequestAccessToServiceRequest) (*api.OutgoingAccessRequest, error) {
	logger := s.logger.With(zap.String("organizationName", request.OrganizationName), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request RequestAccessToService")

	ar := &database.OutgoingAccessRequest{
		OrganizationName: request.OrganizationName,
		ServiceName:      request.ServiceName,
	}

	accessRequest, err := s.configDatabase.CreateOutgoingAccessRequest(ctx, ar)
	if err != nil {
		return nil, err
	}

	response, err := convertDirectoryAccessRequest(accessRequest)
	if err != nil {
		return nil, err
	}

	service, err := s.getService(ctx, logger, request.OrganizationName, request.ServiceName)
	if err != nil {
		return nil, err
	}

	logger = logger.With(zap.Any("service", service))

	logger.Debug("send access request to inway")

	return response, nil
}

func DetermineDirectoryServiceState(inways []*inspectionapi.Inway) api.DirectoryService_State {
	serviceState := api.DirectoryService_unknown

	if len(inways) == 0 {
		return serviceState
	}

	stateMap := map[inspectionapi.Inway_State]int{}

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

func convertDirectoryService(s *inspectionapi.ListServicesResponse_Service) *api.DirectoryService {
	serviceState := DetermineDirectoryServiceState(s.Inways)

	service := &api.DirectoryService{
		ServiceName:          s.Name,
		ApiSpecificationType: s.ApiSpecificationType,
		DocumentationURL:     s.DocumentationUrl,
		PublicSupportContact: s.PublicSupportContact,
		State:                serviceState,
	}

	// @TODO: use Serial Number and org object in api.DirectoryService
	if s.Organization != nil {
		service.OrganizationName = s.Organization.Name
	}

	// @TODO: Use costs object in api.DirectoryService
	if s.Costs != nil {
		service.OneTimeCosts = s.Costs.OneTime
		service.MonthlyCosts = s.Costs.Monthly
		service.RequestCosts = s.Costs.Request
	}

	return service
}

func convertDirectoryAccessRequest(a *database.OutgoingAccessRequest) (*api.OutgoingAccessRequest, error) {
	createdAt, err := ptypes.TimestampProto(a.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := ptypes.TimestampProto(a.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var accessRequestState api.AccessRequestState

	switch a.State {
	case database.OutgoingAccessRequestFailed:
		accessRequestState = api.AccessRequestState_FAILED
	case database.OutgoingAccessRequestCreated:
		accessRequestState = api.AccessRequestState_CREATED
	case database.OutgoingAccessRequestReceived:
		accessRequestState = api.AccessRequestState_RECEIVED
	}

	return &api.OutgoingAccessRequest{
		Id:        uint64(a.ID),
		State:     accessRequestState,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func getLatestAccessRequestAndAccessGrant(ctx context.Context, configDatabase database.ConfigDatabase, organizationName, serviceName string) (*api.OutgoingAccessRequest, *api.AccessProof, error) {
	latestAccessRequest, errDatabase := configDatabase.GetLatestOutgoingAccessRequest(ctx, organizationName, serviceName)
	if errDatabase != nil {
		if !errIsNotFound(errDatabase) {
			return nil, nil, errors.Wrap(errDatabase, "error retrieving latest access request")
		}
	}

	if errIsNotFound(errDatabase) {
		return nil, nil, nil
	}

	latestAccessProof, err := configDatabase.GetAccessProofForOutgoingAccessRequest(ctx, latestAccessRequest.ID)
	if err != nil {
		if !errIsNotFound(err) {
			return nil, nil, errors.Wrap(err, "error retrieving latest access proof")
		}
	}

	var convertedAccessRequest *api.OutgoingAccessRequest

	if latestAccessRequest != nil {
		convertedAccessRequest, err = convertOutgoingAccessRequest(latestAccessRequest)
		if err != nil {
			return nil, nil, errors.Wrap(err, "error converting latest access request")
		}
	}

	var convertedAccessProof *api.AccessProof

	if latestAccessProof != nil {
		convertedAccessProof, err = convertAccessProof(latestAccessProof)
		if err != nil {
			return nil, nil, errors.Wrap(err, "error converting latest access proof")
		}
	}

	return convertedAccessRequest, convertedAccessProof, nil
}

func convertAccessProof(accessProof *database.AccessProof) (*api.AccessProof, error) {
	createdAt, err := ptypes.TimestampProto(accessProof.CreatedAt)
	if err != nil {
		return nil, err
	}

	var revokedAt *timestamp.Timestamp

	if accessProof.RevokedAt.Valid {
		revokedAt, err = ptypes.TimestampProto(accessProof.RevokedAt.Time)
		if err != nil {
			return nil, err
		}
	}

	return &api.AccessProof{
		Id:               uint64(accessProof.ID),
		OrganizationName: accessProof.OutgoingAccessRequest.OrganizationName,
		ServiceName:      accessProof.OutgoingAccessRequest.ServiceName,
		CreatedAt:        createdAt,
		RevokedAt:        revokedAt,
		AccessRequestId:  uint64(accessProof.OutgoingAccessRequest.ID),
	}, nil
}
