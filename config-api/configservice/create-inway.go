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

// CreateInway creates a new inway
// TODO: Limit access to this call so only authorized components can access this endpoint
func (s *ConfigService) CreateInway(ctx context.Context, inway *configapi.Inway) (*configapi.Inway, error) {
	logger := s.logger.With(zap.String("name", inway.Name))
	logger.Info("rpc request CreateInway")

	err := inway.Validate()
	if err != nil {
		logger.Error("invalid inway", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid inway: %s", err))
	}

	err = s.configDatabase.CreateInway(ctx, inway)
	if err != nil {
		logger.Error("error creating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return inway, nil
}
