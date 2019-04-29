package configservice

import (
	"context"

	"go.nlx.io/nlx/config-api/configapi"
)

//ListComponents returns components who have announced them selfs to the config-API
func (s *ConfigService) ListComponents(ctx context.Context, req *configapi.Empty) (*configapi.ListComponentsResponse, error) {
	s.logger.Info("rpc request ListAnnouncements")
	resp := &configapi.ListComponentsResponse{}

	for _, value := range s.connections {
		resp.Components = append(resp.Components, &configapi.ListComponentsResponse_Component{
			Name: value.name,
			Kind: value.kind,
		})
	}
	return resp, nil
}
