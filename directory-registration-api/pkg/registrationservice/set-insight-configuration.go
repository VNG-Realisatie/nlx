// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package registrationservice

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

type SetInsightConfigurationHandler struct {
	logger *zap.Logger

	stmtSetInsightConfiguration *sqlx.Stmt
}

func newSetInsightConfigurationHandler(
	db *sqlx.DB,
	logger *zap.Logger,
) (*SetInsightConfigurationHandler, error) {
	h := &SetInsightConfigurationHandler{
		logger: logger.With(zap.String("handler", "set-insight-configuration")),
	}

	var err error

	// NOTE: We do not have an endpoint yet to create services separately, therefore insert on demand.
	h.stmtSetInsightConfiguration, err = db.Preparex(`
		INSERT INTO directory.organizations (name, insight_log_endpoint, insight_irma_endpoint)
			VALUES ($1, $2, $3)
			ON CONFLICT ON CONSTRAINT organizations_uq_name
				DO UPDATE SET
					insight_log_endpoint = COALESCE(NULLIF(EXCLUDED.insight_log_endpoint, ''), organizations.insight_log_endpoint),
					insight_irma_endpoint = COALESCE(NULLIF(EXCLUDED.insight_irma_endpoint, ''), organizations.insight_irma_endpoint)
			RETURNING id
	`)

	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSetInsightConfiguration")
	}

	return h, nil
}

func (h *SetInsightConfigurationHandler) SetInsightConfiguration(ctx context.Context, req *registrationapi.SetInsightConfigurationRequest) (*registrationapi.Empty, error) {
	h.logger.Info("rpc request SetInsightConfiguration", zap.String("insight api url", req.InsightAPIURL), zap.String("irma server url", req.IrmaServerURL))
	resp := &registrationapi.Empty{}
	organizationName, err := getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, err
	}

	if !validateName(organizationName) {
		h.logger.Info("invalid organization name in setinsightconfigrationrequest", zap.String("organization name", organizationName))
		return nil, status.New(codes.InvalidArgument, "Invalid organization name").Err()
	}

	_, err = h.stmtSetInsightConfiguration.Exec(organizationName, req.InsightAPIURL, req.IrmaServerURL)
	if err != nil {
		statusCode := codes.Internal
		h.logger.Error("failed to execute stmtSetInsightConfiguration", zap.Error(err))
		return nil, status.New(statusCode, FriendlyErrorDatabase).Err()
	}
	return resp, nil
}
