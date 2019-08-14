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

// UpdateInway updates an existing inway
func (s *ConfigService) UpdateInway(ctx context.Context, req *configapi.UpdateInwayRequest) (*configapi.Inway, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	logger := s.logger.With(zap.String("name", req.Name))
	logger.Info("rpc request UpdateInway")

	err := req.Inway.Validate()
	if err != nil {
		logger.Error("invalid inway", zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid inway: %s", err))
	}

	err = s.configDatabase.UpdateInway(ctx, req.Name, req.Inway)
	if err != nil {
		logger.Error("error updating inway in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req.Inway, nil
}
