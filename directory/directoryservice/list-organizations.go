package directoryservice

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/directory/directoryapi"
)

type listOrganizationsHandler struct {
	logger *zap.Logger

	demoEnv    string
	demoDomain string

	stmtSelectOrganizations *sqlx.Stmt
}

func newListOrganizationsHandler(db *sqlx.DB, logger *zap.Logger, demoEnv string, demoDomain string) (*listOrganizationsHandler, error) {
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
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSelectOrganizations")
	}

	return h, nil
}

func (h *listOrganizationsHandler) ListOrganizations(ctx context.Context, req *directoryapi.ListOrganizationsRequest) (*directoryapi.ListOrganizationsResponse, error) {
	fmt.Println("rpc request ListOrganizations()")
	resp := &directoryapi.ListOrganizationsResponse{}

	err := h.stmtSelectOrganizations.Select(&resp.Organizations)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select organizations using stmtSelectOrganizations")
	}

	// Hardcoded
	for _, org := range resp.Organizations {
		switch org.Name {
		case "gemeente", "rdw", "brp":
			org.InsightLogEndpoint = fmt.Sprintf("insight-txlog-api.%s.%s.%s", h.demoEnv, strings.ToLower(org.Name), h.demoDomain)
			org.InsightIrmaEndpoint = fmt.Sprintf("insight-irma-api.%s.%s.%s", h.demoEnv, strings.ToLower(org.Name), h.demoDomain)
		}
	}

	return resp, nil
}
