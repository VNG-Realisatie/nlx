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
func (c *ConfigService) ListInways(ctx context.Context, req *configapi.ListInwaysRequest) (*configapi.ListInwaysResponse, error) {
	// TODO: Limit access to this call so only authorized components can access this endpoint
	c.logger.Info("rpc request ListInways")

	inways, err := c.configDatabase.ListInways(ctx)
	if err != nil {
		c.logger.Error("error getting inway list from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	services, err := c.configDatabase.ListServices(ctx)
	if err != nil {
		c.logger.Error("error getting services  from database", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	if len(services) > 0 {
		for _, i := range inways {
			for _, s := range services {
				if contains(s.Inways, i.Name) {
					i.Services = append(i.Services, &configapi.Inway_Service{Name: s.Name})
				}
			}
		}
	}

	return &configapi.ListInwaysResponse{
		Inways: inways,
	}, nil
}

func contains(s []string, v string) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}

	return false
}
