package directory

import (
	"context"
	"path"

	types "github.com/gogo/protobuf/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/environment"
)

type Service struct {
	logger      *zap.Logger
	environment *environment.Environment

	directoryClient Client
	configDatabase  database.ConfigDatabase
}

var inwayStateToDirectoryStatus = map[InwayState]DirectoryService_Status{
	InwayStateUnknown: DirectoryService_unknown,
	InwayStateUp:      DirectoryService_up,
	InwayStateDown:    DirectoryService_down,
}

func NewDirectoryService(logger *zap.Logger, e *environment.Environment, directoryClient Client, configDatabase database.ConfigDatabase) *Service {
	return &Service{
		logger:          logger,
		environment:     e,
		directoryClient: directoryClient,
		configDatabase:  configDatabase,
	}
}

// ListServices returns all services known to the directory except those with the same organization
func (s Service) ListServices(ctx context.Context, _ *Empty) (*ListServicesResponse, error) {
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

	services := make([]*DirectoryService, len(listServices))

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

	return &ListServicesResponse{Services: services}, nil
}

// GetOrganizationService returns a specific service of and organization
func (s Service) GetOrganizationService(ctx context.Context, request *GetOrganizationServiceRequest) (*DirectoryService, error) {
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

func (s Service) getService(logger *zap.Logger, organizationName, serviceName string) (*InspectionAPIService, error) {
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
func (s Service) RequestAccessToService(ctx context.Context, request *RequestAccessToServiceRequest) (*AccessRequest, error) {
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

	response, err := convertAccessRequest(a)
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

func DetermineDirectoryServiceStatus(inways []*Inway) DirectoryService_Status {
	serviceStatus := DirectoryService_unknown

	if len(inways) == 0 {
		return serviceStatus
	}

	stateMap := map[InwayState]int{}

	for _, i := range inways {
		stateMap[i.State]++
	}

	if len(stateMap) > 1 {
		return DirectoryService_degraded
	}

	for state := range stateMap {
		serviceStatus = inwayStateToDirectoryStatus[state]
	}

	return serviceStatus
}

func convertDirectoryService(s *InspectionAPIService) *DirectoryService {
	serviceStatus := DetermineDirectoryServiceStatus(s.Inways)

	return &DirectoryService{
		ServiceName:          s.Name,
		OrganizationName:     s.OrganizationName,
		APISpecificationType: s.APISpecificationType,
		Status:               serviceStatus,
	}
}

func convertAccessRequest(a *database.AccessRequest) (*AccessRequest, error) {
	createdAt, err := types.TimestampProto(a.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := types.TimestampProto(a.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var accessRequestState AccessRequest_Status

	switch a.Status {
	case database.AccessRequestFailed:
		accessRequestState = AccessRequest_FAILED
	case database.AccessRequestCreated:
		accessRequestState = AccessRequest_CREATED
	case database.AccessRequestReceived:
		accessRequestState = AccessRequest_RECEIVED
	}

	return &AccessRequest{
		Id:        a.ID,
		Status:    accessRequestState,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
