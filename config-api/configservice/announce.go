package configservice

import (
	"context"

	"go.nlx.io/nlx/config-api/configapi"
)

type announcement struct {
	ComponentType string `json:"componentType"`
	FingerPrint   string `json:"fingerPrint"`
}

// Announce is called when a component wants to announce itself the the config API
func (s *ConfigService) Announce(ctx context.Context, req *configapi.AnnounceRequest) (*configapi.AnnounceResponse, error) {
	s.logger.Info("rpc request Announce")

	return &configapi.AnnounceResponse{}, nil
}
