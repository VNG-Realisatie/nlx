// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
)

type listOrganizationsHandler struct {
	logger *zap.Logger

	demoEnv    string
	demoDomain string

	stmtSelectOrganizations *sqlx.Stmt
}

func newListOrganizationsHandler(
	db *sqlx.DB, logger *zap.Logger,
	demoEnv, demoDomain string,
) (*listOrganizationsHandler, error) {
	h := &listOrganizationsHandler{
		logger:     logger.With(zap.String("handler", "list-organizations")),
		demoEnv:    demoEnv,
		demoDomain: demoDomain,
	}

	var err error
	h.stmtSelectOrganizations, err = db.Preparex(`
		SELECT
			name,
			COALESCE(insight_irma_endpoint, '') AS insight_irma_endpoint,
			COALESCE(insight_log_endpoint, '') AS insight_log_endpoint
		FROM directory.organizations
		ORDER BY name
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSelectOrganizations")
	}

	return h, nil
}

func (h *listOrganizationsHandler) ListOrganizations(ctx context.Context, req *inspectionapi.ListOrganizationsRequest) (*inspectionapi.ListOrganizationsResponse, error) {
	h.logger.Info("rpc request ListOrganizations")
	resp := &inspectionapi.ListOrganizationsResponse{}

	err := h.stmtSelectOrganizations.Select(&resp.Organizations)
	if err != nil {
		h.logger.Error("failed to select organizations using stmtSelectOrganizations", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	return resp, nil
}
