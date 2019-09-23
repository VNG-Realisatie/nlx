// nolint:dupl
package configservice

import (
	"context"

	"go.nlx.io/nlx/config-api/configapi"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetInsight returns the insight configuration of a organization
func (s *ConfigService) GetInsightConfiguration(ctx context.Context, req *configapi.Empty) (*configapi.InsightConfiguration, error) {
	s.logger.Info("rpc request GetInsightConfiguration")

	insightConfig, err := s.configDatabase.GetInsightConfiguration(ctx)
	if err != nil {
		s.logger.Error("error getting insight from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if insightConfig == nil {
		s.logger.Warn("insight configuration not found")
		return nil, status.Error(codes.NotFound, "insight configuration not found")
	}

	return insightConfig, nil
}
