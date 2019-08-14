// nolint:dupl
package configservice

import (
	"context"

	"go.nlx.io/nlx/config-api/configapi"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteService deletes a specific service
func (s *ConfigService) DeleteService(ctx context.Context, req *configapi.DeleteServiceRequest) (*configapi.Empty, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteService")

	err := s.configDatabase.DeleteService(ctx, req.Name)

	if err != nil {
		logger.Error("error deleting service in DB", zap.Error(err))
		return &configapi.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &configapi.Empty{}, nil
}
