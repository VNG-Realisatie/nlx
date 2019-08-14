// nolint:dupl
package configservice

import (
	"context"

	"go.nlx.io/nlx/config-api/configapi"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetService returns a specific service
func (s *ConfigService) GetService(ctx context.Context, req *configapi.GetServiceRequest) (*configapi.Service, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request GetService")

	service, err := s.configDatabase.GetService(ctx, req.Name)
	if err != nil {
		logger.Error("error getting service from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if service == nil {
		logger.Warn("service not found")
		return nil, status.Error(codes.NotFound, "service not found")
	}

	return service, nil
}
