package server

import (
	"context"
	"path"

	types "github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/environment"
)

type DirectoryService struct {
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
func (s DirectoryService) ListServices(ctx context.Context, _ *api.Empty) (*api.DirectoryListServicesResponse, error) {
	s.logger.Info("rpc request ListServices")

	resp, err := s.directoryClient.ListServices(ctx, &inspectionapi.ListServicesRequest{})
	if err != nil {
		s.logger.Error("error getting services list from directory", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "directory error")
	}

	latestRequests, err := s.configDatabase.ListAllLatestOutgoingAccessRequests(ctx)
	if err != nil {
		s.logger.Error("error getting access requests", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "database error")
	}

	services := make([]*api.DirectoryService, len(resp.Services))

	for i, service := range resp.Services {
		services[i] = convertDirectoryService(service)

		key := path.Join(service.OrganizationName, service.ServiceName)
		if a, ok := latestRequests[key]; ok {
			latestAccessRequest, err := convertAccessRequest(a)
			if err != nil {
				s.logger.Error("error getting latest access request", zap.Error(err))
				return nil, status.Errorf(codes.Internal, "database error")
			}

			if latestAccessRequest != nil {
				services[i].LatestAccessRequest = latestAccessRequest
			}
		}
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

	latestAccessRequest, err := s.configDatabase.GetLatestOutgoingAccessRequest(ctx, request.OrganizationName, request.ServiceName)
	if err != nil {
		s.logger.Error("error getting access requests", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "database error")
	}

	if latestAccessRequest != nil {
		accessRequest, err := convertAccessRequest(latestAccessRequest)
		if err != nil {
			s.logger.Error("error converting access request", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "database error")
		}

		directoryService.LatestAccessRequest = accessRequest
	}

	return directoryService, nil
}

func (s DirectoryService) getService(ctx context.Context, logger *zap.Logger, organizationName, serviceName string) (*inspectionapi.ListServicesResponse_Service, error) {
	resp, err := s.directoryClient.ListServices(ctx, &inspectionapi.ListServicesRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, "directory not available")
	}

	for _, s := range resp.Services {
		if s.OrganizationName == organizationName && s.ServiceName == serviceName {
			return s, nil
		}
	}

	logger.Warn("service not found")

	return nil, status.Error(codes.NotFound, "service not found")
}

// RequestAccessToService records an access request and sends it to the organization
func (s DirectoryService) RequestAccessToService(ctx context.Context, request *api.RequestAccessToServiceRequest) (*api.AccessRequest, error) {
	logger := s.logger.With(zap.String("organizationName", request.OrganizationName), zap.String("serviceName", request.ServiceName))
	logger.Info("rpc request RequestAccessToService")

	ar := &database.AccessRequest{
		OrganizationName: request.OrganizationName,
		ServiceName:      request.ServiceName,
	}

	a, err := s.configDatabase.CreateAccessRequest(ctx, ar)
	if err != nil {
		return nil, err
	}

	response, err := convertDirectoryAccessRequest(a)
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

	return &api.DirectoryService{
		ServiceName:          s.ServiceName,
		OrganizationName:     s.OrganizationName,
		APISpecificationType: s.ApiSpecificationType,
		State:                serviceState,
	}
}

func convertDirectoryAccessRequest(a *database.AccessRequest) (*api.AccessRequest, error) {
	createdAt, err := types.TimestampProto(a.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := types.TimestampProto(a.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var accessRequestState api.AccessRequestState

	switch a.State {
	case database.AccessRequestFailed:
		accessRequestState = api.AccessRequestState_FAILED
	case database.AccessRequestCreated:
		accessRequestState = api.AccessRequestState_CREATED
	case database.AccessRequestReceived:
		accessRequestState = api.AccessRequestState_RECEIVED
	}

	return &api.AccessRequest{
		Id:        a.ID,
		State:     accessRequestState,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
