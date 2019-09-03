// nolint:dupl
package configservice

import (
	"context"
	"strings"

	"go.nlx.io/nlx/config-api/configapi"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListServices returns a list of services
func (s *ConfigService) ListServices(ctx context.Context, req *configapi.ListServicesRequest) (*configapi.ListServicesResponse, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	s.logger.Info("rpc request ListServices")

	services, err := s.configDatabase.ListServices(ctx)
	if err != nil {
		s.logger.Error("error getting services list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	response := &configapi.ListServicesResponse{}
	if len(req.InwayName) > 0 {
		for _, service := range services {
			for _, inway := range service.Inways {
				if strings.Compare(req.InwayName, inway) == 0 {
					response.Services = append(response.Services, service)
					break
				}
			}
		}
	} else {
		response.Services = services
	}

	return response, nil
}
