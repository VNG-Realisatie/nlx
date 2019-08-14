// nolint:dupl
package configservice

import (
	"context"

	"go.nlx.io/nlx/config-api/configapi"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetInway returns a specific inway
func (s *ConfigService) GetInway(ctx context.Context, req *configapi.GetInwayRequest) (*configapi.Inway, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request GetInway")

	inway, err := s.configDatabase.GetInway(ctx, req.Name)
	if err != nil {
		logger.Error("error getting inway from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if inway == nil {
		logger.Warn("inway not found")
		return nil, status.Error(codes.NotFound, "inway not found")
	}

	return inway, nil
}
