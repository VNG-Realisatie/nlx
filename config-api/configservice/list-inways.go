// nolint:dupl
package configservice

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/config-api/configapi"
)

// ListInways returns a list of inways
func (s *ConfigService) ListInways(ctx context.Context, req *configapi.ListInwaysRequest) (*configapi.ListInwaysResponse, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	s.logger.Info("rpc request ListInways")

	inways, err := s.configDatabase.ListInways(ctx)
	if err != nil {
		s.logger.Error("error getting inway list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &configapi.ListInwaysResponse{
		Inways: inways,
	}, nil
}
