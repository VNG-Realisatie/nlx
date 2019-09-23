// nolint:dupl
package configservice

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

// PutInsight sets the insight configuration of a organization
func (s *ConfigService) PutInsightConfiguration(ctx context.Context, req *configapi.InsightConfiguration) (*configapi.InsightConfiguration, error) {
	s.logger.Info("rpc request PutInsight")

	_, err := s.directoryRegistrationClient.SetInsightConfiguration(ctx, &registrationapi.SetInsightConfigurationRequest{
		InsightAPIURL: req.InsightAPIURL,
		IrmaServerURL: req.IrmaServerURL,
	})

	if err != nil {
		return nil, err
	}

	err = s.configDatabase.PutInsightConfiguration(ctx, req)
	if err != nil {
		s.logger.Error("error updating inway insight config in DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "database error")
	}

	return req, nil
}
