// nolint:dupl
package configservice

import (
	"context"

	"go.nlx.io/nlx/config-api/configapi"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteInway deletes a specific inway
func (s *ConfigService) DeleteInway(ctx context.Context, req *configapi.DeleteInwayRequest) (*configapi.Empty, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request DeleteInway")

	err := s.configDatabase.DeleteInway(ctx, req.Name)

	if err != nil {
		logger.Error("error deleting inway in DB", zap.Error(err))
		return &configapi.Empty{}, status.Error(codes.Internal, "database error")
	}

	return &configapi.Empty{}, nil
}
