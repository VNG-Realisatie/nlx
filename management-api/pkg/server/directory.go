package server

import (
	"context"
	"path"

	types "github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

var inwayStateToDirectoryState = map[directory.InwayState]api.DirectoryService_State{
	directory.InwayStateUnknown: api.DirectoryService_unknown,
	directory.InwayStateUp:      api.DirectoryService_up,
	directory.InwayStateDown:    api.DirectoryService_down,
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

	listServices, err := s.directoryClient.ListServices()
	if err != nil {
		s.logger.Error("error getting services list from directory", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "directory error")
	}

	latestRequests, err := s.configDatabase.ListAllLatestOutgoingAccessRequests(ctx)
	if err != nil {
		s.logger.Error("error getting access requests", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "database error")
	}

	services := make([]*api.DirectoryService, len(listServices))

	for i, service := range listServices {
		services[i] = convertDirectoryService(service)

		key := path.Join(service.OrganizationName, service.Name)
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

	service, err := s.getService(logger, request.OrganizationName, request.ServiceName)
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

func (s DirectoryService) getService(logger *zap.Logger, organizationName, serviceName string) (*directory.InspectionAPIService, error) {
	services, err := s.directoryClient.ListServices()
	if err != nil {
		return nil, status.Error(codes.Internal, "directory not available")
	}

	for _, s := range services {
		if s.OrganizationName == organizationName && s.Name == serviceName {
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

	service, err := s.getService(logger, request.OrganizationName, request.ServiceName)
	if err != nil {
		return nil, err
	}

	logger = logger.With(zap.Any("service", service))

	logger.Debug("send access request to inway")

	return response, nil
}

func DetermineDirectoryServiceState(inways []*directory.Inway) api.DirectoryService_State {
	serviceState := api.DirectoryService_unknown

	if len(inways) == 0 {
		return serviceState
	}

	stateMap := map[directory.InwayState]int{}

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

func convertDirectoryService(s *directory.InspectionAPIService) *api.DirectoryService {
	serviceState := DetermineDirectoryServiceState(s.Inways)

	return &api.DirectoryService{
		ServiceName:          s.Name,
		OrganizationName:     s.OrganizationName,
		APISpecificationType: s.APISpecificationType,
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
