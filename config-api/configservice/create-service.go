// nolint:dupl
package configservice

import (
	"context"
	"fmt"

	"go.nlx.io/nlx/config-api/configapi"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateService creates a new service
func (s *ConfigService) CreateService(ctx context.Context, service *configapi.Service) (*configapi.Service, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	logger := s.logger.With(zap.String("name", service.Name))
	logger.Info("rpc request CreateService")

	err := service.Validate()
	if err != nil {
		logger.Error("invalid service", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid service: %s", err))
	}

	err = s.configDatabase.CreateService(ctx, service)
	if err != nil {
		logger.Error("error creating service in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return service, nil
}
