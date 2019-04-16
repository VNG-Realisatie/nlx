package configservice

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/config-api/configapi"
)

// SetConfig is used to set a config for a component
func (s *ConfigService) SetConfig(ctx context.Context, req *configapi.SetConfigRequest) (*configapi.SetConfigResponse, error) {
	s.logger.Info("rpc request SetConfig")
	_, err := s.etcdCli.Put(context.Background(), fmt.Sprintf("%s", req.ComponentName), req.Config.Config)
	if err != nil {
		s.logger.Error("error writing config to DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return &configapi.SetConfigResponse{}, nil
}
