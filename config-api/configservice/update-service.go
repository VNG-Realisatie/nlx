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

// UpdateService updates an existing service
func (s *ConfigService) UpdateService(ctx context.Context, req *configapi.UpdateServiceRequest) (*configapi.Service, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request UpdateService")

	err := req.Service.Validate()
	if err != nil {
		logger.Error("invalid service", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid service: %s", err))
	}

	err = s.configDatabase.UpdateService(ctx, req.Name, req.Service)
	if err != nil {
		logger.Error("error updating service in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req.Service, nil
}
